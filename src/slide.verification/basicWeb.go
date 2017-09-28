package slide_verification

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"net/http"
)

/**
需要安装：
go get github.com/ugorji/go/codec
go get github.com/gin-gonic/gin
*/
func aa() {

	r := gin.Default()
	//get json数据
	r.GET("/ping", get1)

	//get 传多个参数,pathVariable一样,restful的形式
	r.GET("/ping/param/:name", func(c *gin.Context) {
		log.Printf("pingParam|%s", c)
		param := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"data": param})
	})

	//get 传多个参数,返回json对象
	r.GET("/ping/param", get3)

	//post　最简单的post,发送form-data数据
	r.POST("/ping/post", post1)

	//post 发送json数据,返回json
	r.POST("/ping/post2", post2)

	getGroup := r.Group("/get")

	//统一路径的get
	getGroup.GET("/demo1", get4)

	r.Run(":8080")

}

func get1(c *gin.Context) {
	//c.String(200,"pong")
	log.Printf("ping|%s", c)
	c.JSON(http.StatusOK, gin.H{"data": 123, "value": "111", "response": "welcome to go web env!"})
}

func get3(c *gin.Context) {
	log.Printf("pingParam|%s", c)
	var aa webStruct
	aa.Name = c.DefaultQuery("name", "")
	aa.Age, _ = strconv.Atoi(c.DefaultQuery("age", ""))
	aa.Desc = c.DefaultQuery("desc", "")
	c.JSON(http.StatusOK, gin.H{"data": aa})
}

func get4(c *gin.Context) {
	name := c.Query("name")
	age, _ := strconv.Atoi(c.Query("age"))
	c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
}

func post1(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
}

func post2(c *gin.Context) {
	var dto webStruct
	c.BindJSON(&dto)
	c.JSON(http.StatusOK, gin.H{"name": dto.Name, "age": dto.Age})
}

type webStruct struct {
	Name string
	Age  int
	Desc string
}
