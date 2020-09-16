package models

import (
	"BlogProject/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
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
	SetArticleRowsNum()
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

func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id, title, tags, short, content, author, createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

func UpdateArticle(article Article) (int64, error) {
	// 数据库操作
	return utils.ModifyDB("update article set title=?,tags=?,content=? where id=?", article.Title, article.Tags, article.Content, article.Id)
}

//--------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

// 添加sql语句
func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}

// 查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

// -------按照标签查询---------
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%&" + tag + "&%'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticleWithCon(sql)
}