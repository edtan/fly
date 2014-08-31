// This file was generated by counterfeiter
package fakes

import (
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/concourse/turbine/api/builds"
	"github.com/concourse/turbine/builder"
	"github.com/concourse/turbine/event"

	"sync"
)

type FakeBuilder struct {
	StartStub        func(builds.Build, event.Emitter, <-chan struct{}) (builder.RunningBuild, error)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 builds.Build
		arg2 event.Emitter
		arg3 <-chan struct{}
	}
	startReturns struct {
		result1 builder.RunningBuild
		result2 error
	}
	AttachStub        func(builder.RunningBuild, event.Emitter, <-chan struct{}) (builder.SucceededBuild, error, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		arg1 builder.RunningBuild
		arg2 event.Emitter
		arg3 <-chan struct{}
	}
	attachReturns struct {
		result1 builder.SucceededBuild
		result2 error
		result3 error
	}
	HijackStub        func(builder.RunningBuild, warden.ProcessSpec, warden.ProcessIO) (warden.Process, error)
	hijackMutex       sync.RWMutex
	hijackArgsForCall []struct {
		arg1 builder.RunningBuild
		arg2 warden.ProcessSpec
		arg3 warden.ProcessIO
	}
	hijackReturns struct {
		result1 warden.Process
		result2 error
	}
	CompleteStub        func(builder.SucceededBuild, event.Emitter, <-chan struct{}) (builds.Build, error)
	completeMutex       sync.RWMutex
	completeArgsForCall []struct {
		arg1 builder.SucceededBuild
		arg2 event.Emitter
		arg3 <-chan struct{}
	}
	completeReturns struct {
		result1 builds.Build
		result2 error
	}
}

func (fake *FakeBuilder) Start(arg1 builds.Build, arg2 event.Emitter, arg3 <-chan struct{}) (builder.RunningBuild, error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 builds.Build
		arg2 event.Emitter
		arg3 <-chan struct{}
	}{arg1, arg2, arg3})
	if fake.StartStub != nil {
		return fake.StartStub(arg1, arg2, arg3)
	} else {
		return fake.startReturns.result1, fake.startReturns.result2
	}
}

func (fake *FakeBuilder) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeBuilder) StartArgsForCall(i int) (builds.Build, event.Emitter, <-chan struct{}) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].arg1, fake.startArgsForCall[i].arg2, fake.startArgsForCall[i].arg3
}

func (fake *FakeBuilder) StartReturns(result1 builder.RunningBuild, result2 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 builder.RunningBuild
		result2 error
	}{result1, result2}
}

func (fake *FakeBuilder) Attach(arg1 builder.RunningBuild, arg2 event.Emitter, arg3 <-chan struct{}) (builder.SucceededBuild, error, error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		arg1 builder.RunningBuild
		arg2 event.Emitter
		arg3 <-chan struct{}
	}{arg1, arg2, arg3})
	if fake.AttachStub != nil {
		return fake.AttachStub(arg1, arg2, arg3)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2, fake.attachReturns.result3
	}
}

func (fake *FakeBuilder) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeBuilder) AttachArgsForCall(i int) (builder.RunningBuild, event.Emitter, <-chan struct{}) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].arg1, fake.attachArgsForCall[i].arg2, fake.attachArgsForCall[i].arg3
}

func (fake *FakeBuilder) AttachReturns(result1 builder.SucceededBuild, result2 error, result3 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 builder.SucceededBuild
		result2 error
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuilder) Hijack(arg1 builder.RunningBuild, arg2 warden.ProcessSpec, arg3 warden.ProcessIO) (warden.Process, error) {
	fake.hijackMutex.Lock()
	defer fake.hijackMutex.Unlock()
	fake.hijackArgsForCall = append(fake.hijackArgsForCall, struct {
		arg1 builder.RunningBuild
		arg2 warden.ProcessSpec
		arg3 warden.ProcessIO
	}{arg1, arg2, arg3})
	if fake.HijackStub != nil {
		return fake.HijackStub(arg1, arg2, arg3)
	} else {
		return fake.hijackReturns.result1, fake.hijackReturns.result2
	}
}

func (fake *FakeBuilder) HijackCallCount() int {
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return len(fake.hijackArgsForCall)
}

func (fake *FakeBuilder) HijackArgsForCall(i int) (builder.RunningBuild, warden.ProcessSpec, warden.ProcessIO) {
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return fake.hijackArgsForCall[i].arg1, fake.hijackArgsForCall[i].arg2, fake.hijackArgsForCall[i].arg3
}

func (fake *FakeBuilder) HijackReturns(result1 warden.Process, result2 error) {
	fake.HijackStub = nil
	fake.hijackReturns = struct {
		result1 warden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeBuilder) Complete(arg1 builder.SucceededBuild, arg2 event.Emitter, arg3 <-chan struct{}) (builds.Build, error) {
	fake.completeMutex.Lock()
	defer fake.completeMutex.Unlock()
	fake.completeArgsForCall = append(fake.completeArgsForCall, struct {
		arg1 builder.SucceededBuild
		arg2 event.Emitter
		arg3 <-chan struct{}
	}{arg1, arg2, arg3})
	if fake.CompleteStub != nil {
		return fake.CompleteStub(arg1, arg2, arg3)
	} else {
		return fake.completeReturns.result1, fake.completeReturns.result2
	}
}

func (fake *FakeBuilder) CompleteCallCount() int {
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	return len(fake.completeArgsForCall)
}

func (fake *FakeBuilder) CompleteArgsForCall(i int) (builder.SucceededBuild, event.Emitter, <-chan struct{}) {
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	return fake.completeArgsForCall[i].arg1, fake.completeArgsForCall[i].arg2, fake.completeArgsForCall[i].arg3
}

func (fake *FakeBuilder) CompleteReturns(result1 builds.Build, result2 error) {
	fake.CompleteStub = nil
	fake.completeReturns = struct {
		result1 builds.Build
		result2 error
	}{result1, result2}
}

var _ builder.Builder = new(FakeBuilder)