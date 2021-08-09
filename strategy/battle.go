package strategy

import (
	"github.com/fangker/yu-gi-boy/yglib"
	"github.com/go-vgo/robotgo"
)

func checkHandState(game *yglib.YugiGame) {
	//game.BitMapManager.GameState
	// 设置取样点
	//_checkHandState()
}

func _checkHandState(cBitmap robotgo.CBitmap) {
	robotgo.GetColors(robotgo.ToMMBitmapRef(cBitmap), 0, 0)
}
