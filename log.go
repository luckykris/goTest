package main
import(
	"log"
	"fmt"
	"os"
	//"io/ioutil"
)

func main(){
	logfile,_ := os.OpenFile("/tmp/test.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0);
	logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile|log.Lshortfile|log.LstdFlags)   // LstdFlags==Ldate | Ltime 
	logger.Println("asdasd")
	logger.SetPrefix("[MYTEST]") // it will set the string to the head of each line.
	fmt.Println(logger.Prefix()) // return prefix
	logger.Println("test prefix") 

	logger.Output(0,"[ERROR]") // logger the stack depth of the function file & line.

	fmt.Println(logger.Flags()) // return flags
	logger.SetFlags(1)  //change logger flags
	logger.Println("test flags") 

	logger.Panicln("PANIC") // log and exit
	logger.Fatalln("fatal exit") // log and Exit(1)
}
