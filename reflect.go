package main
import (
	"reflect"
	"fmt"
)

func main(){
	ValueTest()
}
type ValueTestStruct struct{
	a int
	n string
	t ValueTestStruct2
}
type ValueTestStruct2 struct{
        a int
        n string
	i ValueTestInterface
}
type ValueTestInterface interface{
	ValueTestFunc(i int) bool
}

func (x *ValueTestStruct2)ValueTestFunc(i int)bool{
	return x.a>i
}
func ValueTest(){
	var val int = 5
	val2:=ValueTestStruct{a:12,n:"asd"}
	v:=reflect.ValueOf(val) //get a Value type of interface{}
	v2:=reflect.ValueOf(val2) 
	v3:=reflect.ValueOf(&val2) 
	fmt.Println(v)
	fmt.Println(v.Kind()) // return Kind ,the type of v
	fmt.Println(v3.Elem()) //return Value ,v must be Ptr or interface unless it will panic,it return what v poiont to or what v contain
	fmt.Println(v.Type()) //return Type
	fmt.Println(v2.NumField()) //return int, return how many fields in  struct,it panic if v2 not struct.
	fmt.Println(v2.Field(0).Kind()) //return Value, return the field which index is '0' 
	structField_test:= v2.Type().Field(0)
	structField_test.Tag = `test1:"1" test2:"2"` //set Tag of structField, key:"value" ,each key is split by a space charactor
	fmt.Println(structField_test) // return structField ,Example {a main int  0 [0] false} <==> {Name string ,PkgPath string ,Type Type,Tag StructTag,offset uintptr,Index []int,Anonymous bool}
	fmt.Println(structField_test.Tag.Get("test1")) // return string ,get the key`s value ,if key not exsist ,return a empty string
	fmt.Println(v3.Elem().CanSet()) // return bool, v3 must a PTR of struct ,and use Elem to get the struct.
}
