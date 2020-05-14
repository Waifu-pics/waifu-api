package app

import (
	"github.com/gorilla/mux"
	"waifu.pics/app/API"
	"waifu.pics/util"
)

// Router :
func Router(mux *mux.Router, config util.Config) *mux.Router {
	// endPoints := []string{"nsfw", "sfw"}
	endPoints := config.ENDPOINTS

	for _, endP := range endPoints {
		endpoint := endP // Evaluates instantly

		API.SingleImagePoint(mux, endpoint, config)
	}

	return mux
}
