package yglib

import (
	"errors"
	"fmt"
	"github.com/fangker/yu-gi-boy/bm"
	"github.com/fatih/color"
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	log "github.com/sirupsen/logrus"
	"syscall"
	"time"
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
	window        win.HWND
	BitMapManager *bm.BitMapManager
	pos           Pos
	pid           int32
	state         int
	scale         int
}

func (ygg YugiGame) GetPos() Pos {
	var rect win.RECT
	win.GetWindowRect(ygg.window, &rect)
	ygg.pos.x = int(rect.Left)
	ygg.pos.y = int(rect.Top)
	return ygg.pos
}
func NewYugiGame(bmManager *bm.BitMapManager) (ygg *YugiGame, err error) {
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
	ygg.BitMapManager = bmManager
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
	robotgo.ActivePID(ygg.pid)
	// if set scaling get scale for moveMouse
	syscall.NewLazyDLL("user32.dll").NewProc("SetProcessDPIAware").Call()
	ygg.scale = robotgo.Scale()

	YuGiGameEntry = ygg
	return
}
func (ygg *YugiGame) MoveAndClick(x int, y int) {
	robotgo.MoveClick(x, y, "left")
}
func (ygg *YugiGame) MovesAndClick(x int, y int) {
	robotgo.MoveMouseSmooth(x, y, "left")
}
func (ygg *YugiGame) IsBitMapExist(sbe bm.SrcBitMapEntity, args ...interface{}) bool {
	var (
		tolerance = 0.05
	)
	if len(args) > 0 && args[0] != nil {
		time.Sleep(args[0].(time.Duration))
	}
	if len(args) > 1 {
		tolerance = args[1].(float64)
	}
	x, y := robotgo.FindCBitmap(sbe.CBitmap, nil, tolerance)
	return !(x == -1 && y == -1)
}
func (ygg *YugiGame) FindBitMapAndMoveClick(sbe bm.SrcBitMapEntity, args ...interface{}) {
	var (
		tolerance = 0.05
	)
	if len(args) > 0 && args[0] != nil {
		time.Sleep(args[0].(time.Duration))
	}
	if len(args) > 0 && args[1] != nil {
		tolerance = args[1].(float64)
	}
	x, y := robotgo.FindCBitmap(sbe.CBitmap, nil, tolerance)
	log.Debugf("find x: %d y: %d", x, y)
	if x == -1 || y == -1 {
		return
	}
	bmp := robotgo.ToBitmap(robotgo.ToMMBitmapRef(sbe.CBitmap))
	log.Debugf("click x: %d y: %d", x+bmp.Width/2, y+bmp.Height/2)
	ygg.MoveAndClick(x+bmp.Width/2, y+bmp.Height/2)
}

func (ygg *YugiGame) FindBitMapLookForward(sbe bm.SrcBitMapEntity, checkInterval time.Duration, deadTime time.Duration, args ...interface{}) bool {
	var (
		tolerance = 0.05
	)
	if len(args) > 0 && args[0] != nil {
		tolerance = args[0].(float64)
	}
	trigger := time.NewTicker(checkInterval)
	for {
		select {
		case <-trigger.C:
			fmt.Println("xxxx", checkInterval, tolerance)
			x, y := robotgo.FindCBitmap(sbe.CBitmap, nil, tolerance)
			if x != -1 && y != -1 {
				return true
			}
		case <-time.After(deadTime):
			return false
		}
	}
}
