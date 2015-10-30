package main
import (
	//"fmt"
	"runtime/pprof"
	"os"
)
func main(){
	WriteHeapProfile()
}	

func WriteHeapProfile(){
	profFile, _ := os.Create("/tmp/prof.txt")
	pprof.WriteHeapProfile(profFile)
	profFile.Close()
}
