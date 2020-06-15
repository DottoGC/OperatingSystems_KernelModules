package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

)

type Welcome struct {
	Name string
	Time string
}

type ProcArray struct {
	Procesos []Proc
}

type Proc struct {
	Pid string
	Nombre string
	Usuario string
	Estado string
	Porcentaje string
}

type ProcChild struct {
	Name string
	Time string
}

type ramStruct struct {
	Total      float64 `json:"total"`
	Libre      float64  `json:"libre"`
	Porcentaje float64 `json:"porcentaje"`
	Consumo float64 `json:"consumo"`
}

type cpuStruct struct {

	Porcentaje float64 `json:"porcentaje"`
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
	http.HandleFunc("/procs", getProcInfo)
	http.HandleFunc("/cpuPorcentaje", getCpuInfo)

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

//	ramlTotalKb, err1 := strconv.Atoi(total)
	ramLibreKb, err2 := strconv.Atoi(libre)

	if  err2 == nil {

		ramlTotalMb := 7784.8
		ramLibreMb := float64(ramLibreKb) / 1024
		consumo := ramlTotalMb - ramLibreMb
		porcentaje := float64(consumo) * 100 / float64(ramlTotalMb)
		fmt.Println("RAM usada: ", porcentaje, "%")

		ramObj := &ramStruct{ramlTotalMb, math.Round(ramLibreMb*100)/100, math.Round(porcentaje*100)/100, math.Round(consumo*100)/100}
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

func getCpuInfo(w http.ResponseWriter, r *http.Request) {
	var prevIdleTime, prevTotalTime uint64
	var cpuUsage = 0.0
	for i := 0; i < 4; i++ {
		file, err := os.Open("/proc/stat")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		firstLine := scanner.Text()[5:] // get rid of cpu plus 2 spaces
		file.Close()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		split := strings.Fields(firstLine)
		idleTime, _ := strconv.ParseUint(split[3], 10, 64)
		totalTime := uint64(0)
		for _, s := range split {
			u, _ := strconv.ParseUint(s, 10, 64)
			totalTime += u
		}
		if i > 0 {
			deltaIdleTime := idleTime - prevIdleTime
			deltaTotalTime := totalTime - prevTotalTime
			cpuUsage = (1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)) * 100.0
			fmt.Printf("%d : %6.3f\n", i, cpuUsage)
		}

		prevIdleTime = idleTime
		prevTotalTime = totalTime
		time.Sleep(time.Second)
	}


		cpuObj := &cpuStruct{math.Round(cpuUsage*100)/100 }
		jsonResponse, errorjson := json.Marshal(cpuObj)
		if errorjson != nil {
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))


		/*contents, err := ioutil.ReadFile("/proc/stat")
		if err != nil {
			return
		}
		var total = 0
		var idle = 0

		lines := strings.Split(string(contents), "\n")
		for j := 0; j < 5; j++ {
			line := lines[j]
			fmt.Println("Entro aqui 1  ")

			fields := strings.Fields(line)
			if fields[0] == "cpu" {
				numFields := len(fields)
				for i := 1; i < numFields; i++ {

					val, err := strconv.Atoi(fields[i])
					fmt.Println("Entro aqui 2 ")
					if err != nil {
						fmt.Println("Error: ", i, fields[i], err)
					}
					total += val // tally up all the numbers to get total ticks
					if i == 4 {  // idle is the 5th field in the cpu line
						idle = val
					}
				}
			}
		}

		porcentaje := ( total - idle ) / total*/



}


func getProcInfo(w http.ResponseWriter, r *http.Request){

	var procesos = getLinuxProcesses();
	//arrayP := ProcArray{procesos}
	jsonResponse, errorjson := json.Marshal(procesos)
	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResponse))
}


func getLinuxProcesses() []Proc {
	var linuxProcesses []Proc
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	var numberProcs []os.FileInfo
	for _, f := range files {
		_, err := strconv.Atoi(f.Name())
		if err == nil {
			numberProcs = append(numberProcs, f)
		}
	}
	numericFileInfos := numberProcs

	for _, f := range numericFileInfos {

		file, err := os.Open("/proc/" + f.Name() + "/status")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fileContentBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fileContent := fmt.Sprintf("%s", fileContentBytes)

		processInfo := extractLinuxProcessInfo(f.Name(), fileContent)

		linuxProcesses = append(linuxProcesses, processInfo)

	//	linuxProcesses = append(linuxProcesses, processInfo)
		//fmt.Println("asdf")
		//fmt.Printf("%s", fileContent)

		//break
	}

	return linuxProcesses
}

func extractLinuxProcessInfo(pid string, content string) Proc{


	var line = 0
	var saveChars = false
	var value = ""
	var nombre = "";
	var estado = "";
	//var uid = "";
	var porcent = "";


	for _, c := range content {
		if saveChars && c != '\n' {
			value += string(c)
		}
		if c == ':' {
			saveChars = true
		}
		if c == '\n' {

			switch line {
			case 0:
				nombre = strings.TrimSpace(value)
				break
			case 2:
				estado = strings.TrimSpace(value)
				break
			case 8:
				//uid = strings.TrimSpace(value)
				break
			case 28:
				cadena := strings.TrimSpace(value);
				val := strings.Replace(cadena, " kB", "", 1)
				ram, err :=  strconv.ParseFloat(val, 64)
				if err != nil {
					return Proc{}
				}
				ramMb := ram / 1024
				var porcentaje = (ramMb / 7862)*100
				porcent = fmt.Sprintf("%.2f", porcentaje)
				break
			}

			line += 1
			saveChars = false
			value = ""
		}
	}


	return Proc{pid, nombre, "kenia", estado,porcent};
}


func killProcess(pid int) error {
	process, err := os.FindProcess(pid);
	if err != nil {
		return err
	}

	err = process.Signal(syscall.Signal(0)) // if nil then is ok to kill

	if err != nil {
		return err
	}

	err = process.Kill()

	if err != nil {
		return err
	}

	return nil
}

