package bm

import "github.com/fangker/yu-gi-boy/constants"

var battleSrcMap = make(map[string]string)
var gameStateSrcMap = make(map[string]string)

func loadSrcMap() {
	// game state
	gameStateSrcMap[constants.GAME_STATE_START] = "game_state/start_game.bmp"
	// battle
}
