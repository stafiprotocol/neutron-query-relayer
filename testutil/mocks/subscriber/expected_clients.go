// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../internal/subscriber/clients.go
//
// Generated by this command:
//
//	mockgen -source=./../../internal/subscriber/clients.go -destination ./subscriber/expected_clients.go
//

// Package mock_subscriber is a generated GoMock package.
package mock_subscriber

import (
	context "context"
	reflect "reflect"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	query "github.com/neutron-org/neutron-query-relayer/internal/subscriber/querier/client/query"
	gomock "go.uber.org/mock/gomock"
)

// MockRpcHttpClient is a mock of RpcHttpClient interface.
type MockRpcHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockRpcHttpClientMockRecorder
}

// MockRpcHttpClientMockRecorder is the mock recorder for MockRpcHttpClient.
type MockRpcHttpClientMockRecorder struct {
	mock *MockRpcHttpClient
}

// NewMockRpcHttpClient creates a new mock instance.
func NewMockRpcHttpClient(ctrl *gomock.Controller) *MockRpcHttpClient {
	mock := &MockRpcHttpClient{ctrl: ctrl}
	mock.recorder = &MockRpcHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRpcHttpClient) EXPECT() *MockRpcHttpClientMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockRpcHttpClient) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockRpcHttpClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRpcHttpClient)(nil).Start))
}

// Status mocks base method.
func (m *MockRpcHttpClient) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", ctx)
	ret0, _ := ret[0].(*coretypes.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockRpcHttpClientMockRecorder) Status(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRpcHttpClient)(nil).Status), ctx)
}

// Subscribe mocks base method.
func (m *MockRpcHttpClient) Subscribe(ctx context.Context, subscriber, query string, outCapacity ...int) (<-chan coretypes.ResultEvent, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, subscriber, query}
	for _, a := range outCapacity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(<-chan coretypes.ResultEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockRpcHttpClientMockRecorder) Subscribe(ctx, subscriber, query any, outCapacity ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subscriber, query}, outCapacity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockRpcHttpClient)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method.
func (m *MockRpcHttpClient) Unsubscribe(ctx context.Context, subscriber, query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx, subscriber, query)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockRpcHttpClientMockRecorder) Unsubscribe(ctx, subscriber, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockRpcHttpClient)(nil).Unsubscribe), ctx, subscriber, query)
}

// MockRestHttpQuery is a mock of RestHttpQuery interface.
type MockRestHttpQuery struct {
	ctrl     *gomock.Controller
	recorder *MockRestHttpQueryMockRecorder
}

// MockRestHttpQueryMockRecorder is the mock recorder for MockRestHttpQuery.
type MockRestHttpQueryMockRecorder struct {
	mock *MockRestHttpQuery
}

// NewMockRestHttpQuery creates a new mock instance.
func NewMockRestHttpQuery(ctrl *gomock.Controller) *MockRestHttpQuery {
	mock := &MockRestHttpQuery{ctrl: ctrl}
	mock.recorder = &MockRestHttpQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestHttpQuery) EXPECT() *MockRestHttpQueryMockRecorder {
	return m.recorder
}

// NeutronInterchainQueriesRegisteredQueries mocks base method.
func (m *MockRestHttpQuery) NeutronInterchainQueriesRegisteredQueries(params *query.NeutronInterchainQueriesRegisteredQueriesParams, opts ...query.ClientOption) (*query.NeutronInterchainQueriesRegisteredQueriesOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NeutronInterchainQueriesRegisteredQueries", varargs...)
	ret0, _ := ret[0].(*query.NeutronInterchainQueriesRegisteredQueriesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeutronInterchainQueriesRegisteredQueries indicates an expected call of NeutronInterchainQueriesRegisteredQueries.
func (mr *MockRestHttpQueryMockRecorder) NeutronInterchainQueriesRegisteredQueries(params any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeutronInterchainQueriesRegisteredQueries", reflect.TypeOf((*MockRestHttpQuery)(nil).NeutronInterchainQueriesRegisteredQueries), varargs...)
}

// NeutronInterchainQueriesRegisteredQuery mocks base method.
func (m *MockRestHttpQuery) NeutronInterchainQueriesRegisteredQuery(params *query.NeutronInterchainQueriesRegisteredQueryParams, opts ...query.ClientOption) (*query.NeutronInterchainQueriesRegisteredQueryOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NeutronInterchainQueriesRegisteredQuery", varargs...)
	ret0, _ := ret[0].(*query.NeutronInterchainQueriesRegisteredQueryOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeutronInterchainQueriesRegisteredQuery indicates an expected call of NeutronInterchainQueriesRegisteredQuery.
func (mr *MockRestHttpQueryMockRecorder) NeutronInterchainQueriesRegisteredQuery(params any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeutronInterchainQueriesRegisteredQuery", reflect.TypeOf((*MockRestHttpQuery)(nil).NeutronInterchainQueriesRegisteredQuery), varargs...)
}
