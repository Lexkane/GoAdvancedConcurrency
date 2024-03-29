package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)
func main(){
runtime.GOMAXPROCS(4)
mutex:=make(chan bool,1)

f,_:=os.Create("./log.txt")
f.Close()


logCh:=make(chan string,50)
go func (){
	for{
		msg,ok:=<-logCh
		if ok {
			f,_:=os.OpenFile("./logfile.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,os.ModeAppend)
			logTime:=time.Now().Format(time.RFC3339)
			f.WriteString(logTime+ "-" +msg)
			f.Close()
		} else {
				break
		}
		}
	}()



for i:=1;i<10;i++{
	for j:=1;j<10;j++{
		go func(){
			msg:=fmt.Sprintf("%d + %d = %d \n",i,j,i+j)
			logCh<-msg
			fmt.Print(msg)
			<-mutex
		}()
	}
}
	fmt.Scanln()

}
