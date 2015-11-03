package main
import (
	"fmt"
)
func main(){
	Cap()
}
func Cap(){
	a:=make(chan int,10)
	fmt.Println(cap(a))
}
