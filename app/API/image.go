package api

import (
	"context"
	"encoding/json"
	"net/http"
	"waifu.pics/util/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Multi struct {
	API      *API
	Endpoint string
}

// GetImage : Get a single image from the DB
func (multi Multi) GetImage(w http.ResponseWriter, r *http.Request) {
	matchStage := bson.D{{
		Key: "$match", Value: bson.D{
			{Key: "type", Value: multi.Endpoint},
			{Key: "verified", Value: true},
		},
	}}

	sampleStage := bson.D{{
		Key: "$sample", Value: bson.D{
			{Key: "size", Value: 1},
		},
	}}

	mongoRes, err := multi.API.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sampleStage})
	if err != nil {
		web.WriteResp(w, 400, "Server Error!")
		return
	}

	var dumpRes []struct {
		URL string `bson:"file,omitempty"`
	}

	err = mongoRes.All(context.TODO(), &dumpRes)
	if err != nil {
		web.WriteResp(w, 400, "Server Error!")
		return
	}

	type sendRes struct {
		URL string `json:"url"`
	}

	response, _ := json.Marshal(sendRes{URL: multi.API.Config.URL + dumpRes[0].URL})

	web.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}

func (multi Multi) GetManyImage(w http.ResponseWriter, r *http.Request) {
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
			{Key: "type", Value: multi.Endpoint},
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

	mongoRes, err := multi.API.Database.Collection("uploads").Aggregate(context.TODO(), mongo.Pipeline{matchStage, sampleStage})
	if err != nil {
		web.WriteResp(w, 500, "Server Error!")
		return
	}

	// Query non json struct
	var dumpRes []struct {
		URLs string `bson:"file,omitempty"`
	}

	// Dump the query to dumpRes
	err = mongoRes.All(context.TODO(), &dumpRes)
	if err != nil {
		web.WriteResp(w, 500, "Server Error!")
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

	web.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}
