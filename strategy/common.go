package strategy

import (
	"github.com/fangker/yu-gi-boy/constants"
	"github.com/fangker/yu-gi-boy/yglib"
	log "github.com/sirupsen/logrus"
	"time"
)

func tryToMainPage(game *yglib.YugiGame) bool {
	log.Info("尝试进入主页面")
	gbm := game.BitMapManager
	game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_DISCONNECT])
	game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_NOTIFICATION_QUIT])
	game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_NOTIFICATION_QUIT])
	game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_DIALOG_CAMPAIGN_QUIT])
	ok := game.IsBitMapExist(gbm.GameState[constants.GAME_STATE_HOME_PAGE])
	if ok {
		log.Info("进入主页成功")
	} else {
		log.Info("进入主页失败")
	}
	return ok
}

func turnToPvpPage(game *yglib.YugiGame) bool {
	gbm := game.BitMapManager
	var ok bool = true
	if !game.IsBitMapExist(gbm.GameState[constants.GAME_STATE_PVP_READY_PAGE], 1*time.Second) {
		game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_MENU_BOTTOM_PVP_UNCLICK], 3*time.Second)
		game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_MENU_BOTTOM_PVP_CLICK], 3*time.Second)
		ok = game.IsBitMapExist(gbm.GameState[constants.GAME_STATE_MENU_PVP_CASUAL], 3*time.Second)
	}
	if ok {
		log.Info("进入PVP成功")
	} else {
		log.Info("进入PVP失败")
	}
	return ok
}
