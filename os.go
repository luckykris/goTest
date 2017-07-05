package main
import (
	"os"
	"fmt"
	"syscall"
)

func main(){
	//Lstat 获取FileInfo 接口 ，文件基本信息,如果是链接，则返回链接信息
	fi,_:=os.Lstat("/etc/passwd")
	fmt.Println(fi)
	
	//Stat 获取FileInfo 接口 ，文件基本信息.
	fi,_=os.Stat("/etc/passwd")
	fmt.Println(fi)

	//Hostname()
	//PathSeparator()
	//Open()
	//MkdirAll()
	//FindProcess()
	//Create()
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
func FindProcess(){
	process,err:=os.FindProcess(1) //return  Process ,but do nothing.
	if err!=nil{
		panic("Error")
	}
	if err= process.Signal(syscall.Signal(0));err !=nil{ // syscall 0 is like 'ping' to the process,just check if process is exist
		panic("did not exist")
	}
	fmt.Println("process exist")
}
func Create(){
	_,err:=os.Create("/tmp/testCreat.txt") //return *File,error ,create a file with 644 permission
	if err!=nil{
		panic("error")
	}
}

