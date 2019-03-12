package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"king/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["king"]="hello 区块链"
	c.TplName = "king.html"

	//插入操作
	//获取orm对象
	/*o:=orm.NewOrm()
	//执行某个操作函数，增删改查
	//插入操作
	var user models.User
	user.Username="nihao"
	user.Password="1234"
	n,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("插入失败",err)
	}
	beego.Info(n)*/

	//查询操作
	//获取orm查询对象
	/*o:=orm.NewOrm()

	//定义一个要获取数据的结构体对象
	var user models.User

	//给结构体对象赋值，相当于给查询条件赋值
	user.Username="nihao"

	//执行查询操作
	err:=o.Read(&user,"Username")
	if err!=nil{
		beego.Info("查询失败",err)
	}
	beego.Info(user)*/

	//更新操作
	//获取orm对象
	/*o:=orm.NewOrm()
	//定义结构体对象
	var user models.User
	user.Password="1234"
	err:=o.Read(&user,"Password")
	if err!=nil{
		beego.Info("要更新的数据不存在",err)
		return
	}
	user.Password="123"
	n,err:=o.Update(&user)
	if err!=nil{
		beego.Info("更新失败",err)
	}
	beego.Info(n)*/

	//删除操作
	o:=orm.NewOrm()
	var user models.User
	user.Username="nihao"
	err:=o.Read(&user,"Username")
	if err!=nil{
		beego.Info("查询错误",err)
	}

	n,err:=o.Delete(&user)
	if err!=nil{
		beego.Info("删除失败",err)
	}
	beego.Info(n)
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Post(){
	this.Data["king"]="从你的全世界路过"
	this.TplName="king.html"
}

func (this *IndexController) ShowGet(){
	id:=this.GetString(":id")
	beego.Info(id)
	this.Data["nihao"]="这是get请求所对应的方法"
	this.TplName="king.html"

}
