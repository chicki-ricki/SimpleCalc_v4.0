package tools

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	d "smartCalc/domains"
)

type cLogs struct {
	Level   int
	LogFile *os.File
	// strChan *chan string

	errorLog     *log.Logger
	infoLog      *log.Logger
	warningLog   *log.Logger
	debugLog     *log.Logger
	deepDebugLog *log.Logger
	mapCLog      map[int]string
}

var Clg cLogs = *StartLogs(d.Config.LogDir + d.Config.LogFile)

// defer lg.logFile.Close()

// Init logging
func StartLogs(logFileName string) *cLogs {
	var lg cLogs
	var err error

	lg.mapCLog = map[int]string{
		0: "Log don't writing",
		1: "ERROR",
		2: "WARNING",
		3: "INFO",
		4: "DEBAG",
		5: "DEEPDEBAG",
	}

	lg.Level = d.Config.Debug

	if lg.LogFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777); err == nil {
		lg.errorLog = log.New(lg.LogFile, "ERROR\t", log.Ldate|log.Ltime)
		lg.infoLog = log.New(lg.LogFile, "INFO\t", log.Ldate|log.Ltime)
		lg.warningLog = log.New(lg.LogFile, "WARNING\t", log.Ldate|log.Ltime)
		lg.debugLog = log.New(lg.LogFile, "DEBUG\t", log.Ldate|log.Ltime)
		lg.deepDebugLog = log.New(lg.LogFile, "DEEPBUG\t", log.Ldate|log.Ltime)
	} else {
		fmt.Println("cannot open logfile: ", logFileName, "\n logs will be write to terminal")
		lg.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		lg.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
		lg.warningLog = log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime)
		lg.debugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime)
		lg.deepDebugLog = log.New(os.Stdout, "DEEPDEBUG\t", log.Ldate|log.Ltime)
		lg.warningLog.Println("cannot open logfile: ", logFileName, "logs will be write to terminal")
	}
	lg.infoLog.Printf("\n\nSTART PROGRAMM: LOGLEVEL = %d - %s\n%s\n\n", lg.Level,
		lg.mapCLog[lg.Level], d.Config.LogDir+d.Config.LogFile)

	return &lg
}

func DbgPrint(s string) {
	if d.Config.Debug > 3 {
		fmt.Println(s)
		Clg.debugLog.Println(s)
	}
}

func (lg *cLogs) LogPrint(level int, s string) {
	var LogLevel int = 2
	if LogLevel > 0 && LogLevel >= level {
		switch level {
		case 4:
			lg.deepDebugLog.Println(s)
		case 3:
			lg.debugLog.Println(s)
		case 2:
			lg.infoLog.Println(s)
		case 1:
			lg.errorLog.Println(s)
		}
	}
}

func (lg *cLogs) Error(s string) {
	if lg.Level >= 1 {
		lg.errorLog.Println(s)
	}
}

func (lg *cLogs) Warning(s string) {
	if lg.Level >= 2 {
		lg.warningLog.Println(s)
	}
}

func (lg *cLogs) Info(s string) {
	if lg.Level >= 3 {
		lg.infoLog.Println(s)
	}
}

func (lg *cLogs) Debug(s string) {
	if lg.Level >= 4 {
		lg.debugLog.Println(s)
	}
}

func (lg *cLogs) DeepDebug(s string) {
	if lg.Level == 5 {
		lg.deepDebugLog.Println(s)
	}
}

func FileCheck(files []string) {
	var errCount int // count access error

	// Check files for exist
	for _, fileName := range files {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			Clg.Warning(fmt.Sprintf("_FileCheck_ Not available: %s", fileName))
			errCount++
		} else {
			Clg.DeepDebug(fmt.Sprintf("_FileCheck_ Available: %s", fileName))
		}
	}

	// Output Result FileCheck
	if errCount == 0 {
		Clg.Info("_FileCheck_ success all files are available")
	} else {
		Clg.Warning("_FileCheck_ some files not available")
	}
}

// Load Image from file to image object
func LoadImage(fileName string) (im image.Image, err error) {

	fd, err := os.Open(fileName)
	Clg.Debug("_LoadImage_ OPEN File in LoadImage")
	if err != nil {
		Clg.Warning(fmt.Sprint(err))
		return
	}

	im, err = png.Decode(fd)
	Clg.Debug("_LoadImage_ DECODE file in LoadImage")
	if err != nil {
		Clg.Warning(fmt.Sprint("_LoadImage_ ", err))
		return
	}

	if err = fd.Close(); err != nil {
		Clg.Warning(fmt.Sprint("_LoadImage_ ", err))
	}
	return
}

func ExportImageToPng(im image.Image, fileName string) (err error) {

	if f, err := os.Create(fileName); err != nil {
		Clg.warningLog.Println(fmt.Sprintln("_ExportImageToPng_ Create File in ExportImageToPng:", err))
	} else if err = png.Encode(f, im); err != nil {
		Clg.Warning(fmt.Sprintln("_ExportImageToPng_ Encode Image in ExportImageToPng:", err))
		f.Close()
	} else {
		f.Close()
	}
	return
}

// writing data to file
func WriteData(fileName string, data []byte) (err error) {
	// fd, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err = os.WriteFile(fileName, data, 0777); err != nil { // write json([]byte) to file
		Clg.Warning(fmt.Sprint("_WriteData_ Cannot write data to file: ", fileName))
	}
	// fd.Close()
	return
}

// reading data from file
func ReadData(fileName string) (data []byte, err error) {
	if data, err = os.ReadFile(fileName); err != nil {
		Clg.Warning(fmt.Sprint("_ReadData_ Cannot read from file: ", fileName))
	}
	return
}
