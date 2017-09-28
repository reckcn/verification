package slide_verification

import (

	"math/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)


var globalSessions *Manager
//然后在init函数中初始化
func init() {
	globalSessions, _ = NewManager("memory","gosessionid",3600)
}

func main()  {
	for  i := 0;i<=100 ;i++ {
		fmt.Println(rand.Intn(4) + 2)
	}



}

/**
 获取滑动验证码
  */
func getCode(c *gin.Context)  {
	// [minRangeX,MaxRangeX]范围
	positionX := rand.Intn(minRangeX+1) + maxRangeX - minRangeX
	positionY := rand.Intn(minRangeY+1) + maxRangeY - minRangeY
	fmt.Sprintf("getCode|positionX=%s|positionY=%s",positionX,positionY)
	session := globalSessions.SessionStart(c.Writer, c.Request)
	session.Set("code",positionX)
	session.Set("code_errornum",nil)
	var a = [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}


}

/**
校验滑动验证码
 */
func checkSlideVerificationCode()  (right bool,err error){


	return true,nil
}

func minute(A,B int) int {
	right := A - B
	if right == 0 {
		return 0
	}
	return random(right)
}
func random(right int) int {
	rightString := strconv.Itoa(right)
	index := rand.Intn(len(rightString))
	bytes := []byte(rightString)
	if index == 0 {
		bytes[0] = byte(rand.Intn(9)+1)
		i,_ := strconv.Atoi(string(bytes))
		return i
	}
	bytes[index] = byte(rand.Intn(10))
	i,_ :=strconv.Atoi(string(bytes))
	return i
}




