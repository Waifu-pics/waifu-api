package api

import (
	"context"
	"encoding/json"
	"net/http"
	"waifu.pics/util/file"
	"waifu.pics/util/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ListFile : listing all unverified files
func (api API) ListFile(w http.ResponseWriter, r *http.Request) {
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
		web.WriteResp(w, 400, "Invalid credentials!")
		return
	}

	matchStage := bson.D{{
		Key: "$match", Value: bson.D{
			{Key: "verified", Value: false},
		},
	}}

	sortStage := bson.D{{
		Key: "$sort", Value: bson.D{
			{Key: "_id", Value: -1},
		},
	}}

	mongoRes, _ := api.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage})

	var dumpRes []struct {
		File string `bson:"file,omitempty"`
		Type string `bson:"type,omitempty"`
	}

	if err := mongoRes.All(context.TODO(), &dumpRes); err != nil {
		return
	}

	response, _ := json.Marshal(dumpRes)

	web.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}

// VerifyFile : Verifying user uploads
func (api API) VerifyFile(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		IsVerified bool   `json:"isVer"`
		File       string `json:"file"`
		Token      string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&responseData)
	if err != nil {
		web.WriteResp(w, 400, "Invalid JSON!")
		return
	}

	count, _ := api.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
	if count == 0 {
		web.WriteResp(w, 400, "Invalid credentials!")
		return
	}

	count, _ = api.Database.Collection("uploads").CountDocuments(context.TODO(), bson.M{"file": responseData.File})
	if count == 0 {
		web.WriteResp(w, 400, "Invalid file!")
		return
	}

	filter := bson.M{"file": responseData.File}

	if responseData.IsVerified == true {
		update := bson.M{"$set": bson.M{"verified": true}}
		api.Database.Collection("uploads").UpdateOne(context.TODO(), filter, update)

		web.WriteResp(w, 200, "File has been verified!")
	} else {
		if err := file.DeleteFile(responseData.File, api.Config); err != nil {
			web.WriteResp(w, 400, "File could not be deleted!")
			return
		}
		api.Database.Collection("uploads").DeleteOne(context.TODO(), filter)

		web.WriteResp(w, 200, "File has been Deleted")
	}

	defer r.Body.Close()
}
