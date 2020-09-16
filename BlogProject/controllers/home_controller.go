package controllers

import (
	"BlogProject/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	// 获取请求参数
	tag := c.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := c.GetInt("page")
	// 定义一个变量 设置类型为列表元素类型是models.Article
	var artList []models.Article
	if len(tag) > 0 {
		//安装指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
	}
	if page <= 0 {
		page = 1
	}

	artList, _ = models.FindArticleWithPage(page)
	c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	c.Data["HasFooter"] = true
	fmt.Println("IsLogin", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)
	c.TplName = "home.html"
}