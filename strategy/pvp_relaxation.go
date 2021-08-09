package strategy

import (
	"github.com/fangker/yu-gi-boy/constants"
	"github.com/fangker/yu-gi-boy/yglib"
	log "github.com/sirupsen/logrus"
	"time"
)

type PVPRelaxationStrategy struct {
}

func UsePvPRelaxation(game *yglib.YugiGame) {
	gbm := game.BitMapManager
	//tryToMainPage(game)
	turnToPvpPage(game)
	game.FindBitMapAndMoveClick(gbm.GameState[constants.GAME_STATE_MENU_PVP_DUEL], 3*time.Second, 0.3)
	ok := game.FindBitMapLookForward(gbm.Battle[constants.BATTLE_PAGE], time.Second, time.Second*35)
	log.Info("ok===", ok)
}
