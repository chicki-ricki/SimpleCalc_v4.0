package tools

import (
	"fmt"
	"image"
	"image/png"
	"os"

	d "smartCalc/domains"
)

func DbgPrint(s string) {
	if d.Config.Debug {
		fmt.Println(s)
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
