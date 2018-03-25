package handlers

import (
	"crypto/rand"
	"errors"
	"io"
	"log"
	"strconv"
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

// CreateToken Generates a JSON Web Token given an userId (typically an id or an email), and the JWT options
// to set SigningMethod and the keys you can check
// http://github.com/dgrijalva/jwt-go
//
// In case you use an symmetric-key algorithm set PublicKey and PrivateKey equal to the SecretKey ,
func (m *JWTManager) CreateToken(u model.User) (string, error) {

	now := time.Now()
	// set claims
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(m.OP.Expiration).Unix(),
		Subject:   strconv.FormatInt(u.ID, 10),
		Id:        string(generateRandomKey(32)),
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(m.OP.SigningMethod), claims)

	return t.SignedString(m.OP.PrivateKey)
}

// ValidateToken Validates the token that is passed in the request with the Authorization header
// Authorization: Bearer eyJhbGciOiJub25lIn0
//
// Returns the userId, token (base64 encoded), error
func (m *JWTManager) ValidateToken(tokenString string) (int64, error) {
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
				return 0, errors.New("Token Expired, get a new one")
			default:
				log.Printf("[INFO][Auth Middleware] %s", vErr.Error())
				return 0, errors.New("JWT Token ValidationError")
			}
		}
		return 0, errors.New("JWT Token Error Parsing the token or empty token")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		log.Printf("%#v",token.Claims)
		return 0, errors.New("JWT Token is not Valid")
	}
	userID, err := strconv.Atoi(claims.Subject)
	return int64(userID), err
}

func generateRandomKey(strength int) []byte {
	k := make([]byte, strength)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
