package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
)
type dockType struct {
	FacilityId string
	Name string
}

func CreateDockTypeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love coding %s!", r.URL.Path[1:])
	decoder := json.NewDecoder(r.Body)

	var temp dockType
	err := decoder.Decode(&temp)

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, temp.FacilityId)
	fmt.Println(temp.FacilityId)
	fmt.Println(temp.Name)
}