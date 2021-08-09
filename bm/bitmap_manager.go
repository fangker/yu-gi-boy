package bm

import (
	"github.com/go-vgo/robotgo"
	"log"
	"os"
	"path"
	"path/filepath"
)

type SrcBitMapEntity struct {
	CBitmap    robotgo.CBitmap
	ActionName string
	SrcPath    string
}
type srcBitMapType map[string]SrcBitMapEntity
type BitMapManager struct {
	srcPath   string
	Battle    srcBitMapType
	GameState srcBitMapType
}

var BitMapEntry *BitMapManager

func NewBitMapManager(srcPath string) *BitMapManager {
	if srcPath == "" {
		srcPath = path.Join(getCurrentPath(), "./src")
	}
	BitMapEntry = &BitMapManager{srcPath: srcPath}
	BitMapEntry.loadBitmapBySrcMap()
	return BitMapEntry
}

func (bm *BitMapManager) loadBitmapBySrcMap() {
	srcPath := bm.srcPath
	// load fileName map
	sfm := loadSrcFileMap()
	// load bitmap
	bm.GameState = loadBitmapBySrcMap(srcPath, sfm.gameState)
	bm.Battle = loadBitmapBySrcMap(srcPath, sfm.battle)
}
func loadBitmapBySrcMap(srcRootPath string, mpc map[string]string) srcBitMapType {
	res := make(srcBitMapType)
	for k, v := range mpc {
		bit := robotgo.OpenBitmap(path.Join(srcRootPath, filepath.FromSlash(v)))
		res[k] = SrcBitMapEntity{CBitmap: robotgo.CBitmap(bit), ActionName: k, SrcPath: path.Join(srcRootPath, filepath.FromSlash(v))}
	}
	return res
}

func getCurrentPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}
