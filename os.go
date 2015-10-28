package main
import (
	"os"
	"fmt"
)

func main(){
	HostnameTest()
	PathSeparatorTest()
}

func HostnameTest(){
	name,err:=os.Hostname()
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(name)
}
func PathSeparatorTest(){
	fmt.Println(string(os.PathSeparator))  //os.PathSeparator is ascii 
}
