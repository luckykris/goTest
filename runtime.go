package main
import (
	//"fmt"
	"runtime"
	"runtime/pprof"
	"os"
)

func main(){
	//MemProfileRate 是内存概况记录比率，含义是分配了多少内存进行一次打点纪录。
	runtime.MemProfileRate = 512*1024 
	
	//Breakpoint 打断点，进行断点捕捉
	//runtime.Breakpoint()

	//CPUProfile 返回cpu下一个二进制状态栈数据，如果profiling关闭,则返回nil
	println(runtime.CPUProfile())
	
	//Caller 返回文件跟行数信息
	pc,file,line,ok:=runtime.Caller(0)
	println(pc,file,line,ok)
	WriteHeapProfile()
}	

func WriteHeapProfile(){
	profFile, _ := os.Create("/tmp/prof.txt")
	pprof.WriteHeapProfile(profFile)
	profFile.Close()
}
