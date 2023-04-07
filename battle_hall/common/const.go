package common

import "time"

var (
	MAX_SERVICE_TIME          = 1000 // 续约时间
	BATTLE_HALL_SERVER_PREFIX = "/server/battleHallServer/"

	DB_BATTLE_HALL        = "battleHall"
	DB_BATTLE_HALL_SERVER = "battleServer"
	DB_OP_TIME_OUT        = 10 * time.Second

	MAX_SCENE_LIMIT = 1000
)
