package api

import (
	"encoding/json"
	"net/http"

	"github.com/Riku32/waifu.pics/src/server/util/web"
)

func (api API) GetEndpoints(w http.ResponseWriter, r *http.Request) {
	endpoints, _ := json.Marshal(api.Config.Endpoints)

	web.SendJSON(w, 200, endpoints)
}
