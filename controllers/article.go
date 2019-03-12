package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"king/models"
	"math"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}


//展示文章列表页
func (this *ArticleController)ShowArticleList() {
	//session判断
	userName:=this.GetSession("userName")
	if userName==nil{
		this.Redirect("/login",302)
		return
	}

	//获取数据
	//高级查询
	//指定表

	o:=orm.NewOrm()
	qs:=o.QueryTable("Article")//qs---queryseter

	var articles []models.Article
	/*_,err:=qs.All(&articles)
	if err!=nil{
		beego.Info("查询数据错误")
	}*/

	//查询总记录数
	typeName:=this.GetString("select")
	var count int64



	//获取总页数
	pageSize:=2


	//获取页码
	pageIndex,err:=this.GetInt("pageIndex")
	//beego.Info(pageIndex)
	if err!=nil{
		pageIndex=1
	}

	//获取数据
	//起始位置
	start:=(pageIndex-1)*pageSize

	if typeName==""{
		count,_=qs.Count()
	}else {
		count,_=qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()
	}
	pageCount:=math.Ceil(float64(count)/float64(pageSize))

	//limit（）作用是获取数据库部分数据，参数一：获取几条，参数二：从哪条数据开始获取，返回值还是queryseter
	//qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)

	//获取文章类型
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	this.Data["types"]=types

	//根据选中的类型，获取相应类型的文章

	beego.Info(typeName)
	if typeName==""{
		qs.Limit(pageSize,start).All(&articles)
	}else{
		qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
	}


	//数据传递,将数据传递给前端页面
	this.Data["typeName"]=typeName
	this.Data["pageIndex"]=pageIndex
	this.Data["pageCount"]=int(pageCount)
	this.Data["count"]=count
	this.Data["articles"]=articles
	this.TplName="index.html"

}

//展示添加文章页面
func (this *ArticleController)ShowAddArticle() {

	//查询所有文章类型，并展示
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	//传递数据
	this.Data["types"]=types
	this.TplName="add.html"

}

//获取添加文章数据
func (this *ArticleController)HandleAddArticle()  {
	//获取数据
	articleName:=this.GetString("articleName")
	content:=this.GetString("content")

	//校验数据
	if articleName==""||content==""{
		this.Data["errmsg"]="添加数据不完整，请重新输入"
		this.TplName="add.html"
		return
	}

	//beego.Info(articleName,content)
	//处理文件上传
	file,head,err:=this.GetFile("uploadname")
	defer file.Close()

	if err!=nil{
		this.Data["errmsg"]="文件上传失败，请重新上传"
		this.TplName="add.html"
		return
	}
    //需判断
      //文件大小
      if head.Size>5000000 {
		  this.Data["errmsg"]="文件太大，请重新上传"
		  this.TplName="add.html"
		  return
	  }
      //文件格式
      ext:=path.Ext(head.Filename)
      if ext!=".jpg"&& ext!=".png"&& ext!=".jpeg"{
		  this.Data["errmsg"]="文件格式错误，请重新上传"
		  this.TplName="add.html"
		  return
	  }
      //防止重名
      fileName:=time.Now().Format("2006-01-02  15:04:05")+ext

      //存储
	  this.SaveToFile("uploadname","./static/img"+fileName)

	//处理数据
	//插入操作
	  o:=orm.NewOrm()
	  var article models.Article
	  article.ArtiName=articleName
	  article.Acontent=content
	  article.Aimg="/static/img"+fileName

	  //给文章添加类型
	   //获取类型数据
	   typeName:=this.GetString("select")
	   //根据名称查询类型
	   var articleType models.ArticleType
	   articleType.TypeName=typeName
	   o.Read(&articleType,"TypeName")
	   article.ArticleType=&articleType

	  o.Insert(&article)

	//返回页面
	this.Redirect("/article/ShowArticleList",302)
}

