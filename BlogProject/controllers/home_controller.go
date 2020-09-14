package controllers

import (
	"BlogProject/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	fmt.Println("page", page)

	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	fmt.Println("artList", artList)
	c.Data["pageCode"] = 1
	c.Data["HasFooter"] = true
	fmt.Println("IsLogin", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)
	c.TplName = "home.html"
}