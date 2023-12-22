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
}

type CalculateController struct {
	baseController
	cnv convert
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
	var (
		output   string
		upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
	)

	// c.TplName = "calculate/startCalculate.tpl"
	c.TplName = "calculate/startCalculate.html"

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatalf("Cannot setup WebSocket connection: %v\n", err)
	}
	defer ws.Close()

	uname := "_" + fmt.Sprint(rand.Intn(6000))
	tmpGraphImageName := d.Config.TempFileDir + "tempGraph" + uname + ".png"
	defer c.removeTmpGraph(tmpGraphImageName)

	if err := ws.WriteMessage(1, []byte("5 "+uname)); err != nil {
		log.Println("Write message error:", err)
	}
	
	if err := ws.WriteMessage(1, loadHistoryFromModel()); err != nil {
		log.Println("Can not write data from model:", err)
	}

	// Message receive loop.
	for {
		messageType, text, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("ws_text: ", string(text))
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

		// if err := ws.WriteMessage(1, lastHistory(input, output)); err != nil {
		// 	log.Println("Can not write data from model:", err)
		// }

		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
