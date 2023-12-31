// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/controller/controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// GetArticle mocks base method.
func (m *MockController) GetArticle(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetArticle", c)
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockControllerMockRecorder) GetArticle(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockController)(nil).GetArticle), c)
}

// GetArticles mocks base method.
func (m *MockController) GetArticles(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetArticles", c)
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockControllerMockRecorder) GetArticles(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockController)(nil).GetArticles), c)
}
