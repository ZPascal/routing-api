// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"database/sql"
	"sync"

	"code.cloudfoundry.org/routing-api/db"
)

type FakeClient struct {
	AddUniqueIndexStub        func(string, interface{}) error
	addUniqueIndexMutex       sync.RWMutex
	addUniqueIndexArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	addUniqueIndexReturns struct {
		result1 error
	}
	addUniqueIndexReturnsOnCall map[int]struct {
		result1 error
	}
	AutoMigrateStub        func(...interface{}) error
	autoMigrateMutex       sync.RWMutex
	autoMigrateArgsForCall []struct {
		arg1 []interface{}
	}
	autoMigrateReturns struct {
		result1 error
	}
	autoMigrateReturnsOnCall map[int]struct {
		result1 error
	}
	BeginStub        func() db.Client
	beginMutex       sync.RWMutex
	beginArgsForCall []struct {
	}
	beginReturns struct {
		result1 db.Client
	}
	beginReturnsOnCall map[int]struct {
		result1 db.Client
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
	}
	commitReturns struct {
		result1 error
	}
	commitReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(interface{}) (int64, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 interface{}
	}
	createReturns struct {
		result1 int64
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	DeleteStub        func(interface{}, ...interface{}) (int64, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 interface{}
		arg2 []interface{}
	}
	deleteReturns struct {
		result1 int64
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	DropColumnStub        func(string) error
	dropColumnMutex       sync.RWMutex
	dropColumnArgsForCall []struct {
		arg1 string
	}
	dropColumnReturns struct {
		result1 error
	}
	dropColumnReturnsOnCall map[int]struct {
		result1 error
	}
	ExecStub        func(string, ...interface{}) int64
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	execReturns struct {
		result1 int64
	}
	execReturnsOnCall map[int]struct {
		result1 int64
	}
	FindStub        func(interface{}, ...interface{}) error
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 interface{}
		arg2 []interface{}
	}
	findReturns struct {
		result1 error
	}
	findReturnsOnCall map[int]struct {
		result1 error
	}
	FirstStub        func(interface{}, ...interface{}) error
	firstMutex       sync.RWMutex
	firstArgsForCall []struct {
		arg1 interface{}
		arg2 []interface{}
	}
	firstReturns struct {
		result1 error
	}
	firstReturnsOnCall map[int]struct {
		result1 error
	}
	HasTableStub        func(interface{}) bool
	hasTableMutex       sync.RWMutex
	hasTableArgsForCall []struct {
		arg1 interface{}
	}
	hasTableReturns struct {
		result1 bool
	}
	hasTableReturnsOnCall map[int]struct {
		result1 bool
	}
	ModelStub        func(interface{}) db.Client
	modelMutex       sync.RWMutex
	modelArgsForCall []struct {
		arg1 interface{}
	}
	modelReturns struct {
		result1 db.Client
	}
	modelReturnsOnCall map[int]struct {
		result1 db.Client
	}
	RemoveIndexStub        func(string, interface{}) error
	removeIndexMutex       sync.RWMutex
	removeIndexArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	removeIndexReturns struct {
		result1 error
	}
	removeIndexReturnsOnCall map[int]struct {
		result1 error
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct {
	}
	rollbackReturns struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	RowsStub        func(string) (*sql.Rows, error)
	rowsMutex       sync.RWMutex
	rowsArgsForCall []struct {
		arg1 string
	}
	rowsReturns struct {
		result1 *sql.Rows
		result2 error
	}
	rowsReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	SaveStub        func(interface{}) (int64, error)
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		arg1 interface{}
	}
	saveReturns struct {
		result1 int64
		result2 error
	}
	saveReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	UpdateStub        func(string, interface{}) (int64, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	updateReturns struct {
		result1 int64
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	WhereStub        func(interface{}, ...interface{}) db.Client
	whereMutex       sync.RWMutex
	whereArgsForCall []struct {
		arg1 interface{}
		arg2 []interface{}
	}
	whereReturns struct {
		result1 db.Client
	}
	whereReturnsOnCall map[int]struct {
		result1 db.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) AddUniqueIndex(arg1 string, arg2 interface{}) error {
	fake.addUniqueIndexMutex.Lock()
	ret, specificReturn := fake.addUniqueIndexReturnsOnCall[len(fake.addUniqueIndexArgsForCall)]
	fake.addUniqueIndexArgsForCall = append(fake.addUniqueIndexArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("AddUniqueIndex", []interface{}{arg1, arg2})
	fake.addUniqueIndexMutex.Unlock()
	if fake.AddUniqueIndexStub != nil {
		return fake.AddUniqueIndexStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addUniqueIndexReturns
	return fakeReturns.result1
}

func (fake *FakeClient) AddUniqueIndexCallCount() int {
	fake.addUniqueIndexMutex.RLock()
	defer fake.addUniqueIndexMutex.RUnlock()
	return len(fake.addUniqueIndexArgsForCall)
}

func (fake *FakeClient) AddUniqueIndexCalls(stub func(string, interface{}) error) {
	fake.addUniqueIndexMutex.Lock()
	defer fake.addUniqueIndexMutex.Unlock()
	fake.AddUniqueIndexStub = stub
}

func (fake *FakeClient) AddUniqueIndexArgsForCall(i int) (string, interface{}) {
	fake.addUniqueIndexMutex.RLock()
	defer fake.addUniqueIndexMutex.RUnlock()
	argsForCall := fake.addUniqueIndexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) AddUniqueIndexReturns(result1 error) {
	fake.addUniqueIndexMutex.Lock()
	defer fake.addUniqueIndexMutex.Unlock()
	fake.AddUniqueIndexStub = nil
	fake.addUniqueIndexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AddUniqueIndexReturnsOnCall(i int, result1 error) {
	fake.addUniqueIndexMutex.Lock()
	defer fake.addUniqueIndexMutex.Unlock()
	fake.AddUniqueIndexStub = nil
	if fake.addUniqueIndexReturnsOnCall == nil {
		fake.addUniqueIndexReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addUniqueIndexReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AutoMigrate(arg1 ...interface{}) error {
	fake.autoMigrateMutex.Lock()
	ret, specificReturn := fake.autoMigrateReturnsOnCall[len(fake.autoMigrateArgsForCall)]
	fake.autoMigrateArgsForCall = append(fake.autoMigrateArgsForCall, struct {
		arg1 []interface{}
	}{arg1})
	fake.recordInvocation("AutoMigrate", []interface{}{arg1})
	fake.autoMigrateMutex.Unlock()
	if fake.AutoMigrateStub != nil {
		return fake.AutoMigrateStub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.autoMigrateReturns
	return fakeReturns.result1
}

func (fake *FakeClient) AutoMigrateCallCount() int {
	fake.autoMigrateMutex.RLock()
	defer fake.autoMigrateMutex.RUnlock()
	return len(fake.autoMigrateArgsForCall)
}

func (fake *FakeClient) AutoMigrateCalls(stub func(...interface{}) error) {
	fake.autoMigrateMutex.Lock()
	defer fake.autoMigrateMutex.Unlock()
	fake.AutoMigrateStub = stub
}

func (fake *FakeClient) AutoMigrateArgsForCall(i int) []interface{} {
	fake.autoMigrateMutex.RLock()
	defer fake.autoMigrateMutex.RUnlock()
	argsForCall := fake.autoMigrateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) AutoMigrateReturns(result1 error) {
	fake.autoMigrateMutex.Lock()
	defer fake.autoMigrateMutex.Unlock()
	fake.AutoMigrateStub = nil
	fake.autoMigrateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AutoMigrateReturnsOnCall(i int, result1 error) {
	fake.autoMigrateMutex.Lock()
	defer fake.autoMigrateMutex.Unlock()
	fake.AutoMigrateStub = nil
	if fake.autoMigrateReturnsOnCall == nil {
		fake.autoMigrateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.autoMigrateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Begin() db.Client {
	fake.beginMutex.Lock()
	ret, specificReturn := fake.beginReturnsOnCall[len(fake.beginArgsForCall)]
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct {
	}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.beginReturns
	return fakeReturns.result1
}

func (fake *FakeClient) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeClient) BeginCalls(stub func() db.Client) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = stub
}

func (fake *FakeClient) BeginReturns(result1 db.Client) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) BeginReturnsOnCall(i int, result1 db.Client) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	if fake.beginReturnsOnCall == nil {
		fake.beginReturnsOnCall = make(map[int]struct {
			result1 db.Client
		})
	}
	fake.beginReturnsOnCall[i] = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Commit() error {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
	}{})
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if fake.CommitStub != nil {
		return fake.CommitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commitReturns
	return fakeReturns.result1
}

func (fake *FakeClient) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeClient) CommitCalls(stub func() error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeClient) CommitReturns(result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CommitReturnsOnCall(i int, result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Create(arg1 interface{}) (int64, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeClient) CreateCalls(stub func(interface{}) (int64, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeClient) CreateArgsForCall(i int) interface{} {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) CreateReturns(result1 int64, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateReturnsOnCall(i int, result1 int64, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Delete(arg1 interface{}, arg2 ...interface{}) (int64, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 interface{}
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeClient) DeleteCalls(stub func(interface{}, ...interface{}) (int64, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeClient) DeleteArgsForCall(i int) (interface{}, []interface{}) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DeleteReturns(result1 int64, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DeleteReturnsOnCall(i int, result1 int64, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DropColumn(arg1 string) error {
	fake.dropColumnMutex.Lock()
	ret, specificReturn := fake.dropColumnReturnsOnCall[len(fake.dropColumnArgsForCall)]
	fake.dropColumnArgsForCall = append(fake.dropColumnArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DropColumn", []interface{}{arg1})
	fake.dropColumnMutex.Unlock()
	if fake.DropColumnStub != nil {
		return fake.DropColumnStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dropColumnReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DropColumnCallCount() int {
	fake.dropColumnMutex.RLock()
	defer fake.dropColumnMutex.RUnlock()
	return len(fake.dropColumnArgsForCall)
}

func (fake *FakeClient) DropColumnCalls(stub func(string) error) {
	fake.dropColumnMutex.Lock()
	defer fake.dropColumnMutex.Unlock()
	fake.DropColumnStub = stub
}

func (fake *FakeClient) DropColumnArgsForCall(i int) string {
	fake.dropColumnMutex.RLock()
	defer fake.dropColumnMutex.RUnlock()
	argsForCall := fake.dropColumnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) DropColumnReturns(result1 error) {
	fake.dropColumnMutex.Lock()
	defer fake.dropColumnMutex.Unlock()
	fake.DropColumnStub = nil
	fake.dropColumnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DropColumnReturnsOnCall(i int, result1 error) {
	fake.dropColumnMutex.Lock()
	defer fake.dropColumnMutex.Unlock()
	fake.DropColumnStub = nil
	if fake.dropColumnReturnsOnCall == nil {
		fake.dropColumnReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.dropColumnReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Exec(arg1 string, arg2 ...interface{}) int64 {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1
}

func (fake *FakeClient) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeClient) ExecCalls(stub func(string, ...interface{}) int64) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeClient) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ExecReturns(result1 int64) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeClient) ExecReturnsOnCall(i int, result1 int64) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeClient) Find(arg1 interface{}, arg2 ...interface{}) error {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 interface{}
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Find", []interface{}{arg1, arg2})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.findReturns
	return fakeReturns.result1
}

func (fake *FakeClient) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeClient) FindCalls(stub func(interface{}, ...interface{}) error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeClient) FindArgsForCall(i int) (interface{}, []interface{}) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) FindReturns(result1 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) FindReturnsOnCall(i int, result1 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) First(arg1 interface{}, arg2 ...interface{}) error {
	fake.firstMutex.Lock()
	ret, specificReturn := fake.firstReturnsOnCall[len(fake.firstArgsForCall)]
	fake.firstArgsForCall = append(fake.firstArgsForCall, struct {
		arg1 interface{}
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("First", []interface{}{arg1, arg2})
	fake.firstMutex.Unlock()
	if fake.FirstStub != nil {
		return fake.FirstStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.firstReturns
	return fakeReturns.result1
}

func (fake *FakeClient) FirstCallCount() int {
	fake.firstMutex.RLock()
	defer fake.firstMutex.RUnlock()
	return len(fake.firstArgsForCall)
}

func (fake *FakeClient) FirstCalls(stub func(interface{}, ...interface{}) error) {
	fake.firstMutex.Lock()
	defer fake.firstMutex.Unlock()
	fake.FirstStub = stub
}

func (fake *FakeClient) FirstArgsForCall(i int) (interface{}, []interface{}) {
	fake.firstMutex.RLock()
	defer fake.firstMutex.RUnlock()
	argsForCall := fake.firstArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) FirstReturns(result1 error) {
	fake.firstMutex.Lock()
	defer fake.firstMutex.Unlock()
	fake.FirstStub = nil
	fake.firstReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) FirstReturnsOnCall(i int, result1 error) {
	fake.firstMutex.Lock()
	defer fake.firstMutex.Unlock()
	fake.FirstStub = nil
	if fake.firstReturnsOnCall == nil {
		fake.firstReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.firstReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) HasTable(arg1 interface{}) bool {
	fake.hasTableMutex.Lock()
	ret, specificReturn := fake.hasTableReturnsOnCall[len(fake.hasTableArgsForCall)]
	fake.hasTableArgsForCall = append(fake.hasTableArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("HasTable", []interface{}{arg1})
	fake.hasTableMutex.Unlock()
	if fake.HasTableStub != nil {
		return fake.HasTableStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasTableReturns
	return fakeReturns.result1
}

func (fake *FakeClient) HasTableCallCount() int {
	fake.hasTableMutex.RLock()
	defer fake.hasTableMutex.RUnlock()
	return len(fake.hasTableArgsForCall)
}

func (fake *FakeClient) HasTableCalls(stub func(interface{}) bool) {
	fake.hasTableMutex.Lock()
	defer fake.hasTableMutex.Unlock()
	fake.HasTableStub = stub
}

func (fake *FakeClient) HasTableArgsForCall(i int) interface{} {
	fake.hasTableMutex.RLock()
	defer fake.hasTableMutex.RUnlock()
	argsForCall := fake.hasTableArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) HasTableReturns(result1 bool) {
	fake.hasTableMutex.Lock()
	defer fake.hasTableMutex.Unlock()
	fake.HasTableStub = nil
	fake.hasTableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) HasTableReturnsOnCall(i int, result1 bool) {
	fake.hasTableMutex.Lock()
	defer fake.hasTableMutex.Unlock()
	fake.HasTableStub = nil
	if fake.hasTableReturnsOnCall == nil {
		fake.hasTableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasTableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) Model(arg1 interface{}) db.Client {
	fake.modelMutex.Lock()
	ret, specificReturn := fake.modelReturnsOnCall[len(fake.modelArgsForCall)]
	fake.modelArgsForCall = append(fake.modelArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Model", []interface{}{arg1})
	fake.modelMutex.Unlock()
	if fake.ModelStub != nil {
		return fake.ModelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.modelReturns
	return fakeReturns.result1
}

func (fake *FakeClient) ModelCallCount() int {
	fake.modelMutex.RLock()
	defer fake.modelMutex.RUnlock()
	return len(fake.modelArgsForCall)
}

func (fake *FakeClient) ModelCalls(stub func(interface{}) db.Client) {
	fake.modelMutex.Lock()
	defer fake.modelMutex.Unlock()
	fake.ModelStub = stub
}

func (fake *FakeClient) ModelArgsForCall(i int) interface{} {
	fake.modelMutex.RLock()
	defer fake.modelMutex.RUnlock()
	argsForCall := fake.modelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) ModelReturns(result1 db.Client) {
	fake.modelMutex.Lock()
	defer fake.modelMutex.Unlock()
	fake.ModelStub = nil
	fake.modelReturns = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) ModelReturnsOnCall(i int, result1 db.Client) {
	fake.modelMutex.Lock()
	defer fake.modelMutex.Unlock()
	fake.ModelStub = nil
	if fake.modelReturnsOnCall == nil {
		fake.modelReturnsOnCall = make(map[int]struct {
			result1 db.Client
		})
	}
	fake.modelReturnsOnCall[i] = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) RemoveIndex(arg1 string, arg2 interface{}) error {
	fake.removeIndexMutex.Lock()
	ret, specificReturn := fake.removeIndexReturnsOnCall[len(fake.removeIndexArgsForCall)]
	fake.removeIndexArgsForCall = append(fake.removeIndexArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("RemoveIndex", []interface{}{arg1})
	fake.removeIndexMutex.Unlock()
	if fake.RemoveIndexStub != nil {
		return fake.RemoveIndexStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeIndexReturns
	return fakeReturns.result1
}

func (fake *FakeClient) RemoveIndexCallCount() int {
	fake.removeIndexMutex.RLock()
	defer fake.removeIndexMutex.RUnlock()
	return len(fake.removeIndexArgsForCall)
}

func (fake *FakeClient) RemoveIndexCalls(stub func(string, interface{}) error) {
	fake.removeIndexMutex.Lock()
	defer fake.removeIndexMutex.Unlock()
	fake.RemoveIndexStub = stub
}

func (fake *FakeClient) RemoveIndexArgsForCall(i int) string {
	fake.removeIndexMutex.RLock()
	defer fake.removeIndexMutex.RUnlock()
	argsForCall := fake.removeIndexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) RemoveIndexReturns(result1 error) {
	fake.removeIndexMutex.Lock()
	defer fake.removeIndexMutex.Unlock()
	fake.RemoveIndexStub = nil
	fake.removeIndexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RemoveIndexReturnsOnCall(i int, result1 error) {
	fake.removeIndexMutex.Lock()
	defer fake.removeIndexMutex.Unlock()
	fake.RemoveIndexStub = nil
	if fake.removeIndexReturnsOnCall == nil {
		fake.removeIndexReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeIndexReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Rollback() error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct {
	}{})
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if fake.RollbackStub != nil {
		return fake.RollbackStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.rollbackReturns
	return fakeReturns.result1
}

func (fake *FakeClient) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeClient) RollbackCalls(stub func() error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *FakeClient) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) RollbackReturnsOnCall(i int, result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Rows(arg1 string) (*sql.Rows, error) {
	fake.rowsMutex.Lock()
	ret, specificReturn := fake.rowsReturnsOnCall[len(fake.rowsArgsForCall)]
	fake.rowsArgsForCall = append(fake.rowsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Rows", []interface{}{arg1})
	fake.rowsMutex.Unlock()
	if fake.RowsStub != nil {
		return fake.RowsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.rowsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) RowsCallCount() int {
	fake.rowsMutex.RLock()
	defer fake.rowsMutex.RUnlock()
	return len(fake.rowsArgsForCall)
}

func (fake *FakeClient) RowsCalls(stub func(string) (*sql.Rows, error)) {
	fake.rowsMutex.Lock()
	defer fake.rowsMutex.Unlock()
	fake.RowsStub = stub
}

func (fake *FakeClient) RowsArgsForCall(i int) string {
	fake.rowsMutex.RLock()
	defer fake.rowsMutex.RUnlock()
	argsForCall := fake.rowsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) RowsReturns(result1 *sql.Rows, result2 error) {
	fake.rowsMutex.Lock()
	defer fake.rowsMutex.Unlock()
	fake.RowsStub = nil
	fake.rowsReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RowsReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.rowsMutex.Lock()
	defer fake.rowsMutex.Unlock()
	fake.RowsStub = nil
	if fake.rowsReturnsOnCall == nil {
		fake.rowsReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.rowsReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Save(arg1 interface{}) (int64, error) {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Save", []interface{}{arg1})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.saveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeClient) SaveCalls(stub func(interface{}) (int64, error)) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = stub
}

func (fake *FakeClient) SaveArgsForCall(i int) interface{} {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	argsForCall := fake.saveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) SaveReturns(result1 int64, result2 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SaveReturnsOnCall(i int, result1 int64, result2 error) {
	fake.saveMutex.Lock()
	defer fake.saveMutex.Unlock()
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Update(arg1 string, arg2 interface{}) (int64, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeClient) UpdateCalls(stub func(string, interface{}) (int64, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeClient) UpdateArgsForCall(i int) (string, interface{}) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) UpdateReturns(result1 int64, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateReturnsOnCall(i int, result1 int64, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Where(arg1 interface{}, arg2 ...interface{}) db.Client {
	fake.whereMutex.Lock()
	ret, specificReturn := fake.whereReturnsOnCall[len(fake.whereArgsForCall)]
	fake.whereArgsForCall = append(fake.whereArgsForCall, struct {
		arg1 interface{}
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Where", []interface{}{arg1, arg2})
	fake.whereMutex.Unlock()
	if fake.WhereStub != nil {
		return fake.WhereStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.whereReturns
	return fakeReturns.result1
}

func (fake *FakeClient) WhereCallCount() int {
	fake.whereMutex.RLock()
	defer fake.whereMutex.RUnlock()
	return len(fake.whereArgsForCall)
}

func (fake *FakeClient) WhereCalls(stub func(interface{}, ...interface{}) db.Client) {
	fake.whereMutex.Lock()
	defer fake.whereMutex.Unlock()
	fake.WhereStub = stub
}

func (fake *FakeClient) WhereArgsForCall(i int) (interface{}, []interface{}) {
	fake.whereMutex.RLock()
	defer fake.whereMutex.RUnlock()
	argsForCall := fake.whereArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) WhereReturns(result1 db.Client) {
	fake.whereMutex.Lock()
	defer fake.whereMutex.Unlock()
	fake.WhereStub = nil
	fake.whereReturns = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) WhereReturnsOnCall(i int, result1 db.Client) {
	fake.whereMutex.Lock()
	defer fake.whereMutex.Unlock()
	fake.WhereStub = nil
	if fake.whereReturnsOnCall == nil {
		fake.whereReturnsOnCall = make(map[int]struct {
			result1 db.Client
		})
	}
	fake.whereReturnsOnCall[i] = struct {
		result1 db.Client
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addUniqueIndexMutex.RLock()
	defer fake.addUniqueIndexMutex.RUnlock()
	fake.autoMigrateMutex.RLock()
	defer fake.autoMigrateMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.dropColumnMutex.RLock()
	defer fake.dropColumnMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.firstMutex.RLock()
	defer fake.firstMutex.RUnlock()
	fake.hasTableMutex.RLock()
	defer fake.hasTableMutex.RUnlock()
	fake.modelMutex.RLock()
	defer fake.modelMutex.RUnlock()
	fake.removeIndexMutex.RLock()
	defer fake.removeIndexMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	fake.rowsMutex.RLock()
	defer fake.rowsMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.whereMutex.RLock()
	defer fake.whereMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Client = new(FakeClient)
