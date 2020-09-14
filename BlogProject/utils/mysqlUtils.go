package utils

func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content varchar(255),
		createtime int(10)
	);`
	ModifyDB(sql)
}