package route

import (
	db "api/db"
	"api/model"
	"encoding/json"
	"io/ioutil"
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

	if userName.String() != "" && password.String() != "" && user.Password == password.String() {
		response, _ := json.Marshal("match")
		w.Write(response)
	} else {
		response, _ := json.Marshal("unmatch")
		w.Write(response)
	}
}
