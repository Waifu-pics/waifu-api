package image

import (
	"strings"

	"github.com/Riku32/waifu.pics/src/api"
	"github.com/labstack/echo"
)

// GetImage : Get a single image from the DB
func (i route) GetImage(c echo.Context) error {
	files, err := i.Options.Database.GetFiles(i.Endpoint, i.Nsfw, nil, 1)
	if err != nil || len(files) == 0 {
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	return c.JSON(200, ResImage{URL: i.Options.Config.Web.Cdn + files[0]})
}

// GetManyImage : Get many images from the database
func (i route) GetManyImage(c echo.Context) error {
	body := new(ReqManyImages)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	var list []string
	for _, v := range body.Exclude {
		list = append(list, strings.TrimPrefix(v, i.Options.Config.Web.Cdn))
	}

	filenames, err := i.Options.Database.GetFiles(i.Endpoint, i.Nsfw, list, 30)
	if err != nil || len(filenames) == 0 {
		return c.JSON(400, api.Basic{Message: api.ErrServer})
	}

	var response ResManyImages
	for _, filename := range filenames {
		response.Files = append(response.Files, i.Options.Config.Web.Cdn+filename)
	}

	return c.JSON(200, response)
}
