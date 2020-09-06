package api

// Basic : basic response
type Basic struct {
	Message string `json:"message"`
}

// ImageEndpoint : Image endpoint data
type ImageEndpoint struct {
	Nsfw bool   `json:"nsfw"`
	Type string `json:"type"`
}
