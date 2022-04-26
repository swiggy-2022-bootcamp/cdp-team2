package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/cart/cmd/server/docs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server/httphandlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.WithField("Error: ", err).Fatalf("Server quitting...")
	}
}

// init function sets the logging configurations and also, initializes the public and private
// keys for JWT authentication.
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

	publicKey := []byte("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCTd0I0EhRo7Kjh7oOch2w/hOZRBNgsi4eYa/PoHGE6EqihrVzRf6iQ/snmyfn+nTxeGdQ+StiLBl6eBcSJh0mrtfKvB41wm4Vvh+YUhvf+5Bl6ifaKHeLJl1d32gi/c+ZgSg3B/yVm5hZ8i3s/oud6qqk/+t/wU0MwOZlle06q7QIDAQAB")
	verifykey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}

	util.VerifyKey = verifykey
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
