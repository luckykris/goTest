package main
import (
	"fmt"
)
const(
	A=1<<iota
	B
	C
)
const(
	D=1
	E=iota //iota is index of the 'const zone'
	F
)

func main(){
	Const()
	//Cap()
}
func Const(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
func Cap(){
	a:=make(chan int,10)
	fmt.Println(cap(a))
}
