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
	runtime.Breakpoint()
	WriteHeapProfile()
}	

func WriteHeapProfile(){
	profFile, _ := os.Create("/tmp/prof.txt")
	pprof.WriteHeapProfile(profFile)
	profFile.Close()
}
