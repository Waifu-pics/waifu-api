package util

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckAuth : Function to check if the token is valid
func CheckAuth(w http.ResponseWriter, r *http.Request) error {
	var responseData struct {
		Token string `json:"token"`
	}

	json.NewDecoder(r.Body).Decode(&responseData)

	count, _ := Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
	if count == 0 {
		return errors.New("auth: invalid token")
	}

	return nil
}
