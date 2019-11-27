package easy

import (
	"github.com/Shan4Kites/GoShanPersonnal/unittesting/easy/mocks"
	"github.com/Shan4Kites/GoShanPersonnal/unittesting/easy/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMultiLegLoad_SingleLeg(t *testing.T) {
	trackingClientMock := new(mocks.TrackingClient)
	service := Service{trackingClient: trackingClientMock}
	load := model.Load{Stops:[]string{"s1","s2"}}
	trackingClientMock.On("GetLoad", "tracking-id1").Return(load)

	res := service.isMultiLegLoad("tracking-id1")

	assert.False(t, res)
}

func TestIsMultiLegLoad_MultiLeg(t *testing.T) {
	trackingClientMock := new(mocks.TrackingClient)
	service := Service{trackingClient: trackingClientMock}
	load := model.Load{Stops:[]string{"s1","s2", "s3"}}
	trackingClientMock.On("GetLoad", "tracking-id1").Return(load)

	res := service.isMultiLegLoad("tracking-id1")

	assert.True(t, res)
}