package models

import (
	"BlogProject/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type Article struct {
	Id  int
	Title string
	Tags string
	Short string
	Content string
	Author string
	Createtime int64
}

// -------数据处理-------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

// ---------数据库操作--------
// 插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values(?,?,?,?,?,?)", article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

// -----查询文章------
// 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	// 从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("----------->page", page)
	return QueryArticleWithPage(page, num)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面咩有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) ([]Article, error) {
	// 组成需要查询的sql语句
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	// 获取数据库数据
	fmt.Println("sql", sql)
	rows, err := utils.QueryDB(sql)
	fmt.Println("rows",  rows)
	if err != nil {
		return nil, err
	}
	var artList []Article
	// 循环赋值
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}
