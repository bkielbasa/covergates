// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/code-devel-cover/CodeCover/core (interfaces: ReportStore,CoverageService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	core "github.com/code-devel-cover/CodeCover/core"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockReportStore is a mock of ReportStore interface
type MockReportStore struct {
	ctrl     *gomock.Controller
	recorder *MockReportStoreMockRecorder
}

// MockReportStoreMockRecorder is the mock recorder for MockReportStore
type MockReportStoreMockRecorder struct {
	mock *MockReportStore
}

// NewMockReportStore creates a new mock instance
func NewMockReportStore(ctrl *gomock.Controller) *MockReportStore {
	mock := &MockReportStore{ctrl: ctrl}
	mock.recorder = &MockReportStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReportStore) EXPECT() *MockReportStoreMockRecorder {
	return m.recorder
}

// CreateComment mocks base method
func (m *MockReportStore) CreateComment(arg0 *core.Report, arg1 *core.ReportComment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockReportStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockReportStore)(nil).CreateComment), arg0, arg1)
}

// Find mocks base method
func (m *MockReportStore) Find(arg0 *core.Report) (*core.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*core.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockReportStoreMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockReportStore)(nil).Find), arg0)
}

// FindComment mocks base method
func (m *MockReportStore) FindComment(arg0 *core.Report, arg1 int) (*core.ReportComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindComment", arg0, arg1)
	ret0, _ := ret[0].(*core.ReportComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindComment indicates an expected call of FindComment
func (mr *MockReportStoreMockRecorder) FindComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindComment", reflect.TypeOf((*MockReportStore)(nil).FindComment), arg0, arg1)
}

// Finds mocks base method
func (m *MockReportStore) Finds(arg0 *core.Report) ([]*core.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finds", arg0)
	ret0, _ := ret[0].([]*core.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Finds indicates an expected call of Finds
func (mr *MockReportStoreMockRecorder) Finds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finds", reflect.TypeOf((*MockReportStore)(nil).Finds), arg0)
}

// Upload mocks base method
func (m *MockReportStore) Upload(arg0 *core.Report) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockReportStoreMockRecorder) Upload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockReportStore)(nil).Upload), arg0)
}

// MockCoverageService is a mock of CoverageService interface
type MockCoverageService struct {
	ctrl     *gomock.Controller
	recorder *MockCoverageServiceMockRecorder
}

// MockCoverageServiceMockRecorder is the mock recorder for MockCoverageService
type MockCoverageServiceMockRecorder struct {
	mock *MockCoverageService
}

// NewMockCoverageService creates a new mock instance
func NewMockCoverageService(ctrl *gomock.Controller) *MockCoverageService {
	mock := &MockCoverageService{ctrl: ctrl}
	mock.recorder = &MockCoverageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoverageService) EXPECT() *MockCoverageServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockCoverageService) Find(arg0 context.Context, arg1 core.ReportType, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCoverageServiceMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCoverageService)(nil).Find), arg0, arg1, arg2)
}

// Open mocks base method
func (m *MockCoverageService) Open(arg0 context.Context, arg1 core.ReportType, arg2 string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockCoverageServiceMockRecorder) Open(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockCoverageService)(nil).Open), arg0, arg1, arg2)
}

// Report mocks base method
func (m *MockCoverageService) Report(arg0 context.Context, arg1 core.ReportType, arg2 io.Reader) (*core.CoverageReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Report", arg0, arg1, arg2)
	ret0, _ := ret[0].(*core.CoverageReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Report indicates an expected call of Report
func (mr *MockCoverageServiceMockRecorder) Report(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockCoverageService)(nil).Report), arg0, arg1, arg2)
}

// TrimFileNames mocks base method
func (m *MockCoverageService) TrimFileNames(arg0 context.Context, arg1 *core.CoverageReport, arg2 core.FileNameFilters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrimFileNames", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TrimFileNames indicates an expected call of TrimFileNames
func (mr *MockCoverageServiceMockRecorder) TrimFileNames(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrimFileNames", reflect.TypeOf((*MockCoverageService)(nil).TrimFileNames), arg0, arg1, arg2)
}
