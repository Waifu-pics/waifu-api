package fun

import "github.com/Riku32/waifu.pics/src/api"

// ReqGenerate : generate request body
type ReqGenerate struct {
	Endpoint api.ImageEndpoint `json:"endpoint"`
	Text     struct {
		Bottom string `json:"bottom"`
		Top    string `json:"top"`
	} `json:"text"`
}
