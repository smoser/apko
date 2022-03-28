// Code generated by counterfeiter. DO NOT EDIT.
package buildfakes

import (
	"sync"

	"chainguard.dev/apko/pkg/build"
	"chainguard.dev/apko/pkg/build/types"
	"chainguard.dev/apko/pkg/exec"
	"chainguard.dev/apko/pkg/s6"
)

type FakeBuildImplementation struct {
	BuildImageStub        func(*build.Options, *types.ImageConfiguration, *exec.Executor, *s6.Context) error
	buildImageMutex       sync.RWMutex
	buildImageArgsForCall []struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
		arg3 *exec.Executor
		arg4 *s6.Context
	}
	buildImageReturns struct {
		result1 error
	}
	buildImageReturnsOnCall map[int]struct {
		result1 error
	}
	BuildTarballStub        func(*build.Options) (string, error)
	buildTarballMutex       sync.RWMutex
	buildTarballArgsForCall []struct {
		arg1 *build.Options
	}
	buildTarballReturns struct {
		result1 string
		result2 error
	}
	buildTarballReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	FixateApkWorldStub        func(*build.Options, *exec.Executor) error
	fixateApkWorldMutex       sync.RWMutex
	fixateApkWorldArgsForCall []struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}
	fixateApkWorldReturns struct {
		result1 error
	}
	fixateApkWorldReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateSBOMStub        func(*build.Options) error
	generateSBOMMutex       sync.RWMutex
	generateSBOMArgsForCall []struct {
		arg1 *build.Options
	}
	generateSBOMReturns struct {
		result1 error
	}
	generateSBOMReturnsOnCall map[int]struct {
		result1 error
	}
	InitApkDBStub        func(*build.Options, *exec.Executor) error
	initApkDBMutex       sync.RWMutex
	initApkDBArgsForCall []struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}
	initApkDBReturns struct {
		result1 error
	}
	initApkDBReturnsOnCall map[int]struct {
		result1 error
	}
	InitApkKeyringStub        func(*build.Options, *types.ImageConfiguration) error
	initApkKeyringMutex       sync.RWMutex
	initApkKeyringArgsForCall []struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}
	initApkKeyringReturns struct {
		result1 error
	}
	initApkKeyringReturnsOnCall map[int]struct {
		result1 error
	}
	InitApkWorldStub        func(*build.Options, *types.ImageConfiguration) error
	initApkWorldMutex       sync.RWMutex
	initApkWorldArgsForCall []struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}
	initApkWorldReturns struct {
		result1 error
	}
	initApkWorldReturnsOnCall map[int]struct {
		result1 error
	}
	InitializeApkStub        func(*build.Options, *types.ImageConfiguration, *exec.Executor) error
	initializeApkMutex       sync.RWMutex
	initializeApkArgsForCall []struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
		arg3 *exec.Executor
	}
	initializeApkReturns struct {
		result1 error
	}
	initializeApkReturnsOnCall map[int]struct {
		result1 error
	}
	InstallBusyboxSymlinksStub        func(*build.Options, *exec.Executor) error
	installBusyboxSymlinksMutex       sync.RWMutex
	installBusyboxSymlinksArgsForCall []struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}
	installBusyboxSymlinksReturns struct {
		result1 error
	}
	installBusyboxSymlinksReturnsOnCall map[int]struct {
		result1 error
	}
	MutateAccountsStub        func(*build.Options, *types.ImageConfiguration) error
	mutateAccountsMutex       sync.RWMutex
	mutateAccountsArgsForCall []struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}
	mutateAccountsReturns struct {
		result1 error
	}
	mutateAccountsReturnsOnCall map[int]struct {
		result1 error
	}
	NormalizeApkScriptsTarStub        func(*build.Options) error
	normalizeApkScriptsTarMutex       sync.RWMutex
	normalizeApkScriptsTarArgsForCall []struct {
		arg1 *build.Options
	}
	normalizeApkScriptsTarReturns struct {
		result1 error
	}
	normalizeApkScriptsTarReturnsOnCall map[int]struct {
		result1 error
	}
	RefreshStub        func(*build.Options) (*s6.Context, *exec.Executor, error)
	refreshMutex       sync.RWMutex
	refreshArgsForCall []struct {
		arg1 *build.Options
	}
	refreshReturns struct {
		result1 *s6.Context
		result2 *exec.Executor
		result3 error
	}
	refreshReturnsOnCall map[int]struct {
		result1 *s6.Context
		result2 *exec.Executor
		result3 error
	}
	ValidateImageConfigurationStub        func(*types.ImageConfiguration) error
	validateImageConfigurationMutex       sync.RWMutex
	validateImageConfigurationArgsForCall []struct {
		arg1 *types.ImageConfiguration
	}
	validateImageConfigurationReturns struct {
		result1 error
	}
	validateImageConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	WriteSupervisionTreeStub        func(*s6.Context, *types.ImageConfiguration) error
	writeSupervisionTreeMutex       sync.RWMutex
	writeSupervisionTreeArgsForCall []struct {
		arg1 *s6.Context
		arg2 *types.ImageConfiguration
	}
	writeSupervisionTreeReturns struct {
		result1 error
	}
	writeSupervisionTreeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildImplementation) BuildImage(arg1 *build.Options, arg2 *types.ImageConfiguration, arg3 *exec.Executor, arg4 *s6.Context) error {
	fake.buildImageMutex.Lock()
	ret, specificReturn := fake.buildImageReturnsOnCall[len(fake.buildImageArgsForCall)]
	fake.buildImageArgsForCall = append(fake.buildImageArgsForCall, struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
		arg3 *exec.Executor
		arg4 *s6.Context
	}{arg1, arg2, arg3, arg4})
	stub := fake.BuildImageStub
	fakeReturns := fake.buildImageReturns
	fake.recordInvocation("BuildImage", []interface{}{arg1, arg2, arg3, arg4})
	fake.buildImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) BuildImageCallCount() int {
	fake.buildImageMutex.RLock()
	defer fake.buildImageMutex.RUnlock()
	return len(fake.buildImageArgsForCall)
}

