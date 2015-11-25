package main
import (
	"fmt"
	"regexp"
)

var TestString=[]byte(`2015-11-14 21:35:41,570 [nioEventLoopGroup-2-14] DEBUG com.roiland.dcs.component.gb04.P040101Handler - [E2020R10002990FAK] KEY:0055:VALUE:30.0 2015-11-14 21:35:41,570 [nioEventLoopGroup-2-14]`)

func main(){
//	Compile()
//	Find()
//	SubexpNames()
//	NumSubexp()
//	FindAllStringSubmatch()
	FindSubmatchIndex()
}

func Compile()*regexp.Regexp{
	r,err:=regexp.Compile(`(?P<Timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3})`)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return r

}
func Find(){
	r:=Compile()
	fmt.Println(string(r.Find(TestString)))
	
}
func SubexpNames(){
	r:=Compile()
	fmt.Println(r.SubexpNames())
}
func NumSubexp(){
	r:=Compile()
	fmt.Println(r.NumSubexp())
}
func FindAllStringSubmatch(){
	r:=Compile()
	fmt.Println(r.FindAllStringSubmatch(string(TestString),10000))
	for i,k:=range r.FindAllStringSubmatch(string(TestString),1){
		fmt.Printf(" index------------:%d\n",i)
		for _,j:=range k{
			fmt.Printf(" value:%s\n",j)
		}	
	}
}
func FindSubmatchIndex(){
	r:=Compile()
	fmt.Println(r.FindSubmatchIndex(TestString))
	for _,k:=range r.FindSubmatchIndex(TestString){
                fmt.Printf(" index------------:%d\n",k)
        }
}
