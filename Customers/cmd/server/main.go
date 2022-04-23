package main

import (
	"customers/internal/server"
	log "github.com/sirupsen/logrus"
	"fmt"
	"runtime"
	"strings"
 )

func init() {
	// setting logger configurations
	log.SetReportCaller(true)

	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
}


func main(){
	if err:=server.RunServer();err!=nil{
		log.WithField("Error:",err).Fatalf("Server quiting....")
	}
}


func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}


