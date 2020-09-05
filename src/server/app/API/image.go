package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Riku32/waifu.pics/src/server/util/web"
)

type Multi struct {
	API      *API
	Nsfw     bool
	Endpoint string
}

// GetImage : Get a single image from the DB
func (multi Multi) GetImage(w http.ResponseWriter, r *http.Request) {
	files, err := multi.API.Database.GetFiles(multi.Endpoint, multi.Nsfw, nil, 1)
	if err != nil || len(files) == 0 {
		web.SendJSON(w, 500, web.NewMessage(servererror))
		return
	}

	type sendRes struct {
		URL string `json:"url"`
	}

	response, _ := json.Marshal(sendRes{URL: multi.API.Config.Web.CdnURL + files[0]})

	web.SendJSON(w, 200, response)

	defer r.Body.Close()
}

// GetManyImage : Get many images from the database
func (multi Multi) GetManyImage(w http.ResponseWriter, r *http.Request) {
	var excludeDat struct {
		Exclude []string `json:"exclude"`
	}

	err := json.NewDecoder(r.Body).Decode(&excludeDat)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	var list []string
	for _, v := range excludeDat.Exclude {
		list = append(list, strings.TrimPrefix(v, multi.API.Config.Web.CdnURL))
	}

	files, err := multi.API.Database.GetFiles(multi.Endpoint, multi.Nsfw, excludeDat.Exclude, 30)
	if err != nil || len(files) == 0 {
		web.SendJSON(w, 400, web.NewMessage(servererror))
		return
	}

	var urls []string
	for _, v := range files {
		urls = append(urls, multi.API.Config.Web.CdnURL+v)
	}

	// Response json struct
	type sendRes struct {
		Files []string `json:"files"`
	}

	response, _ := json.Marshal(sendRes{Files: urls})

	web.SendJSON(w, 200, response)

	defer r.Body.Close()
}
