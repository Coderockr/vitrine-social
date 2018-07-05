package handlers

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"

	jwt "github.com/dgrijalva/jwt-go"
)

type (
	// JWTOptions contains the needed information to validate and generate jwt tokens
	JWTOptions struct {
		SigningMethod string
		PublicKey     []byte
		PrivateKey    []byte
		Expiration    time.Duration
	}

	// JWTManager actually validate and generate the tokens
	JWTManager struct {
		OP JWTOptions
	}
)

type tokenFormat struct {
	UserID      int64     `json:"userId"`
	Permissions *[]string `json:"permissions"`
}

// CreateToken Generates a JSON Web Token given an userId (typically an id or an email), and the JWT options
// to set SigningMethod and the keys you can check
// http://github.com/dgrijalva/jwt-go
//
// In case you use an symmetric-key algorithm set PublicKey and PrivateKey equal to the SecretKey ,
func (m *JWTManager) CreateToken(u model.User, permissions *[]string) (string, error) {

	b, _ := json.Marshal(tokenFormat{
		UserID:      u.ID,
		Permissions: permissions,
	})

	now := time.Now()
	// set claims
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(m.OP.Expiration).Unix(),
		Subject:   string(b),
		Id:        string(generateRandomKey(32)),
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(m.OP.SigningMethod), claims)

	return t.SignedString(m.OP.PrivateKey)
}

// ValidateToken Validates the token that is passed in the request with the Authorization header
// Authorization: Bearer eyJhbGciOiJub25lIn0
//
// Returns the userId, token (base64 encoded), error
func (m *JWTManager) ValidateToken(tokenString string) (*model.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		switch token.Method {
		case jwt.SigningMethodHS256:
			return m.OP.PrivateKey, nil
		case jwt.SigningMethodRS256:
			return m.OP.PublicKey, nil
		default:
			return nil, errors.New("JWT Token is not Valid")
		}
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				return nil, errors.New("Token Expired, get a new one")
			default:
				log.Printf("[INFO][Auth Middleware] %s", vErr.Error())
				return nil, errors.New("JWT Token ValidationError")
			}
		}
		return nil, errors.New("JWT Token Error Parsing the token or empty token")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("JWT Token is not Valid")
	}

	if userID, err := strconv.Atoi(claims.Subject); err == nil {
		return &model.Token{UserID: int64(userID)}, nil
	}

	v := tokenFormat{}

	err = json.NewDecoder(strings.NewReader(claims.Subject)).Decode(&v)
	if err != nil {
		log.Printf("[INFO][Auth Middleware] TokenManager was not able to decode the Subject: %s", claims.Subject)
		return nil, errors.New("JWT token has a unknown subject format")
	}

	t := model.Token{
		UserID:      v.UserID,
		Permissions: make(map[string]bool),
		Token:       tokenString,
	}
	if v.Permissions != nil {
		for _, p := range *v.Permissions {
			t.Permissions[p] = true
		}
	}

	return &t, err
}

func generateRandomKey(strength int) []byte {
	k := make([]byte, strength)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
