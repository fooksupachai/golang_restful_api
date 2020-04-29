package controller

import (
	"net/http"

	m "github.com/fooksupachai/golang_restful_api/model"
	u "github.com/fooksupachai/golang_restful_api/utils"
)

// LiveChat websocket api
func LiveChat(w http.ResponseWriter, r *http.Request) {

	ws, err := u.NewWebSocket(w, r)

	if err != nil {
		res := u.Message("error")
		res["data"] = "Connot create socket server"
		u.Response(w, res)
	}

	ws.On("message", func(e *m.Event) {

		ws.Out <- (&m.Event{
			Name: "response",
			Data: e.Data,
		}).Raw()
	})
}
