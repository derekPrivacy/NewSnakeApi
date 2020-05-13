package route

import (
	"encoding/json"
	"log"
	"net/http"

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

	var cRes *wsModel.CRes

	switch cReq.Type {

	case "updateAvatar":

		UpdateAvatar(cReq)

		break

	case "spawnFood":

		cRes = SpawnFood(cReq)

		res2B, _ := json.Marshal(cRes)
		log.Printf("whats in you cRes " + string(res2B))

		if len(cRes.Room.Data) >= 2 {

			for _, c := range cRes.Room.ConnectionArr {
				c.Connection.WriteJSON(cRes.Room)
			}
		}

		break
	case "gameFinish":

		break

	}

}
