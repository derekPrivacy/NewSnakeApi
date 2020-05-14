package route

import (
	wsModel "api/model/socket"
	"log"
	"sync"
)

var mutex sync.Mutex

func cleanUp(wg *sync.WaitGroup) {
	defer wg.Done()
	if r := recover(); r != nil {
		log.Println("recover in clean up: ", r)
	}
	mutex.Unlock()
}

func SayHello(cReq *wsModel.CReq) {
	// defer cleanUp(wg)

	// // defer wg.Done()

	// mutex.Lock()
	// // defer mutex.Unlock()

	log.Printf("my g is inside")

	var cRes *wsModel.CRes = &wsModel.CRes{}

	var reqRoomId = cReq.RoomID

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {

			cRes = &wsModel.CRes{Room: wsModel.RoomArr[index]}
			break
		}
	}

	if len(cRes.Room.Data) >= 2 {
		for _, c := range cRes.Room.ConnectionArr {
			c.Connection.WriteJSON(cRes.Room)
		}
	}
}
