package main
import (
	"os"
	"fmt"
)

func main(){
	Hostname()
	PathSeparator()
	Open()
	MkdirAll()
}

func Hostname(){
	name,err:=os.Hostname()
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(name)
}
func PathSeparator(){
	fmt.Println(string(os.PathSeparator))  //os.PathSeparator is ascii 
}
func Open(){
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
func MkdirAll(){
	err:=os.MkdirAll("/tmp/a/b/", 0755) // same as shell "mkdir -p /tmp/a/b"
	if err!=nil{
		panic("Error")
	}
}
