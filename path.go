package main
import (
	"fmt"
	"path/filepath"
)

func main(){
	FromSlash()
	Join()
}

func FromSlash(){
	name:=filepath.FromSlash("/var/cache/hekad") // if os.PathSeparator == '/' return path ,else replace '/' with os.PathSeparator.
	fmt.Println(name)
}
func Join(){
	fullpath:=filepath.Join("/a/","/b") // combine the path
	fmt.Println(fullpath)
}
