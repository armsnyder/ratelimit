// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/asnyder/ratelimit/src/redis (interfaces: RateLimitCache,Pool,Connection,Response,TimeSource,JitterRandSource)

package mock_redis

import (
	gomock "github.com/golang/mock/gomock"
	ratelimit "github.com/asnyder/ratelimit/proto/envoy/service/ratelimit/v2"
	config "github.com/asnyder/ratelimit/src/config"
	redis "github.com/asnyder/ratelimit/src/redis"
	context "golang.org/x/net/context"
)

// Mock of RateLimitCache interface
type MockRateLimitCache struct {
	ctrl     *gomock.Controller
	recorder *_MockRateLimitCacheRecorder
}

// Recorder for MockRateLimitCache (not exported)
type _MockRateLimitCacheRecorder struct {
	mock *MockRateLimitCache
}

func NewMockRateLimitCache(ctrl *gomock.Controller) *MockRateLimitCache {
	mock := &MockRateLimitCache{ctrl: ctrl}
	mock.recorder = &_MockRateLimitCacheRecorder{mock}
	return mock
}

func (_m *MockRateLimitCache) EXPECT() *_MockRateLimitCacheRecorder {
	return _m.recorder
}

func (_m *MockRateLimitCache) DoLimit(_param0 context.Context, _param1 *ratelimit.RateLimitRequest, _param2 []*config.RateLimit) []*ratelimit.RateLimitResponse_DescriptorStatus {
	ret := _m.ctrl.Call(_m, "DoLimit", _param0, _param1, _param2)
	ret0, _ := ret[0].([]*ratelimit.RateLimitResponse_DescriptorStatus)
	return ret0
}

func (_mr *_MockRateLimitCacheRecorder) DoLimit(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoLimit", arg0, arg1, arg2)
}

// Mock of Pool interface
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *_MockPoolRecorder
}

// Recorder for MockPool (not exported)
type _MockPoolRecorder struct {
	mock *MockPool
}

func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &_MockPoolRecorder{mock}
	return mock
}

func (_m *MockPool) EXPECT() *_MockPoolRecorder {
	return _m.recorder
}

func (_m *MockPool) Get() redis.Connection {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(redis.Connection)
	return ret0
}

func (_mr *_MockPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockPool) Put(_param0 redis.Connection) {
	_m.ctrl.Call(_m, "Put", _param0)
}

func (_mr *_MockPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockConnectionRecorder
}

// Recorder for MockConnection (not exported)
type _MockConnectionRecorder struct {
	mock *MockConnection
}

func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &_MockConnectionRecorder{mock}
	return mock
}

func (_m *MockConnection) EXPECT() *_MockConnectionRecorder {
	return _m.recorder
}

func (_m *MockConnection) PipeAppend(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "PipeAppend", _s...)
}

func (_mr *_MockConnectionRecorder) PipeAppend(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PipeAppend", _s...)
}

func (_m *MockConnection) PipeResponse() redis.Response {
	ret := _m.ctrl.Call(_m, "PipeResponse")
	ret0, _ := ret[0].(redis.Response)
	return ret0
}

func (_mr *_MockConnectionRecorder) PipeResponse() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PipeResponse")
}

// Mock of Response interface
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *_MockResponseRecorder
}

// Recorder for MockResponse (not exported)
type _MockResponseRecorder struct {
	mock *MockResponse
}

func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &_MockResponseRecorder{mock}
	return mock
}

func (_m *MockResponse) EXPECT() *_MockResponseRecorder {
	return _m.recorder
}

func (_m *MockResponse) Int() int64 {
	ret := _m.ctrl.Call(_m, "Int")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockResponseRecorder) Int() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Int")
}

// Mock of TimeSource interface
type MockTimeSource struct {
	ctrl     *gomock.Controller
	recorder *_MockTimeSourceRecorder
}

// Recorder for MockTimeSource (not exported)
type _MockTimeSourceRecorder struct {
	mock *MockTimeSource
}

func NewMockTimeSource(ctrl *gomock.Controller) *MockTimeSource {
	mock := &MockTimeSource{ctrl: ctrl}
	mock.recorder = &_MockTimeSourceRecorder{mock}
	return mock
}

func (_m *MockTimeSource) EXPECT() *_MockTimeSourceRecorder {
	return _m.recorder
}

func (_m *MockTimeSource) UnixNow() int64 {
	ret := _m.ctrl.Call(_m, "UnixNow")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockTimeSourceRecorder) UnixNow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnixNow")
}

// Mock of JitterRandSource interface
type MockJitterRandSource struct {
	ctrl     *gomock.Controller
	recorder *_MockJitterRandSourceRecorder
}

// Recorder for MockJitterRandSource (not exported)
type _MockJitterRandSourceRecorder struct {
	mock *MockJitterRandSource
}

func NewMockJitterRandSource(ctrl *gomock.Controller) *MockJitterRandSource {
	mock := &MockJitterRandSource{ctrl: ctrl}
	mock.recorder = &_MockJitterRandSourceRecorder{mock}
	return mock
}

func (_m *MockJitterRandSource) EXPECT() *_MockJitterRandSourceRecorder {
	return _m.recorder
}

func (_m *MockJitterRandSource) Int63() int64 {
	ret := _m.ctrl.Call(_m, "Int63")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockJitterRandSourceRecorder) Int63() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Int63")
}

func (_m *MockJitterRandSource) Seed(_param0 int64) {
	_m.ctrl.Call(_m, "Seed", _param0)
}

func (_mr *_MockJitterRandSourceRecorder) Seed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Seed", arg0)
}
