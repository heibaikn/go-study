package controllers
import (
	"fmt"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	fmt.Printf("string")
	sum :=0
	for i:=0;i<10 ; i++ {
		sum +=i
	}
	fmt.Print(sum)
}