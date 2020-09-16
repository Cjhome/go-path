package controllers

import (
	"BlogProject/models"
	"fmt"
)

type AlbumController struct {
	BaseController
}

func (c *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		fmt.Println(nil)
	}
	c.Data["Album"] = albums
	c.TplName = "album.html"
}
