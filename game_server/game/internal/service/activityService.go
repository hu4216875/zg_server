package service

type ActivityService struct {
	IService
}

func NewActivityService() *ActivityService {
	return &ActivityService{}
}

func (t *ActivityService) OnInit() {
}

func (t *ActivityService) OnDestory() {

}
