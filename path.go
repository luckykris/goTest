package main
import (
	"fmt"
	"path/filepath"
)

func main(){
	FromSlashTest()
}

func FromSlashTest(){
	name:=filepath.FromSlash("/var/cache/hekad") // if os.PathSeparator == '/' return path ,else replace '/' with os.PathSeparator.
	fmt.Println(name)
}
