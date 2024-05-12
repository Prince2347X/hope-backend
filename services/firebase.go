package services

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseServices struct{}

func (FirebaseServices) VerifyAuthToken(idToken string) (string, error) {
	// Initialize the Firebase Admin SDK
	options := option.WithCredentialsFile("secrets/hope2347x-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		return "", err
	}

	// Get the Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return "", err
	}

	// Verify the ID token
	token, err := authClient.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		return "", err
	}

	// Token is valid, return the UID
	return token.UID, nil
}
