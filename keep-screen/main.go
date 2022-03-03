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
	X1      int `json:"x1"`
	X2      int `json:"x2"`
	Y1      int `json:"y1"`
	Y2      int `json:"y2"`
	Mode    int `json:"mode"`
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
	if flag%2 == 1 {
		return pos + 1
	} else {
		return pos - 1
	}
}

func keepScreen(config Config) {
	x0, y0 := robotgo.GetMousePos()
	flag := 1

	for {
		x1, y1 := robotgo.GetMousePos()
		if x0 == x1 && y0 == y1 {
			x0, y0 = movePos(x1, flag), movePos(y1, flag)
			robotgo.Move(x0, y0)
			fmt.Printf("Moved: (%v, %v)\n", x0, y0)
			flag = flag%2 + 1
		}
		x0, y0 = robotgo.GetMousePos()
		time.Sleep(time.Duration(config.Timeout) * time.Second)
	}
}

func keepClick(config Config) {

	for {
		robotgo.Move(config.X1, config.Y1)
		robotgo.Click()
		time.Sleep(time.Duration(config.Timeout) * time.Second)
		robotgo.Move(config.X2, config.Y2)
		robotgo.Click()
		time.Sleep(time.Duration(config.Timeout) * time.Second)
	}
}

func main() {
	config := readConfig("config.json")

	if config.Mode == 1 {
		keepScreen(config)
	} else if config.Mode == 2 {
		keepClick(config)
	}

	fmt.Println("End")
}
