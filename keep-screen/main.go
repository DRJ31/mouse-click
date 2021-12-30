package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Timeout int `json:"timeout"`
}

func readConfig(path string) Config {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)

	var result Config
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func movePos(pos, flag int) int {
	if flag % 2 == 1 {
		return pos + 1
	} else {
		return pos - 1
	}
}

func main() {
	config := readConfig("config.json")
	x0, y0 := robotgo.GetMousePos()
	flag := 1

	for {
		x1, y1 := robotgo.GetMousePos()
		if x0 == x1 && y0 == y1 {
			x0, y0 = movePos(x1, flag), movePos(y1, flag)
			robotgo.Move(x0, y0)
			fmt.Printf("Moved: (%v, %v)\n", x0, y0)
			flag = flag % 2 + 1
		}
		x0, y0 = robotgo.GetMousePos()
		time.Sleep(time.Duration(config.Timeout) * time.Second)
	}
}
