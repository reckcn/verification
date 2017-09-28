package main

import (
	"fmt"
	"math/rand"
	"slide.verification"
)

/**
滑动验证码的实现
  刷新每次返回两张图片,两张图片的url,以及这种图片对应的唯一token
 当校验成功,返回true,以及对应的token
当校验失败,返回false,以及对应的token,同时自动再调用获取新的滑动验证码的接口

 */



var globalSessions *slide_verification.Manager
//然后在init函数中初始化
func init() {
	globalSessions, _ = slide_verification.NewManager("memory","gosessionid",3600)
}

func main()  {
	for  i := 0;i<=100 ;i++ {
		fmt.Println(rand.Intn(4) + 2)
	}


}