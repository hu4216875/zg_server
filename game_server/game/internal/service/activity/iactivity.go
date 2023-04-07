package activity

type IActivity interface {
	Update()
	GetReward()
	DailyRefresh()
	GetActivityId()
}
