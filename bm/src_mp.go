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
	gameStateSrcMap[constants.GAME_STATE_NOTIFICATION] = "game_state/notification.png"
	gameStateSrcMap[constants.GAME_STATE_NOTIFICATION_QUIT] = "game_state/notification_quit.png"
	gameStateSrcMap[constants.GAME_STATE_DIALOG_CAMPAIGN_QUIT] = "game_state/dialog_campaign_quit.png"
	// home
	gameStateSrcMap[constants.GAME_STATE_HOME_PAGE] = "game_state/home_page.png"
	gameStateSrcMap[constants.GAME_STATE_MENU_BOTTOM_PVP_CLICK] = "game_state/menu_bottom_pvp_click.png"
	gameStateSrcMap[constants.GAME_STATE_MENU_BOTTOM_PVP_UNCLICK] = "game_state/menu_bottom_pvp_unclick.png"
	gameStateSrcMap[constants.GAME_STATE_MENU_PVP_CASUAL] = "game_state/menu_pvp_casual.png"
	gameStateSrcMap[constants.GAME_STATE_MENU_PVP_DUEL] = "game_state/menu_pvp_duel.png"
	gameStateSrcMap[constants.GAME_STATE_PVP_READY_PAGE] = "game_state/pvp_ready_page.png"
	// battle
	battleSrcMap[constants.BATTLE_PAGE] = "battle/battle_page.png"
	sfm := srcFileMap{}
	sfm.gameState = gameStateSrcMap
	sfm.battle = battleSrcMap
	return sfm
}
