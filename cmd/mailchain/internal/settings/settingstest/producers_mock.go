// Code generated by MockGen. DO NOT EDIT.
// Source: producers.go

// Package settingstest is a generated GoMock package.
package settingstest

import (
	gomock "github.com/golang/mock/gomock"
	settings "github.com/mailchain/mailchain/cmd/mailchain/internal/settings"
	output "github.com/mailchain/mailchain/cmd/mailchain/internal/settings/output"
	mailbox "github.com/mailchain/mailchain/internal/mailbox"
	nameservice "github.com/mailchain/mailchain/nameservice"
	sender "github.com/mailchain/mailchain/sender"
	reflect "reflect"
)

// MockSupporter is a mock of Supporter interface
type MockSupporter struct {
	ctrl     *gomock.Controller
	recorder *MockSupporterMockRecorder
}

// MockSupporterMockRecorder is the mock recorder for MockSupporter
type MockSupporterMockRecorder struct {
	mock *MockSupporter
}

// NewMockSupporter creates a new mock instance
func NewMockSupporter(ctrl *gomock.Controller) *MockSupporter {
	mock := &MockSupporter{ctrl: ctrl}
	mock.recorder = &MockSupporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSupporter) EXPECT() *MockSupporterMockRecorder {
	return m.recorder
}

// Supports mocks base method
func (m *MockSupporter) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockSupporterMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockSupporter)(nil).Supports))
}

// MockNetworkClient is a mock of NetworkClient interface
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// ProduceNameServiceDomain mocks base method
func (m *MockNetworkClient) ProduceNameServiceDomain(dns *settings.DomainNameServices) (nameservice.ForwardLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceNameServiceDomain", dns)
	ret0, _ := ret[0].(nameservice.ForwardLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProduceNameServiceDomain indicates an expected call of ProduceNameServiceDomain
func (mr *MockNetworkClientMockRecorder) ProduceNameServiceDomain(dns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceNameServiceDomain", reflect.TypeOf((*MockNetworkClient)(nil).ProduceNameServiceDomain), dns)
}

// ProduceNameServiceAddress mocks base method
func (m *MockNetworkClient) ProduceNameServiceAddress(ans *settings.AddressNameServices) (nameservice.ReverseLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceNameServiceAddress", ans)
	ret0, _ := ret[0].(nameservice.ReverseLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProduceNameServiceAddress indicates an expected call of ProduceNameServiceAddress
func (mr *MockNetworkClientMockRecorder) ProduceNameServiceAddress(ans interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceNameServiceAddress", reflect.TypeOf((*MockNetworkClient)(nil).ProduceNameServiceAddress), ans)
}

// ProduceSender mocks base method
func (m *MockNetworkClient) ProduceSender(senders *settings.Senders) (sender.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceSender", senders)
	ret0, _ := ret[0].(sender.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProduceSender indicates an expected call of ProduceSender
func (mr *MockNetworkClientMockRecorder) ProduceSender(senders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceSender", reflect.TypeOf((*MockNetworkClient)(nil).ProduceSender), senders)
}

// ProduceReceiver mocks base method
func (m *MockNetworkClient) ProduceReceiver(receivers *settings.Receivers) (mailbox.Receiver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceReceiver", receivers)
	ret0, _ := ret[0].(mailbox.Receiver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProduceReceiver indicates an expected call of ProduceReceiver
func (mr *MockNetworkClientMockRecorder) ProduceReceiver(receivers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceReceiver", reflect.TypeOf((*MockNetworkClient)(nil).ProduceReceiver), receivers)
}

// ProducePublicKeyFinders mocks base method
func (m *MockNetworkClient) ProducePublicKeyFinders(publicKeyFinders *settings.PublicKeyFinders) (mailbox.PubKeyFinder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducePublicKeyFinders", publicKeyFinders)
	ret0, _ := ret[0].(mailbox.PubKeyFinder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProducePublicKeyFinders indicates an expected call of ProducePublicKeyFinders
func (mr *MockNetworkClientMockRecorder) ProducePublicKeyFinders(publicKeyFinders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducePublicKeyFinders", reflect.TypeOf((*MockNetworkClient)(nil).ProducePublicKeyFinders), publicKeyFinders)
}

// Disabled mocks base method
func (m *MockNetworkClient) Disabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Disabled indicates an expected call of Disabled
func (mr *MockNetworkClientMockRecorder) Disabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disabled", reflect.TypeOf((*MockNetworkClient)(nil).Disabled))
}

// Kind mocks base method
func (m *MockNetworkClient) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind
func (mr *MockNetworkClientMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockNetworkClient)(nil).Kind))
}

// Output mocks base method
func (m *MockNetworkClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockNetworkClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockNetworkClient)(nil).Output))
}

// MockSenderClient is a mock of SenderClient interface
type MockSenderClient struct {
	ctrl     *gomock.Controller
	recorder *MockSenderClientMockRecorder
}

// MockSenderClientMockRecorder is the mock recorder for MockSenderClient
type MockSenderClientMockRecorder struct {
	mock *MockSenderClient
}

// NewMockSenderClient creates a new mock instance
func NewMockSenderClient(ctrl *gomock.Controller) *MockSenderClient {
	mock := &MockSenderClient{ctrl: ctrl}
	mock.recorder = &MockSenderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSenderClient) EXPECT() *MockSenderClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockSenderClient) Produce() (sender.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(sender.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockSenderClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockSenderClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockSenderClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockSenderClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockSenderClient)(nil).Supports))
}

