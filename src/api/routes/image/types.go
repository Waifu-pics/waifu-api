package image

// ResImage : send a single image url
type ResImage struct {
	URL string `json:"url"`
}

// ResManyImages : send many image urls
type ResManyImages struct {
	Files []string `json:"files"`
}

// ReqManyImages : body for many images
type ReqManyImages struct {
	Exclude []string `json:"exclude"`
}
