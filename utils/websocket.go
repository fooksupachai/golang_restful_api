package utils

import (
	"net/http"

	"github.com/gorilla/websocket"

	m "github.com/fooksupachai/golang_restful_api/model"
)

const maxMagSize = 5000

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket for structure of websocket server
type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]m.EventHandler
}

// NewWebSocket for create new socket server
func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err.Error())
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]m.EventHandler),
	}

	go ws.Reader()
	go ws.Writer()

	return ws, nil
}

// Reader function to recieve data fron socket
func (ws *WebSocket) Reader() {

	defer func() {
		ws.Conn.Close()
	}()

	ws.Conn.SetReadLimit(maxMagSize)

	for {
		_, message, err := ws.Conn.ReadMessage()

		if err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {

				SocketMessage(err)
			}

			break
		}

		event, err := m.NewEventFromRaw(message)

		if err != nil {
			SocketMessage(err)
		} else {
			SocketMessage(err)
		}

		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		}
	}
}

// Writer function to write data into web socket
func (ws *WebSocket) Writer() {

	for {

		select {

		case message, ok := <-ws.Out:

			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := ws.Conn.NextWriter(websocket.TextMessage)

			if err != nil {
				SocketMessage(err)
			}
			w.Write(message)
			w.Close()
		}
	}
}

// On function to return socket server
func (ws *WebSocket) On(event string, action m.EventHandler) *WebSocket {
	ws.Events[event] = action
	return ws
}
