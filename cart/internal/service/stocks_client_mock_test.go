// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package service

//go:generate minimock -i route256.ozon.ru/project/cart/internal/service.StocksClient -o stocks_client_mock_test.go -n StocksClientMock -p service

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	desc "route256.ozon.ru/project/cart/external/stocks/gen/api/orders/v1"
)

// StocksClientMock implements StocksClient
type StocksClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateOrder          func(ctx context.Context, order *desc.Order) (u1 uint64, err error)
	inspectFuncCreateOrder   func(ctx context.Context, order *desc.Order)
	afterCreateOrderCounter  uint64
	beforeCreateOrderCounter uint64
	CreateOrderMock          mStocksClientMockCreateOrder

	funcGetStocksInfo          func(ctx context.Context, SKU uint32) (u1 uint64, err error)
	inspectFuncGetStocksInfo   func(ctx context.Context, SKU uint32)
	afterGetStocksInfoCounter  uint64
	beforeGetStocksInfoCounter uint64
	GetStocksInfoMock          mStocksClientMockGetStocksInfo
}

// NewStocksClientMock returns a mock for StocksClient
func NewStocksClientMock(t minimock.Tester) *StocksClientMock {
	m := &StocksClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateOrderMock = mStocksClientMockCreateOrder{mock: m}
	m.CreateOrderMock.callArgs = []*StocksClientMockCreateOrderParams{}

	m.GetStocksInfoMock = mStocksClientMockGetStocksInfo{mock: m}
	m.GetStocksInfoMock.callArgs = []*StocksClientMockGetStocksInfoParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mStocksClientMockCreateOrder struct {
	mock               *StocksClientMock
	defaultExpectation *StocksClientMockCreateOrderExpectation
	expectations       []*StocksClientMockCreateOrderExpectation

	callArgs []*StocksClientMockCreateOrderParams
	mutex    sync.RWMutex
}

// StocksClientMockCreateOrderExpectation specifies expectation struct of the StocksClient.CreateOrder
type StocksClientMockCreateOrderExpectation struct {
	mock    *StocksClientMock
	params  *StocksClientMockCreateOrderParams
	results *StocksClientMockCreateOrderResults
	Counter uint64
}

// StocksClientMockCreateOrderParams contains parameters of the StocksClient.CreateOrder
type StocksClientMockCreateOrderParams struct {
	ctx   context.Context
	order *desc.Order
}

// StocksClientMockCreateOrderResults contains results of the StocksClient.CreateOrder
type StocksClientMockCreateOrderResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for StocksClient.CreateOrder
func (mmCreateOrder *mStocksClientMockCreateOrder) Expect(ctx context.Context, order *desc.Order) *mStocksClientMockCreateOrder {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("StocksClientMock.CreateOrder mock is already set by Set")
	}

	if mmCreateOrder.defaultExpectation == nil {
		mmCreateOrder.defaultExpectation = &StocksClientMockCreateOrderExpectation{}
	}

	mmCreateOrder.defaultExpectation.params = &StocksClientMockCreateOrderParams{ctx, order}
	for _, e := range mmCreateOrder.expectations {
		if minimock.Equal(e.params, mmCreateOrder.defaultExpectation.params) {
			mmCreateOrder.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateOrder.defaultExpectation.params)
		}
	}

	return mmCreateOrder
}

// Inspect accepts an inspector function that has same arguments as the StocksClient.CreateOrder
func (mmCreateOrder *mStocksClientMockCreateOrder) Inspect(f func(ctx context.Context, order *desc.Order)) *mStocksClientMockCreateOrder {
	if mmCreateOrder.mock.inspectFuncCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("Inspect function is already set for StocksClientMock.CreateOrder")
	}

	mmCreateOrder.mock.inspectFuncCreateOrder = f

	return mmCreateOrder
}

// Return sets up results that will be returned by StocksClient.CreateOrder
func (mmCreateOrder *mStocksClientMockCreateOrder) Return(u1 uint64, err error) *StocksClientMock {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("StocksClientMock.CreateOrder mock is already set by Set")
	}

	if mmCreateOrder.defaultExpectation == nil {
		mmCreateOrder.defaultExpectation = &StocksClientMockCreateOrderExpectation{mock: mmCreateOrder.mock}
	}
	mmCreateOrder.defaultExpectation.results = &StocksClientMockCreateOrderResults{u1, err}
	return mmCreateOrder.mock
}

