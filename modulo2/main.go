package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"io/ioutil"
	"strings"
)

type Welcome struct {
	Name string
	Time string
}

type ramStruct struct {
	Total      int `json:"total"`
	Libre      int  `json:"libre"`
	Porcentaje int `json:"porcentaje"`
	Consumo int `json:"consumo"`
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("static/index.html"))
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	http.HandleFunc("/memoria", getMemInfo)
	http.ListenAndServe(":8080", nil)

}

func getMemInfo(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	line := lines[0]
	total := strings.Replace(string(line)[10:24], " ", "", -1)
	fmt.Println("Total de RAM: " + total)

	line2 := lines[1]
	libre := strings.Replace(string(line2)[10:24], " ", "", -1)
	fmt.Println("RAM Libre: " + libre)

	ramlTotalKb, err1 := strconv.Atoi(total)
	ramLibreKb, err2 := strconv.Atoi(libre)

	if err1 == nil && err2 == nil {

		ramlTotalMb := ramlTotalKb / 1024
		ramLibreMb := ramLibreKb / 1024
		consumo := ramlTotalMb - ramLibreMb
		porcentaje := consumo * 100 / ramlTotalMb
		fmt.Println("RAM usada: ", porcentaje, "%")

		ramObj := &ramStruct{ramlTotalMb, ramLibreMb, porcentaje, consumo}
		jsonResponse, errorjson := json.Marshal(ramObj)
		if errorjson != nil {
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))

	}
}

