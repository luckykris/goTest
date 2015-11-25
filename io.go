package main
import (
	"fmt"
	"io/ioutil"
)

func main(){
	ReadDir()
	ReadFile()
}

func ReadDir(){
	fiList,_:=ioutil.ReadDir("/tmp/") // return []os.FileInfo
	for _,fi :=range fiList{
		fmt.Println(fi.Name())
	}
}
func ReadFile(){
	content,err:=ioutil.ReadFile("/root/test.txt") //return []byte
	if err!=nil{
		panic("read failed")
	}
	fmt.Println(string(content))
}
