package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/master", get_master)
	http.HandleFunc("/slave", get_slave)

	fmt.Println(v)
	http.ListenAndServe(":80", nil)
}

func get_master(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://my-master-server/master")
	if err != nil {
		w.Write([]byte("/my-master-server" + err.Error()))
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("/my-master-server" + err.Error()))
		return
	}
	w.Write([]byte("l am my-master-server"))
}

func get_slave(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://my-slave-server/slave")
	if err != nil {
		w.Write([]byte("my-slave-server" + err.Error()))
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("my-slave-server" + err.Error()))
		return
	}
	w.Write([]byte("l am my-slave-server"))
}
