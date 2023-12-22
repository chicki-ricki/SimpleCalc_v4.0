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
			t.Clg.Warning(fmt.Sprintf("Cannot remove tempfile"))
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
	t.Clg.Info("_Start_ \n\nRUN SESSION")

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		t.Clg.Error(fmt.Sprintf("_Start_ Cannot setup WebSocket connection: %v\n", err))
		log.Fatalf("Cannot setup WebSocket connection: %v\n", err)
	}
	defer ws.Close()

	uname := "_" + fmt.Sprint(rand.Intn(6000))
	tmpGraphImageName := d.Config.TempFileDir + "tempGraph" + uname + ".png"
	t.Clg.Info(fmt.Sprintf("_Start_ uname = %s; tmpGraphImageName := %s", uname, tmpGraphImageName))
	defer c.removeTmpGraph(tmpGraphImageName)

	if err := ws.WriteMessage(1, []byte("5 "+uname)); err != nil {
		t.Clg.Warning(fmt.Sprintf("_Start_ Write message error: %v", err))
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
		if string(text) == "clearHistory" {
			// clearHistory()
			if err := ws.WriteMessage(1, clearHistory()); err != nil {
				t.Clg.Warning(fmt.Sprint("_Start_ Can not write clear hisory data:", err))
			}
			continue
		}
		t.Clg.Info(fmt.Sprintf("_Start_ Message from user %s: %s", uname, string(text)))
		input, er := c.cnv.UIToModel(string(text))
		t.Clg.DeepDebug(fmt.Sprint("_Start_ input: ", input))
		if er {
			output = ("Error string")
		} else {
			modelsOutput := m.ModelCalc.GetCalcResult(input)
			t.Clg.DeepDebug(fmt.Sprint("_Start_ modelsOutput After equation:", modelsOutput))
			if modelsOutput.Mode == 2 && !modelsOutput.Err &&
				t.ExportImageToPng(modelsOutput.ModelGraphResult.GraphImage, tmpGraphImageName) != nil {
				t.Clg.Warning("_Start_ cannot write tempGraph image to disk")
			}
			output = (c.cnv.ModelToUI(modelsOutput))
		}
		t.Clg.Info(fmt.Sprintf("_Start_ Message To user%s: %s", uname, output))
		if err := ws.WriteMessage(messageType, []byte(output)); err != nil {
			t.Clg.Warning(fmt.Sprint("_Start_ Write message error:", err))

			return
		}

		if err := ws.WriteMessage(1, lastHistory(input, output)); err != nil {
			t.Clg.Warning(fmt.Sprint("_Start_ Can not write data from model:", err))
		}

		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
