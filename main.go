package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"text/template"
)

type Weather struct {
	Status Condition `json:status`
}

type Condition struct {
	Water int `json:Water`
	Wind  int `json:Wind`
}

// func doEvery(d time.Duration, f func(time.Time)) {
// 	for x := range time.Tick(d) {
// 		f(x)

// 	}
// }

func GetCurrentStatus(w http.ResponseWriter, GetAll Weather) {

	if GetAll.Status.Wind <= 6 {
		fmt.Fprintln(w, "Wind :Aman")
	}
	if GetAll.Status.Wind >= 7 && GetAll.Status.Wind <= 15 {
		fmt.Fprintln(w, "Wind :Siaga")
	}
	if GetAll.Status.Wind > 15 {
		fmt.Fprintln(w, "Wind :Bahaya")
	}

	if GetAll.Status.Water <= 5 {
		fmt.Fprintln(w, "Water :Aman")
	}
	if GetAll.Status.Water >= 6 && GetAll.Status.Water <= 8 {
		fmt.Fprintln(w, "Water :Siaga")
	}
	if GetAll.Status.Water > 8 {
		fmt.Fprintln(w, "Water :Bahaya")
	}
}

// func GetCurrentStatus(t time.Time) {

// 	// encoder.Encode(string(GetStatusData))
// 	var (
// 		water     = rand.Intn(15)
// 		wind      = rand.Intn(15)
// 		CurrStats = " "
// 	)
// 	if water < 5 && wind < 6 {
// 		CurrStats = "Aman"
// 	} else if (water >= 6 && water <= 8) && (wind >= 7 && wind <= 15) {
// 		CurrStats = "Siaga"
// 	} else {
// 		CurrStats = "Bahaya"
// 	}

// 	// GetWaterAndWind := map[string]int{"Water": water, "Wind": wind}
// 	// GetStatus := map[string]string{"Status": CurrStats}
// 	GetAllData := Status{CurrStats, water, wind}

// 	GetData, _ := json.Marshal(GetAllData)
// 	// GetStatusData, _ := json.Marshal(GetStatus)
// 	fmt.Println(GetAllData)
// 	// fmt.Println(string(mapB))

// 	file, _ := os.OpenFile("big_encode.json", os.O_CREATE, os.ModePerm)
// 	defer file.Close()
// 	encoder := json.NewEncoder(file)
// 	encoder.Encode(string(GetData))
// }

func dataJson() {

	var (
		wind  = rand.Intn(100)
		water = rand.Intn(100)
	)
	data := Weather{
		Status: Condition{Wind: wind,
			Water: water},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	//write the file
	_ = ioutil.WriteFile("big_encode.json", file, 0644)
}

func foo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("big_encode.json")
	if err != nil {
		fmt.Print(err)
	}

	var get Weather

	err = json.Unmarshal(data, &get)
	if err != nil {
		fmt.Println("error:", err)
	}

	dataJson()

	fmt.Fprintln(w, "The Weather is :")

	fmt.Fprintln(w, "wind :", get.Status.Wind, "kmph")
	fmt.Fprintln(w, "water :", get.Status.Water, "m")
	GetCurrentStatus(w, get)
}

func main() {

	// var x = Status{}
	// fmt.Println(x)

	http.HandleFunc("/", foo)

	http.ListenAndServe(":3000", nil)
	// doEvery(1000*time.Millisecond, GetCurrentStatus)
}
