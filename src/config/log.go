package config

import (
	"log"
	"os"
	"path/filepath"
)

var Logger map[string]*log.Logger

func setLog(key string) *log.Logger {
	if ConfigVal.PathLog[key] == nil {
		log.Fatalf("please ensure you have proper setup config file for %s", key)
	}
	logPath := ConfigVal.PathLog[key].Path

	isFileFolderExist, err := IsFileFolderExist(logPath)
	if err != nil {
		log.Fatalf("error opening file %s", err.Error())
	}
	if !isFileFolderExist {
		os.MkdirAll(filepath.Dir(logPath), 0777)
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	l := log.New(os.Stdout, "", 0)
	// defer f.Close()
	l.SetOutput(f)

	return l
}

func IsFileFolderExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
