package jrlog

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jianliangan/flowmanage/libs/tools/unit"
)

func Loginit(console bool) error {
	syspath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	//local, _ := time.LoadLocation("Local")
	isexits, _ := unit.PathExists(syspath + "/log/node.log")
	log.Println(syspath, ",,,,,,,,,,,")
	if !isexits {
		isexits, _ := unit.PathExists(syspath + "/log/")
		if !isexits {
			err = os.Mkdir(syspath+"/log/", 0777)
			if err != nil {
				return err
			}
		}
	}
	var f *os.File
	f, err = os.OpenFile(syspath+"/log/node.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	if console {
		writers := []io.Writer{
			f,
			os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.SetOutput(f)
	}

	return nil
}
func Logfatal(v ...interface{}) {
	log.Fatal(v...)
}
func Logprintln(v ...interface{}) {
	log.Println(v...)
}
