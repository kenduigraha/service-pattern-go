// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/irahardianto/service-pattern-go/models"

// IPlayerService is an autogenerated mock type for the IPlayerService type
type IPlayerService struct {
	mock.Mock
}

// GetPlayerMessage provides a mock function with given fields:
func (_m *IPlayerService) GetPlayerMessage() models.MessageModel {
	ret := _m.Called()

	var r0 models.MessageModel
	if rf, ok := ret.Get(0).(func() models.MessageModel); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.MessageModel)
	}

	return r0
}

// GetScores provides a mock function with given fields: player1Name, player2Name
func (_m *IPlayerService) GetScores(player1Name string, player2Name string) (string, error) {
	ret := _m.Called(player1Name, player2Name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(player1Name, player2Name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(player1Name, player2Name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
