package current

import (
	"brostools-api-person/lib/log"
	//"path"
	"os"
	"runtime"
)

// __FILE__
func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

// __DIR__
func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Infof(nil, "Error dir: "+dir)
		return ""
	}
	return dir
}
