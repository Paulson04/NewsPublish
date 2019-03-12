package main

import (
	"github.com/astaxie/beego"
	_ "king/routers"
	_"king/models"
)

func main() {
	beego.AddFuncMap("prepage",ShowPrePage)
	beego.AddFuncMap("nextpage",ShowNextPage)
	beego.Run()
}

//试图函数

//后台定义
//在视图里定义函数名
//在beego.run之前，关联视图与后台 参数一：视图函数名称。参数二；后台函数名
func ShowPrePage(pageIndex int)int {
	if pageIndex==1{
		return pageIndex
	}
	return pageIndex-1
}

func ShowNextPage(pageIndex int ,pageCount int)int  {
	if pageIndex==pageCount{
		return pageIndex
	}
	return pageIndex+1

}