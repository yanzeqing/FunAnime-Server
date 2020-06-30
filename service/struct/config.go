package serviceStruct

import (
	"encoding/json"
	"github.com/yanzeqing/FunAnime-Server/util/common"
)

type BasicRouter struct {
	Name string `json:"title"`
	Path string `json:"router"`
}

type CarouselInfo struct {
	Image    string `json:"image"`
	TrueLink string `json:"true_img"`
	VideoId  int64  `json:"video_id"`
}

type MainConfig struct {
	CarouselImg []*CarouselInfo `json:"carousel_img"`
	SearchText  string          `json:"search_text"`
	Router      []*BasicRouter  `json:"router"`
}

func (mc *MainConfig) ToJson() string {
	byteStr, _ := json.Marshal(mc)
	return string(byteStr)
}

func (mv *MainConfig) FromJson(str string) error {
	return json.Unmarshal([]byte(str), mv)
}

func (mv *MainConfig) BuildImgLink() {
	for _, info := range mv.CarouselImg {
		info.TrueLink = common.BuildImageLink(info.Image)
	}
}
