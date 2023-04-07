package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/db"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
	"strings"
	"time"
)

var (
	TaskDao = &taskDao{}
)

type taskDao struct {
}

func (a *taskDao) LoadTasks(accountId int64) (msg.ErrCode, *model.AccountTask) {
	collection := db.GetGlobalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_TASK)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"accountid": accountId})
	if result == nil {
		return msg.ErrCode_SYSTEM_ERROR, nil
	}

	var ret model.AccountTask
	err := result.Decode(&ret)
	if err != nil {
		if strings.Contains(publicconst.MONGO_NO_RESULT, err.Error()) {
			accountItem := model.NewAccountTask(accountId)
			if _, err = collection.InsertOne(ctx, accountItem); err != nil {
				log.Error("LoadTasks accountId:%v err:%v", accountId, err)
			}
		} else {
			log.Error("LoadTasks accountId:%v err:%v", accountId, err)
			return msg.ErrCode_SYSTEM_ERROR, nil
		}
	}
	return msg.ErrCode_SUCC, &ret
}

func (a *taskDao) AddTask(accountId int64, task *model.Task) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_TASK)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.D{{"$addToSet", bson.D{{"tasks", task}}}}); err != nil {
		log.Error("AddTask err:%v", err)
		return msg.ErrCode_SYSTEM_ERROR
	}
	return msg.ErrCode_SUCC
}

func (a *taskDao) UpdateTask(accountId int64, task *model.Task) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_TASK)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	filter := bson.M{"accountid": accountId, "tasks.taskid": task.TaskId}
	update := bson.M{"tasks.$[item].taskvalue": task.TaskValue, "tasks.$[item].taskstate": task.TaskState, "tasks.$[item].updatetime": curTime}
	arrayFilter := bson.M{"item.taskid": task.TaskId}
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
		log.Error("account:%v UpdateTask error:%v", accountId, res.Err())
	}
	return msg.ErrCode_SUCC
}

func (a *taskDao) DeleteTask(accountId int64, taskId uint32) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_TASK)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.M{"$pull": bson.M{"tasks": bson.M{"taskid": taskId}}}); err != nil {
		log.Error("account:%v DeleteTask err:%v", accountId, err)
		return msg.ErrCode_SYSTEM_ERROR
	}
	return msg.ErrCode_SUCC
}
