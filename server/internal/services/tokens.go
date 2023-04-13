package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type tokensService struct {
	TokenKeys     models.TokenKeys
	SigningMethod jwt.SigningMethod
}

func NewTokensService(token_keys models.TokenKeys, signingMethod jwt.SigningMethod) *tokensService {
	return &tokensService{
		TokenKeys:     token_keys,
		SigningMethod: signingMethod,
	}
}

func (ss *tokensService) SignTokenWithCustomClaims(token *jwt.Token) (string, error) {

	// token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, models.CustomJWTClaims{
	// 	CustomClaims: *customClaims,
	// })

	signedToken, err := token.SignedString(ss.TokenKeys.PrivateKey)
	if err != nil {
		fmt.Println("error occurred, err: ", err)
		return "", err
	}

	return signedToken, nil

}

func (ss *tokensService) validateSignedToken(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != ss.SigningMethod.Alg() {
		fmt.Println("error during method comparison")
		return nil, errors.New("signing methods are not identical")
	}

	return ss.TokenKeys.PublicKey, nil
}

func (ss *tokensService) ParseSignedToken(signedToken string) (*jwt.Token, error) {
	parsed_token, err := jwt.ParseWithClaims(
		signedToken,
		&models.CustomJWTClaims{},
		ss.validateSignedToken,
		jwt.WithLeeway(5*time.Second),
	)

	if err != nil {
		fmt.Println("error here: ", err)
		return parsed_token, err

	}

	return parsed_token, err

}
