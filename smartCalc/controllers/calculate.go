package controllers

import (
	"fmt"
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"

	// d "smartCalc/domains"

	m "smartCalc/model"
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

func (c *CalculateController) Calculate() {
	c.TplName = "calculate/startCalculate.tpl"
}

func (c *CalculateController) Start() {
	var output string
	c.TplName = "calculate/startCalculate.tpl"
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
			output = (c.cnv.ModelToUI((m.ModelCalc.GetCalcResult(input))))
		}
		if err := ws.WriteMessage(messageType, []byte(output)); err != nil {
			log.Println("Write message error:", err)
			return
		}
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
