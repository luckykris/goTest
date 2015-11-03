package main
import (
	"fmt"
	"strings"
)

func main(){
//	HasPrefix()
//	HasSuffix()
//	TrimSpace()
//	Index()
	ToLower()
}

func HasPrefix(){
	fmt.Println(strings.HasPrefix("aaa.qweasdas.toml","aaA."))
}
func HasSuffix(){
	fmt.Println(strings.HasSuffix("aaa.qweasdas.toml",".toml"))
}
func TrimSpace(){
	fmt.Println(strings.TrimSpace("   AAA    ")) // remove all leading or tailing white space  
}
func Index(){
	fmt.Println(strings.Index("abcdc","c"))
}
func ToLower(){
	s:=strings.ToLower("ABC")
	fmt.Println(s)
}
