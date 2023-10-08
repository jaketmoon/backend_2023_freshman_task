package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type TODO struct {
	Content    string `json:"content"`
	Done       bool   `json:"done"`
	LoginTime  uint64 `json:"login_time"`
	LogoutTime uint64 `json:"logout_time"`
}
type UserInformation struct {
	Name           string `json:"name"`
	PassWord       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	IdentityNumber string `json:"identity_number"`
}

func Register(c *gin.Context) {
	var userinformation UserInformation
	c.BindJSON(&userinformation)
	users = append(users, userinformation)
	c.JSON(200, gin.H{"状态": "ok", "已成功添加用户为": userinformation})
}

func Tianjia(c *gin.Context) {
	var todo TODO
	c.BindJSON(&todo) //添加TODO，接受前端传来的json数据
	todos = append(todos, todo)
	c.JSON(200, gin.H{"状态": "ok", "已成功添加值为": todo})
}
func Shanchu(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))
	deleted := todos[index]
	todos = append(todos[:index], todos[index+1:]...)
	c.JSON(200, gin.H{"状态": "ok", "已成功删除数据": deleted})
}

func Xiugai(c *gin.Context) { //地址+回调函数
	index, _ := strconv.Atoi(c.Param("index"))
	var todo TODO
	xiugaied := todos[index]
	c.BindJSON(&todo)
	todos[index] = todo
	c.JSON(200, gin.H{"状态": "ok", "已成功修改掉": xiugaied, "新的值为": todo})
}

func Huoqu(c *gin.Context) {
	c.JSON(200, gin.H{"状态": "ok", "当前获取内容为": todos})
}

func Chaxun(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))
	c.JSON(200, gin.H{"状态": "ok", "查询值为": todos[index]})
}

// 切片
var todos []TODO
var users []UserInformation

func main() {
	r := gin.Default() //创建一个默认路由
	//绑定路由规则和函数，访问index的路由，将有对应的函数去处理
	//用户注册
	r.POST("/users", Register)
	//新增TODO
	r.POST("/todo", Tianjia)

	//删除TODO
	r.DELETE("/todo/:index", Shanchu)

	//修改TODO
	r.PUT("/todo/:index", Xiugai)

	//获取TODO
	r.GET("/todo", Huoqu)

	//查询TODO
	r.GET("/todo/:index", Chaxun)

	r.Run(":8080") //运行
}