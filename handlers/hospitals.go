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
	mapsServices := services.GoogleMapsServices{}

	// Verifying the ID token
	uid, err := firebaseServices.VerifyAuthToken(idToken);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized request", Message: err.Error()})
		return
	}

	// Token is valid
	fmt.Printf("User is authenticated. UID: %s", uid)

	// MARK: TODO: - Fetching hospitals using the Maps API
	hospitals := mapsServices.FetchHospitals(0.0, 0.0)
	ctx.IndentedJSON(http.StatusOK, hospitals)
}
