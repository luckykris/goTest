package main
import (
	"time"
	"fmt"
)

func main(){
	StringTest()
	ParseDurationTest()
}

func StringTest(){
	a:=187*time.Second
	fmt.Println(a.String())
}
func ParseDurationTest(){
	duration,err:=time.ParseDuration("5h34m2s1ms1us1ns")
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(duration.String())
}