// Set uses given function f to mock the StocksClient.CreateOrder method
func (mmCreateOrder *mStocksClientMockCreateOrder) Set(f func(ctx context.Context, order *desc.Order) (u1 uint64, err error)) *StocksClientMock {
	if mmCreateOrder.defaultExpectation != nil {
		mmCreateOrder.mock.t.Fatalf("Default expectation is already set for the StocksClient.CreateOrder method")
	}

	if len(mmCreateOrder.expectations) > 0 {
		mmCreateOrder.mock.t.Fatalf("Some expectations are already set for the StocksClient.CreateOrder method")
	}

	mmCreateOrder.mock.funcCreateOrder = f
	return mmCreateOrder.mock
}

// When sets expectation for the StocksClient.CreateOrder which will trigger the result defined by the following
// Then helper
func (mmCreateOrder *mStocksClientMockCreateOrder) When(ctx context.Context, order *desc.Order) *StocksClientMockCreateOrderExpectation {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("StocksClientMock.CreateOrder mock is already set by Set")
	}

	expectation := &StocksClientMockCreateOrderExpectation{
		mock:   mmCreateOrder.mock,
		params: &StocksClientMockCreateOrderParams{ctx, order},
	}
	mmCreateOrder.expectations = append(mmCreateOrder.expectations, expectation)
	return expectation
}

// Then sets up StocksClient.CreateOrder return parameters for the expectation previously defined by the When method
func (e *StocksClientMockCreateOrderExpectation) Then(u1 uint64, err error) *StocksClientMock {
	e.results = &StocksClientMockCreateOrderResults{u1, err}
	return e.mock
}

// CreateOrder implements StocksClient
func (mmCreateOrder *StocksClientMock) CreateOrder(ctx context.Context, order *desc.Order) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmCreateOrder.beforeCreateOrderCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateOrder.afterCreateOrderCounter, 1)

	if mmCreateOrder.inspectFuncCreateOrder != nil {
		mmCreateOrder.inspectFuncCreateOrder(ctx, order)
	}

	mm_params := StocksClientMockCreateOrderParams{ctx, order}

	// Record call args
	mmCreateOrder.CreateOrderMock.mutex.Lock()
	mmCreateOrder.CreateOrderMock.callArgs = append(mmCreateOrder.CreateOrderMock.callArgs, &mm_params)
	mmCreateOrder.CreateOrderMock.mutex.Unlock()

	for _, e := range mmCreateOrder.CreateOrderMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmCreateOrder.CreateOrderMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateOrder.CreateOrderMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateOrder.CreateOrderMock.defaultExpectation.params
		mm_got := StocksClientMockCreateOrderParams{ctx, order}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateOrder.t.Errorf("StocksClientMock.CreateOrder got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateOrder.CreateOrderMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateOrder.t.Fatal("No results are set for the StocksClientMock.CreateOrder")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmCreateOrder.funcCreateOrder != nil {
		return mmCreateOrder.funcCreateOrder(ctx, order)
	}
	mmCreateOrder.t.Fatalf("Unexpected call to StocksClientMock.CreateOrder. %v %v", ctx, order)
	return
}

// CreateOrderAfterCounter returns a count of finished StocksClientMock.CreateOrder invocations
func (mmCreateOrder *StocksClientMock) CreateOrderAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateOrder.afterCreateOrderCounter)
}

// CreateOrderBeforeCounter returns a count of StocksClientMock.CreateOrder invocations
func (mmCreateOrder *StocksClientMock) CreateOrderBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateOrder.beforeCreateOrderCounter)
}

// Calls returns a list of arguments used in each call to StocksClientMock.CreateOrder.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateOrder *mStocksClientMockCreateOrder) Calls() []*StocksClientMockCreateOrderParams {
	mmCreateOrder.mutex.RLock()

	argCopy := make([]*StocksClientMockCreateOrderParams, len(mmCreateOrder.callArgs))
	copy(argCopy, mmCreateOrder.callArgs)

	mmCreateOrder.mutex.RUnlock()

	return argCopy
}