func (fake *FakeBuildImplementation) BuildImageCalls(stub func(*build.Options, *types.ImageConfiguration, *exec.Executor, *s6.Context) error) {
	fake.buildImageMutex.Lock()
	defer fake.buildImageMutex.Unlock()
	fake.BuildImageStub = stub
}

func (fake *FakeBuildImplementation) BuildImageArgsForCall(i int) (*build.Options, *types.ImageConfiguration, *exec.Executor, *s6.Context) {
	fake.buildImageMutex.RLock()
	defer fake.buildImageMutex.RUnlock()
	argsForCall := fake.buildImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBuildImplementation) BuildImageReturns(result1 error) {
	fake.buildImageMutex.Lock()
	defer fake.buildImageMutex.Unlock()
	fake.BuildImageStub = nil
	fake.buildImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) BuildImageReturnsOnCall(i int, result1 error) {
	fake.buildImageMutex.Lock()
	defer fake.buildImageMutex.Unlock()
	fake.BuildImageStub = nil
	if fake.buildImageReturnsOnCall == nil {
		fake.buildImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.buildImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) BuildTarball(arg1 *build.Options) (string, error) {
	fake.buildTarballMutex.Lock()
	ret, specificReturn := fake.buildTarballReturnsOnCall[len(fake.buildTarballArgsForCall)]
	fake.buildTarballArgsForCall = append(fake.buildTarballArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.BuildTarballStub
	fakeReturns := fake.buildTarballReturns
	fake.recordInvocation("BuildTarball", []interface{}{arg1})
	fake.buildTarballMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildImplementation) BuildTarballCallCount() int {
	fake.buildTarballMutex.RLock()
	defer fake.buildTarballMutex.RUnlock()
	return len(fake.buildTarballArgsForCall)
}

func (fake *FakeBuildImplementation) BuildTarballCalls(stub func(*build.Options) (string, error)) {
	fake.buildTarballMutex.Lock()
	defer fake.buildTarballMutex.Unlock()
	fake.BuildTarballStub = stub
}

func (fake *FakeBuildImplementation) BuildTarballArgsForCall(i int) *build.Options {
	fake.buildTarballMutex.RLock()
	defer fake.buildTarballMutex.RUnlock()
	argsForCall := fake.buildTarballArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildImplementation) BuildTarballReturns(result1 string, result2 error) {
	fake.buildTarballMutex.Lock()
	defer fake.buildTarballMutex.Unlock()
	fake.BuildTarballStub = nil
	fake.buildTarballReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildImplementation) BuildTarballReturnsOnCall(i int, result1 string, result2 error) {
	fake.buildTarballMutex.Lock()
	defer fake.buildTarballMutex.Unlock()
	fake.BuildTarballStub = nil
	if fake.buildTarballReturnsOnCall == nil {
		fake.buildTarballReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.buildTarballReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildImplementation) FixateApkWorld(arg1 *build.Options, arg2 *exec.Executor) error {
	fake.fixateApkWorldMutex.Lock()
	ret, specificReturn := fake.fixateApkWorldReturnsOnCall[len(fake.fixateApkWorldArgsForCall)]
	fake.fixateApkWorldArgsForCall = append(fake.fixateApkWorldArgsForCall, struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}{arg1, arg2})
	stub := fake.FixateApkWorldStub
	fakeReturns := fake.fixateApkWorldReturns
	fake.recordInvocation("FixateApkWorld", []interface{}{arg1, arg2})
	fake.fixateApkWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) FixateApkWorldCallCount() int {
	fake.fixateApkWorldMutex.RLock()
	defer fake.fixateApkWorldMutex.RUnlock()
	return len(fake.fixateApkWorldArgsForCall)
}

func (fake *FakeBuildImplementation) FixateApkWorldCalls(stub func(*build.Options, *exec.Executor) error) {
	fake.fixateApkWorldMutex.Lock()
	defer fake.fixateApkWorldMutex.Unlock()
	fake.FixateApkWorldStub = stub
}

func (fake *FakeBuildImplementation) FixateApkWorldArgsForCall(i int) (*build.Options, *exec.Executor) {
	fake.fixateApkWorldMutex.RLock()
	defer fake.fixateApkWorldMutex.RUnlock()
	argsForCall := fake.fixateApkWorldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) FixateApkWorldReturns(result1 error) {
	fake.fixateApkWorldMutex.Lock()
	defer fake.fixateApkWorldMutex.Unlock()
	fake.FixateApkWorldStub = nil
	fake.fixateApkWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) FixateApkWorldReturnsOnCall(i int, result1 error) {
	fake.fixateApkWorldMutex.Lock()
	defer fake.fixateApkWorldMutex.Unlock()
	fake.FixateApkWorldStub = nil
	if fake.fixateApkWorldReturnsOnCall == nil {
		fake.fixateApkWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.fixateApkWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) GenerateSBOM(arg1 *build.Options) error {
	fake.generateSBOMMutex.Lock()
	ret, specificReturn := fake.generateSBOMReturnsOnCall[len(fake.generateSBOMArgsForCall)]
	fake.generateSBOMArgsForCall = append(fake.generateSBOMArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.GenerateSBOMStub
	fakeReturns := fake.generateSBOMReturns
	fake.recordInvocation("GenerateSBOM", []interface{}{arg1})
	fake.generateSBOMMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) GenerateSBOMCallCount() int {
	fake.generateSBOMMutex.RLock()
	defer fake.generateSBOMMutex.RUnlock()
	return len(fake.generateSBOMArgsForCall)
}

func (fake *FakeBuildImplementation) GenerateSBOMCalls(stub func(*build.Options) error) {
	fake.generateSBOMMutex.Lock()
	defer fake.generateSBOMMutex.Unlock()
	fake.GenerateSBOMStub = stub
}

func (fake *FakeBuildImplementation) GenerateSBOMArgsForCall(i int) *build.Options {
	fake.generateSBOMMutex.RLock()
	defer fake.generateSBOMMutex.RUnlock()
	argsForCall := fake.generateSBOMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildImplementation) GenerateSBOMReturns(result1 error) {
	fake.generateSBOMMutex.Lock()
	defer fake.generateSBOMMutex.Unlock()
	fake.GenerateSBOMStub = nil
	fake.generateSBOMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) GenerateSBOMReturnsOnCall(i int, result1 error) {
	fake.generateSBOMMutex.Lock()
	defer fake.generateSBOMMutex.Unlock()
	fake.GenerateSBOMStub = nil
	if fake.generateSBOMReturnsOnCall == nil {
		fake.generateSBOMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateSBOMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkDB(arg1 *build.Options, arg2 *exec.Executor) error {
	fake.initApkDBMutex.Lock()
	ret, specificReturn := fake.initApkDBReturnsOnCall[len(fake.initApkDBArgsForCall)]
	fake.initApkDBArgsForCall = append(fake.initApkDBArgsForCall, struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}{arg1, arg2})
	stub := fake.InitApkDBStub
	fakeReturns := fake.initApkDBReturns
	fake.recordInvocation("InitApkDB", []interface{}{arg1, arg2})
	fake.initApkDBMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) InitApkDBCallCount() int {
	fake.initApkDBMutex.RLock()
	defer fake.initApkDBMutex.RUnlock()
	return len(fake.initApkDBArgsForCall)
}

func (fake *FakeBuildImplementation) InitApkDBCalls(stub func(*build.Options, *exec.Executor) error) {
	fake.initApkDBMutex.Lock()
	defer fake.initApkDBMutex.Unlock()
	fake.InitApkDBStub = stub
}

func (fake *FakeBuildImplementation) InitApkDBArgsForCall(i int) (*build.Options, *exec.Executor) {
	fake.initApkDBMutex.RLock()
	defer fake.initApkDBMutex.RUnlock()
	argsForCall := fake.initApkDBArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) InitApkDBReturns(result1 error) {
	fake.initApkDBMutex.Lock()
	defer fake.initApkDBMutex.Unlock()
	fake.InitApkDBStub = nil
	fake.initApkDBReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkDBReturnsOnCall(i int, result1 error) {
	fake.initApkDBMutex.Lock()
	defer fake.initApkDBMutex.Unlock()
	fake.InitApkDBStub = nil
	if fake.initApkDBReturnsOnCall == nil {
		fake.initApkDBReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initApkDBReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkKeyring(arg1 *build.Options, arg2 *types.ImageConfiguration) error {
	fake.initApkKeyringMutex.Lock()
	ret, specificReturn := fake.initApkKeyringReturnsOnCall[len(fake.initApkKeyringArgsForCall)]
	fake.initApkKeyringArgsForCall = append(fake.initApkKeyringArgsForCall, struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.InitApkKeyringStub
	fakeReturns := fake.initApkKeyringReturns
	fake.recordInvocation("InitApkKeyring", []interface{}{arg1, arg2})
	fake.initApkKeyringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) InitApkKeyringCallCount() int {
	fake.initApkKeyringMutex.RLock()
	defer fake.initApkKeyringMutex.RUnlock()
	return len(fake.initApkKeyringArgsForCall)
}

func (fake *FakeBuildImplementation) InitApkKeyringCalls(stub func(*build.Options, *types.ImageConfiguration) error) {
	fake.initApkKeyringMutex.Lock()
	defer fake.initApkKeyringMutex.Unlock()
	fake.InitApkKeyringStub = stub
}

func (fake *FakeBuildImplementation) InitApkKeyringArgsForCall(i int) (*build.Options, *types.ImageConfiguration) {
	fake.initApkKeyringMutex.RLock()
	defer fake.initApkKeyringMutex.RUnlock()
	argsForCall := fake.initApkKeyringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) InitApkKeyringReturns(result1 error) {
	fake.initApkKeyringMutex.Lock()
	defer fake.initApkKeyringMutex.Unlock()
	fake.InitApkKeyringStub = nil
	fake.initApkKeyringReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkKeyringReturnsOnCall(i int, result1 error) {
	fake.initApkKeyringMutex.Lock()
	defer fake.initApkKeyringMutex.Unlock()
	fake.InitApkKeyringStub = nil
	if fake.initApkKeyringReturnsOnCall == nil {
		fake.initApkKeyringReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initApkKeyringReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkWorld(arg1 *build.Options, arg2 *types.ImageConfiguration) error {
	fake.initApkWorldMutex.Lock()
	ret, specificReturn := fake.initApkWorldReturnsOnCall[len(fake.initApkWorldArgsForCall)]
	fake.initApkWorldArgsForCall = append(fake.initApkWorldArgsForCall, struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.InitApkWorldStub
	fakeReturns := fake.initApkWorldReturns
	fake.recordInvocation("InitApkWorld", []interface{}{arg1, arg2})
	fake.initApkWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) InitApkWorldCallCount() int {
	fake.initApkWorldMutex.RLock()
	defer fake.initApkWorldMutex.RUnlock()
	return len(fake.initApkWorldArgsForCall)
}

func (fake *FakeBuildImplementation) InitApkWorldCalls(stub func(*build.Options, *types.ImageConfiguration) error) {
	fake.initApkWorldMutex.Lock()
	defer fake.initApkWorldMutex.Unlock()
	fake.InitApkWorldStub = stub
}

func (fake *FakeBuildImplementation) InitApkWorldArgsForCall(i int) (*build.Options, *types.ImageConfiguration) {
	fake.initApkWorldMutex.RLock()
	defer fake.initApkWorldMutex.RUnlock()
	argsForCall := fake.initApkWorldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) InitApkWorldReturns(result1 error) {
	fake.initApkWorldMutex.Lock()
	defer fake.initApkWorldMutex.Unlock()
	fake.InitApkWorldStub = nil
	fake.initApkWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitApkWorldReturnsOnCall(i int, result1 error) {
	fake.initApkWorldMutex.Lock()
	defer fake.initApkWorldMutex.Unlock()
	fake.InitApkWorldStub = nil
	if fake.initApkWorldReturnsOnCall == nil {
		fake.initApkWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initApkWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitializeApk(arg1 *build.Options, arg2 *types.ImageConfiguration, arg3 *exec.Executor) error {
	fake.initializeApkMutex.Lock()
	ret, specificReturn := fake.initializeApkReturnsOnCall[len(fake.initializeApkArgsForCall)]
	fake.initializeApkArgsForCall = append(fake.initializeApkArgsForCall, struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
		arg3 *exec.Executor
	}{arg1, arg2, arg3})
	stub := fake.InitializeApkStub
	fakeReturns := fake.initializeApkReturns
	fake.recordInvocation("InitializeApk", []interface{}{arg1, arg2, arg3})
	fake.initializeApkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) InitializeApkCallCount() int {
	fake.initializeApkMutex.RLock()
	defer fake.initializeApkMutex.RUnlock()
	return len(fake.initializeApkArgsForCall)
}

func (fake *FakeBuildImplementation) InitializeApkCalls(stub func(*build.Options, *types.ImageConfiguration, *exec.Executor) error) {
	fake.initializeApkMutex.Lock()
	defer fake.initializeApkMutex.Unlock()
	fake.InitializeApkStub = stub
}

func (fake *FakeBuildImplementation) InitializeApkArgsForCall(i int) (*build.Options, *types.ImageConfiguration, *exec.Executor) {
	fake.initializeApkMutex.RLock()
	defer fake.initializeApkMutex.RUnlock()
	argsForCall := fake.initializeApkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBuildImplementation) InitializeApkReturns(result1 error) {
	fake.initializeApkMutex.Lock()
	defer fake.initializeApkMutex.Unlock()
	fake.InitializeApkStub = nil
	fake.initializeApkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InitializeApkReturnsOnCall(i int, result1 error) {
	fake.initializeApkMutex.Lock()
	defer fake.initializeApkMutex.Unlock()
	fake.InitializeApkStub = nil
	if fake.initializeApkReturnsOnCall == nil {
		fake.initializeApkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeApkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinks(arg1 *build.Options, arg2 *exec.Executor) error {
	fake.installBusyboxSymlinksMutex.Lock()
	ret, specificReturn := fake.installBusyboxSymlinksReturnsOnCall[len(fake.installBusyboxSymlinksArgsForCall)]
	fake.installBusyboxSymlinksArgsForCall = append(fake.installBusyboxSymlinksArgsForCall, struct {
		arg1 *build.Options
		arg2 *exec.Executor
	}{arg1, arg2})
	stub := fake.InstallBusyboxSymlinksStub
	fakeReturns := fake.installBusyboxSymlinksReturns
	fake.recordInvocation("InstallBusyboxSymlinks", []interface{}{arg1, arg2})
	fake.installBusyboxSymlinksMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinksCallCount() int {
	fake.installBusyboxSymlinksMutex.RLock()
	defer fake.installBusyboxSymlinksMutex.RUnlock()
	return len(fake.installBusyboxSymlinksArgsForCall)
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinksCalls(stub func(*build.Options, *exec.Executor) error) {
	fake.installBusyboxSymlinksMutex.Lock()
	defer fake.installBusyboxSymlinksMutex.Unlock()
	fake.InstallBusyboxSymlinksStub = stub
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinksArgsForCall(i int) (*build.Options, *exec.Executor) {
	fake.installBusyboxSymlinksMutex.RLock()
	defer fake.installBusyboxSymlinksMutex.RUnlock()
	argsForCall := fake.installBusyboxSymlinksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinksReturns(result1 error) {
	fake.installBusyboxSymlinksMutex.Lock()
	defer fake.installBusyboxSymlinksMutex.Unlock()
	fake.InstallBusyboxSymlinksStub = nil
	fake.installBusyboxSymlinksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) InstallBusyboxSymlinksReturnsOnCall(i int, result1 error) {
	fake.installBusyboxSymlinksMutex.Lock()
	defer fake.installBusyboxSymlinksMutex.Unlock()
	fake.InstallBusyboxSymlinksStub = nil
	if fake.installBusyboxSymlinksReturnsOnCall == nil {
		fake.installBusyboxSymlinksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installBusyboxSymlinksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) MutateAccounts(arg1 *build.Options, arg2 *types.ImageConfiguration) error {
	fake.mutateAccountsMutex.Lock()
	ret, specificReturn := fake.mutateAccountsReturnsOnCall[len(fake.mutateAccountsArgsForCall)]
	fake.mutateAccountsArgsForCall = append(fake.mutateAccountsArgsForCall, struct {
		arg1 *build.Options
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.MutateAccountsStub
	fakeReturns := fake.mutateAccountsReturns
	fake.recordInvocation("MutateAccounts", []interface{}{arg1, arg2})
	fake.mutateAccountsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) MutateAccountsCallCount() int {
	fake.mutateAccountsMutex.RLock()
	defer fake.mutateAccountsMutex.RUnlock()
	return len(fake.mutateAccountsArgsForCall)
}

func (fake *FakeBuildImplementation) MutateAccountsCalls(stub func(*build.Options, *types.ImageConfiguration) error) {
	fake.mutateAccountsMutex.Lock()
	defer fake.mutateAccountsMutex.Unlock()
	fake.MutateAccountsStub = stub
}

func (fake *FakeBuildImplementation) MutateAccountsArgsForCall(i int) (*build.Options, *types.ImageConfiguration) {
	fake.mutateAccountsMutex.RLock()
	defer fake.mutateAccountsMutex.RUnlock()
	argsForCall := fake.mutateAccountsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) MutateAccountsReturns(result1 error) {
	fake.mutateAccountsMutex.Lock()
	defer fake.mutateAccountsMutex.Unlock()
	fake.MutateAccountsStub = nil
	fake.mutateAccountsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) MutateAccountsReturnsOnCall(i int, result1 error) {
	fake.mutateAccountsMutex.Lock()
	defer fake.mutateAccountsMutex.Unlock()
	fake.MutateAccountsStub = nil
	if fake.mutateAccountsReturnsOnCall == nil {
		fake.mutateAccountsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mutateAccountsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTar(arg1 *build.Options) error {
	fake.normalizeApkScriptsTarMutex.Lock()
	ret, specificReturn := fake.normalizeApkScriptsTarReturnsOnCall[len(fake.normalizeApkScriptsTarArgsForCall)]
	fake.normalizeApkScriptsTarArgsForCall = append(fake.normalizeApkScriptsTarArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.NormalizeApkScriptsTarStub
	fakeReturns := fake.normalizeApkScriptsTarReturns
	fake.recordInvocation("NormalizeApkScriptsTar", []interface{}{arg1})
	fake.normalizeApkScriptsTarMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTarCallCount() int {
	fake.normalizeApkScriptsTarMutex.RLock()
	defer fake.normalizeApkScriptsTarMutex.RUnlock()
	return len(fake.normalizeApkScriptsTarArgsForCall)
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTarCalls(stub func(*build.Options) error) {
	fake.normalizeApkScriptsTarMutex.Lock()
	defer fake.normalizeApkScriptsTarMutex.Unlock()
	fake.NormalizeApkScriptsTarStub = stub
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTarArgsForCall(i int) *build.Options {
	fake.normalizeApkScriptsTarMutex.RLock()
	defer fake.normalizeApkScriptsTarMutex.RUnlock()
	argsForCall := fake.normalizeApkScriptsTarArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTarReturns(result1 error) {
	fake.normalizeApkScriptsTarMutex.Lock()
	defer fake.normalizeApkScriptsTarMutex.Unlock()
	fake.NormalizeApkScriptsTarStub = nil
	fake.normalizeApkScriptsTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) NormalizeApkScriptsTarReturnsOnCall(i int, result1 error) {
	fake.normalizeApkScriptsTarMutex.Lock()
	defer fake.normalizeApkScriptsTarMutex.Unlock()
	fake.NormalizeApkScriptsTarStub = nil
	if fake.normalizeApkScriptsTarReturnsOnCall == nil {
		fake.normalizeApkScriptsTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.normalizeApkScriptsTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) Refresh(arg1 *build.Options) (*s6.Context, *exec.Executor, error) {
	fake.refreshMutex.Lock()
	ret, specificReturn := fake.refreshReturnsOnCall[len(fake.refreshArgsForCall)]
	fake.refreshArgsForCall = append(fake.refreshArgsForCall, struct {
		arg1 *build.Options
	}{arg1})
	stub := fake.RefreshStub
	fakeReturns := fake.refreshReturns
	fake.recordInvocation("Refresh", []interface{}{arg1})
	fake.refreshMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBuildImplementation) RefreshCallCount() int {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return len(fake.refreshArgsForCall)
}

func (fake *FakeBuildImplementation) RefreshCalls(stub func(*build.Options) (*s6.Context, *exec.Executor, error)) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = stub
}

func (fake *FakeBuildImplementation) RefreshArgsForCall(i int) *build.Options {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	argsForCall := fake.refreshArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildImplementation) RefreshReturns(result1 *s6.Context, result2 *exec.Executor, result3 error) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = nil
	fake.refreshReturns = struct {
		result1 *s6.Context
		result2 *exec.Executor
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildImplementation) RefreshReturnsOnCall(i int, result1 *s6.Context, result2 *exec.Executor, result3 error) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = nil
	if fake.refreshReturnsOnCall == nil {
		fake.refreshReturnsOnCall = make(map[int]struct {
			result1 *s6.Context
			result2 *exec.Executor
			result3 error
		})
	}
	fake.refreshReturnsOnCall[i] = struct {
		result1 *s6.Context
		result2 *exec.Executor
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildImplementation) ValidateImageConfiguration(arg1 *types.ImageConfiguration) error {
	fake.validateImageConfigurationMutex.Lock()
	ret, specificReturn := fake.validateImageConfigurationReturnsOnCall[len(fake.validateImageConfigurationArgsForCall)]
	fake.validateImageConfigurationArgsForCall = append(fake.validateImageConfigurationArgsForCall, struct {
		arg1 *types.ImageConfiguration
	}{arg1})
	stub := fake.ValidateImageConfigurationStub
	fakeReturns := fake.validateImageConfigurationReturns
	fake.recordInvocation("ValidateImageConfiguration", []interface{}{arg1})
	fake.validateImageConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) ValidateImageConfigurationCallCount() int {
	fake.validateImageConfigurationMutex.RLock()
	defer fake.validateImageConfigurationMutex.RUnlock()
	return len(fake.validateImageConfigurationArgsForCall)
}

func (fake *FakeBuildImplementation) ValidateImageConfigurationCalls(stub func(*types.ImageConfiguration) error) {
	fake.validateImageConfigurationMutex.Lock()
	defer fake.validateImageConfigurationMutex.Unlock()
	fake.ValidateImageConfigurationStub = stub
}

func (fake *FakeBuildImplementation) ValidateImageConfigurationArgsForCall(i int) *types.ImageConfiguration {
	fake.validateImageConfigurationMutex.RLock()
	defer fake.validateImageConfigurationMutex.RUnlock()
	argsForCall := fake.validateImageConfigurationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildImplementation) ValidateImageConfigurationReturns(result1 error) {
	fake.validateImageConfigurationMutex.Lock()
	defer fake.validateImageConfigurationMutex.Unlock()
	fake.ValidateImageConfigurationStub = nil
	fake.validateImageConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) ValidateImageConfigurationReturnsOnCall(i int, result1 error) {
	fake.validateImageConfigurationMutex.Lock()
	defer fake.validateImageConfigurationMutex.Unlock()
	fake.ValidateImageConfigurationStub = nil
	if fake.validateImageConfigurationReturnsOnCall == nil {
		fake.validateImageConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateImageConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) WriteSupervisionTree(arg1 *s6.Context, arg2 *types.ImageConfiguration) error {
	fake.writeSupervisionTreeMutex.Lock()
	ret, specificReturn := fake.writeSupervisionTreeReturnsOnCall[len(fake.writeSupervisionTreeArgsForCall)]
	fake.writeSupervisionTreeArgsForCall = append(fake.writeSupervisionTreeArgsForCall, struct {
		arg1 *s6.Context
		arg2 *types.ImageConfiguration
	}{arg1, arg2})
	stub := fake.WriteSupervisionTreeStub
	fakeReturns := fake.writeSupervisionTreeReturns
	fake.recordInvocation("WriteSupervisionTree", []interface{}{arg1, arg2})
	fake.writeSupervisionTreeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildImplementation) WriteSupervisionTreeCallCount() int {
	fake.writeSupervisionTreeMutex.RLock()
	defer fake.writeSupervisionTreeMutex.RUnlock()
	return len(fake.writeSupervisionTreeArgsForCall)
}

func (fake *FakeBuildImplementation) WriteSupervisionTreeCalls(stub func(*s6.Context, *types.ImageConfiguration) error) {
	fake.writeSupervisionTreeMutex.Lock()
	defer fake.writeSupervisionTreeMutex.Unlock()
	fake.WriteSupervisionTreeStub = stub
}

func (fake *FakeBuildImplementation) WriteSupervisionTreeArgsForCall(i int) (*s6.Context, *types.ImageConfiguration) {
	fake.writeSupervisionTreeMutex.RLock()
	defer fake.writeSupervisionTreeMutex.RUnlock()
	argsForCall := fake.writeSupervisionTreeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildImplementation) WriteSupervisionTreeReturns(result1 error) {
	fake.writeSupervisionTreeMutex.Lock()
	defer fake.writeSupervisionTreeMutex.Unlock()
	fake.WriteSupervisionTreeStub = nil
	fake.writeSupervisionTreeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) WriteSupervisionTreeReturnsOnCall(i int, result1 error) {
	fake.writeSupervisionTreeMutex.Lock()
	defer fake.writeSupervisionTreeMutex.Unlock()
	fake.WriteSupervisionTreeStub = nil
	if fake.writeSupervisionTreeReturnsOnCall == nil {
		fake.writeSupervisionTreeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeSupervisionTreeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildImageMutex.RLock()
	defer fake.buildImageMutex.RUnlock()
	fake.buildTarballMutex.RLock()
	defer fake.buildTarballMutex.RUnlock()
	fake.fixateApkWorldMutex.RLock()
	defer fake.fixateApkWorldMutex.RUnlock()
	fake.generateSBOMMutex.RLock()
	defer fake.generateSBOMMutex.RUnlock()
	fake.initApkDBMutex.RLock()
	defer fake.initApkDBMutex.RUnlock()
	fake.initApkKeyringMutex.RLock()
	defer fake.initApkKeyringMutex.RUnlock()
	fake.initApkWorldMutex.RLock()
	defer fake.initApkWorldMutex.RUnlock()
	fake.initializeApkMutex.RLock()
	defer fake.initializeApkMutex.RUnlock()
	fake.installBusyboxSymlinksMutex.RLock()
	defer fake.installBusyboxSymlinksMutex.RUnlock()
	fake.mutateAccountsMutex.RLock()
	defer fake.mutateAccountsMutex.RUnlock()
	fake.normalizeApkScriptsTarMutex.RLock()
	defer fake.normalizeApkScriptsTarMutex.RUnlock()
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	fake.validateImageConfigurationMutex.RLock()
	defer fake.validateImageConfigurationMutex.RUnlock()
	fake.writeSupervisionTreeMutex.RLock()
	defer fake.writeSupervisionTreeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildImplementation) recordInvocation(key string, args []interface{}) {
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
