package main
import(
	"fmt"
)
//constants 静态变量声明
const (
	A string="aaa"
)

// variable 变量声明
var nil interface{}

func main(){
	//done all
	//append 追加slice的元素到slice尾部,向slice1种追加元素2,3
	slice1:=[]int{1}
	slice:=append(slice1,2,3)
	fmt.Println(slice)

	//cap 获取容量，参数可以为array,pointer,slice,nil,channel
	fmt.Println(cap(slice1))
	
	//len 获取元素个数，参数可以为array,pointer,slice,nil,channel
	fmt.Println(len(slice1))

	//close 关闭channel
	channel1:=make(chan interface{},2)
	close(channel1)
	
	//complex 把2个浮点型变成复数complex型，float32 float64分别转换为complex32 complex64
	var float1 float32 = 1.0
	var float2 float32 = 0.2
	complex1:= complex(float1,float2)
	fmt.Println(complex1)
	
	//real 返回复数的实数部分
	real1:=real(complex1)
	fmt.Println(real1)
	
	//imag 返回复数的虚数部分
	imag1:=imag(complex1)
	fmt.Println(imag1)
	
	//copy 传入两个slice参数， 函数把第二个slice里的元素copy到第一个参数的slice里，函数返回值时copy的slice个数
	slice2:=make([]int,5)
	slice3:=[]int{1,2,3,4}
	num:=copy(slice2,slice3)
	fmt.Println(num)
	fmt.Println(slice2)

	//delete 删除map的元素，如果map为nil或者删除key不存在则无影响
	map1:=map[int]int{}
	map1[1]=2
	delete(map1,1)
	value1,ok:=map1[1]
	fmt.Println(ok)
	fmt.Println(value1)
}
