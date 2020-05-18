package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"waifu.pics/util"
)

// ListFile : listing all unverified files
func ListFile(w http.ResponseWriter, r *http.Request) {
	var responseData struct {
		Token string `json:"token"`
	}

	json.NewDecoder(r.Body).Decode(&responseData)

	count, _ := util.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
	if count == 0 {
		util.WriteResp(w, 400, "Invalid credentials!")
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

	mongoRes, _ := util.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage})

	var dumpRes []struct {
		File string `bson:"file,omitempty"`
		Type string `bson:"type,omitempty"`
	}

	if err := mongoRes.All(context.TODO(), &dumpRes); err != nil {
		return
	}

	response, _ := json.Marshal(dumpRes)

	util.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}

// VerifyFile : Verifying user uploads
func VerifyFile(mux *mux.Router, conf util.Config) {
	mux.HandleFunc("/api/admin/verify", func(w http.ResponseWriter, r *http.Request) {
		var responseData struct {
			IsVerified bool   `json:"isVer"`
			File       string `json:"file"`
			Token      string `json:"token"`
		}

		json.NewDecoder(r.Body).Decode(&responseData)

		count, _ := util.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": responseData.Token})
		if count == 0 {
			util.WriteResp(w, 400, "Invalid credentials!")
			return
		}

		count, _ = util.Database.Collection("uploads").CountDocuments(context.TODO(), bson.M{"file": responseData.File})
		if count == 0 {
			util.WriteResp(w, 400, "Invalid file!")
			return
		}

		filter := bson.M{"file": responseData.File}

		if responseData.IsVerified == true {
			update := bson.M{"$set": bson.M{"verified": true}}
			util.Database.Collection("uploads").UpdateOne(context.TODO(), filter, update)

			util.WriteResp(w, 200, "File has been verified!")
		} else {
			if err := util.DeleteFile(responseData.File, conf); err != nil {
				return
			}
			util.Database.Collection("uploads").DeleteOne(context.TODO(), filter)

			util.WriteResp(w, 200, "File has been Deleted")
		}

		defer r.Body.Close()
	})
}
