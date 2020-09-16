package controllers

import (
	"BlogProject/models"
	"fmt"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

// 点击删除后重定向到首页
func (c *DeleteArticleController) Get() {
	artID, _ := c.GetInt("id")
	fmt.Println("删除 id:", artID)
	_, err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/", 302)
}