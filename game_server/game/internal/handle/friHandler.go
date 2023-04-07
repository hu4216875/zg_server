package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestAddFriHandle 请求加好友
func RequestAddFriHandle(args interface{}, playerData *common.PlayerData) {
	request := args.(*msg.RequestAddFri)
	err := service.ServMgr.GetFriService().AddFriend(playerData.AccountInfo.AccountId, request.AccountId, request.DestGsId)
	if err != msg.ErrCode_SUCC {
		retMsg := &msg.ResponseAddFri{
			Result: uint32(err),
		}
		playerData.PlayerAgent.WriteMsg(retMsg)
	}
}
