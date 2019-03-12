package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"king/models"
)

type Usercontroller struct {
	beego.Controller
}


//显示注册页面
func (this *Usercontroller)ShowRegister()  {
     //指定视图
     this.TplName="register.html"
}

//处理注册数据
func (this *Usercontroller) HandlePost(){
	//获取数据
	userName:=this.GetString("userName")
	pwd:=this.GetString("password")
	//beego.Info(userName,pwd)
	//校验数据
    if userName=="" || pwd ==""{
    	this.Data["errmsg"]="注册数据不完整，请重新输入"
    	//beego.Info("注册数据不完整，请重新输入")
    	this.TplName="register.html"
		return

	}
	// 处理数据
	 o:=orm.NewOrm()
	 var user models.User
	 user.Username=userName
	 user.Password=pwd

	 o.Insert(&user)
	// 返回页面
	// this.Ctx.WriteString("注册成功")
	this.Redirect("/login",302)
}

//显示登陆页面
func (this *Usercontroller)ShowLogin()  {
	userName:=this.Ctx.GetCookie("userName")
	if userName==""{
		this.Data["userName"]=""
		this.Data["checked"]=""
	}else{
		this.Data["userName"]=userName
		this.Data["checked"]="checked"
	}
	this.TplName="login.html"
}


/*func (this *Usercontroller)HandleLogin() {
    //获取数据
    userName:=this.GetString("userName")
    pwd:=this.GetString("password")
    //校验数据
    if userName=="" ||  pwd==""{
    	this.Data["errmsg"]="登陆数据不完整"
    	this.TplName="login.html"
		return
	}

    //操作数据
    o:=orm.NewOrm()
    var user models.User
    user.Username=userName
    err:=o.Read(&user,"Name")
    if err!= nil{
		this.Data["errmsg"]="用户名不存在"
		this.TplName="login.html"
		return
	}
    if user.Password!=pwd{
		this.Data["errmsg"]="密码错误"
		this.TplName="login.html"
		return
	}
    //返回页面
    this.Ctx.WriteString("登陆成功")

}*/

func (this *Usercontroller)HandleLogin(){
	userName:=this.GetString("userName")
	pwd:=this.GetString("password")
	//beego.Info("0000000000000000000000000000000000")

	if userName==""|| pwd==""{
		this.Data["errmsg"]="登陆数据不完整"
		this.TplName="login.html"
		return
	}
	//beego.Info("hhhhhhhhhhhhhhhhhhhhh")
	//查询数据库，判断用户名和密码是否正确
	o:=orm.NewOrm()
	var user models.User
	//给对象赋值
	user.Username=userName
	err:=o.Read(&user,"Username")
	if err!=nil{
		this.Data["errmsg"]="用户名输入错误，请重新输入"
		this.TplName="login.html"
		return
	}
	if user.Password!=pwd{
		this.Data["errmsg"]="密码错误，请重新输入"
		this.TplName="login.html"
		return
	}
	//this.Ctx.WriteString("登陆成功")

	//记住用户名
	data:=this.GetString("remember")
	if data=="on"{
		this.Ctx.SetCookie("userName",userName,1000)
	}else {
		this.Ctx.SetCookie("userName",userName,-1)
	}

	this.SetSession("userName",userName)
	this.Redirect("/article/ShowArticleList",302)
}

//退出登录
func (this *Usercontroller)LogOut()  {
	//删除session
	this.DelSession("userName")

	//跳转登陆页面
	this.Redirect("/login",302)

}