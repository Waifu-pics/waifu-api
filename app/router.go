package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"waifu.pics/util"
)

// PageData : Sorting data type
type PageData struct {
	Links []string `json:"exclude"`
}

type DB struct {
	db *util.Mongo
}

func (glob *DB) getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing")
	// glob.db.Database.Collection("test").InsertOne(glob.db.Context, bson.D{
	// 	{Key: "hi", Value: "HELO"},
	// })
}

// Router : Test
func Router(mux *mux.Router, db *util.Mongo, config util.Config) *mux.Router {
	glob := DB{db: db}
	mux.HandleFunc("/api/test", glob.getTest).Methods("GET")

	return mux
}
