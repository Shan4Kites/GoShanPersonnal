package unittesting

type Load struct {
	stops []string
}

//It calls tracking service
func GetLoad(trackingId string) Load {
	return Load{stops: []string{}}
}

type Service struct {
}

func (Service) isMultiLegLoad(trackingId string) bool {
	load := GetLoad(trackingId)
	return len(load.stops) > 2
}