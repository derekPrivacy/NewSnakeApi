package route

import (
	wsModel "api/model/socket"
	"log"
)

func GameOver(cReq *wsModel.CReq) {

	log.Printf("my g game over")

	var reqRoomId = cReq.RoomID

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {

			// a := []string{"A", "B", "C", "D", "E"}
			// i := 2

			// wsModel.RoomArr

			copy(wsModel.RoomArr[index:], wsModel.RoomArr[index+1:])
			wsModel.RoomArr[len(wsModel.RoomArr)-1] = wsModel.Room{}
			wsModel.RoomArr = wsModel.RoomArr[:len(wsModel.RoomArr)-1]

			break
		}
	}
}
