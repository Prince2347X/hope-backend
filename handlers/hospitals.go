package handlers

import (
	"fmt"
	"net/http"

	"github.com/Prince2347X/hope-backend/helpers"
	"github.com/Prince2347X/hope-backend/models"
	"github.com/Prince2347X/hope-backend/services"
	"github.com/gin-gonic/gin"
)

func HandleHospitals(ctx *gin.Context) {

	// Extracting the ID token from the Authorization header
	idToken := helpers.ExtractBearerToken(ctx.GetHeader("Authorization"))

	firebaseServices := services.FirebaseServices{}

	// Verifying the ID token
	uid, err := firebaseServices.VerifyAuthToken(idToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized request", Message: err.Error()})
		return
	}

	// Token is valid
	fmt.Printf("User is authenticated. UID: %s", uid)

	// MARK: TODO: - Fetching hospitals using the Maps API

	mapsServices := services.GoogleMapsServices{}

	hospitals, err := mapsServices.FetchHospitals(0.0, 0.0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal server error", Message: err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, hospitals)
}
