// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/utils/utils.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "pet/pkg/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUtility is a mock of Utility interface.
type MockUtility struct {
	ctrl     *gomock.Controller
	recorder *MockUtilityMockRecorder
}

// MockUtilityMockRecorder is the mock recorder for MockUtility.
type MockUtilityMockRecorder struct {
	mock *MockUtility
}

// NewMockUtility creates a new mock instance.
func NewMockUtility(ctrl *gomock.Controller) *MockUtility {
	mock := &MockUtility{ctrl: ctrl}
	mock.recorder = &MockUtilityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUtility) EXPECT() *MockUtilityMockRecorder {
	return m.recorder
}

// GetRandomArticle mocks base method.
func (m *MockUtility) GetRandomArticle() models.Article {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomArticle")
	ret0, _ := ret[0].(models.Article)
	return ret0
}

// GetRandomArticle indicates an expected call of GetRandomArticle.
func (mr *MockUtilityMockRecorder) GetRandomArticle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomArticle", reflect.TypeOf((*MockUtility)(nil).GetRandomArticle))
}

// GetUUID mocks base method.
func (m *MockUtility) GetUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUUID indicates an expected call of GetUUID.
func (mr *MockUtilityMockRecorder) GetUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUUID", reflect.TypeOf((*MockUtility)(nil).GetUUID))
}

// Random mocks base method.
func (m *MockUtility) Random(min, max int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Random", min, max)
	ret0, _ := ret[0].(int)
	return ret0
}

// Random indicates an expected call of Random.
func (mr *MockUtilityMockRecorder) Random(min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Random", reflect.TypeOf((*MockUtility)(nil).Random), min, max)
}

// RandomAuthor mocks base method.
func (m *MockUtility) RandomAuthor() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomAuthor")
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomAuthor indicates an expected call of RandomAuthor.
func (mr *MockUtilityMockRecorder) RandomAuthor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomAuthor", reflect.TypeOf((*MockUtility)(nil).RandomAuthor))
}

// RandomContent mocks base method.
func (m *MockUtility) RandomContent() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomContent")
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomContent indicates an expected call of RandomContent.
func (mr *MockUtilityMockRecorder) RandomContent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomContent", reflect.TypeOf((*MockUtility)(nil).RandomContent))
}

// RandomTitle mocks base method.
func (m *MockUtility) RandomTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// RandomTitle indicates an expected call of RandomTitle.
func (mr *MockUtilityMockRecorder) RandomTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomTitle", reflect.TypeOf((*MockUtility)(nil).RandomTitle))
}
