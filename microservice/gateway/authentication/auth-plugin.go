package authentication

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	AccessTokenType       = "access"
	RefreshTokenType      = "refresh"
	OtpTokenType          = "otp"
	OtpResendType         = "otp-resend"
	EmailTokenType        = "email"
	ChildModeOtpTokenType = "child-otp"
	WebOtpTokenType       = "web-otp"
)

type TokenPayloadType string

var TokenPayloadKey TokenPayloadType = "auth-payload"
var pluginName = "auth-handler"

func init() {
	fmt.Println("Custom Plugin Loaded!")
}

func RegisterHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {

	config, ok := extra[pluginName+"-value"].(map[string]interface{})
	if !ok {
		return h, errors.New("configuration not found")
	}

	secretKey := "WBmppP/4PnivFd3XGXwlYpVwAI7dtanMTiWmuetXOmQ="
	authPlugin := NewAuthPlugin([]byte(secretKey))
	passCustomData, _ := config["config"].(map[string]interface{})
	header, _ := passCustomData["header"].(string)
	value, _ := passCustomData["value"].(string)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		requestID := req.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			req.Header.Set("X-Request-ID", requestID)
		}

		fmt.Printf("Request ID: %s\n", requestID)

		if header != "" && value != "" {
			req.Header.Set(header, value)
			fmt.Printf("Custom Header Added: %s: %s\n", header, value)
		}

		fmt.Println("Header After->", req.Header)

		// Token validation logic here
		payload, err := authPlugin.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}
		req.Header.Set("SID", payload.Sid.String())
		h.ServeHTTP(w, req)
	}), nil
}

// AuthPlugin structure for token validation
type AuthPlugin struct {
	secretKey []byte
}

func NewAuthPlugin(secretKey []byte) *AuthPlugin {
	fmt.Println("Secret key loaded:", secretKey)
	return &AuthPlugin{secretKey: secretKey}
}

func main() {}
