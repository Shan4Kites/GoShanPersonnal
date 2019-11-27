package easy

import "github.com/Shan4Kites/GoShanPersonnal/unittesting/easy/model"

type TrackingClient interface {
	GetLoad(string) model.Load
}

type TrackingService struct {
}
//It calls tracking service
func (t TrackingService) GetLoad(trackingId string) model.Load {
	return model.Load{Stops: []string{}}
}

type Service struct {
	trackingClient TrackingClient
}

func NewService() Service {
	return Service{trackingClient: TrackingService{}}
}

func (s Service) isMultiLegLoad(trackingId string) bool {
	load := s.trackingClient.GetLoad(trackingId)
	return len(load.Stops) > 2
}