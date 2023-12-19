package tools

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	d "smartCalc/domains"
)

type Logs struct {
	logFile    *os.File
	infoLog    *log.Logger
	errorLog   *log.Logger
	debugLog   *log.Logger
	deepBugLog *log.Logger
}

var lg Logs = *StartLogs("smartCalc.log")

// Init logging
func StartLogs(logFileName string) *Logs {
	var lg Logs
	var err error

	if lg.logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777); err == nil {
		lg.infoLog = log.New(lg.logFile, "INFO\t", log.Ldate|log.Ltime)
		lg.errorLog = log.New(lg.logFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		lg.debugLog = log.New(lg.logFile, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
		lg.deepBugLog = log.New(lg.logFile, "DEEPBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
		return &lg
	} else {
		fmt.Println("cannot open logfile: ", logFileName, "\n logs will be write to terminal")
		lg.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		lg.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		lg.debugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
		lg.deepBugLog = log.New(os.Stdout, "DEEPBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return &lg
}

func DbgPrint(s string) {
	if d.Config.Debug {
		fmt.Println(s)
		lg.debugLog.Println(s)
	}
}

func LogPrint(s string, level int) {
	var LogLevel int = 2
	if LogLevel > 0 && LogLevel >= level {
		switch level {
		case 4:
			lg.deepBugLog.Println(s)
		case 3:
			lg.debugLog.Println(s)
		case 2:
			lg.infoLog.Println(s)
		case 1:
			lg.errorLog.Println(s)
		}
	}
}

// Load Image from file to image object
func LoadImage(fileName string) (im image.Image, err error) {

	fd, err := os.Open(fileName)
	DbgPrint(fmt.Sprint("OPEN File in LoadImage"))
	if err != nil {
		DbgPrint(fmt.Sprint(err))
		return
	}

	im, err = png.Decode(fd)
	DbgPrint(fmt.Sprint("DECODE file in LoadImage"))
	if err != nil {
		DbgPrint(fmt.Sprint(err))
		return
	}

	err = fd.Close()
	if err != nil {
		DbgPrint(fmt.Sprint(err))
	}
	return
}

func ExportImageToPng(im image.Image, fileName string) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		DbgPrint(fmt.Sprintln("Create File in ExportImageToPng:", err))
		return
	}
	if err = png.Encode(f, im); err != nil {
		f.Close()
		DbgPrint(fmt.Sprintln("Encode Image in ExportImageToPng:", err))
		return
	}
	return
}

// writing data to file
func WriteData(fileName string, data []byte) (err error) {
	fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	err = os.WriteFile(fileName, data, 0777) // write json([]byte) to file
	fd.Close()
	return
}

// reading data from file
func ReadData(fileName string) (data []byte, err error) {

	fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		return
	}
	data, err = os.ReadFile(fileName)

	fd.Close()
	return
}
