package service

import (
	"github.com/name5566/leaf/log"
	"math"
	"ore_server/base"
	"ore_server/common"
	"ore_server/conf"
	"ore_server/game/internal/dao"
	"ore_server/game/internal/model"
	"ore_server/msg"
	"ore_server/template"
	"strconv"
	"time"
)

var (
	oreInfos []*model.OreInfo
)

func InitOreService() {
	registService(int64(common.MAX_SERVICE_TIME))
	oreInfos = dao.OreDao.LoadOreInfo()
	// 获取所有的矿洞配置没有的则需要加上
	ores := template.GetOreTempalte().GetAll()
	for i := 0; i < len(ores); i++ {
		if getOre(ores[i].Id) == nil {
			data := model.NewOreInfo(ores[i].Id, ores[i].Total)
			dao.OreDao.AddOre(data)
			oreInfos = append(oreInfos, data)
		}
	}
}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.ORE_SERVER_PREFIX + conf.Server.TCPAddr
	strServerId := strconv.Itoa(int(conf.Server.ServerId))
	register := base.NewServiceRegister(endpoint, common.ORE_SERVER_PREFIX, serviceKey, strServerId)
	err := register.Register(ttl)
	if err != nil {
		log.Debug("registService %v\n", err)
	}
}

func GetOreInfo(oreId uint32) (msg.ErrCode, uint32, uint32) {
	ore := getOre(oreId)
	if ore == nil {
		return msg.ErrCode_NO_ORE_RESOURCE, 0, 0
	}
	return msg.ErrCode_SUCC, ore.Total, ore.EndTime
}

func AddOrePlayer(accountId int64, oreId, serverId, speed uint32) (msg.ErrCode, *model.OreInfo, *model.OrePlayer) {
	ore := getOre(oreId)
	if ore == nil {
		return msg.ErrCode_NO_ORE_RESOURCE, nil, nil
	}

	// 没有资源了
	if !canEnterOre(oreId) {
		return msg.ErrCode_NO_ORE_RESOURCE, nil, nil
	}

	if existOrePlayer(ore, accountId) {
		return msg.ErrCode_HAS_START_ORE, nil, nil
	}
	return msg.ErrCode_SUCC, ore, addPlayer(oreId, accountId, serverId, speed)
}

func UpdateOrePlayer(accountId int64, oreId, serverId, newspeed uint32) (msg.ErrCode, uint32, *model.OreInfo) {
	ore := getOre(oreId)
	if ore == nil {
		return msg.ErrCode_NO_ORE_RESOURCE, 0, nil
	}

	if !existOrePlayer(ore, accountId) {
		return msg.ErrCode_NO_START_ORE, 0, nil
	}

	return msg.ErrCode_SUCC, updatePlayer(oreId, accountId, newspeed), nil
}

func SettlePlayer(accountId int64, oreId uint32) (msg.ErrCode, uint32, *model.OreInfo) {
	ore := getOre(oreId)
	if ore == nil {
		return msg.ErrCode_NO_ORE_RESOURCE, 0, nil
	}

	if !existOrePlayer(ore, accountId) {
		return msg.ErrCode_NO_START_ORE, 0, nil
	}
	return msg.ErrCode_SUCC, settlePlayer(oreId, accountId), ore
}

// addPlayer 添加玩家
func addPlayer(oreId uint32, accountId int64, serverId, speed uint32) *model.OrePlayer {
	player := model.NewOrePlayer(accountId, serverId, speed, uint32(time.Now().Unix()))
	oreInfo := addOrePlayer(oreId, player)
	dao.OreDao.AddOrePlayer(oreInfo, player)
	return player
}

// addOrePlayer 增加挖矿的人
func addOrePlayer(oreId uint32, player *model.OrePlayer) *model.OreInfo {
	ore := getOre(oreId)
	if ore == nil {
		return nil
	}
	ore.Players = append(ore.Players, player)

	// 重新计算总时间
	var total int64
	var totalSpeed int64 = 0

	total = int64(ore.Total)
	for i := 0; i < len(ore.Players); i++ {
		total += int64(ore.Players[i].Speed) * int64(ore.Players[i].StartTime)
		totalSpeed += int64(ore.Players[i].Speed)
	}
	ore.EndTime = uint32(total / totalSpeed)
	return ore
}

