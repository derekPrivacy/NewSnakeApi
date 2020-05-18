package route

import (
	wsModel "api/model/socket"
	"fmt"
	"log"
	"strconv"
)

func SpawnFood(cReq *wsModel.CReq) {

	var reqRoomId = cReq.RoomID
	roomId := strconv.Itoa(reqRoomId)

	var reqRoom wsModel.Room

	var foodType = cReq.FoodType

	// food one
	var foodOneX = cReq.FoodOneX
	var foodOneY = cReq.FoodOneY

	xxOne := fmt.Sprintf("%f", foodOneX)
	yyOne := fmt.Sprintf("%f", foodOneY)

	log.Println("food route called " + foodType + " " + xxOne + " " + yyOne + " " + roomId)

	//

	// food two

	var foodTwoX = cReq.FoodTwoX
	var foodTwoY = cReq.FoodTwoY

	xxTwo := fmt.Sprintf("%f", foodTwoX)
	yyTwo := fmt.Sprintf("%f", foodTwoY)

	log.Println("food route called " + foodType + " " + xxTwo + " " + yyTwo + " " + roomId)

	//

	for index, room := range wsModel.RoomArr {
		if room.ID == reqRoomId {

			log.Println("food here")

			wsModel.RoomArr[index].Food = wsModel.Food{
				FoodType: foodType,
				FoodOneX: foodOneX, FoodOneY: foodOneY,
				FoodTwoX: foodTwoX, FoodTwoY: foodTwoY,
			}

			reqRoom = wsModel.RoomArr[index]
			break
		}
	}

	if len(reqRoom.ConnectionArr) >= 2 {
		for _, c := range reqRoom.ConnectionArr {
			obj := &FoodRes{Food: reqRoom.Food}

			c.Connection.WriteJSON(obj)
		}
	}

}

type FoodRes struct {
	Food wsModel.Food
}
