package route

import (
	wsModel "api/model/socket"
	"log"

	"github.com/gorilla/websocket"
)

func GameFinish(cReq *wsModel.CReq, conn *websocket.Conn) *wsModel.CRes {

	var cRes *wsModel.CRes = &wsModel.CRes{}

	// req:{roomId, positionX,positionY}
	//check request room id
	// reassign food

	var positionX = cReq.FoodPositionX
	var positionY = cReq.FoodPositionY
	log.Printf("%f", positionX)
	log.Printf("%f", positionY)

	var reqRoomId = cReq.RoomID
	log.Printf("%d", reqRoomId)

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {

			log.Println("foudn u ya")

			wsModel.RoomArr[index].Food = wsModel.Food{PositionX: positionX, PositionY: positionY}

			cRes = &wsModel.CRes{Room: wsModel.RoomArr[index]}
			break
		}
	}

	log.Println("updated cRes room id is")
	log.Println(cRes.Room.ID)
	log.Println(cRes.Room.Food.PositionX)
	log.Println(cRes.Room.Food.PositionY)

	return cRes
}
