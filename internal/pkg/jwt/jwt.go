package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Verifier interface {
		Verify(string) (*Claims, error)
	}

	Signer interface {
		Sign(claims Claims) (string, error)
	}

	SignVerifier interface {
		Signer
		Verifier
	}

	Claims struct {
		jwt.StandardClaims
		Roles     []string
		AvatarURL string
		Name      string
		LastName  string
		FirstName string
		UserID    string
	}

	Config struct {
		JWTSecret string `envconfig:"JWT_SECRET" default:"a star in the sky..."`
	}

	Generator struct {
		Config        Config
		SigningMethod jwt.SigningMethod
	}

	StandardClaims = jwt.StandardClaims
)

const (
	// DefaultIssuer is default issuer name
	DefaultIssuer = "robusta"
	// DefaultLifeTime is default life time of a token
	DefaultLifeTime = time.Hour * 24
)

var (
	// ErrInvalidToken report that the JWT token is invalid
	ErrInvalidToken = errors.New("invalid token")
)

// New return new instance of JWTGenerator with configuration provided via env variable
func New(conf Config) *Generator {
	return &Generator{
		Config:        conf,
		SigningMethod: jwt.SigningMethodHS256,
	}
}

// NewWithConfig return new instance from config
func NewWithConfig(conf Config) *Generator {
	return &Generator{
		Config:        conf,
		SigningMethod: jwt.SigningMethodHS256,
	}
}

// Sign generate JWT token base on the given claims
func (g *Generator) Sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(g.SigningMethod, claims)
	v, err := token.SignedString([]byte(g.Config.JWTSecret))
	return v, err
}

// Verify verify if the given token is valid
func (g *Generator) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.Config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
