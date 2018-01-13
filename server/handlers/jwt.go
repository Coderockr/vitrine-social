package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTOptions struct {
	SigningMethod string
	PublicKey     []byte
	PrivateKey    []byte
	Expiration    time.Duration
}

// generateJWTToken Generates a JSON Web Token given an userId (typically an id or an email), and the JWT options
// to set SigningMethod and the keys you can check
// http://github.com/dgrijalva/jwt-go
//
// In case you use an symmetric-key algorithm set PublicKey and PrivateKey equal to the SecretKey ,
func generateJWTToken(userID int64, op JWTOptions) (string, error) {

	now := time.Now()
	// set claims
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(op.Expiration).Unix(),
		Subject:   strconv.Itoa(int(userID)),
		Id:        string(generateRandomKey(32)),
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(op.SigningMethod), claims)

	return t.SignedString(op.PrivateKey)
}

// validateToken Validates the token that is passed in the request with the Authorization header
// Authorization: Bearer eyJhbGciOiJub25lIn0
//
// Returns the userId, token (base64 encoded), error
func validateToken(r *http.Request, options JWTOptions) (int64, string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0, "", errors.New("Unauthorized")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		switch token.Method {
		case jwt.SigningMethodHS256:
			return options.PrivateKey, nil
		case jwt.SigningMethodRS256:
			return options.PublicKey, nil
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
				return 0, "", errors.New("Token Expired, get a new one")
			default:
				return 0, "", errors.New("JWT Token ValidationError")
			}
		}
		return 0, "", errors.New("JWT Token Error Parsing the token or empty token")
	}
	claims, ok := token.Claims.(jwt.StandardClaims)
	if !ok || !token.Valid {
		return 0, "", errors.New("JWT Token is not Valid")
	}
	r.Header.Add("user_id", claims.Subject)
	userID, err := strconv.Atoi(claims.Subject)
	return int64(userID), token.Raw, err
}
