// Code generated by MockGen. DO NOT EDIT.
// Source: ./traffictarget_translator.go

// Package mock_traffictarget is a generated GoMock package.
package mock_traffictarget

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// TranslateTrafficTargets mocks base method
func (m *MockTranslator) TranslateTrafficTargets(services v1sets.ServiceSet, workloads v1alpha2sets.WorkloadSet, meshes v1alpha2sets.MeshSet) v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateTrafficTargets", services, workloads, meshes)
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// TranslateTrafficTargets indicates an expected call of TranslateTrafficTargets
func (mr *MockTranslatorMockRecorder) TranslateTrafficTargets(services, workloads, meshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateTrafficTargets", reflect.TypeOf((*MockTranslator)(nil).TranslateTrafficTargets), services, workloads, meshes)
}
