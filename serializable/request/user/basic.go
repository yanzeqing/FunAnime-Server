package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yanzeqing/FunAnime-Server/middleware/token"
	"github.com/yanzeqing/FunAnime-Server/util/common"
	"github.com/yanzeqing/FunAnime-Server/util/errno"
)

type BasicUser struct {
	UserInfo *token.UserInfo `json:"user_info"`
}

func (bu *BasicUser) GetUserInfo(ctx *gin.Context) error {
	userInfo := GetUserInfoFromContext(ctx)
	if userInfo == nil {
		common.EchoFailedJson(ctx, errno.Uncertified)
		return errors.New("user_not_login")
	}

	bu.UserInfo = userInfo

	return nil
}

func GetUserInfoFromContext(ctx *gin.Context) *token.UserInfo {
	uInfo, ok := ctx.Get("userInfo")
	if !ok {
		fmt.Println(uInfo)
		return nil
	}

	userInfo, ok := uInfo.(*token.UserInfo)
	if !ok {
		fmt.Println(userInfo)
		return nil
	}

	return userInfo
}
