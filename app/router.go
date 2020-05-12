package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"waifu.pics/util"
)

// PageData : Sorting data type
type PageData struct {
	Links []string `json:"exclude"`
}

// GetTest : Test route
func GetTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing")
}

// Router : Test
func Router(mux *mux.Router, db *mongo.Client, config util.Config) *mux.Router {
	mux.HandleFunc("/api/test", GetTest).Methods("GET")

	return mux
}
