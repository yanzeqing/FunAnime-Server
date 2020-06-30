package responseVideo

import (
	"encoding/json"
	"github.com/yanzeqing/FunAnime-Server/model"
	barrage "github.com/yanzeqing/FunAnime-Server/service/websocket"
)

func BuildBarrageArrayResp(modelList []*model.FaBarrage) []*barrage.BarrageType {
	resp := make([]*barrage.BarrageType, len(modelList))
	for i, faBarrage := range modelList {
		barr := new(barrage.BarrageType)
		_ = json.Unmarshal([]byte(faBarrage.BarrageText), barr)

		resp[i] = barr
	}

	return resp
}
