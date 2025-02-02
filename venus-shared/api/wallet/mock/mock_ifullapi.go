// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/venus/venus-shared/api/wallet (interfaces: IFullAPI)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	auth "github.com/filecoin-project/go-jsonrpc/auth"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	internal "github.com/filecoin-project/venus/venus-shared/internal"
	types "github.com/filecoin-project/venus/venus-shared/types"
	wallet "github.com/filecoin-project/venus/venus-shared/types/wallet"
	gomock "github.com/golang/mock/gomock"
)

// MockIFullAPI is a mock of IFullAPI interface.
type MockIFullAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIFullAPIMockRecorder
}

// MockIFullAPIMockRecorder is the mock recorder for MockIFullAPI.
type MockIFullAPIMockRecorder struct {
	mock *MockIFullAPI
}

// NewMockIFullAPI creates a new mock instance.
func NewMockIFullAPI(ctrl *gomock.Controller) *MockIFullAPI {
	mock := &MockIFullAPI{ctrl: ctrl}
	mock.recorder = &MockIFullAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFullAPI) EXPECT() *MockIFullAPIMockRecorder {
	return m.recorder
}

// AddMethodIntoKeyBind mocks base method.
func (m *MockIFullAPI) AddMethodIntoKeyBind(arg0 context.Context, arg1 string, arg2 []string) (*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMethodIntoKeyBind", arg0, arg1, arg2)
	ret0, _ := ret[0].(*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMethodIntoKeyBind indicates an expected call of AddMethodIntoKeyBind.
func (mr *MockIFullAPIMockRecorder) AddMethodIntoKeyBind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMethodIntoKeyBind", reflect.TypeOf((*MockIFullAPI)(nil).AddMethodIntoKeyBind), arg0, arg1, arg2)
}

// AddMsgTypeIntoKeyBind mocks base method.
func (m *MockIFullAPI) AddMsgTypeIntoKeyBind(arg0 context.Context, arg1 string, arg2 []int) (*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMsgTypeIntoKeyBind", arg0, arg1, arg2)
	ret0, _ := ret[0].(*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMsgTypeIntoKeyBind indicates an expected call of AddMsgTypeIntoKeyBind.
func (mr *MockIFullAPIMockRecorder) AddMsgTypeIntoKeyBind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMsgTypeIntoKeyBind", reflect.TypeOf((*MockIFullAPI)(nil).AddMsgTypeIntoKeyBind), arg0, arg1, arg2)
}

// AddNewAddress mocks base method.
func (m *MockIFullAPI) AddNewAddress(arg0 context.Context, arg1 []address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewAddress indicates an expected call of AddNewAddress.
func (mr *MockIFullAPIMockRecorder) AddNewAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewAddress", reflect.TypeOf((*MockIFullAPI)(nil).AddNewAddress), arg0, arg1)
}

// AddSupportAccount mocks base method.
func (m *MockIFullAPI) AddSupportAccount(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSupportAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSupportAccount indicates an expected call of AddSupportAccount.
func (mr *MockIFullAPIMockRecorder) AddSupportAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSupportAccount", reflect.TypeOf((*MockIFullAPI)(nil).AddSupportAccount), arg0, arg1)
}

// AuthNew mocks base method.
func (m *MockIFullAPI) AuthNew(arg0 context.Context, arg1 []auth.Permission) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthNew", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthNew indicates an expected call of AuthNew.
func (mr *MockIFullAPIMockRecorder) AuthNew(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthNew", reflect.TypeOf((*MockIFullAPI)(nil).AuthNew), arg0, arg1)
}

// AuthVerify mocks base method.
func (m *MockIFullAPI) AuthVerify(arg0 context.Context, arg1 string) ([]auth.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthVerify", arg0, arg1)
	ret0, _ := ret[0].([]auth.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthVerify indicates an expected call of AuthVerify.
func (mr *MockIFullAPIMockRecorder) AuthVerify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthVerify", reflect.TypeOf((*MockIFullAPI)(nil).AuthVerify), arg0, arg1)
}

// ContainWallet mocks base method.
func (m *MockIFullAPI) ContainWallet(arg0 context.Context, arg1 address.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainWallet", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainWallet indicates an expected call of ContainWallet.
func (mr *MockIFullAPIMockRecorder) ContainWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainWallet", reflect.TypeOf((*MockIFullAPI)(nil).ContainWallet), arg0, arg1)
}

// GetGroupByName mocks base method.
func (m *MockIFullAPI) GetGroupByName(arg0 context.Context, arg1 string) (*wallet.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupByName", arg0, arg1)
	ret0, _ := ret[0].(*wallet.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupByName indicates an expected call of GetGroupByName.
func (mr *MockIFullAPIMockRecorder) GetGroupByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupByName", reflect.TypeOf((*MockIFullAPI)(nil).GetGroupByName), arg0, arg1)
}

// GetKeyBindByName mocks base method.
func (m *MockIFullAPI) GetKeyBindByName(arg0 context.Context, arg1 string) (*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyBindByName", arg0, arg1)
	ret0, _ := ret[0].(*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyBindByName indicates an expected call of GetKeyBindByName.
func (mr *MockIFullAPIMockRecorder) GetKeyBindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyBindByName", reflect.TypeOf((*MockIFullAPI)(nil).GetKeyBindByName), arg0, arg1)
}

// GetKeyBinds mocks base method.
func (m *MockIFullAPI) GetKeyBinds(arg0 context.Context, arg1 address.Address) ([]*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyBinds", arg0, arg1)
	ret0, _ := ret[0].([]*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyBinds indicates an expected call of GetKeyBinds.
func (mr *MockIFullAPIMockRecorder) GetKeyBinds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyBinds", reflect.TypeOf((*MockIFullAPI)(nil).GetKeyBinds), arg0, arg1)
}

// GetMethodTemplateByName mocks base method.
func (m *MockIFullAPI) GetMethodTemplateByName(arg0 context.Context, arg1 string) (*wallet.MethodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMethodTemplateByName", arg0, arg1)
	ret0, _ := ret[0].(*wallet.MethodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMethodTemplateByName indicates an expected call of GetMethodTemplateByName.
func (mr *MockIFullAPIMockRecorder) GetMethodTemplateByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMethodTemplateByName", reflect.TypeOf((*MockIFullAPI)(nil).GetMethodTemplateByName), arg0, arg1)
}

// GetMsgTypeTemplate mocks base method.
func (m *MockIFullAPI) GetMsgTypeTemplate(arg0 context.Context, arg1 string) (*wallet.MsgTypeTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgTypeTemplate", arg0, arg1)
	ret0, _ := ret[0].(*wallet.MsgTypeTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMsgTypeTemplate indicates an expected call of GetMsgTypeTemplate.
func (mr *MockIFullAPIMockRecorder) GetMsgTypeTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgTypeTemplate", reflect.TypeOf((*MockIFullAPI)(nil).GetMsgTypeTemplate), arg0, arg1)
}

// GetWalletTokenInfo mocks base method.
func (m *MockIFullAPI) GetWalletTokenInfo(arg0 context.Context, arg1 string) (*wallet.GroupAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletTokenInfo", arg0, arg1)
	ret0, _ := ret[0].(*wallet.GroupAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletTokenInfo indicates an expected call of GetWalletTokenInfo.
func (mr *MockIFullAPIMockRecorder) GetWalletTokenInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletTokenInfo", reflect.TypeOf((*MockIFullAPI)(nil).GetWalletTokenInfo), arg0, arg1)
}

// GetWalletTokensByGroup mocks base method.
func (m *MockIFullAPI) GetWalletTokensByGroup(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletTokensByGroup", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletTokensByGroup indicates an expected call of GetWalletTokensByGroup.
func (mr *MockIFullAPIMockRecorder) GetWalletTokensByGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletTokensByGroup", reflect.TypeOf((*MockIFullAPI)(nil).GetWalletTokensByGroup), arg0, arg1)
}

// ListGroups mocks base method.
func (m *MockIFullAPI) ListGroups(arg0 context.Context, arg1, arg2 int) ([]*wallet.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*wallet.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockIFullAPIMockRecorder) ListGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockIFullAPI)(nil).ListGroups), arg0, arg1, arg2)
}

// ListKeyBinds mocks base method.
func (m *MockIFullAPI) ListKeyBinds(arg0 context.Context, arg1, arg2 int) ([]*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeyBinds", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeyBinds indicates an expected call of ListKeyBinds.
func (mr *MockIFullAPIMockRecorder) ListKeyBinds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeyBinds", reflect.TypeOf((*MockIFullAPI)(nil).ListKeyBinds), arg0, arg1, arg2)
}

// ListMethodTemplates mocks base method.
func (m *MockIFullAPI) ListMethodTemplates(arg0 context.Context, arg1, arg2 int) ([]*wallet.MethodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMethodTemplates", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*wallet.MethodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMethodTemplates indicates an expected call of ListMethodTemplates.
func (mr *MockIFullAPIMockRecorder) ListMethodTemplates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMethodTemplates", reflect.TypeOf((*MockIFullAPI)(nil).ListMethodTemplates), arg0, arg1, arg2)
}

// ListMsgTypeTemplates mocks base method.
func (m *MockIFullAPI) ListMsgTypeTemplates(arg0 context.Context, arg1, arg2 int) ([]*wallet.MsgTypeTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMsgTypeTemplates", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*wallet.MsgTypeTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMsgTypeTemplates indicates an expected call of ListMsgTypeTemplates.
func (mr *MockIFullAPIMockRecorder) ListMsgTypeTemplates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMsgTypeTemplates", reflect.TypeOf((*MockIFullAPI)(nil).ListMsgTypeTemplates), arg0, arg1, arg2)
}

// Lock mocks base method.
func (m *MockIFullAPI) Lock(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock.
func (mr *MockIFullAPIMockRecorder) Lock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockIFullAPI)(nil).Lock), arg0, arg1)
}

// LockState mocks base method.
func (m *MockIFullAPI) LockState(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockState", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// LockState indicates an expected call of LockState.
func (mr *MockIFullAPIMockRecorder) LockState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockState", reflect.TypeOf((*MockIFullAPI)(nil).LockState), arg0)
}

// LogList mocks base method.
func (m *MockIFullAPI) LogList(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogList", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogList indicates an expected call of LogList.
func (mr *MockIFullAPIMockRecorder) LogList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogList", reflect.TypeOf((*MockIFullAPI)(nil).LogList), arg0)
}

// LogSetLevel mocks base method.
func (m *MockIFullAPI) LogSetLevel(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogSetLevel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogSetLevel indicates an expected call of LogSetLevel.
func (mr *MockIFullAPIMockRecorder) LogSetLevel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogSetLevel", reflect.TypeOf((*MockIFullAPI)(nil).LogSetLevel), arg0, arg1, arg2)
}

// NewGroup mocks base method.
func (m *MockIFullAPI) NewGroup(arg0 context.Context, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewGroup indicates an expected call of NewGroup.
func (mr *MockIFullAPIMockRecorder) NewGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGroup", reflect.TypeOf((*MockIFullAPI)(nil).NewGroup), arg0, arg1, arg2)
}

// NewKeyBindCustom mocks base method.
func (m *MockIFullAPI) NewKeyBindCustom(arg0 context.Context, arg1 string, arg2 address.Address, arg3 []int, arg4 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewKeyBindCustom", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewKeyBindCustom indicates an expected call of NewKeyBindCustom.
func (mr *MockIFullAPIMockRecorder) NewKeyBindCustom(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewKeyBindCustom", reflect.TypeOf((*MockIFullAPI)(nil).NewKeyBindCustom), arg0, arg1, arg2, arg3, arg4)
}

// NewKeyBindFromTemplate mocks base method.
func (m *MockIFullAPI) NewKeyBindFromTemplate(arg0 context.Context, arg1 string, arg2 address.Address, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewKeyBindFromTemplate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewKeyBindFromTemplate indicates an expected call of NewKeyBindFromTemplate.
func (mr *MockIFullAPIMockRecorder) NewKeyBindFromTemplate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewKeyBindFromTemplate", reflect.TypeOf((*MockIFullAPI)(nil).NewKeyBindFromTemplate), arg0, arg1, arg2, arg3, arg4)
}

// NewMethodTemplate mocks base method.
func (m *MockIFullAPI) NewMethodTemplate(arg0 context.Context, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMethodTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewMethodTemplate indicates an expected call of NewMethodTemplate.
func (mr *MockIFullAPIMockRecorder) NewMethodTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMethodTemplate", reflect.TypeOf((*MockIFullAPI)(nil).NewMethodTemplate), arg0, arg1, arg2)
}

// NewMsgTypeTemplate mocks base method.
func (m *MockIFullAPI) NewMsgTypeTemplate(arg0 context.Context, arg1 string, arg2 []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMsgTypeTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewMsgTypeTemplate indicates an expected call of NewMsgTypeTemplate.
func (mr *MockIFullAPIMockRecorder) NewMsgTypeTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMsgTypeTemplate", reflect.TypeOf((*MockIFullAPI)(nil).NewMsgTypeTemplate), arg0, arg1, arg2)
}

// NewStToken mocks base method.
func (m *MockIFullAPI) NewStToken(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStToken indicates an expected call of NewStToken.
func (mr *MockIFullAPIMockRecorder) NewStToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStToken", reflect.TypeOf((*MockIFullAPI)(nil).NewStToken), arg0, arg1)
}

// RemoveGroup mocks base method.
func (m *MockIFullAPI) RemoveGroup(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGroup indicates an expected call of RemoveGroup.
func (mr *MockIFullAPIMockRecorder) RemoveGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroup", reflect.TypeOf((*MockIFullAPI)(nil).RemoveGroup), arg0, arg1)
}

// RemoveKeyBind mocks base method.
func (m *MockIFullAPI) RemoveKeyBind(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveKeyBind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveKeyBind indicates an expected call of RemoveKeyBind.
func (mr *MockIFullAPIMockRecorder) RemoveKeyBind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveKeyBind", reflect.TypeOf((*MockIFullAPI)(nil).RemoveKeyBind), arg0, arg1)
}

// RemoveKeyBindByAddress mocks base method.
func (m *MockIFullAPI) RemoveKeyBindByAddress(arg0 context.Context, arg1 address.Address) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveKeyBindByAddress", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveKeyBindByAddress indicates an expected call of RemoveKeyBindByAddress.
func (mr *MockIFullAPIMockRecorder) RemoveKeyBindByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveKeyBindByAddress", reflect.TypeOf((*MockIFullAPI)(nil).RemoveKeyBindByAddress), arg0, arg1)
}

// RemoveMethodFromKeyBind mocks base method.
func (m *MockIFullAPI) RemoveMethodFromKeyBind(arg0 context.Context, arg1 string, arg2 []string) (*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMethodFromKeyBind", arg0, arg1, arg2)
	ret0, _ := ret[0].(*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMethodFromKeyBind indicates an expected call of RemoveMethodFromKeyBind.
func (mr *MockIFullAPIMockRecorder) RemoveMethodFromKeyBind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMethodFromKeyBind", reflect.TypeOf((*MockIFullAPI)(nil).RemoveMethodFromKeyBind), arg0, arg1, arg2)
}

// RemoveMethodTemplate mocks base method.
func (m *MockIFullAPI) RemoveMethodTemplate(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMethodTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMethodTemplate indicates an expected call of RemoveMethodTemplate.
func (mr *MockIFullAPIMockRecorder) RemoveMethodTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMethodTemplate", reflect.TypeOf((*MockIFullAPI)(nil).RemoveMethodTemplate), arg0, arg1)
}

// RemoveMsgTypeFromKeyBind mocks base method.
func (m *MockIFullAPI) RemoveMsgTypeFromKeyBind(arg0 context.Context, arg1 string, arg2 []int) (*wallet.KeyBind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMsgTypeFromKeyBind", arg0, arg1, arg2)
	ret0, _ := ret[0].(*wallet.KeyBind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMsgTypeFromKeyBind indicates an expected call of RemoveMsgTypeFromKeyBind.
func (mr *MockIFullAPIMockRecorder) RemoveMsgTypeFromKeyBind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMsgTypeFromKeyBind", reflect.TypeOf((*MockIFullAPI)(nil).RemoveMsgTypeFromKeyBind), arg0, arg1, arg2)
}

// RemoveMsgTypeTemplate mocks base method.
func (m *MockIFullAPI) RemoveMsgTypeTemplate(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMsgTypeTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMsgTypeTemplate indicates an expected call of RemoveMsgTypeTemplate.
func (mr *MockIFullAPIMockRecorder) RemoveMsgTypeTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMsgTypeTemplate", reflect.TypeOf((*MockIFullAPI)(nil).RemoveMsgTypeTemplate), arg0, arg1)
}

// RemoveStToken mocks base method.
func (m *MockIFullAPI) RemoveStToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStToken indicates an expected call of RemoveStToken.
func (mr *MockIFullAPIMockRecorder) RemoveStToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStToken", reflect.TypeOf((*MockIFullAPI)(nil).RemoveStToken), arg0, arg1)
}

// ScopeWallet mocks base method.
func (m *MockIFullAPI) ScopeWallet(arg0 context.Context) (*wallet.AddressScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScopeWallet", arg0)
	ret0, _ := ret[0].(*wallet.AddressScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScopeWallet indicates an expected call of ScopeWallet.
func (mr *MockIFullAPIMockRecorder) ScopeWallet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScopeWallet", reflect.TypeOf((*MockIFullAPI)(nil).ScopeWallet), arg0)
}

// SetPassword mocks base method.
func (m *MockIFullAPI) SetPassword(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockIFullAPIMockRecorder) SetPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockIFullAPI)(nil).SetPassword), arg0, arg1)
}

