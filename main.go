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

var SupportV = map[string][]int{
	"v1": []int{2, 3, 4, 7, 8, 9, 10, 11, 14, 15, 17, 18, 22, 23, 24, 25, 27},
	"v2": []int{3, 5, 7, 11, 13, 15, 19, 21, 23, 29, 31, 33, 35, 37, 8, 10, 12, 16, 18, 22, 24, 26, 32, 36, 38, 40},
	"v3": []int{3, 5, 7, 11, 13, 15, 19, 21, 23, 29, 31, 33, 35, 37, 8, 10, 12, 16, 18, 22, 24, 26, 32, 36, 38, 40},
}
var OffResult string = "light set Off"
var OnResult string = "light set On"

func check(key string, pin int) bool {
	keys := reflect.ValueOf(SupportV).MapKeys()

	for _, value := range keys {
		if key == reflect.Value(value).String() {
			for _, value1 := range SupportV[key] {
				if pin == value1 {
					return true
				}
			}
		}
	}
	return false
}

// On function use for feed the corrent.
func On(w http.ResponseWriter, req *http.Request) {
	paramPin := req.URL.Query().Get("pin")
	paramVersion := req.URL.Query().Get("ver")
	valuePin, error := strconv.Atoi(paramPin)
	if error != nil {
		os.Exit(1)
	}
	if check(paramVersion, valuePin) {
		var pin = rpio.Pin(valuePin)
		pin.Output()
		pin.High()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, OnResult)
		fmt.Println(OnResult)

	} else {
		fmt.Fprintf(w, "error")
	}

}

func Test(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is working")

}

// Off function use for block the corrent.
func Off(w http.ResponseWriter, req *http.Request) {
	paramPin := req.URL.Query().Get("pin")
	paramVersion := req.URL.Query().Get("ver")
	valuePin, error := strconv.Atoi(paramPin)
	if error != nil {
		os.Exit(1)
	}
	if check(paramVersion, valuePin) {
		var pin = rpio.Pin(valuePin)
		pin.Output()
		pin.Low()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, OffResult)
		fmt.Println(OffResult)
	} else {
		fmt.Fprintf(w, "error")
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
	router.HandleFunc("/on/", On).Methods("GET")
	router.HandleFunc("/off/", Off).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
