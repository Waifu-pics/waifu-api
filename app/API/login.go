package api

import (
	"context"
	"encoding/json"
	"net/http"
	"waifu.pics/util/crypto"
	"waifu.pics/util/web"

	"go.mongodb.org/mongo-driver/bson"
)

// AdminLogin : admin login endpoint
func (api API) AdminLogin(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&responseData)
	if err != nil {
		web.WriteResp(w, 400, "Invalid JSON!")
		return
	}

	count, _ := api.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"username": responseData.Username})
	if count == 0 {
		web.WriteResp(w, 400, "Incorrect credentials!")
		return
	}

	var queryUser struct {
		Password string `bson:"password,omitempty"`
		Token    string `bson:"token,omitempty"`
	}

	if err := api.Database.Collection("admins").FindOne(context.TODO(), bson.M{"username": responseData.Username}).Decode(&queryUser); err != nil {
		web.WriteResp(w, 400, "Error!")
		return
	}

	authValid, err := crypto.ComparePassword(responseData.Password, queryUser.Password)
	if err != nil {
		web.WriteResp(w, 400, "Error!")
		return
	}

	if authValid {
		web.WriteResp(w, 200, queryUser.Token)
		return
	}

	web.WriteResp(w, 400, "Incorrect credentials!")

	defer r.Body.Close()
}

// AdminVerify : check if the token is valid
func (api API) AdminVerify(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&responseData)
	if err != nil {
		web.WriteResp(w, 400, "Invalid JSON!")
		return
	}

	count, _ := api.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
	if count == 0 {
		web.WriteResp(w, 400, "Invalid token!")
		return
	}

	web.WriteResp(w, 200, "Token is valid!")
	defer r.Body.Close()
}
