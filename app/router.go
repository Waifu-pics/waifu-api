package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"waifu.pics/util"
)

// PageData : Sorting data type
type PageData struct {
	Links []string `json:"exclude"`
}

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing")

	util.Database.Collection("test").InsertOne(context.TODO(), bson.D{
		{Key: "test", Value: "test"},
	})

	// route.mongo.Collection("test").InsertOne(ctx, bson.D{
	// 	{Key: "test", Value: "test"},
	// })
}

// Router : Test
func Router(mux *mux.Router, config util.Config) *mux.Router {
	mux.HandleFunc("/api/test", getTest).Methods("GET")

	return mux
}
