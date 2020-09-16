package database

// FileData : file information from database
type FileData struct {
	Uploaded string
	Name     string
	Type     string
	Nsfw     bool
	Verified bool
}
