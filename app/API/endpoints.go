package api

import (
	"encoding/json"
	"net/http"
	"waifu.pics/util/web"
)

func (api API) GetEndpoints(w http.ResponseWriter, r *http.Request) {
	endpoints, _ := json.Marshal(api.Config.ENDPOINTS)

	web.WriteResp(w, 200, string(endpoints))
}
