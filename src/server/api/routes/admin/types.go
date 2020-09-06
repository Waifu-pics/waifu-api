package admin

// Credentials : user credentials for sign in
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Search : admin search images
type Search struct {
	Endpoint string `json:"endpoint"`
	Query    string `json:"query"`
	Verified bool   `json:"verified"`
	Nsfw     bool   `json:"nsfw"`
}

// File : file data for listing
type File struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Filelist : list of files
type Filelist struct {
	Files []File `json:"files"`
}
