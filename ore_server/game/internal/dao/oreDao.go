package dao

import (
	"errors"
	"fmt"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"ore_server/common"
	"ore_server/db"
	"ore_server/game/internal/model"
	"time"
)

var (
	OreDao oreDao
)

type oreDao struct {
}

// LoadOreInfo 加载所有矿洞
func (o *oreDao) LoadOreInfo() []*model.OreInfo {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		log.Error("LoadOreInfo err:%v", currErr)
		return nil
	}
	defer cur.Close(ctx)

	var oreInfo []*model.OreInfo
	if err := cur.All(ctx, &oreInfo); err != nil {
		log.Error("LoadOreInfo err:%v", err)
		return nil
	}
	return oreInfo
}

// AddOre 添加矿洞
func (o *oreDao) AddOre(data *model.OreInfo) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, data); err != nil {
		return errors.New(fmt.Sprintf("AddOre err:%v", err))
	}

	oreLog := model.NewOreLog(data.OreDistId)
	logCollection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_LOG_COLLECTION)
	opts := options.UpdateOptions{}
	opts.SetUpsert(true)
	if _, err := logCollection.UpdateOne(ctx, bson.M{"oreid": data.OreDistId}, bson.D{{"$set", oreLog}}, &opts); err != nil {
		return errors.New(fmt.Sprintf("AddOreDistrict err:%v", err))
	}
	return nil
}

// AddOrePlayer 添加矿洞玩家
func (o *oreDao) AddOrePlayer(ore *model.OreInfo, player *model.OrePlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	filter := bson.M{"oredistid": ore.OreDistId}
	update := bson.D{{"$addToSet", bson.D{{"players", player}}},
		{"$set", bson.D{{"endtime", ore.EndTime}, {"updatetime", curTime}}}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		log.Error("AddOrePlayer err:%v", err)
		return errors.New(fmt.Sprintf("AddOrePlayer err:%v", err))
	}
	return nil
}

// RemoveOrePlayer 移除矿洞玩家
func (o *oreDao) RemoveOrePlayer(ore *model.OreInfo, player *model.OrePlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	update := bson.D{{"$set", bson.D{{"endtime", ore.EndTime}, {"total", ore.Total}, {"updatetime", curTime}}}}
	if _, err := collection.UpdateOne(ctx, bson.D{{"oredistid", ore.OreDistId}}, update); err != nil {
		log.Error("Account:%v RemoveOrePlayer error:%v", player.AccountId, err)
		return err
	}

	if _, err := collection.UpdateOne(ctx, bson.M{"oredistid": ore.OreDistId}, bson.M{"$pull": bson.M{"players": bson.M{"accountid": player.AccountId}}}); err != nil {
		log.Error("Account:%v RemoveOrePlayer err:%v", player.AccountId, err)
		return err
	}
	return nil
}

// UpdateOrePlayer 更新矿洞玩家
func (o *oreDao) UpdateOrePlayer(ore *model.OreInfo, player *model.OrePlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	update := bson.D{{"$set", bson.D{{"endtime", ore.EndTime}, {"total", ore.Total}, {"updatetime", curTime}}}}

	if _, err := collection.UpdateOne(ctx, bson.D{{"oredistid", ore.OreDistId}}, update); err != nil {
		log.Error("Account:%v UpdateOreDistrictPlayer error:%v", player.AccountId, err)
		return err
	}

	filter := bson.M{"oredistid": ore.OreDistId, "players.accountid": player.AccountId}
	update = bson.D{{"players.$[item].speed", player.Speed}, {"players.$[item].starttime", player.StartTime}}
	arrayFilter := bson.M{"item.accountid": player.AccountId}
	res := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update},
		options.FindOneAndUpdate().SetArrayFilters(
			options.ArrayFilters{
				Filters: []interface{}{
					arrayFilter,
				},
			},
		),
	)
	if res.Err() != nil {
		log.Error("Account:%v UpdateOreDistrictPlayer error:%v", player.AccountId, res.Err())
	}
	return nil
}

// UpdateOreRecord 更新矿洞记录
func (o *oreDao) UpdateOreRecord(oreId uint32, accountId int64, num uint32) {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.ORE_DB_LOG_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()
	record := model.NewOreRecord(accountId, num)
	if _, err := collection.UpdateOne(ctx, bson.M{"oreid": oreId}, bson.D{{"$addToSet", bson.D{{"records", record}}}}); err != nil {
		log.Error("UpdateOreRecord err:%v", err)
	}
}
