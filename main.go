package main

import (
	"errors"
	"fmt"
	"github.com/fangker/yu-gi-boy/bm"
	"github.com/fangker/yu-gi-boy/config"
	"github.com/fatih/color"
	"github.com/go-vgo/robotgo"
	win "github.com/lxn/win"
	"syscall"
)

var green = color.New(color.FgGreen).PrintlnFunc()
var fgreen = color.New(color.FgGreen).PrintfFunc()
var red = color.New(color.FgHiRed).PrintlnFunc()

const GAME_WINDOW = "Yu-Gi-Oh! DUEL LINKS"
const GAME_PROCCESS = "dlpc"

var YuGiGameEntry *YugiGame

type Pos struct {
	x int
	y int
}

type YugiGame struct {
	window win.HWND
	pos    Pos
	pid    int32
	state  int
}

func (ygg YugiGame) GetPos() Pos {
	var rect win.RECT
	win.GetWindowRect(ygg.window, &rect)
	ygg.pos.x = int(rect.Left)
	ygg.pos.y = int(rect.Top)
	return ygg.pos
}
func NewYugiGame() (ygg *YugiGame, err error) {
	green("检查游戏主体是否打开")
	wp, _ := syscall.UTF16PtrFromString(GAME_WINDOW)
	hwnd := win.FindWindow(nil, wp)
	if hwnd == 0 {
		return nil, errors.New("游戏未打开")
	}

	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)
	posionX := int(rect.Left)
	posionY := int(rect.Top)
	hwndWidth := int(rect.Right - rect.Left)
	hwndHeight := int(rect.Bottom - rect.Top)
	fgreen("检测窗口位置 Px: %d Py: %d width: %d height:%d \n", posionX, posionY, hwndWidth, hwndHeight)

	ygg = &YugiGame{}
	ygg.window = hwnd
	ygg.pos = Pos{x: posionX, y: posionY}
	fpid, err := robotgo.FindIds(GAME_PROCCESS)
	if fpid[0] == 0 {
		return nil, errors.New("未检测到pid")
	}
	fgreen("检测到pid %d 激活窗口 \n", fpid[0])

	ygg.pid = fpid[0]
	la := robotgo.GetActive() // getting C.MData of active window
	robotgo.SetActive(la)     // Trying to set window active again with it's C.MData

	YuGiGameEntry = ygg
	return
}
func findAndCalculateRelativePos() {

}

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	_, err := NewYugiGame()
	if err != nil {
		red(err.Error())
	}
	config := config.LoadConfig("./config.yaml")
	bm.NewBitMapManager(config.SrcPath)
	select {}
}

func findGameWindow() {
}
