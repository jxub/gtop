package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

// a top-like program written in golang

func main() {
	GetAllPids()
	http.HandleFunc("/", GtopHandler)
}

func GtopHandler(w http.ResponseWriter, r *http.Request) {
	g := new(Gtop)
	json.NewEncoder(w).Encode(&g)
	return
}

type Gtop struct {
	Processes []Process
}

type Process struct {
	pid     int
	user    string
	pr      int
	ni      int
	virt    int
	res     int
	shr     int
	s       string
	cpuPct  float32
	memPct  float32
	hourSnc int
	order   string
}

func GetAllPids() {
	all, err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	dirs := make([]int, 20)
	for _, f := range all {
		if f.IsDir() {
			dir, _ := strconv.Atoi(f.Name())
			if err == nil {
				dirs = append(dirs, dir)
			}
		}
	}
	stuff := make(map[int][]string)
	for dir := range dirs {
		path := "/proc/" + strconv.Itoa(dir) + "/status"
		cmd := exec.Command("cat", path)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			l := scanner.Text()
			// fmt.Println(l)
			stuff[dir] = append(stuff[dir], l)
		}
	}
	fmt.Println(stuff[22])

}
