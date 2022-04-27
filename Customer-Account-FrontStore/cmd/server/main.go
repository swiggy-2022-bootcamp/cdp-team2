package main

import (
	server1 "customer-account/internal/server"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"fmt"
)
 
func Init() {
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
	Init()
	server1.RunServer()
}


func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