// Unlock mocks base method.
func (m *MockIFullAPI) Unlock(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock.
func (mr *MockIFullAPIMockRecorder) Unlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockIFullAPI)(nil).Unlock), arg0, arg1)
}

// Verify mocks base method.
func (m *MockIFullAPI) Verify(arg0 context.Context, arg1 address.Address, arg2 types.MsgType, arg3 *internal.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockIFullAPIMockRecorder) Verify(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockIFullAPI)(nil).Verify), arg0, arg1, arg2, arg3)
}

// VerifyPassword mocks base method.
func (m *MockIFullAPI) VerifyPassword(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyPassword indicates an expected call of VerifyPassword.
func (mr *MockIFullAPIMockRecorder) VerifyPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPassword", reflect.TypeOf((*MockIFullAPI)(nil).VerifyPassword), arg0, arg1)
}

// Version mocks base method.
func (m *MockIFullAPI) Version(arg0 context.Context) (types.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockIFullAPIMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockIFullAPI)(nil).Version), arg0)
}

// WalletDelete mocks base method.
func (m *MockIFullAPI) WalletDelete(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalletDelete indicates an expected call of WalletDelete.
func (mr *MockIFullAPIMockRecorder) WalletDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletDelete", reflect.TypeOf((*MockIFullAPI)(nil).WalletDelete), arg0, arg1)
}

