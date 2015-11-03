package main
import (
	"time"
	"fmt"
)

func main(){
	//String()
	//ParseDuration()
	//Tick()
	//NewTicker()
	Now()
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
func Tick(){
	intger:=1
	b:=make(chan int)
	t:=time.Duration(time.Duration(intger)) * time.Second
	ticker:=time.Tick(t)
	select{
		case<-ticker:
			fmt.Println("tick")
		case<-b:
			fmt.Println("tok")
	}
}
func NewTicker(){
	ticker := time.NewTicker(time.Duration(5) * time.Second)
	ticker2 := time.NewTicker(time.Duration(8) * time.Second)
	go func(){ time.Sleep(time.Duration(2) * time.Second)
		   ticker.Stop()
		}()
	select{
		case<-ticker.C:
			fmt.Println("tick")
		case<-ticker2.C:
			fmt.Println("tick2")
	}
}
func Now(){
	t:=time.Now()
	t2:=time.Now().Add(time.Duration(2) * time.Minute)
	fmt.Println(t.Format("2006-01-02T15:04:05Z07:00"))
	fmt.Println(t.Add(time.Duration(2) * time.Minute).Format("2006-01-02T15:04:05Z07:00"))
	if t.Before(t2){
		fmt.Println(true)
	}
	if !t.After(t2){
		fmt.Println(true)
	}
}
