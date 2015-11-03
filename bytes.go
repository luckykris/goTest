package main
import (
	"bytes"
	"fmt"
)

func main(){
	//NewBuffer()
	//WriteByte()
	WriteString()
}

func NewBuffer()*bytes.Buffer{
	b:=make([]byte,64)
	buf:=bytes.NewBuffer(b)
	return buf
}
func WriteByte(){
	buf:=NewBuffer()
	buf.WriteByte(76)
	fmt.Println(string(76))
}
func WriteString(){
	buf:=NewBuffer()
	buf.WriteString(`76`)
	fmt.Println(string(`76`))
}
