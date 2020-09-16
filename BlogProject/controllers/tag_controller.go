package controllers

import (
	"BlogProject/models"
	"fmt"
)

type TagsController struct {
	BaseController
}

func (c *TagsController) Get() {
	// 查询数据库获取所有文章的标签，组成列表返回
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(tags)
	fmt.Println(models.HandleTagsListData(tags))
	// 赋值Tags 统计每个标签的文章数，组成列表返回
	c.Data["Tags"] = models.HandleTagsListData(tags)
	c.TplName = "tags.html"
}
