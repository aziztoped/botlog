// Package conf reads all main configs on botlog app files and map it into corresponding modules
package config

import (
	"log"
	"os"
	"strings"
	"time"

	gcfg "gopkg.in/gcfg.v1"
	logging "gopkg.in/tokopedia/logging.v1"
)

var ConfigVal config

var TimeStart time.Time

type config struct {
	PathLog map[string]*pathLog
}

type pathLog struct {
	Path string
}

func Init() {
	appName := strings.Replace(os.Args[0], "./", "", -1)
	cfgenv := os.Getenv("TKPENV")
	if cfgenv == "" {
		log.Printf("[%s] No environment set. Using 'development'.", appName)
		log.Printf("[%s] Use 'export TKPENV=[development|alpha|staging|production]' to change.", appName)
		cfgenv = "development"
	}

	var ok bool
	types := []string{"main"}
	for _, v := range types {
		ok = logging.ReadModuleConfig(&ConfigVal, "/etc/botlog", v) || logging.ReadModuleConfig(&ConfigVal, "files/etc/botlog", v)
		if !ok {
			// when the app is run with -e switch, this message will automatically be redirected to the log file specified
			log.Fatalln("Failed to read main config")
		}
	}

	keys := []string{"product"}

	Logger = make(map[string]*log.Logger)
	for _, v := range keys {
		Logger[v] = setLog(v)
	}
}

func ReadConfig(filePath string) (bool, config) {
	var c config
	if err := gcfg.ReadFileInto(&c, filePath); err != nil {
		log.Printf("%s\n", err)
		return false, c
	}
	return true, c
}
