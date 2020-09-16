package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"

	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	// 注册表
	fmt.Println("InitMysql....")
	// driverName unknown driver "" (forgotten import?) 獲取配置数据库名称错误
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		db = db1
		fmt.Println("db", db)
		// 创建users表
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	fmt.Println("sql", sql)
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// 时间戳格式化
func SwitchTimeStampToData(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}