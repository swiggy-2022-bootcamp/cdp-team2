package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	_ "github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server/httphandlers"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
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

	// read the key files before starting http handlers
	verifyKeyByte, err := ioutil.ReadFile(util.PubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
	util.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
