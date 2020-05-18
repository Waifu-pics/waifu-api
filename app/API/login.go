package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"waifu.pics/util"
)

// AdminLogin : admin login endpoint
func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&responseData)

	count, _ := util.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"username": responseData.Username})
	if count == 0 {
		util.WriteResp(w, 400, "Incorrect credentials!")
		return
	}

	var queryUser struct {
		Password string `bson:"password,omitempty"`
		Token    string `bson:"token,omitempty"`
	}

	if err := util.Database.Collection("admins").FindOne(context.TODO(), bson.M{"username": responseData.Username}).Decode(&queryUser); err != nil {
		util.WriteResp(w, 400, "Error!")
		return
	}

	authValid, err := util.ComparePassword(responseData.Password, queryUser.Password)
	if err != nil {
		util.WriteResp(w, 400, "Error!")
		fmt.Println(err)
		return
	}

	if authValid {
		util.WriteResp(w, 200, queryUser.Token)
		return
	}

	util.WriteResp(w, 400, "Incorrect credentials!")

	defer r.Body.Close()
}

// AdminVerify : check if the token is valid
func AdminVerify(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		Token string `json:"token"`
	}

	json.NewDecoder(r.Body).Decode(&responseData)

	count, _ := util.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
	if count == 0 {
		util.WriteResp(w, 400, "Invalid token!")
		return
	}

	util.WriteResp(w, 200, "Token is valid!")
	defer r.Body.Close()
}
