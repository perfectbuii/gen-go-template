// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/perfectbuii/gen-go-template/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// SearchService is an autogenerated mock type for the SearchService type
type SearchService struct {
	mock.Mock
}

// TeamHub provides a mock function with given fields: ctx, q
func (_m *SearchService) TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error) {
	ret := _m.Called(ctx, q)

	var r0 []*models.Team
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.Team); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Team)
		}
	}

	var r1 []*models.Hub
	if rf, ok := ret.Get(1).(func(context.Context, string) []*models.Hub); ok {
		r1 = rf(ctx, q)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*models.Hub)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewSearchService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSearchService creates a new instance of SearchService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSearchService(t mockConstructorTestingTNewSearchService) *SearchService {
	mock := &SearchService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
