package route

import (
	db "api/db"
	"api/model"
	wsModel "api/model/socket"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func ValidateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	userName := gjson.Get(
		string(body),
		"user",
	)

	password := gjson.Get(
		string(body),
		"password",
	)

	var user model.GoUser

	db.DB.Where("user_name = ?", userName.String()).First(&user)

	if checkDoubleLogin(userName.String()) == true {
		response, _ := json.Marshal("doubleLogin")
		w.Write(response)
	} else if userName.String() != "" && password.String() != "" && user.Password == password.String() {
		response, _ := json.Marshal("match")
		w.Write(response)
	} else {
		response, _ := json.Marshal("unmatch")
		w.Write(response)
	}

}

func checkDoubleLogin(userName string) bool {

	doubleLogin := false

	if len(wsModel.RoomArr) != 0 && len(wsModel.RoomArr[0].Data) != 0 {
		for _, data := range wsModel.RoomArr[0].Data {
			if userName == data {
				doubleLogin = true
			}
			break
		}
	}

	if doubleLogin {
		log.Printf("yes")
		return true
	} else {
		log.Printf("no")
		return false
	}

}
