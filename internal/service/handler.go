// sentiric-identity-service/internal/service/handler.go
package service

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	userv1 "github.com/sentiric/sentiric-contracts/gen/go/sentiric/user/v1"
	"github.com/sentiric/sentiric-identity-service/internal/config"
	"github.com/sentiric/sentiric-identity-service/internal/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IdentityHandler struct {
	userv1.UnimplementedIdentityServiceServer
	cfg          *config.Config
	log          zerolog.Logger
	tokenManager *token.JWTManager
}

func NewIdentityHandler(cfg *config.Config, log zerolog.Logger) *IdentityHandler {
	// JWTManager'ı konfigürasyondan oluştur
	tm := token.NewJWTManager(cfg.JWTSecret, time.Hour*24)
	return &IdentityHandler{
		cfg:          cfg,
		log:          log,
		tokenManager: tm,
	}
}

// Authenticate, kullanıcı adı ve şifre ile kimlik doğrulaması yapar ve JWT üretir.
func (h *IdentityHandler) Authenticate(ctx context.Context, req *userv1.AuthenticateRequest) (*userv1.AuthenticateResponse, error) {
	log := h.log.With().Str("rpc", "Authenticate").Str("username", req.GetUsername()).Logger()

	// TODO: 1. User Service'e şifre kontrolü için gRPC çağrısı (veya doğrudan DB/Redis kontrolü)
	if req.GetUsername() != "admin" || req.GetPassword() != "adminpass" {
		log.Warn().Msg("Geçersiz kimlik bilgisi (Placeholder)")
		return nil, status.Errorf(codes.Unauthenticated, "Kullanıcı adı veya şifre hatalı.")
	}

	// Simüle edilmiş kullanıcı kimlikleri
	userID := "user-123"
	tenantID := "tenant-sentiric"

	// 2. Token oluştur
	accessToken, err := h.tokenManager.Generate(userID, tenantID)
	if err != nil {
		log.Error().Err(err).Msg("JWT oluşturma hatası")
		return nil, status.Errorf(codes.Internal, "Token oluşturulamadı.")
	}

	log.Info().Str("user_id", userID).Msg("Kimlik doğrulama başarılı, token oluşturuldu.")
	return &userv1.AuthenticateResponse{
		AccessToken: accessToken,
		UserId:      userID,
	}, nil
}

// AuthorizeToken, gelen JWT'yi doğrular ve yük (claims) bilgilerini döndürür.
func (h *IdentityHandler) AuthorizeToken(ctx context.Context, req *userv1.AuthorizeTokenRequest) (*userv1.AuthorizeTokenResponse, error) {
	log := h.log.With().Str("rpc", "AuthorizeToken").Logger()

	claims, err := h.tokenManager.Verify(req.GetAccessToken())
	if err != nil {
		log.Warn().Err(err).Msg("Token doğrulama başarısız.")
		return &userv1.AuthorizeTokenResponse{IsValid: false}, nil
	}

	log.Debug().Str("user_id", claims.UserID).Msg("Token doğrulama başarılı.")
	return &userv1.AuthorizeTokenResponse{
		IsValid:  true,
		UserId:   claims.UserID,
		TenantId: claims.TenantID,
	}, nil
}
