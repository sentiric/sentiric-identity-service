// sentiric-identity-service/internal/token/jwt.go
package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey string
	tokenDuration time.Duration
}

type UserClaims struct {
	UserID string `json:"user_id"`
	TenantID string `json:"tenant_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

func (m *JWTManager) Generate(userID string, tenantID string) (string, error) {
	claims := UserClaims{
		UserID: userID,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "sentiric-identity-service",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Token'ı imzala
	signedToken, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", fmt.Errorf("token imzalanırken hata oluştu: %w", err)
	}
	
	return signedToken, nil
}

func (m *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("beklenmedik imzalama metodu: %s", t.Header["alg"])
			}
			return []byte(m.secretKey), nil
		},
	)
	
	if err != nil {
		return nil, fmt.Errorf("token doğrulanamadı: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("token claims tipi hatalı")
	}
	
	return claims, nil
}