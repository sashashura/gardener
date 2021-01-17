// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/operation/botanist/extensions/backupentry (interfaces: BackupEntry)

// Package backupentry is a generated GoMock package.
package backupentry

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// MockBackupEntry is a mock of BackupEntry interface.
type MockBackupEntry struct {
	ctrl     *gomock.Controller
	recorder *MockBackupEntryMockRecorder
}

// MockBackupEntryMockRecorder is the mock recorder for MockBackupEntry.
type MockBackupEntryMockRecorder struct {
	mock *MockBackupEntry
}

// NewMockBackupEntry creates a new mock instance.
func NewMockBackupEntry(ctrl *gomock.Controller) *MockBackupEntry {
	mock := &MockBackupEntry{ctrl: ctrl}
	mock.recorder = &MockBackupEntryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupEntry) EXPECT() *MockBackupEntryMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockBackupEntry) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockBackupEntryMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockBackupEntry)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockBackupEntry) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockBackupEntryMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockBackupEntry)(nil).Destroy), arg0)
}

// Migrate mocks base method.
func (m *MockBackupEntry) Migrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockBackupEntryMockRecorder) Migrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockBackupEntry)(nil).Migrate), arg0)
}

// Restore mocks base method.
func (m *MockBackupEntry) Restore(arg0 context.Context, arg1 *v1alpha1.ShootState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockBackupEntryMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockBackupEntry)(nil).Restore), arg0, arg1)
}

// SetBackupBucketProviderStatus mocks base method.
func (m *MockBackupEntry) SetBackupBucketProviderStatus(arg0 *runtime.RawExtension) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBackupBucketProviderStatus", arg0)
}

// SetBackupBucketProviderStatus indicates an expected call of SetBackupBucketProviderStatus.
func (mr *MockBackupEntryMockRecorder) SetBackupBucketProviderStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackupBucketProviderStatus", reflect.TypeOf((*MockBackupEntry)(nil).SetBackupBucketProviderStatus), arg0)
}

// SetProviderConfig mocks base method.
func (m *MockBackupEntry) SetProviderConfig(arg0 *runtime.RawExtension) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProviderConfig", arg0)
}

// SetProviderConfig indicates an expected call of SetProviderConfig.
func (mr *MockBackupEntryMockRecorder) SetProviderConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProviderConfig", reflect.TypeOf((*MockBackupEntry)(nil).SetProviderConfig), arg0)
}

// SetRegion mocks base method.
func (m *MockBackupEntry) SetRegion(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRegion", arg0)
}

// SetRegion indicates an expected call of SetRegion.
func (mr *MockBackupEntryMockRecorder) SetRegion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRegion", reflect.TypeOf((*MockBackupEntry)(nil).SetRegion), arg0)
}

// SetType mocks base method.
func (m *MockBackupEntry) SetType(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetType", arg0)
}

// SetType indicates an expected call of SetType.
func (mr *MockBackupEntryMockRecorder) SetType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetType", reflect.TypeOf((*MockBackupEntry)(nil).SetType), arg0)
}

// Wait mocks base method.
func (m *MockBackupEntry) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockBackupEntryMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockBackupEntry)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockBackupEntry) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockBackupEntryMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockBackupEntry)(nil).WaitCleanup), arg0)
}

// WaitMigrate mocks base method.
func (m *MockBackupEntry) WaitMigrate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMigrate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitMigrate indicates an expected call of WaitMigrate.
func (mr *MockBackupEntryMockRecorder) WaitMigrate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMigrate", reflect.TypeOf((*MockBackupEntry)(nil).WaitMigrate), arg0)
}
