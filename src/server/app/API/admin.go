package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Riku32/waifu.pics/util/file"
	"github.com/Riku32/waifu.pics/util/web"
)

// ListFile : listing files
func (api API) ListFile(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Endpoint string `json:"endpoint"`
		Query    string `json:"query"`
		Verified bool   `json:"verified"`
		Nsfw     bool   `json:"nsfw"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	if !CheckValid(ImageRoute{
		Type: res.Endpoint,
		Nsfw: res.Nsfw,
	}, api.Config) {
		web.SendJSON(w, 400, web.NewMessage("invalid type"))
		return
	}

	files, err := api.Database.GetFilesAdmin(res.Endpoint, res.Query, res.Verified, res.Nsfw)
	if err != nil {
		web.SendJSON(w, 500, web.NewMessage("error getting files"))
		return
	}

	type File struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	var response struct {
		Files []File `json:"files"`
	}

	for _, v := range files {
		response.Files = append(response.Files, File{
			Name: v,
			URL:  api.Config.Web.CdnURL + v,
		})
	}

	jsonres, err := json.Marshal(response)
	if err != nil {
		web.SendJSON(w, 500, web.NewMessage(errorencode))
		return
	}

	web.SendJSON(w, 200, jsonres)

	defer r.Body.Close()
}

// VerifyFile : Verifying user uploads
func (api API) VerifyFile(w http.ResponseWriter, r *http.Request) {
	var res struct {
		File []string `json:"file"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	defer r.Body.Close()

	var errcount int

	for _, v := range res.File {
		err = api.Database.VerifyFile(v)
		if err != nil {
			errcount++
		}
	}

	web.SendJSON(w, 200, web.NewMessage(fmt.Sprintf("files have been verified with %d errors", errcount)))
	return
}

// DeleteFile : delete a file from the API
func (api API) DeleteFile(w http.ResponseWriter, r *http.Request) {
	var res struct {
		File []string `json:"file"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	defer r.Body.Close()

	var errcount int

	for _, v := range res.File {
		err = api.Database.DeleteFile(v)
		if err != nil {
			errcount++
		}
		if err := file.DeleteFile(v, api.Config); err != nil {
			errcount++
		}
	}

	web.SendJSON(w, 200, web.NewMessage(fmt.Sprintf("Files have been deleted with %d errors!", errcount)))
	return
}
