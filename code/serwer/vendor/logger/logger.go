package logger

import (
	//	"fmt"
	"log"
	"os"
)

var (
	Trace      *log.Logger
	Info       *log.Logger
	Warning    *log.Logger
	Error      *log.Logger
	FatalError *log.Logger
)

func Start(filepatch string) {
	LogConn, err := os.OpenFile(filepatch,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	f := log.Ldate //| log.LstdFlags | log.Lshortfile
	ftracer := log.Ldate | log.LstdFlags | log.Lshortfile
	Trace = log.New(LogConn, "Trace: ", ftracer)
	Info = log.New(LogConn, "Info: ", f)
	Warning = log.New(LogConn, "Warning: ", f)
	Error = log.New(LogConn, "Error: ", f)
	FatalError = log.New(LogConn, "FatalError: ", f)
	//	fmt.Println(FatalError.Prefix())

}
