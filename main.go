package main

import "net/http"
import "fmt"
import "encoding/json"

func main() {
	http.HandleFunc ("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if !Auth(w, r)						{ return }
	if !AllowOnlyGET(w, r)		 { return }

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJson(w, SelectStudent(id))
		return
	}

	OutputJson(w, GetStudents())
}

func OutputJson(w http.ResponseWriter, o interface{}){
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}