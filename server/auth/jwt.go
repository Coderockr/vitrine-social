package auth

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrTokenExpired    = errors.New("Token Expired, get a new one")
	ErrTokenValidation = errors.New("JWT Token ValidationError")
	ErrTokenParse      = errors.New("JWT Token Error Parsing the token or empty token")
	ErrTokenInvalid    = errors.New("JWT Token is not Valid")

	logOn = true
)

type Options struct {
	SigningMethod string
	PublicKey     string
	PrivateKey    string
	Expiration    time.Duration
}

// generateJWTToken Generates a JSON Web Token given an userId (typically an id or an email), and the JWT options
// to set SigningMethod and the keys you can check
// http://github.com/dgrijalva/jwt-go
//
// In case you use an symmetric-key algorithm set PublicKey and PrivateKey equal to the SecretKey ,
func generateJWTToken(userID int64, op Options) (string, error) {

	now := time.Now()
	// set claims
	claims := jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(op.Expiration).Unix(),
		Subject:   strconv.Itoa(int(userID)),
		Id:        string(generateRandomKey(32)),
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod(op.SigningMethod), claims)

	tokenString, err := t.SignedString([]byte(op.PrivateKey))
	if err != nil {
		logError("ERROR: GenerateJWTToken: %v\n", err)
	}
	return tokenString, err

}

// validateToken Validates the token that is passed in the request with the Authorization header
// Authorization: Bearer eyJhbGciOiJub25lIn0
//
// Returns the userId, token (base64 encoded), error
func validateToken(r *http.Request, publicKey string) (int64, string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0, "", errors.New("Unauthorized")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return 0, "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return 0, "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return 0, "", ErrTokenParse
	}
	claims, ok := token.Claims.(jwt.StandardClaims)
	if ok && token.Valid {
		r.Header.Add("user_id", claims.Subject)
	} else {
		return 0, "", ErrTokenInvalid
	}
	userID, err := strconv.Atoi(claims.Subject)
	return int64(userID), token.Raw, err
}

func logError(format string, err interface{}) {
	if logOn && err != nil {
		log.Printf(format, err)
	}
}
