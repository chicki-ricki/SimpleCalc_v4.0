package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"

	// d "smartCalc/domains"

	d "smartCalc/domains"
	m "smartCalc/model"
	t "smartCalc/tools"
)

type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	// i18n.Locate // For i18n usage when process data and render template.
}

type CalculateController struct {
	baseController
	cnv convert
	mod m.CalcModel
}

type MessageToUI struct {
	mode   int
	result string
}

func (c *CalculateController) Calculate() {
	// c.TplName = "calculate/startCalculate.tpl"
	c.TplName = "calculate/startCalculate.html"
}

func (c *CalculateController) removeTmpGraph(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		if os.Remove(fileName) != nil {
			fmt.Println("Cannot remove tempfile")
		}
	}
}

func (c *CalculateController) Start() {
	var output string
	// c.TplName = "calculate/startCalculate.tpl"
	c.TplName = "calculate/startCalculate.html"
	fmt.Println("start function are going")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatalf("Cannot setup WebSocket connection: %v\n", err)
	}
	defer ws.Close()

	uname := "_" + fmt.Sprint(rand.Intn(6000))
	tmpGraphImageName := d.Config.TempFileDir + "tempGraph" + uname + ".png"
	defer os.Remove(tmpGraphImageName)

	if err := ws.WriteMessage(1, []byte("5 "+uname)); err != nil {
		log.Println("Write message error:", err)
	}
	// Message receive loop.
	for {
		messageType, text, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("ws: ", string(text))
		input, er := c.cnv.UIToModel(string(text))
		if er {
			output = ("Error string")
		} else {
			modelsOutput := m.ModelCalc.GetCalcResult(input)

			if modelsOutput.Mode == 2 && t.ExportImageToPng(modelsOutput.ModelGraphResult.GraphImage, tmpGraphImageName) != nil {
				fmt.Println("cannot write tempGraph image to disk")
			}
			output = (c.cnv.ModelToUI(modelsOutput))
		}
		if err := ws.WriteMessage(messageType, []byte(output)); err != nil {
			log.Println("Write message error:", err)
			return
		}
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