//显示文章详情
func (this *ArticleController)ArticleDetail(){
	//获取数据
	id,err:=this.GetInt("articleId")
	//数据校验
	if err!=nil{
		beego.Info("获取数据失败")
		return
	}
	//数据处理
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	//o.Read(&article)
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("Id",id).One(&article)


	article.Acount += 1

	o.Update(&article)

	//多对多插入浏览记录
	//获取orm对象
	//获取操作数据
	//获取多对多操作对象
	//获取插入对象
	//插入
	m2m:=o.QueryM2M(&article,"Users")
	userName:=this.GetSession("userName")
	if userName==nil{
		this.Redirect("/login",302)
		return
	}
	var user models.User
	user.Username=userName.(string)
	o.Read(&user,"userName")

	m2m.Add(user)

	//多对多查询
	//o.LoadRelated(&article,"Users")
	var users []models.User
	o.QueryTable("User").Filter("Articles__Article__Id",id).Distinct().All(&users)

	this.Data["users"]=users
	//返回视图页面
	this.Data["article"]=article
	this.TplName="content.html"

}

//编辑文章
func (this *ArticleController)ShowUpdateArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	//校验数据
	if err!=nil{
		beego.Info("获取数据失败")
		return
	}
	//处理数据
	   //查询相应文章，做更新
	   o:=orm.NewOrm()
	   var article models.Article
	   article.Id=id
	   o.Read(&article)
	//返回视图
	this.Data["article"]=article
	this.TplName="update.html"

}


//封装上传文件函数
func UpdateFile(this *beego.Controller,filePath string) string {
	//处理文件上传
	file,head,err:=this.GetFile(filePath)
	if head.Filename==""{
		return "NoAimg"
	}

	if err!=nil{
		this.Data["errmsg"]="文件上传失败，请重新上传"
		this.TplName="add.html"
		return ""
	}
	defer file.Close()
	//需判断
	//文件大小
	if head.Size>5000000 {
		this.Data["errmsg"]="文件太大，请重新上传"
		this.TplName="add.html"
		return ""
	}
	//文件格式
	ext:=path.Ext(head.Filename)
	if ext!=".jpg"&& ext!=".png"&& ext!=".jpeg"{
		this.Data["errmsg"]="文件格式错误，请重新上传"
		this.TplName="add.html"
		return ""
	}
	//防止重名
	fileName:=time.Now().Format("2006-01-02  15:04:05")+ext

	//存储
	this.SaveToFile(filePath,"./static/img"+fileName)
	return "/static/img"+fileName

}


//处理编辑界面数据
func (this *ArticleController)HandleUpdateArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	articleName:=this.GetString("articleName")
	content:=this.GetString("content")
	filePath:=UpdateFile(&this.Controller,"uploadname")
	//数据校验
	if err!=nil || articleName==""||content==""  || filePath==""{
		beego.Info("请求错误")
		return
	}
	//数据处理
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	err=o.Read(&article)
	if err!=nil{
		beego.Info("更新的文章不存在")
		return
	}
	//beego.Info("66666666666666666")
	article.ArtiName=articleName
	article.Acontent=content
	if filePath!="NoAimg"{
		article.Aimg=filePath
	}

	o.Update(&article)
	//返回视图
	this.Redirect("/article/ShowArticleList",302)

}

//删除文章
func (this *ArticleController)DeleteArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	//校验数据
	if err!=nil{
		beego.Info("删除文章请求路径失败")
		return
	}
	//处理数据

	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.Delete(&article)
	//返回视图
	this.Redirect("/article/ShowArticleList",302)

}

//添加分类
func (this *ArticleController)ShowAddType() {
	//查询
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	this.Data["types"]=types
	this.TplName="addType.html"
}

//处理添加类型数据
func (this *ArticleController)HandleAddType()  {
	//获取数据
	typeName:=this.GetString("typeName")
	//校验数据
	if typeName==""{
		beego.Info("信息不完整，请重新输入")
		return
	}
	//处理数据
	o:=orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName=typeName
	o.Insert(&articleType)
	//返回视图
	this.Redirect("/article/addType",302)


}
















