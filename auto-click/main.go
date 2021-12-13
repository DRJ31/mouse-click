package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Sleep       int      `json:"sleep"`
	Offset      int      `json:"offset"`
	StartPos    Position `json:"start_pos"`
	PicPos      Picture  `json:"pic_pos"`
	TagPos      Position `json:"tag_pos"`
	BtnPos      Position `json:"btn_pos"`
	RandomRange Range    `json:"random_range"`
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

func calculatePos(ratioX, ratioY float64, size Picture) Position {
	var position Position
	position.X = size.X1 + int(math.Floor(float64(size.X2-size.X1)*ratioX))
	position.Y = size.Y1 + int(math.Floor(float64(size.Y2-size.Y1)*ratioY))
	return position
}

func main() {
	config := readConfig("config.json")
	ran := config.RandomRange
	var times int

	fmt.Print("请输入循环次数：")
	_, err := fmt.Scanln(&times)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sleep for %d seconds.", config.Sleep)
	robotgo.Sleep(config.Sleep)
	//robotgo.MoveClick(config.StartPos.X, config.StartPos.Y)

	for i := 0; i < times; i++ {
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y)
		// Person
		pos := calculatePos(0.3, 0.2, config.PicPos)
		robotgo.MoveClick(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		pos = calculatePos(0.6, 0.7, config.PicPos)
		robotgo.Move(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		robotgo.Click("left")
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y)

		// Helmet
		pos = calculatePos(0.31, 0.21, config.PicPos)
		robotgo.MoveClick(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		pos = calculatePos(0.59, 0.3, config.PicPos)
		robotgo.Move(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		robotgo.Click("left")
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y+config.Offset)

		// Clothes
		pos = calculatePos(0.31, 0.31, config.PicPos)
		robotgo.MoveClick(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		pos = calculatePos(0.59, 0.69, config.PicPos)
		robotgo.Move(pos.X+randInt(ran.Min, ran.Max), pos.Y+randInt(ran.Min, ran.Max))
		robotgo.Click("left")
		robotgo.MoveClick(config.TagPos.X, config.TagPos.Y+config.Offset*2)

		robotgo.MoveClick(config.BtnPos.X, config.BtnPos.Y)
		fmt.Printf("Loop %d finish, sleep for %d seconds.\n", i+1, config.Sleep)
		robotgo.Sleep(config.Sleep)
	}
}
