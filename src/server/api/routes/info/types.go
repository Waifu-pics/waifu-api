package info

// FileData : file information
type FileData struct {
	Uploaded string `json:"uploaded"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Type     string `json:"type"`
	Nsfw     bool   `json:"nsfw"`
	Verified bool   `json:"verified"`
}
