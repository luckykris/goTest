package main
import (
	"fmt"
	"strings"
)

func main(){
	HasPrefix()
	HasSuffix()
	TrimSpace()
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
