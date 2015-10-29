package main
import (
	"time"
	"fmt"
)

func main(){
	String()
	ParseDuration()
}

func String(){
	a:=187*time.Second
	fmt.Println(a.String())
}
func ParseDuration(){
	duration,err:=time.ParseDuration("5h34m2s1ms1us1ns")
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(duration.String())
}
