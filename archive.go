package main
import(
	"archive/tar"
	"os"
	"fmt"
	"io/ioutil"
	"io"
)

func main(){
	//FileInfoHeader  通过os.FileInfo生产header 并且可以打一个link的标签名 信息os.FileInfo
	fi,_:=os.Lstat("/tmp/1.txt")
	h,_:=tar.FileInfoHeader(fi,"asd")
	fmt.Println(h.Linkname)
	
	//FileInfo 从header里获取文件信息os.FileInfo
	fi=h.FileInfo()
	fmt.Println(fi)
	
	//NewWriter 生生一个tar压缩的writer
	fd, _ := os.OpenFile("/tmp/archive_test_encode.tar", os.O_RDWR|os.O_CREATE, 0755)
	writer:=tar.NewWriter(fd)
	
	//WriteHeader 把头信息写到tar头里
	writer.WriteHeader(h)

	//Write 把内容写进压缩包
	fd2,_ :=os.Open("/tmp/1.txt")
	b,_:=ioutil.ReadAll(fd2)
	writer.Write(b)

	//Flush 刷已写完的数据
	writer.Flush()

	//Close 关闭写入tar包
	writer.Close()
	fd.Close()

	//NewReader 生产一个度tar的reader
	fd, _ = os.Open("/tmp/archive_test_encode.tar")
	r:=tar.NewReader(fd)
	
	//Read方法是io.reader接口的继承 必须逐步都到参数的[]byte里，费劲所以一般不适用，都是实用 io.Copy，ioutil的ReadAll全部读取 
	
	
	//Next 挨个读取头信息，并且从reader缓冲字段获取数据
	for tmp_h,err:=r.Next();err!=io.EOF;tmp_h,err= r.Next(){
		tmp_fd,_ :=os.Create(tmp_h.Name)
		defer tmp_fd.Close()
		io.Copy(tmp_fd,r)
	}
}

