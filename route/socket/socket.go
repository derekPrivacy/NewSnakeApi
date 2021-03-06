package route

import (
	"log"
	"net/http"
	"sync"

	wsModel "api/model/socket"

	"github.com/gorilla/websocket"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var cReq *wsModel.CReq = &wsModel.CReq{}

	var Upgrader websocket.Upgrader

	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	jsonErr := conn.ReadJSON(cReq)
	if jsonErr != nil {
		log.Println(jsonErr)
		return
	}

	switch cReq.Type {

	case "updateAvatar":
		var wg sync.WaitGroup

		wg.Add(1)

		go UpdateAvatar(cReq, &wg)

		wg.Wait()
		break

	}

}
