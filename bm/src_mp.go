package bm

import "github.com/fangker/yu-gi-boy/constants"

type srcFileMapType map[string]string

var battleSrcMap = make(srcFileMapType)
var gameStateSrcMap = make(srcFileMapType)

type srcFileMap struct {
	battle    srcFileMapType
	gameState srcFileMapType
}

func loadSrcFileMap() srcFileMap {
	// game state
	gameStateSrcMap[constants.GAME_STATE_DISCONNECT] = "game_state/disconnection.png"
	// battle
	sfm := srcFileMap{}
	sfm.gameState = gameStateSrcMap
	return sfm
}
