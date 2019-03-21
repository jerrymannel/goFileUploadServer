package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getHandler(_res http.ResponseWriter, _req *http.Request) {
	log.Println("*** NEW REQUEST AT /***")
	log.Println(_req.URL)
	log.Println(_req.Method)
	for i, h := range _req.Header {
		log.Printf("%s \t:: %s", i, h[0])
		_res.Header().Add(i, h[0])
	}
	data, _ := ioutil.ReadAll(_req.Body)
	log.Println(string(data))
	_res.Write(data)
	log.Println("------------------------------")
}

func uploadHandler(_res http.ResponseWriter, _req *http.Request) {
	log.Println("*** NEW REQUEST UPLOADHANDLER 1 ***")
	log.Println(_req.URL)
	log.Println(_req.Method)
	for i, h := range _req.Header {
		log.Printf("%s \t:: %s", i, h[0])
	}
	f, err := os.OpenFile("./uploads/"+_req.Header["Odp-File-Name"][0], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(_req.Body)
	f.Write(data)
	log.Println("File saved under - " + "./uploads/" + _req.Header["Odp-File-Name"][0])
	_res.Header().Add("Content-type", "application/octet-stream")
	log.Println("Content-type : " + _res.Header().Get("Content-type"))
	log.Println("Sending data contents of the file :" + f.Name())
	_, err = _res.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("------------------------------")
}

func uploadHandler2(_res http.ResponseWriter, _req *http.Request) {
	log.Println("*** NEW REQUEST UPLOADHANDLER 2 ***")
	log.Println(_req.URL)
	log.Println(_req.Method)
	for i, h := range _req.Header {
		log.Printf("%s \t:: %s", i, h[0])
	}
	f, err := os.OpenFile("./uploads2/"+_req.Header["Odp-File-Name"][0], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(_req.Body)
	f.Write(data)
	log.Println("File saved under - " + "./uploads2/" + _req.Header["Odp-File-Name"][0])
	_res.Header().Add("Content-type", "application/json")
	log.Println("Content-type : " + _res.Header().Get("Content-type"))
	log.Println("Sending data : " + f.Name())
	_, err = _res.Write([]byte(`{"fileName": ` + _req.Header["Odp-File-Name"][0] + `}`))
	log.Println("Sending data : " + `{"fileName": ` + _req.Header["Odp-File-Name"][0] + `}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("------------------------------")
}

func echoHandler(_res http.ResponseWriter, _req *http.Request) {
	log.Println("*** ECHO INVOKED ***")
	log.Println(_req.URL)
	log.Println(_req.RemoteAddr)
	log.Println(_req.Method)
	for i, h := range _req.Header {
		log.Printf("%s \t:: %s", i, h[0])
	}
	data, _ := ioutil.ReadAll(_req.Body)
	log.Println(string(data))
	_res.Write(data)
	log.Println("------------------------------")
}

func errorHandler(_res http.ResponseWriter, _req *http.Request) {
	log.Println("*** ERROR INVOKED ***")
	log.Println(_req.URL)
	log.Println(_req.RemoteAddr)
	log.Println(_req.Method)
	for i, h := range _req.Header {
		log.Printf("%s \t:: %s", i, h[0])
	}
	data, _ := ioutil.ReadAll(_req.Body)
	log.Println(string(data))
	_res.WriteHeader(400)
	_res.Header().Add("Content-type", "application/json")
	_res.Write([]byte(`{"message": "Oops there was an error"}`))
	log.Println("------------------------------")
}

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/upload2", uploadHandler2)
	err := http.ListenAndServe(":32100", nil)
	if err != nil {
		panic(err)
	}
}
