package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio"
)

var supportV = map[string][]int{
	"v1": []int{7, 11, 13, 15, 12, 16, 18, 22},
	"v2": []int{3, 5, 7, 11, 13, 15, 19, 21, 23, 29, 31, 33, 35, 37, 8, 10, 12, 16, 18, 22, 24, 26, 32, 36, 38, 40},
	"v3": []int{3, 5, 7, 11, 13, 15, 19, 21, 23, 29, 31, 33, 35, 37, 8, 10, 12, 16, 18, 22, 24, 26, 32, 36, 38, 40},
}

func check(key string, pin int) bool {
	keys := reflect.ValueOf(supportV).MapKeys()

	for _, value := range keys {
		if key == reflect.Value(value).String() {
			for _, value1 := range supportV[key] {
				if pin == value1 {
					return true
				}
			}
		}
	}
	return false
}
func On(w http.ResponseWriter, req *http.Request) {
	paramPin := req.URL.Query().Get("pin")
	paramVersion := req.URL.Query().Get("version")
	valuePin, error := strconv.Atoi(paramPin)
	if error == nil {
		if check(paramVersion, valuePin) {
			var pin = rpio.Pin(valuePin)
			pin.High()
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Fan Set to On")
		}
	}

}

func Test(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is working")

}

func Off(w http.ResponseWriter, req *http.Request) {
	paramPin := req.URL.Query().Get("pin")
	paramVersion := req.URL.Query().Get("version")
	valuePin, error := strconv.Atoi(paramPin)
	if error == nil {
		if check(paramVersion, valuePin) {
			var pin = rpio.Pin(valuePin)
			pin.Low()
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Fan Set to On")
		}
	}
}

func main() {

	if runtime.GOARCH != "arm" {
		fmt.Println("This program must run into a Raspberry Pi")
		os.Exit(1)
	}

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()
	router := mux.NewRouter()
	router.HandleFunc("/test/", Test).Methods("GET")
	router.HandleFunc("/fan-on/", On).Methods("GET")
	router.HandleFunc("/fan-off/", Off).Methods("GET")
	log.Fatal(http.ListenAndServe(":4044", router))
}