// Output mocks base method
func (m *MockSenderClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockSenderClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockSenderClient)(nil).Output))
}

// MockReceiverClient is a mock of ReceiverClient interface
type MockReceiverClient struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverClientMockRecorder
}

// MockReceiverClientMockRecorder is the mock recorder for MockReceiverClient
type MockReceiverClientMockRecorder struct {
	mock *MockReceiverClient
}

// NewMockReceiverClient creates a new mock instance
func NewMockReceiverClient(ctrl *gomock.Controller) *MockReceiverClient {
	mock := &MockReceiverClient{ctrl: ctrl}
	mock.recorder = &MockReceiverClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceiverClient) EXPECT() *MockReceiverClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockReceiverClient) Produce() (mailbox.Receiver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(mailbox.Receiver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockReceiverClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockReceiverClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockReceiverClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockReceiverClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockReceiverClient)(nil).Supports))
}

// Output mocks base method
func (m *MockReceiverClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockReceiverClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockReceiverClient)(nil).Output))
}

// MockPublicKeyFinderClient is a mock of PublicKeyFinderClient interface
type MockPublicKeyFinderClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicKeyFinderClientMockRecorder
}

// MockPublicKeyFinderClientMockRecorder is the mock recorder for MockPublicKeyFinderClient
type MockPublicKeyFinderClientMockRecorder struct {
	mock *MockPublicKeyFinderClient
}

// NewMockPublicKeyFinderClient creates a new mock instance
func NewMockPublicKeyFinderClient(ctrl *gomock.Controller) *MockPublicKeyFinderClient {
	mock := &MockPublicKeyFinderClient{ctrl: ctrl}
	mock.recorder = &MockPublicKeyFinderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicKeyFinderClient) EXPECT() *MockPublicKeyFinderClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockPublicKeyFinderClient) Produce() (mailbox.PubKeyFinder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(mailbox.PubKeyFinder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockPublicKeyFinderClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockPublicKeyFinderClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockPublicKeyFinderClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockPublicKeyFinderClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockPublicKeyFinderClient)(nil).Supports))
}

// Output mocks base method
func (m *MockPublicKeyFinderClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockPublicKeyFinderClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockPublicKeyFinderClient)(nil).Output))
}

// MockNameServiceAddressClient is a mock of NameServiceAddressClient interface
type MockNameServiceAddressClient struct {
	ctrl     *gomock.Controller
	recorder *MockNameServiceAddressClientMockRecorder
}

// MockNameServiceAddressClientMockRecorder is the mock recorder for MockNameServiceAddressClient
type MockNameServiceAddressClientMockRecorder struct {
	mock *MockNameServiceAddressClient
}

// NewMockNameServiceAddressClient creates a new mock instance
func NewMockNameServiceAddressClient(ctrl *gomock.Controller) *MockNameServiceAddressClient {
	mock := &MockNameServiceAddressClient{ctrl: ctrl}
	mock.recorder = &MockNameServiceAddressClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNameServiceAddressClient) EXPECT() *MockNameServiceAddressClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockNameServiceAddressClient) Produce() (nameservice.ReverseLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(nameservice.ReverseLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockNameServiceAddressClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockNameServiceAddressClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockNameServiceAddressClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockNameServiceAddressClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockNameServiceAddressClient)(nil).Supports))
}

// Output mocks base method
func (m *MockNameServiceAddressClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockNameServiceAddressClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockNameServiceAddressClient)(nil).Output))
}

// MockNameServiceDomainClient is a mock of NameServiceDomainClient interface
type MockNameServiceDomainClient struct {
	ctrl     *gomock.Controller
	recorder *MockNameServiceDomainClientMockRecorder
}

// MockNameServiceDomainClientMockRecorder is the mock recorder for MockNameServiceDomainClient
type MockNameServiceDomainClientMockRecorder struct {
	mock *MockNameServiceDomainClient
}

// NewMockNameServiceDomainClient creates a new mock instance
func NewMockNameServiceDomainClient(ctrl *gomock.Controller) *MockNameServiceDomainClient {
	mock := &MockNameServiceDomainClient{ctrl: ctrl}
	mock.recorder = &MockNameServiceDomainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNameServiceDomainClient) EXPECT() *MockNameServiceDomainClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockNameServiceDomainClient) Produce() (nameservice.ForwardLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(nameservice.ForwardLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockNameServiceDomainClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockNameServiceDomainClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockNameServiceDomainClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockNameServiceDomainClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockNameServiceDomainClient)(nil).Supports))
}

// Output mocks base method
func (m *MockNameServiceDomainClient) Output() output.Element {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].(output.Element)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockNameServiceDomainClientMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockNameServiceDomainClient)(nil).Output))
}

// MockSentClient is a mock of SentClient interface
type MockSentClient struct {
	ctrl     *gomock.Controller
	recorder *MockSentClientMockRecorder
}

// MockSentClientMockRecorder is the mock recorder for MockSentClient
type MockSentClientMockRecorder struct {
	mock *MockSentClient
}

// NewMockSentClient creates a new mock instance
func NewMockSentClient(ctrl *gomock.Controller) *MockSentClient {
	mock := &MockSentClient{ctrl: ctrl}
	mock.recorder = &MockSentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSentClient) EXPECT() *MockSentClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockSentClient) Produce(client string) (sender.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", client)
	ret0, _ := ret[0].(sender.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockSentClientMockRecorder) Produce(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockSentClient)(nil).Produce), client)
}
