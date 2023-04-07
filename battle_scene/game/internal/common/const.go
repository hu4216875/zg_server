package common

type ServiceId uint32

var (
	BATTLE_HALL_SERVER_PREFIX  = "/server/battleHallServer/"
	MAX_SERVICE_TIME           = 10
	BATTLE_SCENE_SERVER_PREFIX = "/server/sceneServer/"
)

const (
	InterClientSerivce ServiceId = iota
	SceneService
	ServerService
	PlayerServgice
)
