// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"

// WorkflowHistoryRepository is an autogenerated mock type for the WorkflowHistoryRepository type
type WorkflowHistoryRepository struct {
	mock.Mock
}

// AddWorkflowHistory provides a mock function with given fields: wf
func (_m *WorkflowHistoryRepository) AddWorkflowHistory(wf *v1alpha1.Workflow) error {
	ret := _m.Called(wf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Workflow) error); ok {
		r0 = rf(wf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWorkflowHistory provides a mock function with given fields: namespace, uid
func (_m *WorkflowHistoryRepository) GetWorkflowHistory(namespace string, uid string) (*v1alpha1.Workflow, error) {
	ret := _m.Called(namespace, uid)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Workflow); ok {
		r0 = rf(namespace, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflowHistory provides a mock function with given fields: limit, offset
func (_m *WorkflowHistoryRepository) ListWorkflowHistory(limit int, offset int) ([]v1alpha1.Workflow, error) {
	ret := _m.Called(limit, offset)

	var r0 []v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(int, int) []v1alpha1.Workflow); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
