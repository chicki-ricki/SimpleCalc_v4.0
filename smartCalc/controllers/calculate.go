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

// type MessageToUI struct {
// 	mode   int
// 	result string
// }

func (c *CalculateController) Calculate() {
	// c.TplName = "calculate/startCalculate.tpl"
	c.TplName = "calculate/startCalculate.html"
}

func (c *CalculateController) removeTmpGraph(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		if os.Remove(fileName) != nil {
			t.Clg.Warning("Cannot remove tempfile")
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

	//Create user name
	uname := "_" + fmt.Sprint(rand.Intn(6000))

	// Compose temp graph file name for user
	tmpGraphImageName := d.Config.TempFileDir + "tempGraph" + uname + ".png"
	t.Clg.Info(fmt.Sprintf("_Start_ uname = %s; tmpGraphImageName := %s", uname, tmpGraphImageName))

	// Remove temp graph file after session
	defer c.removeTmpGraph(tmpGraphImageName)

	// Send User name to UI via ws
	if err := ws.WriteMessage(1, []byte("5 "+uname)); err != nil {
		t.Clg.Warning(fmt.Sprintf("_Start_ Write message error: %v", err))
	}

	// First load history from model
	if err := ws.WriteMessage(1, loadHistoryFromModel()); err != nil {
		log.Println("Can not write data from model:", err)
	}

	// Message receive loop
	for {
		// Waiting message from user
		messageType, text, err := ws.ReadMessage()
		if err != nil {
			return
		}

		// Handle clearHistory command case
		if string(text) == "clearHistory" {
			if err := ws.WriteMessage(1, clearHistory()); err != nil {
				t.Clg.Warning(fmt.Sprint("_Start_ Can not write clear hisory data:", err))
			} else {
				t.Clg.Info(fmt.Sprintf("_Start_ Command from user %s: ClearHistory - success", uname))
			}
			continue
		}

		t.Clg.Info(fmt.Sprintf("_Start_ Message from user %s: %s", uname, string(text)))

		// Convert message to input modeldata
		input, er := c.cnv.UIToModel(string(text))
		t.Clg.DeepDebug(fmt.Sprint("_Start_ input: ", input))
		if er {
			output = ("Error string")
		} else {
			//request to Model for result
			modelsOutput := m.ModelCalc.GetCalcResult(input)
			t.Clg.DeepDebug(fmt.Sprint("_Start_ modelsOutput After equation:", modelsOutput))

			// write graph image to file with users name
			if modelsOutput.Mode == 2 && !modelsOutput.Err &&
				t.ExportImageToPng(modelsOutput.ModelGraphResult.GraphImage, tmpGraphImageName) != nil {
				t.Clg.Warning("_Start_ cannot write tempGraph image to disk")
			}

			// Convert ModelsOutput to stringresult for UI
			output = (c.cnv.ModelToUI(modelsOutput))
		}

		// Send Result to UI
		if err := ws.WriteMessage(messageType, []byte(output)); err != nil {
			t.Clg.Warning(fmt.Sprintf("_Start_ Write message To user%serror:", err))
			return
		} else {
			t.Clg.Info(fmt.Sprintf("_Start_ Message To user%s: %s", uname, output))
		}

		// Create and send curent history item to UI
		if err := ws.WriteMessage(1, lastHistory(input, output)); err != nil {
			t.Clg.Warning(fmt.Sprint("_Start_ Can not write data from model:", err))
		}

		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
