package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"waifu.pics/util/web"
)

type Multi struct {
	API      *API
	Endpoint string
}

// GetImage : Get a single image from the DB
func (multi Multi) GetImage(w http.ResponseWriter, r *http.Request) {
	files, err := multi.API.Database.GetFiles(multi.Endpoint, nil, 1)
	if err != nil || len(files) == 0 {
		web.WriteResp(w, 400, "Error")
		return
	}

	type sendRes struct {
		URL string `json:"url"`
	}

	response, _ := json.Marshal(sendRes{URL: multi.API.Config.URL + files[0]})

	web.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}

// GetManyImage : Get many images from the database
func (multi Multi) GetManyImage(w http.ResponseWriter, r *http.Request) {
	var excludeDat struct {
		Exclude []string `json:"exclude"`
	}

	err := json.NewDecoder(r.Body).Decode(&excludeDat)
	if err != nil {
		web.WriteResp(w, 400, "Invalid JSON!")
		return
	}

	files, err := multi.API.Database.GetFiles(multi.Endpoint, excludeDat.Exclude, 30)
	if err != nil || len(files) == 0 {
		fmt.Println(err)
		web.WriteResp(w, 400, "Error!")
		return
	}

	// Response json struct
	type sendRes struct {
		Files []string `json:"files"`
	}

	response, _ := json.Marshal(sendRes{Files: files})

	web.WriteResp(w, 200, string(response))

	defer r.Body.Close()
}
