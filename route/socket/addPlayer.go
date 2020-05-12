package route

import (
	wsModel "api/model/socket"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func AddPlayer(cReq *wsModel.CReq, conn *websocket.Conn) {
	var cRes *wsModel.CRes = &wsModel.CRes{}

	connection := wsModel.Connection{Player: cReq.Data, Connection: conn}

	var reqRoomId = cReq.RoomID

	var reqRoom wsModel.Room

	var roomExist bool = false

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {
			wsModel.RoomArr[index].ConnectionArr = append(room.ConnectionArr, connection)

			wsModel.RoomArr[index].Data = append(room.Data, cReq.Data)

			roomExist = true
			reqRoom = wsModel.RoomArr[index]

			break
		}
	}

	if !roomExist {

		newRoom := wsModel.Room{ID: reqRoomId}
		newRoom.ConnectionArr = append(newRoom.ConnectionArr, connection)
		newRoom.Data = append(newRoom.Data, cReq.Data)
		wsModel.RoomArr = append(wsModel.RoomArr, newRoom)
		reqRoom = newRoom

	}

	cRes = &wsModel.CRes{Room: reqRoom}

	log.Println("our cRes room id is")
	log.Println(cRes.Room.ID)
	log.Println(cRes.Room.Data)

	res, _ := json.Marshal(cRes)
	log.Printf("whats in you cRes " + string(res))

	if len(cRes.Room.Data) >= 2 {

		for _, c := range cRes.Room.ConnectionArr {
			c.Connection.WriteJSON(cRes.Room)
		}
	}

}
