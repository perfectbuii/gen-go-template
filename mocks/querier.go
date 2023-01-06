// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/perfectbuii/gen-go-template/internal/models"
	mock "github.com/stretchr/testify/mock"

	pgtype "github.com/jackc/pgtype"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateHub provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateHub(ctx context.Context, arg models.CreateHubParams) (*models.Hub, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateHubParams) *models.Hub); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CreateHubParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTeam provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateTeam(ctx context.Context, arg models.CreateTeamParams) (*models.Team, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateTeamParams) *models.Team); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CreateTeamParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateUser(ctx context.Context, arg models.CreateUserParams) (*models.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateUserParams) *models.User); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHub provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteHub(ctx context.Context, id pgtype.Text) (*models.Hub, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.Hub); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTeam provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteTeam(ctx context.Context, id pgtype.Text) (*models.Team, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.Team); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteUser(ctx context.Context, id pgtype.Text) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindHubByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindHubByID(ctx context.Context, id pgtype.Text) (*models.Hub, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.Hub); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTeamByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindTeamByID(ctx context.Context, id pgtype.Text) (*models.Team, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.Team); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindUserByID(ctx context.Context, id pgtype.Text) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.Text) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgtype.Text) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListHub provides a mock function with given fields: ctx, arg
func (_m *Querier) GetListHub(ctx context.Context, arg models.GetListHubParams) ([]*models.Hub, error) {
	ret := _m.Called(ctx, arg)

	var r0 []*models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, models.GetListHubParams) []*models.Hub); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetListHubParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListTeam provides a mock function with given fields: ctx, arg
func (_m *Querier) GetListTeam(ctx context.Context, arg models.GetListTeamParams) ([]*models.Team, error) {
	ret := _m.Called(ctx, arg)

	var r0 []*models.Team
	if rf, ok := ret.Get(0).(func(context.Context, models.GetListTeamParams) []*models.Team); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetListTeamParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListUser provides a mock function with given fields: ctx, arg
func (_m *Querier) GetListUser(ctx context.Context, arg models.GetListUserParams) ([]*models.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.GetListUserParams) []*models.User); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetListUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHub provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateHub(ctx context.Context, arg models.UpdateHubParams) (*models.Hub, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, models.UpdateHubParams) *models.Hub); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Hub)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UpdateHubParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTeam provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateTeam(ctx context.Context, arg models.UpdateTeamParams) (*models.Team, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(context.Context, models.UpdateTeamParams) *models.Team); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UpdateTeamParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateUser(ctx context.Context, arg models.UpdateUserParams) (*models.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.UpdateUserParams) *models.User); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UpdateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQuerier interface {
	mock.TestingT
	Cleanup(func())
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuerier(t mockConstructorTestingTNewQuerier) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
