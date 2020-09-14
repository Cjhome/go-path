package models

import (
	"BlogProject/utils"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id int
	Title string
	Tags []TagLink
	Short string
	Content string
	Author string
	CreateTime string
	// 查看文章的地址
	Link string

	// 修改文章的地址
	UpdateLink string
	DeleteLink string
	// 记录是否登录
	IsLogin bool
}

// 标签链接
type TagLink struct {
	TagName string
	TagUrl string
}

func MakeHomeBlocks(article []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range article {
		// 将数据库model转换为首页模版所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTags(art.Tags)
		fmt.Println("tag--->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink =  "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin
		fmt.Println(homeParam)
		// 处理变量
		// parseFile 解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		// 就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("html -->", htmlHome)
	return template.HTML(htmlHome)
}

// 将tags字符串转化成首页模版所需要的数据结构
func createTags(tags string) []TagLink {
	var tagLink []TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}