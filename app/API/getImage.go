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

// SingleImagePoint : Get a single image from the DB
func SingleImagePoint(mux *mux.Router, endpoint string, conf util.Config) {
	mux.HandleFunc("/api/img/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
		matchStage := bson.D{{
			Key: "$match", Value: bson.D{
				{Key: "type", Value: endpoint},
				{Key: "verified", Value: true},
			},
		}}

		sampleStage := bson.D{{
			Key: "$sample", Value: bson.D{
				{Key: "size", Value: 1},
			},
		}}

		mongoRes, err := util.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sampleStage})

		var dumpRes []struct {
			URL string `bson:"file,omitempty"`
		}

		if err = mongoRes.All(context.TODO(), &dumpRes); err != nil {
			panic(err)
		}

		type sendRes struct {
			URL string `json:"url"`
		}

		response, _ := json.Marshal(sendRes{URL: conf.URL + dumpRes[0].URL})

		util.WriteResp(w, 200, string(response))

		defer r.Body.Close()

	}).Methods("GET")
}

// ManyImagePoint : Get many images from the DB, created for use with frontend
func ManyImagePoint(mux *mux.Router, endpoint string, conf util.Config) {
	mux.HandleFunc("/api/many/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
		var excludeDat struct {
			Exclude []string `json:"exclude"`
		}

		json.NewDecoder(r.Body).Decode(&excludeDat)

		// Turn the exclude slice into bson.A interface
		result := bson.A{}
		for _, image := range excludeDat.Exclude {
			result = append(result, image)
		}

		matchStage := bson.D{{
			Key: "$match", Value: bson.D{
				{Key: "type", Value: endpoint},
				{Key: "verified", Value: true},
				{Key: "file", Value: bson.D{
					{Key: "$nin", Value: result},
				}},
			},
		}}

		sampleStage := bson.D{{
			Key: "$sample", Value: bson.D{
				{Key: "size", Value: 30},
			},
		}}

		mongoRes, err := util.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sampleStage})

		// Query non json struct
		var dumpRes []struct {
			URLs string `bson:"file,omitempty"`
		}

		// Dump the query to dumpRes
		if err = mongoRes.All(context.TODO(), &dumpRes); err != nil {
			return
		}

		// Response json struct
		type sendRes struct {
			Data []string `json:"data"`
		}

		// Add all URLs to new var
		var urls = make([]string, len(dumpRes))
		for i, d := range dumpRes {
			urls[i] = d.URLs
		}

		response, _ := json.Marshal(sendRes{Data: urls})

		util.WriteResp(w, 200, string(response))

		defer r.Body.Close()

	}).Methods("POST")
}
