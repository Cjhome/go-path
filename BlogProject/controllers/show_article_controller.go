package controllers

import (
	"BlogProject/models"
	"BlogProject/utils"
	"fmt"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (c *ShowArticleController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)
	// 获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	c.Data["Title"] = art.Title
	c.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	c.TplName = "show_article.html"
}
