package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type websocketHandler struct {
	l *log.Logger
}

func NewWebsocketHandler(l *log.Logger) *websocketHandler {
	return &websocketHandler{l}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (w *websocketHandler) reader(conn *websocket.Conn) {
	for {
		messageType, data, err := conn.ReadMessage()

		if err != nil {
			w.l.Println(err)
			return
		}

		w.l.Println(string(data))

		if err := conn.WriteMessage(messageType, data); err != nil {
			w.l.Println(err)
			return
		}
	}
}

func (w *websocketHandler) WebsocketEndpoint(rw http.ResponseWriter, r *http.Request) {
	//Upgrade fonksiyonu HTTP sunucu bağlantısını Websocket protokolüne yükseltiyor
	ws, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		w.l.Println(err)
		return
	}

	w.l.Println("İstemci başarıyla bağlandı!")

	w.reader(ws)
}
