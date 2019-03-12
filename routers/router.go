package routers

import (
    "king/controllers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

func init() {
    beego.InsertFilter("/aritcle/*",beego.BeforeExec,Filter)
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.IndexController{},"get:ShowGet;post:Post") //指定路由
	//beego.Router("/index/?:id", &controllers.IndexController{},"get:ShowGet;post:Post")

    beego.Router("/register",&controllers.Usercontroller{},"get:ShowRegister;post:HandlePost")
    beego.Router("/login",&controllers.Usercontroller{},"get:ShowLogin;post:HandleLogin")

    //显示文章列表页
    beego.Router("/article/ShowArticleList",&controllers.ArticleController{},"get:ShowArticleList")

    beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")

    //显示文章详情
    beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ArticleDetail")

    //编辑文章
    beego.Router("/article/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")

    //删除文章
    beego.Router("/article/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")

    //添加分类
    beego.Router("/article/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")

    //退出登录
    beego.Router("/article/logout",&controllers.Usercontroller{},"get:LogOut")

}

var Filter = func(ctx *context.Context) {
    userName:=ctx.Input.Session("userName")
    if userName==nil{
        ctx.Redirect(302,"/login")
        return
    }

}
