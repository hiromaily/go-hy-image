package main

import (
	"encoding/json"
	"errors"
	"flag"
	"image"
	"image/color"
	//"image/draw"
	//"image/jpeg"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
)

type Config struct {
	BlockSize int     `json:"blockSize"`
	Rgba      []uint8 `json:"rgba"`
}

var (
	jsonPath = flag.String("j", "", "Json file path")
)

var usage = `Usage: %s [options...]
Options:
  -j  Json file path.

e.g.:
  gohy -j ./jsons/config.json
`

//var (
//	white color.Color = color.RGBA{255, 255, 255, 255}
//	black color.Color = color.RGBA{0, 0, 0, 255}
//	blue  color.Color = color.RGBA{0, 0, 255, 255}
//)

func init() {
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	if *jsonPath == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
}

func main() {
	jsonByte, err := loadJSONFile(*jsonPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to call loadJSONFile()")
		return
	}

	var conf Config
	err = json.Unmarshal(jsonByte, &conf)
	if err != nil {
		log.Error().Err(err).Msg("failed to call json.Unmarshal()")
		return
	}

	//
	createAllImage(&conf)
}

func loadJSONFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("nothing JSON file")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func createAllImage(conf *Config) {
	objWidth := conf.BlockSize * 21
	objHeight := conf.BlockSize * 21

	col := color.RGBA{conf.Rgba[0], conf.Rgba[1], conf.Rgba[2], conf.Rgba[3]}
	//alpha := color.RGBA{0, 0, 0, 0}

	img := image.NewRGBA(image.Rect(0, 0, objWidth, objHeight))

	var drawFlg bool
	for y := 0; y < objHeight; y++ {
		for x := 0; x < objWidth; x++ {
			switch {
			//outline
			case 0 <= y && y < (conf.BlockSize*1):
				drawFlg = true
				//fallthrough
			case (conf.BlockSize*20) <= y && y < (conf.BlockSize*21):
				drawFlg = true
			case 0 <= x && x < (conf.BlockSize*1):
				drawFlg = true
			case (conf.BlockSize*20) <= x && x < (conf.BlockSize*21):
				drawFlg = true
			//h
			case (conf.BlockSize*2) <= x && x < (conf.BlockSize*6) && (conf.BlockSize*2) <= y && y < (conf.BlockSize*19):
				drawFlg = true
			case (conf.BlockSize*6) <= x && x < (conf.BlockSize*10) && (conf.BlockSize*11) <= y && y < (conf.BlockSize*14):
				drawFlg = true
			case (conf.BlockSize*10) <= x && x < (conf.BlockSize*14) && (conf.BlockSize*11) <= y && y < (conf.BlockSize*19):
				drawFlg = true
			//y
			case (conf.BlockSize*7) <= x && x < (conf.BlockSize*11) && (conf.BlockSize*2) <= y && y < (conf.BlockSize*10):
				drawFlg = true
			case (conf.BlockSize*11) <= x && x < (conf.BlockSize*15) && (conf.BlockSize*7) <= y && y < (conf.BlockSize*10):
				drawFlg = true
			case (conf.BlockSize*15) <= x && x < (conf.BlockSize*19) && (conf.BlockSize*2) <= y && y < (conf.BlockSize*19):
				drawFlg = true
			default:
				drawFlg = false
			}

			if drawFlg {
				img.Set(x, y, col)
				//}else{
				//	img.Set(x, y, alpha)
			}
		}
	}

	// Save to out.png
	f, _ := os.OpenFile("hy.png", os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	png.Encode(f, img)
}