func getOre(oreId uint32) *model.OreInfo {
	for i := 0; i < len(oreInfos); i++ {
		if oreInfos[i].OreDistId == oreId {
			return oreInfos[i]
		}
	}
	return nil
}

func canEnterOre(oreId uint32) bool {
	if calcCurOreTotal(oreId) <= 0 {
		return false
	}
	return true
}

func calcCurOreTotal(oreId uint32) uint32 {
	data := getOre(oreId)
	if data == nil {
		return 0
	}
	curTime := uint32(time.Now().Unix())
	var temp uint32 = 0
	for i := 0; i < len(data.Players); i++ {
		temp += (curTime - data.Players[i].StartTime) * data.Players[i].Speed
	}
	if data.Total >= temp {
		data.Total = data.Total - temp
		return data.Total
	}
	return 0
}

func existOrePlayer(data *model.OreInfo, accountId int64) bool {
	for i := 0; i < len(data.Players); i++ {
		if data.Players[i].AccountId == accountId {
			return true
		}
	}
	return false
}

func calcOreEndTime(ore *model.OreInfo) uint32 {
	var total int64
	var totalSpeed int64
	total = int64(ore.Total)
	for i := 0; i < len(ore.Players); i++ {
		total += int64(ore.Players[i].Speed) * int64(ore.Players[i].StartTime)
		totalSpeed += int64(ore.Players[i].Speed)
	}

	var endTime uint32
	if totalSpeed > 0 {
		endTime = uint32(total / totalSpeed)
	} else {
		endTime = uint32(math.MaxInt32)
	}
	return endTime
}

// settlePlayer 结算玩家(离线，离开矿区)
func settlePlayer(oreId uint32, accountId int64) uint32 {
	ore := getOre(oreId)
	if ore == nil {
		return 0
	}

	var player *model.OrePlayer = nil
	for i := 0; i < len(ore.Players); i++ {
		if ore.Players[i].AccountId == accountId {
			player = ore.Players[i]
			ore.Players = append(ore.Players[:i], ore.Players[i+1:]...)
			break
		}
	}

	if player == nil {
		return 0
	}

	curTime := uint32(time.Now().Unix())
	num := player.Speed * (curTime - player.StartTime)
	if num > ore.Total {
		num = ore.Total
		ore.Total = 0
	} else {
		ore.EndTime = calcOreEndTime(ore)
		ore.Total -= num
	}

	dao.OreDao.RemoveOrePlayer(ore, player)
	if num > 0 {
		dao.OreDao.UpdateOreRecord(ore.OreDistId, player.AccountId, num)
	}
	return num
}

// updatePlayer 修改玩家
func updatePlayer(oreId uint32, accountId int64, newSpeed uint32) uint32 {
	ore := getOre(oreId)
	if ore == nil {
		log.Error("updatePlayer oreId:%v, accountId:%v newSpeed:%v ore null", oreId, accountId, newSpeed)
		return 0
	}

	var player *model.OrePlayer = nil
	for i := 0; i < len(ore.Players); i++ {
		if ore.Players[i].AccountId == accountId {
			player = ore.Players[i]
			break
		}
	}

	if player == nil {
		return 0
	}

	curTime := uint32(time.Now().Unix())
	num := player.Speed * (curTime - player.StartTime)
	if num > ore.Total {
		num = ore.Total
		ore.Total = 0
	} else {
		ore.Total -= num
	}
	ore.EndTime = calcOreEndTime(ore)
	player.Speed = newSpeed
	player.StartTime = uint32(time.Now().Unix())
	dao.OreDao.UpdateOrePlayer(ore, player)
	dao.OreDao.UpdateOreRecord(ore.OreDistId, player.AccountId, uint32(num))
	return num
}
