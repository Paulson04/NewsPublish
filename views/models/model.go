package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_"github.com/astaxie/beego/orm"
	"time"
)
type User struct {
	Id int
	Username string
	Password string
	Articles []*Article `orm:"reverse(many)"`
}
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string  `orm:"size(100)"`

	ArticleType * ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"rel(m2m)"`
}

//类型表
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init(){
	/*conn,err:=sql.Open("mysql","root:123456789@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err!=nil{
		beego.Info("连接错误",err)
		beego.Error("连接错误",err)
		return
	}
	defer conn.Close()

	res,err:=conn.Exec("create table paulson(username varchar (40),password varchar (40))")
	if err!=nil{
		beego.Error("创建表失败",err)
		beego.Info("创建表失败",err)
	}
	beego.Info(res)

	//conn.Exec("insert into paulson( username ,password) values (?,?)","nihao","haode")

	res,err:=conn.Query("select username from paulson")

	for res.Next(){
		var name string
		res.Scan(&name)

		beego.Info(name)
	}*/

	//获取连接对象 参数：别名，驱动，连接字符串
	orm.RegisterDataBase("default","mysql","root:123456789@tcp(127.0.0.1:3306)/test?utf8")

	//向数据库中注册表
	orm.RegisterModel(new(User))

	orm.RegisterModel(new(Article),new(ArticleType))
	//生成表
	orm.RunSyncdb("default",false,true)


}