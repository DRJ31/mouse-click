package main

import (
	"encoding/json"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Offset int      `json:"offset"`
	TagPos Position `json:"tag_pos"`
	BtnPos Position `json:"btn_pos"`
}

type Picture struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
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

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

func calculatePos(ratio float64, size Picture) Position {
	var position Position
	if ratio > 0 {
		if math.Abs(float64(size.Y2-size.Y1)) > math.Abs(float64(size.X2-size.X1)) {
			position.X = size.X2
			position.Y = size.Y1 + int(math.Floor(float64(size.Y2-size.Y1)*ratio))
		} else {
			position.X = size.X1 + int(math.Floor(float64(size.X2-size.X1)*ratio))
			position.Y = size.Y2
		}
	} else {
		if math.Abs(float64(size.Y2-size.Y1)) > math.Abs(float64(size.X2-size.X1)) {
			position.X = size.X2
			position.Y = size.Y2
		} else {
			position.X = size.X2 + int(math.Floor(float64(size.X2-size.X1)*ratio))
			position.Y = size.Y2
		}
	}
	return position
}

func main() {
	config := readConfig("config.json")

	for {
		var x1, y1, x2, y2 int
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y)

		robotgo.EventHook(hook.KeyDown, []string{"c"}, func(e hook.Event) {
			robotgo.Click()
			x1, y1 = robotgo.GetMousePos()
		})
		robotgo.EventHook(hook.KeyDown, []string{"z"}, func(e hook.Event) {
			robotgo.Click()
			x2, y2 = robotgo.GetMousePos()
			robotgo.MoveClick(config.TagPos.X, config.TagPos.Y)
			robotgo.EventEnd()
		})
		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)

		// Helmet
		pos := Position{
			X: x1,
			Y: y1,
		}
		robotgo.MoveClick(pos.X, pos.Y)
		pos = calculatePos(0.25, Picture{
			X1: x1,
			Y1: y1,
			X2: x2,
			Y2: y2,
		})
		robotgo.Move(pos.X, pos.Y)
		robotgo.Click("left")
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y+config.Offset)

		// Clothes
		if math.Abs(float64(y2-y1)) > math.Abs(float64(x2-x1)) {
			pos.X = x1
			pos.Y = y1 + int(math.Floor(float64(y2-y1)*0.25))
		} else {
			pos.X = x2
			pos.Y = y1
		}
		robotgo.MoveClick(pos.X, pos.Y)
		pos = calculatePos(-0.75, Picture{
			X1: x1,
			Y1: y1,
			X2: x2,
			Y2: y2,
		})
		robotgo.Move(pos.X, pos.Y)
		robotgo.Click("left")
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y+config.Offset*2)

		robotgo.MoveClick(config.BtnPos.X, config.BtnPos.Y)
	}
}
