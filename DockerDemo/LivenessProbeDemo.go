package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Liveness(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(struct {
		Keyword string `json:"keyword"`
	}{
		Keyword: "CUHK",
	})
	w.Header().Set("Content-Type", "application/json")
	_, _ = io.WriteString(w, string(res))
}

func main() {
	http.HandleFunc("/_ping", Liveness)
	err := http.ListenAndServe(":4128", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
