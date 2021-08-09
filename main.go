package main

import (
	"fmt"
	"github.com/fangker/yu-gi-boy/bm"
	"github.com/fangker/yu-gi-boy/config"
	"github.com/fangker/yu-gi-boy/strategy"
	"github.com/fangker/yu-gi-boy/yglib"
	"github.com/fatih/color"
	"github.com/go-vgo/robotgo"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var green = color.New(color.FgGreen).PrintlnFunc()
var fgreen = color.New(color.FgGreen).PrintfFunc()
var red = color.New(color.FgHiRed).PrintlnFunc()

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	// config
	yggConfig := config.LoadConfig("./config.yaml")
	// set logger
	initLogger()
	// load bm
	bitManager := bm.NewBitMapManager(yggConfig.SrcPath)
	// init ygg
	ygg, err := yglib.NewYugiGame(bitManager)
	if err != nil {
		red(err.Error())
	}
	time.Sleep(3 * time.Second)
	log.Info("OK,ssss ssss")
	strategy.UsePvPRelaxation(ygg)
	log.Info("OK,ssssssss")
	timeTicker := time.NewTicker(1 * time.Second)
	for {
		<-timeTicker.C
		//robotgo.MoveClick(561, 424, "left")
	}
	select {}
}

func findGameWindow() {
}
func initLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
