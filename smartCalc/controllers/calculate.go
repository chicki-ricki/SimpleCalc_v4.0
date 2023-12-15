package controllers

import (
	"fmt"
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	// i18n.Locate // For i18n usage when process data and render template.
}

type CalculateController struct {
	baseController
}

func (c *CalculateController) Calculate() {
	c.TplName = "calculate/startCalculate.tpl"
}

func (c *CalculateController) Start() {
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
	messageType, text, err := ws.ReadMessage()
	if err != nil {
		log.Println("Read message error:", err)
		return
	}
	fmt.Println("ws: ", string(text))
	text1 := []byte("this from calculate")
	if err := ws.WriteMessage(messageType, text1); err != nil {
		log.Println("Write message error:", err)
		return
	}
	defer ws.Close()
}
