package firebase

import (
	"context"
	"errors"
	"fmt"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofrs/uuid"
	"google.golang.org/api/option"
)

var (
	client *auth.Client
)

type FirebasePayload struct {
	UID      string    // Firebase UID
	Email    string    // Firebase Email
	Fullname string    // Firebase Full Name
	Phone    string    // Firebase Phone Number
	UserId   uuid.UUID // Your internal DB UUID (populated later)
	Provider string    // Firebase sign-in provider (e.g., password, google, phone)
	Role     string    // Your internal DB role
}

func InitFirebaseAuth(filePath string) error {
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, filePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// Return the error instead of killing the server process here
		return fmt.Errorf("error initializing firebase app: %w", err)
	}

	_ = app

	return nil
}

func VerifyFirebaseIDToken(ctx context.Context, idtoken string) (*FirebasePayload, error) {
	if client == nil {
		return &FirebasePayload{}, errors.New("firebase auth client not initialized")
	}

	token, err := client.VerifyIDToken(context.Background(), idtoken)
	if err != nil {
		return &FirebasePayload{}, err
	}

	if token.Expires < time.Now().Unix() {
		return &FirebasePayload{}, errors.New("token expired")
	}

	claims := token.Claims

	return &FirebasePayload{
		UID:      token.UID,
		Email:    getClaimString(claims, "email"),
		Fullname: getClaimString(claims, "name"),
		Phone:    getClaimString(claims, "phone_number"),
		Provider: getSignInProvider(claims),
	}, nil
}

func getClaimString(claims map[string]interface{}, key string) string {
	if val, ok := claims[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getSignInProvider(claims map[string]interface{}) string {
	if firebaseClaim, ok := claims["firebase"].(map[string]interface{}); ok {
		if provider, ok := firebaseClaim["sign_in_provider"].(string); ok {
			return provider
		}
	}
	return ""
}
