package main
import (
	"os"
	"fmt"
)

func main(){
	//HostnameTest()
	//PathSeparatorTest()
	OpenTest()
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
func OpenTest(){
	fd, err := os.Open("/tmp") //return *File open a exist file or directory
	if err != nil {
		fmt.Printf("Error opening  file: %s", err.Error())
	}
	fi,err:=fd.Stat() //return *FileInfo
	if err != nil {
		fmt.Printf("Error fetch file info: %s", err.Error())
	}
	fmt.Println(fi.IsDir())
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode()) // return type uint32
	fmt.Println(fi.ModTime()) // return time.Time
	fmt.Println(fi.Sys())
}
