package models

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Token JWT claim struct
type Token struct {
	UserName string
	Roles    []Role
	jwt.StandardClaims
}

func generateToken(userName string, roles []Role) (string, error) {
	tk := &Token{
		UserName: userName,
		Roles:    roles,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	return token.SignedString([]byte(os.Getenv("TOKEN_PASWORD")))

}

// Parse Token
func (claim *Token) Parse(bearerToken string) error {
	//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if bearerToken == "" {
		return errors.New("Access Restricted, Missing Auth key")
	}

	splitted := strings.Split(bearerToken, " ")
	if len(splitted) != 2 {
		return errors.New("Access restricted, Invalid key")
	}

	tokenPart := splitted[1]
	token, err := jwt.ParseWithClaims(tokenPart, claim, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("TOKEN_PASWORD")), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Invalid key, Server not recognize the key")
	}
	return nil
}