// MinimockCreateOrderDone returns true if the count of the CreateOrder invocations corresponds
// the number of defined expectations
func (m *StocksClientMock) MinimockCreateOrderDone() bool {
	for _, e := range m.CreateOrderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateOrderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateOrder != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateOrderInspect logs each unmet expectation
func (m *StocksClientMock) MinimockCreateOrderInspect() {
	for _, e := range m.CreateOrderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StocksClientMock.CreateOrder with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateOrderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		if m.CreateOrderMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StocksClientMock.CreateOrder")
		} else {
			m.t.Errorf("Expected call to StocksClientMock.CreateOrder with params: %#v", *m.CreateOrderMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateOrder != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		m.t.Error("Expected call to StocksClientMock.CreateOrder")
	}
}

type mStocksClientMockGetStocksInfo struct {
	mock               *StocksClientMock
	defaultExpectation *StocksClientMockGetStocksInfoExpectation
	expectations       []*StocksClientMockGetStocksInfoExpectation

	callArgs []*StocksClientMockGetStocksInfoParams
	mutex    sync.RWMutex
}

// StocksClientMockGetStocksInfoExpectation specifies expectation struct of the StocksClient.GetStocksInfo
type StocksClientMockGetStocksInfoExpectation struct {
	mock    *StocksClientMock
	params  *StocksClientMockGetStocksInfoParams
	results *StocksClientMockGetStocksInfoResults
	Counter uint64
}

// StocksClientMockGetStocksInfoParams contains parameters of the StocksClient.GetStocksInfo
type StocksClientMockGetStocksInfoParams struct {
	ctx context.Context
	SKU uint32
}

// StocksClientMockGetStocksInfoResults contains results of the StocksClient.GetStocksInfo
type StocksClientMockGetStocksInfoResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for StocksClient.GetStocksInfo
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) Expect(ctx context.Context, SKU uint32) *mStocksClientMockGetStocksInfo {
	if mmGetStocksInfo.mock.funcGetStocksInfo != nil {
		mmGetStocksInfo.mock.t.Fatalf("StocksClientMock.GetStocksInfo mock is already set by Set")
	}

	if mmGetStocksInfo.defaultExpectation == nil {
		mmGetStocksInfo.defaultExpectation = &StocksClientMockGetStocksInfoExpectation{}
	}

	mmGetStocksInfo.defaultExpectation.params = &StocksClientMockGetStocksInfoParams{ctx, SKU}
	for _, e := range mmGetStocksInfo.expectations {
		if minimock.Equal(e.params, mmGetStocksInfo.defaultExpectation.params) {
			mmGetStocksInfo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetStocksInfo.defaultExpectation.params)
		}
	}

	return mmGetStocksInfo
}

// Inspect accepts an inspector function that has same arguments as the StocksClient.GetStocksInfo
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) Inspect(f func(ctx context.Context, SKU uint32)) *mStocksClientMockGetStocksInfo {
	if mmGetStocksInfo.mock.inspectFuncGetStocksInfo != nil {
		mmGetStocksInfo.mock.t.Fatalf("Inspect function is already set for StocksClientMock.GetStocksInfo")
	}

	mmGetStocksInfo.mock.inspectFuncGetStocksInfo = f

	return mmGetStocksInfo
}

// Return sets up results that will be returned by StocksClient.GetStocksInfo
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) Return(u1 uint64, err error) *StocksClientMock {
	if mmGetStocksInfo.mock.funcGetStocksInfo != nil {
		mmGetStocksInfo.mock.t.Fatalf("StocksClientMock.GetStocksInfo mock is already set by Set")
	}

	if mmGetStocksInfo.defaultExpectation == nil {
		mmGetStocksInfo.defaultExpectation = &StocksClientMockGetStocksInfoExpectation{mock: mmGetStocksInfo.mock}
	}
	mmGetStocksInfo.defaultExpectation.results = &StocksClientMockGetStocksInfoResults{u1, err}
	return mmGetStocksInfo.mock
}

// Set uses given function f to mock the StocksClient.GetStocksInfo method
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) Set(f func(ctx context.Context, SKU uint32) (u1 uint64, err error)) *StocksClientMock {
	if mmGetStocksInfo.defaultExpectation != nil {
		mmGetStocksInfo.mock.t.Fatalf("Default expectation is already set for the StocksClient.GetStocksInfo method")
	}

	if len(mmGetStocksInfo.expectations) > 0 {
		mmGetStocksInfo.mock.t.Fatalf("Some expectations are already set for the StocksClient.GetStocksInfo method")
	}

	mmGetStocksInfo.mock.funcGetStocksInfo = f
	return mmGetStocksInfo.mock
}

// When sets expectation for the StocksClient.GetStocksInfo which will trigger the result defined by the following
// Then helper
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) When(ctx context.Context, SKU uint32) *StocksClientMockGetStocksInfoExpectation {
	if mmGetStocksInfo.mock.funcGetStocksInfo != nil {
		mmGetStocksInfo.mock.t.Fatalf("StocksClientMock.GetStocksInfo mock is already set by Set")
	}

	expectation := &StocksClientMockGetStocksInfoExpectation{
		mock:   mmGetStocksInfo.mock,
		params: &StocksClientMockGetStocksInfoParams{ctx, SKU},
	}
	mmGetStocksInfo.expectations = append(mmGetStocksInfo.expectations, expectation)
	return expectation
}

