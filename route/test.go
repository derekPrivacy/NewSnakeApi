package route

import (
	"encoding/json"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	response, _ := json.Marshal("test ok")
	w.Write(response)

	log.Printf("called")

}
