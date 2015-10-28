package main
import (
	"os"
	"fmt"
)

func main(){
	HostnameTest()
}

func HostnameTest(){
	name,err:=os.Hostname()
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(name)
}
