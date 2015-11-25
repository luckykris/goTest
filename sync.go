package main
import (
	"sync"
	"fmt"
)

func main(){
	//pool()
	once()
}

func pool(){
	pool:=new(sync.Pool)
	pool.New=func()interface{}{
		return "None" //when you use Get() and there is no item in pool,it will be call and return the result as the return of Get().
	}
	pool.Put("t3")
	pool.Put("t2")
	pool.Put("t1")
	for i:=0;i<9;i++{
		a:=pool.Get()
 		fmt.Println(a)
	}
}

func t(){
	fmt.Println(1)
}
func once(){
	o:=new(sync.Once) //only do the function once
	o.Do(t)
	o.Do(t)
	o.Do(t)
}
