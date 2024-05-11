package services

import (
	"context"

	firebase "firebase.google.com/go"
)

func VerifyAuthToken(idToken string) (string, error) {
	// Initialize the Firebase Admin SDK
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return "", err
	}

	// Get the Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return "", err
	}

	// Verify the ID token
	token, err := authClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}

	// Token is valid, return the UID
	return token.UID, nil
}
