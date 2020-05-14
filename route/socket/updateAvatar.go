package route

import (
	wsModel "api/model/socket"
	"log"
	"sync"
)

var mutex sync.Mutex

func cleanUpUpdate(wg *sync.WaitGroup) {
	defer wg.Done()

	if r := recover(); r != nil {
		log.Println("recover in clean up: ", r)
	}

	mutex.Unlock()

}

func UpdateAvatar(cReq *wsModel.CReq, wg *sync.WaitGroup) {
	defer cleanUpUpdate(wg)

	mutex.Lock()

	// var cRes *wsModel.CRes = &wsModel.CRes{}

	// req:{roomId, avatarId, positionX ,positionY,direction,bodyLength}
	//check request room id
	// update the avatar in the avatar array

	var avatarId = cReq.AvatarID

	var positionX = cReq.PositionX
	var positionY = cReq.PositionY

	var direction = cReq.Direction
	log.Println("im updating avatar and the direction send is " + direction)

	var bodyLength = cReq.BodyLength

	var reqRoomId = cReq.RoomID

	var reqRoom wsModel.Room

	var avatarExist bool = false

	log.Println(len(wsModel.RoomArr))

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {
			log.Printf("found u ya ")
			for avatarIndex, avatar := range wsModel.RoomArr[index].Avatar {
				if avatar.ID == avatarId {
					wsModel.RoomArr[index].Avatar[avatarIndex].PositionX = positionX
					wsModel.RoomArr[index].Avatar[avatarIndex].PositionY = positionY
					wsModel.RoomArr[index].Avatar[avatarIndex].BodyLength = bodyLength
					wsModel.RoomArr[index].Avatar[avatarIndex].Direction = direction
					reqRoom = wsModel.RoomArr[index]

					avatarExist = true

					break
				}
			}

			if !avatarExist {
				log.Printf("didn't found u ")

				newAvatar := wsModel.Avatar{ID: avatarId, PositionX: positionX, PositionY: positionY, BodyLength: bodyLength, Direction: direction}

				wsModel.RoomArr[index].Avatar = append(wsModel.RoomArr[index].Avatar, newAvatar)
				reqRoom = wsModel.RoomArr[index]

			}

			break
		}
	}

	if len(reqRoom.Data) >= 2 {
		for _, c := range reqRoom.ConnectionArr {
			c.Connection.WriteJSON(reqRoom)
		}
	}

	// cRes = &wsModel.CRes{Room: reqRoom}

	// log.Println("updated cRes room id is")
	// log.Println(cRes.Room.ID)

	// for index, _ := range cRes.Room.Avatar {
	// 	log.Println(cRes.Room.Avatar[index].ID)
	// 	log.Println(cRes.Room.Avatar[index].PositionX)
	// 	log.Println(cRes.Room.Avatar[index].PositionY)
	// 	log.Println(cRes.Room.Avatar[index].BodyLength)
	// 	log.Println(cRes.Room.Avatar[index].Direction)
	// }

	// return cRes
}
