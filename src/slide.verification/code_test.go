package slide_verification

import (
	"fmt"
	"math/rand"
	"testing"
)

func Testa(t *testing.T)()  {
	for  i := 0;i<=100 ;i++ {
		fmt.Println(rand.Intn(4) + 2)
	}

}
