package bm

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"os"
	"path"
	"path/filepath"
)

type BitMapManager struct {
	Battle    map[string]robotgo.CBitmap
	GameState map[string]robotgo.CBitmap
}

var BitMapEntry *BitMapManager

func NewBitMapManager(srcPath string) {
	if srcPath == "" {
		srcPath = path.Join(getCurrentPath(), "./src")
	}
	BitMapEntry = &BitMapManager{}
	loadSrcMap()
	fmt.Println(srcPath, gameStateSrcMap)
	// load bitmap
}

//func loadBitmap(mpc map[string]string) {
//	res := new(map[string]robotgo.CBitmap)
//	for k, v := range mpc {
//		res[k] = robotgo.OpenBitmap(path.Join())
//	}
//
//}

func getCurrentPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}
