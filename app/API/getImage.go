package API

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"waifu.pics/util"
)

// SingleImagePoint : Endpoint
func SingleImagePoint(mux *mux.Router, endpoint string, conf util.Config) {
	mux.HandleFunc("/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
		matchStage := bson.D{{"$match", bson.D{{"type", endpoint}, {"verified", true}}}}
		sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

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

		fmt.Fprintf(w, string(response))

		defer r.Body.Close()

	}).Methods("GET")
}
