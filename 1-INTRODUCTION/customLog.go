package main 

import (
	"fmt"
	"os"
	"path"
	"log"
)

func customLog(){
	LOGFILE:=path.Join(os.TempDir(),"mcalog.txt")
	fmt.Println("LOGFILE path:",LOGFILE)

	f,err:=os.OpenFile(LOGFILE,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err !=nil{
		fmt.Println("Error opening file:",err)
	}
	_,err=f.WriteString("This is a log entry.\n")
	if err!=nil{
		fmt.Println("Error writing to file:",err)
		return
	}
	fmt.Println("Log written successfully")
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go")
}