// Then sets up StocksClient.GetStocksInfo return parameters for the expectation previously defined by the When method
func (e *StocksClientMockGetStocksInfoExpectation) Then(u1 uint64, err error) *StocksClientMock {
	e.results = &StocksClientMockGetStocksInfoResults{u1, err}
	return e.mock
}

// GetStocksInfo implements StocksClient
func (mmGetStocksInfo *StocksClientMock) GetStocksInfo(ctx context.Context, SKU uint32) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmGetStocksInfo.beforeGetStocksInfoCounter, 1)
	defer mm_atomic.AddUint64(&mmGetStocksInfo.afterGetStocksInfoCounter, 1)

	if mmGetStocksInfo.inspectFuncGetStocksInfo != nil {
		mmGetStocksInfo.inspectFuncGetStocksInfo(ctx, SKU)
	}

	mm_params := StocksClientMockGetStocksInfoParams{ctx, SKU}

	// Record call args
	mmGetStocksInfo.GetStocksInfoMock.mutex.Lock()
	mmGetStocksInfo.GetStocksInfoMock.callArgs = append(mmGetStocksInfo.GetStocksInfoMock.callArgs, &mm_params)
	mmGetStocksInfo.GetStocksInfoMock.mutex.Unlock()

	for _, e := range mmGetStocksInfo.GetStocksInfoMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetStocksInfo.GetStocksInfoMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetStocksInfo.GetStocksInfoMock.defaultExpectation.Counter, 1)
		mm_want := mmGetStocksInfo.GetStocksInfoMock.defaultExpectation.params
		mm_got := StocksClientMockGetStocksInfoParams{ctx, SKU}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetStocksInfo.t.Errorf("StocksClientMock.GetStocksInfo got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetStocksInfo.GetStocksInfoMock.defaultExpectation.results
		if mm_results == nil {
			mmGetStocksInfo.t.Fatal("No results are set for the StocksClientMock.GetStocksInfo")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetStocksInfo.funcGetStocksInfo != nil {
		return mmGetStocksInfo.funcGetStocksInfo(ctx, SKU)
	}
	mmGetStocksInfo.t.Fatalf("Unexpected call to StocksClientMock.GetStocksInfo. %v %v", ctx, SKU)
	return
}

// GetStocksInfoAfterCounter returns a count of finished StocksClientMock.GetStocksInfo invocations
func (mmGetStocksInfo *StocksClientMock) GetStocksInfoAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStocksInfo.afterGetStocksInfoCounter)
}

// GetStocksInfoBeforeCounter returns a count of StocksClientMock.GetStocksInfo invocations
func (mmGetStocksInfo *StocksClientMock) GetStocksInfoBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetStocksInfo.beforeGetStocksInfoCounter)
}

// Calls returns a list of arguments used in each call to StocksClientMock.GetStocksInfo.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetStocksInfo *mStocksClientMockGetStocksInfo) Calls() []*StocksClientMockGetStocksInfoParams {
	mmGetStocksInfo.mutex.RLock()

	argCopy := make([]*StocksClientMockGetStocksInfoParams, len(mmGetStocksInfo.callArgs))
	copy(argCopy, mmGetStocksInfo.callArgs)

	mmGetStocksInfo.mutex.RUnlock()

	return argCopy
}

// MinimockGetStocksInfoDone returns true if the count of the GetStocksInfo invocations corresponds
// the number of defined expectations
func (m *StocksClientMock) MinimockGetStocksInfoDone() bool {
	for _, e := range m.GetStocksInfoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStocksInfoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStocksInfoCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStocksInfo != nil && mm_atomic.LoadUint64(&m.afterGetStocksInfoCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetStocksInfoInspect logs each unmet expectation
func (m *StocksClientMock) MinimockGetStocksInfoInspect() {
	for _, e := range m.GetStocksInfoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StocksClientMock.GetStocksInfo with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetStocksInfoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetStocksInfoCounter) < 1 {
		if m.GetStocksInfoMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StocksClientMock.GetStocksInfo")
		} else {
			m.t.Errorf("Expected call to StocksClientMock.GetStocksInfo with params: %#v", *m.GetStocksInfoMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetStocksInfo != nil && mm_atomic.LoadUint64(&m.afterGetStocksInfoCounter) < 1 {
		m.t.Error("Expected call to StocksClientMock.GetStocksInfo")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StocksClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateOrderInspect()

			m.MinimockGetStocksInfoInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StocksClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *StocksClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateOrderDone() &&
		m.MinimockGetStocksInfoDone()
}