// WalletExport mocks base method.
func (m *MockIFullAPI) WalletExport(arg0 context.Context, arg1 address.Address) (*types.KeyInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletExport", arg0, arg1)
	ret0, _ := ret[0].(*types.KeyInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletExport indicates an expected call of WalletExport.
func (mr *MockIFullAPIMockRecorder) WalletExport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletExport", reflect.TypeOf((*MockIFullAPI)(nil).WalletExport), arg0, arg1)
}

// WalletHas mocks base method.
func (m *MockIFullAPI) WalletHas(arg0 context.Context, arg1 address.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletHas", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletHas indicates an expected call of WalletHas.
func (mr *MockIFullAPIMockRecorder) WalletHas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletHas", reflect.TypeOf((*MockIFullAPI)(nil).WalletHas), arg0, arg1)
}

// WalletImport mocks base method.
func (m *MockIFullAPI) WalletImport(arg0 context.Context, arg1 *types.KeyInfo) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletImport", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletImport indicates an expected call of WalletImport.
func (mr *MockIFullAPIMockRecorder) WalletImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletImport", reflect.TypeOf((*MockIFullAPI)(nil).WalletImport), arg0, arg1)
}

// WalletList mocks base method.
func (m *MockIFullAPI) WalletList(arg0 context.Context) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletList", arg0)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletList indicates an expected call of WalletList.
func (mr *MockIFullAPIMockRecorder) WalletList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletList", reflect.TypeOf((*MockIFullAPI)(nil).WalletList), arg0)
}

// WalletNew mocks base method.
func (m *MockIFullAPI) WalletNew(arg0 context.Context, arg1 types.KeyType) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletNew", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletNew indicates an expected call of WalletNew.
func (mr *MockIFullAPIMockRecorder) WalletNew(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletNew", reflect.TypeOf((*MockIFullAPI)(nil).WalletNew), arg0, arg1)
}

// WalletSign mocks base method.
func (m *MockIFullAPI) WalletSign(arg0 context.Context, arg1 address.Address, arg2 []byte, arg3 types.MsgMeta) (*crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletSign", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletSign indicates an expected call of WalletSign.
func (mr *MockIFullAPIMockRecorder) WalletSign(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletSign", reflect.TypeOf((*MockIFullAPI)(nil).WalletSign), arg0, arg1, arg2, arg3)
}
