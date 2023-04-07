package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/msg"
	"server/publicconst"
	"server/template"
)

func NewOreConn(serverId uint32, serverAddr string) *common.TcpClient {
	cli := common.NewTcpClient(serverId, serverAddr, nil)
	if err := cli.Conn(); err != nil {
		log.Error("serverid:%v serverAddr:%v err:%v", serverId, serverAddr, err)
		return nil
	}
	initOreHandler(cli)

	oreId := template.GetSystemItemTemplate().GetOreId()
	reqMsg := &msg.RequestInterGetOreInfo{
		ServerId: conf.Server.ServerId,
		OreId:    oreId,
	}
	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterGetOreInfo), reqMsg)
	log.Debug("NewOreConn serverId:%v\n", serverId)
	return cli
}

func initOreHandler(cli *common.TcpClient) {
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterGetOreInfo), responseInterGetOreInfo)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterAddOrePlayer), responseInterAddOrePlayer)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterUpdateOrePlayer), responseInterUpdateOrePlayer)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterSettleOrePlayer), responseInterSettleOrePlayer)
}

func responseInterGetOreInfo(data []byte) {
	res := &msg.ResponseInterGetOreInfo{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterGetOreInfo err:%v", err)
		return
	}

	accountId := res.AccountId
	if playData := common.PlayerMgr.FindPlayerData(accountId); playData != nil {
		ackMsg := &msg.ResponseOreTotal{
			Result: res.Result,
			OreId:  res.OreId,
			Total:  res.Total,
		}
		playData.PlayerAgent.WriteMsg(ackMsg)
	}
	ServMgr.GetOreService().updateOreInfo(res.OreId, res.Total, res.EndTime)
	log.Debug("responseInterGetOreInfo result:%v", res.Result)
}

func responseInterAddOrePlayer(data []byte) {
	res := &msg.ResponseInterAddOrePlayer{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("ResponseInterAddOrePlayer err:%v", err)
		return
	}

	accountId := res.AccountId
	if playData := common.PlayerMgr.FindPlayerData(accountId); playData != nil {
		ackMsg := &msg.ResponseStartOre{
			Result:    res.Result,
			Total:     res.Total,
			EndTime:   res.EndTime,
			StartTime: res.StartTime,
		}

		if res.Result == int32(msg.ErrCode_SUCC) {
			playData.AccountInfo.OreInfo.OreId = res.OreId
			playData.AccountInfo.OreInfo.Speed = res.Speed
			playData.AccountInfo.OreInfo.StartTime = res.StartTime
			ServMgr.GetOreService().updateOreInfo(res.OreId, res.Total, res.EndTime)

			dao.OreInfoDao.UpdateOreInfo(accountId, playData.AccountInfo.OreInfo)
		}
		playData.PlayerAgent.WriteMsg(ackMsg)
	}
	log.Debug("responseInterAddOrePlayer result:%v", res.Result)
}

func responseInterUpdateOrePlayer(data []byte) {
	res := &msg.ResponseInterUpdateOrePlayer{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("ResponseInterUpdateOrePlayer err:%v", err)
		return
	}

	accountId := res.AccountId
	if playData := common.PlayerMgr.FindPlayerData(accountId); playData != nil {
		ackMsg := &msg.ResponseUpgradeOreSpeed{
			Result:  res.Result,
			Total:   res.Total,
			EndTime: res.EndTime,
		}

		if res.Result == int32(msg.ErrCode_SUCC) {
			playData.AccountInfo.OreInfo.OreId = res.OreId
			playData.AccountInfo.OreInfo.Speed = res.Speed
			playData.AccountInfo.OreInfo.StartTime = res.StartTime
			dao.OreInfoDao.UpdateOreInfo(accountId, playData.AccountInfo.OreInfo)
			ServMgr.GetOreService().updateOreInfo(res.OreId, res.Total, res.EndTime)

			if res.Num > 0 {
				oreItemId := template.GetSystemItemTemplate().GetOreItemId()
				ServMgr.GetItemService().AddItem(accountId, oreItemId, res.Num, publicconst.OreAddItem)
			}
		}
		playData.PlayerAgent.WriteMsg(ackMsg)
	}
	log.Debug("responseInterAddOrePlayer result:%v", res.Result)
}

func responseInterSettleOrePlayer(data []byte) {
	res := &msg.ResponseInterSettleOrePlayer{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterSettleOrePlayer err:%v", err)
		return
	}

	accountId := res.AccountId
	if playData := common.PlayerMgr.FindPlayerData(accountId); playData != nil {
		if res.Result == uint32(msg.ErrCode_SUCC) {
			playData.AccountInfo.OreInfo.StartTime = 0
			playData.AccountInfo.OreInfo.Speed = 0
			dao.OreInfoDao.UpdateOreInfo(accountId, playData.AccountInfo.OreInfo)
			ServMgr.GetOreService().updateOreInfo(res.OreId, res.Total, res.EndTime)
			if res.Num > 0 {
				oreItemId := template.GetSystemItemTemplate().GetOreItemId()
				ServMgr.GetItemService().AddItem(accountId, oreItemId, res.Num, publicconst.OreAddItem)
			}
		}

		ackMsg := &msg.ResponseEndOre{
			Result:  int32(res.Result),
			Total:   res.Total,
			EndTime: res.EndTime,
		}
		playData.PlayerAgent.WriteMsg(ackMsg)
	}
	log.Debug("responseInterSettleOrePlayer result:%v", res.Result)
}
