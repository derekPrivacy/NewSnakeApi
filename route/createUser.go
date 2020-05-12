package route

import (
	db "api/db"
	"api/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	log.Printf(string(body))

	userName := gjson.Get(
		string(body),
		"user",
	)

	password := gjson.Get(
		string(body),
		"password",
	)

	userObj := model.GoUser{UserName: userName.String(), Password: password.String()}

	row := db.DB.Create(&userObj)

	if row.Error == nil {
		response, _ := json.Marshal("created")
		w.Write(response)
	} else {
		response, _ := json.Marshal(row.Error)
		w.Write(response)
	}

}
