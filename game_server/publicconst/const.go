package publicconst

import "time"

type PlayerState int
type ServiceId int
type ItemSource int

type TaskState int

type TaskType uint32

var (
	MAX_USERID_LEN = 10

	GLOBAL_DB_NAME               = "global"
	GLOBAL_ACCOUNT_COLLECTION    = "account"
	GLOBAL_SERVERINFO_COLLECTION = "serverInfo"
	LOG_DB_NAME                  = "log"

	LOCAL_DB_NAME            = "local"
	LOCAL_ACCOUNT_COLLECTION = "account"
	LOACL_ITEM               = "item"
	LOCAL_TASK               = "task"
	DB_OP_TIME_OUT           = 20 * time.Second

	CLIENT_HEART_INTERVAL = 10 // 客户端心跳间隔(s)
	MAX_CLIENT_HERART_NUM = 3  // 最大心跳包数量

	MAX_UPDATE_ORE_TOTAL_TIME = 10 // 更新矿洞总量时间

	MONGO_NO_RESULT = "mongo: no documents in result"

	MAX_RECYCLE_PLAYER_DATA = 3600 // 玩家数据保留1小时

	REFRESH_ORE_INTEVAL = 10 // 刷新矿洞总量间隔

	GS_SERVER_PREFIX = "/server/gs/"
	MAX_SERVER_TTL   = 30

	BATTLE_HALL_SERVER_PREFIX  = "/server/battleHallServer/"
	BATTLE_SCENE_SERVER_PREFIX = "/server/sceneServer/"
	ORE_SERVER_PREFIX          = "/server/oreServer/"
	MESSAGE_SERVER_PREFIX      = "/server/messageServer/"
)

const (
	Logining PlayerState = iota // 登录中
	Online                      // 在线
	Offline                     // 离线
)

const (
	GMService ServiceId = iota
	ItemService
	OreService
	AccountService
	InterClientService
	BattelService
	FriService
	TaskService
	ActivityService
)

const (
	OreAddItem ItemSource = iota // 挖矿获得
	GMAddItem                    // gm 获得
	OreUpgradeSpeed
)

const (
	TASK_ACCEPT   TaskState = iota
	TASK_COMPLETE           // 完成未领奖
	TASK_DONE               // 完成领奖
)

const (
	DAILY_TASK     TaskType = iota // 日常任务
	WEEKLY_TASK                    // 周长任务
	CHALLENGE_TASK                 // 挑战任务
)
