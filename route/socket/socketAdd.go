package route

import (
	"log"
	"net/http"

	wsModel "api/model/socket"

	"github.com/gorilla/websocket"
)

func SocketAdd(w http.ResponseWriter, r *http.Request) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("websocketAdd called how many times")

	var cReq *wsModel.CReq = &wsModel.CReq{}

	var Upgrader websocket.Upgrader

	Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}

	jsonErr := conn.ReadJSON(cReq)
	if jsonErr != nil {
		log.Println(jsonErr)

	}

	switch cReq.Type {
	case "hello":

		SayHello(cReq)

		break
	case "addPlayer":

		AddPlayer(cReq, conn)
		break

	case "spawnFood":

		SpawnFood(cReq)
		break

	case "gameOver":
		GameOver(cReq)
		break

	}

}
