// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"gitlabhost.rtp.raleigh.ibm.com/spectrum-plugin/core"
)

type FakeSpectrumClient struct {
	CreateFilesetStub        func(fileset *core.Fileset) error
	createFilesetMutex       sync.RWMutex
	createFilesetArgsForCall []struct {
		fileset *core.Fileset
	}
	createFilesetReturns struct {
		result1 error
	}
	RemoveFilesetStub        func(fileset *core.Fileset) error
	removeFilesetMutex       sync.RWMutex
	removeFilesetArgsForCall []struct {
		fileset *core.Fileset
	}
	removeFilesetReturns struct {
		result1 error
	}
	LinkFilesetStub        func(fileset *core.Fileset, path string) error
	linkFilesetMutex       sync.RWMutex
	linkFilesetArgsForCall []struct {
		fileset *core.Fileset
		path    string
	}
	linkFilesetReturns struct {
		result1 error
	}
	UnlinkFilesetStub        func(fileset *core.Fileset) error
	unlinkFilesetMutex       sync.RWMutex
	unlinkFilesetArgsForCall []struct {
		fileset *core.Fileset
	}
	unlinkFilesetReturns struct {
		result1 error
	}
	ListFilesetsStub        func(filesystem string) ([]core.Fileset, error)
	listFilesetsMutex       sync.RWMutex
	listFilesetsArgsForCall []struct {
		filesystem string
	}
	listFilesetsReturns struct {
		result1 []core.Fileset
		result2 error
	}
	ListFilesetStub        func(filesystem string, fileset string) (*core.Fileset, error)
	listFilesetMutex       sync.RWMutex
	listFilesetArgsForCall []struct {
		filesystem string
		fileset    string
	}
	listFilesetReturns struct {
		result1 *core.Fileset
		result2 error
	}
}

func (fake *FakeSpectrumClient) CreateFileset(fileset *core.Fileset) error {
	fake.createFilesetMutex.Lock()
	fake.createFilesetArgsForCall = append(fake.createFilesetArgsForCall, struct {
		fileset *core.Fileset
	}{fileset})
	fake.createFilesetMutex.Unlock()
	if fake.CreateFilesetStub != nil {
		return fake.CreateFilesetStub(fileset)
	} else {
		return fake.createFilesetReturns.result1
	}
}

func (fake *FakeSpectrumClient) CreateFilesetCallCount() int {
	fake.createFilesetMutex.RLock()
	defer fake.createFilesetMutex.RUnlock()
	return len(fake.createFilesetArgsForCall)
}

func (fake *FakeSpectrumClient) CreateFilesetArgsForCall(i int) *core.Fileset {
	fake.createFilesetMutex.RLock()
	defer fake.createFilesetMutex.RUnlock()
	return fake.createFilesetArgsForCall[i].fileset
}

func (fake *FakeSpectrumClient) CreateFilesetReturns(result1 error) {
	fake.CreateFilesetStub = nil
	fake.createFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumClient) RemoveFileset(fileset *core.Fileset) error {
	fake.removeFilesetMutex.Lock()
	fake.removeFilesetArgsForCall = append(fake.removeFilesetArgsForCall, struct {
		fileset *core.Fileset
	}{fileset})
	fake.removeFilesetMutex.Unlock()
	if fake.RemoveFilesetStub != nil {
		return fake.RemoveFilesetStub(fileset)
	} else {
		return fake.removeFilesetReturns.result1
	}
}

func (fake *FakeSpectrumClient) RemoveFilesetCallCount() int {
	fake.removeFilesetMutex.RLock()
	defer fake.removeFilesetMutex.RUnlock()
	return len(fake.removeFilesetArgsForCall)
}

func (fake *FakeSpectrumClient) RemoveFilesetArgsForCall(i int) *core.Fileset {
	fake.removeFilesetMutex.RLock()
	defer fake.removeFilesetMutex.RUnlock()
	return fake.removeFilesetArgsForCall[i].fileset
}

func (fake *FakeSpectrumClient) RemoveFilesetReturns(result1 error) {
	fake.RemoveFilesetStub = nil
	fake.removeFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumClient) LinkFileset(fileset *core.Fileset, path string) error {
	fake.linkFilesetMutex.Lock()
	fake.linkFilesetArgsForCall = append(fake.linkFilesetArgsForCall, struct {
		fileset *core.Fileset
		path    string
	}{fileset, path})
	fake.linkFilesetMutex.Unlock()
	if fake.LinkFilesetStub != nil {
		return fake.LinkFilesetStub(fileset, path)
	} else {
		return fake.linkFilesetReturns.result1
	}
}

func (fake *FakeSpectrumClient) LinkFilesetCallCount() int {
	fake.linkFilesetMutex.RLock()
	defer fake.linkFilesetMutex.RUnlock()
	return len(fake.linkFilesetArgsForCall)
}

func (fake *FakeSpectrumClient) LinkFilesetArgsForCall(i int) (*core.Fileset, string) {
	fake.linkFilesetMutex.RLock()
	defer fake.linkFilesetMutex.RUnlock()
	return fake.linkFilesetArgsForCall[i].fileset, fake.linkFilesetArgsForCall[i].path
}

func (fake *FakeSpectrumClient) LinkFilesetReturns(result1 error) {
	fake.LinkFilesetStub = nil
	fake.linkFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumClient) UnlinkFileset(fileset *core.Fileset) error {
	fake.unlinkFilesetMutex.Lock()
	fake.unlinkFilesetArgsForCall = append(fake.unlinkFilesetArgsForCall, struct {
		fileset *core.Fileset
	}{fileset})
	fake.unlinkFilesetMutex.Unlock()
	if fake.UnlinkFilesetStub != nil {
		return fake.UnlinkFilesetStub(fileset)
	} else {
		return fake.unlinkFilesetReturns.result1
	}
}

func (fake *FakeSpectrumClient) UnlinkFilesetCallCount() int {
	fake.unlinkFilesetMutex.RLock()
	defer fake.unlinkFilesetMutex.RUnlock()
	return len(fake.unlinkFilesetArgsForCall)
}

func (fake *FakeSpectrumClient) UnlinkFilesetArgsForCall(i int) *core.Fileset {
	fake.unlinkFilesetMutex.RLock()
	defer fake.unlinkFilesetMutex.RUnlock()
	return fake.unlinkFilesetArgsForCall[i].fileset
}

func (fake *FakeSpectrumClient) UnlinkFilesetReturns(result1 error) {
	fake.UnlinkFilesetStub = nil
	fake.unlinkFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumClient) ListFilesets(filesystem string) ([]core.Fileset, error) {
	fake.listFilesetsMutex.Lock()
	fake.listFilesetsArgsForCall = append(fake.listFilesetsArgsForCall, struct {
		filesystem string
	}{filesystem})
	fake.listFilesetsMutex.Unlock()
	if fake.ListFilesetsStub != nil {
		return fake.ListFilesetsStub(filesystem)
	} else {
		return fake.listFilesetsReturns.result1, fake.listFilesetsReturns.result2
	}
}

func (fake *FakeSpectrumClient) ListFilesetsCallCount() int {
	fake.listFilesetsMutex.RLock()
	defer fake.listFilesetsMutex.RUnlock()
	return len(fake.listFilesetsArgsForCall)
}

func (fake *FakeSpectrumClient) ListFilesetsArgsForCall(i int) string {
	fake.listFilesetsMutex.RLock()
	defer fake.listFilesetsMutex.RUnlock()
	return fake.listFilesetsArgsForCall[i].filesystem
}

func (fake *FakeSpectrumClient) ListFilesetsReturns(result1 []core.Fileset, result2 error) {
	fake.ListFilesetsStub = nil
	fake.listFilesetsReturns = struct {
		result1 []core.Fileset
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumClient) ListFileset(filesystem string, fileset string) (*core.Fileset, error) {
	fake.listFilesetMutex.Lock()
	fake.listFilesetArgsForCall = append(fake.listFilesetArgsForCall, struct {
		filesystem string
		fileset    string
	}{filesystem, fileset})
	fake.listFilesetMutex.Unlock()
	if fake.ListFilesetStub != nil {
		return fake.ListFilesetStub(filesystem, fileset)
	} else {
		return fake.listFilesetReturns.result1, fake.listFilesetReturns.result2
	}
}

func (fake *FakeSpectrumClient) ListFilesetCallCount() int {
	fake.listFilesetMutex.RLock()
	defer fake.listFilesetMutex.RUnlock()
	return len(fake.listFilesetArgsForCall)
}

func (fake *FakeSpectrumClient) ListFilesetArgsForCall(i int) (string, string) {
	fake.listFilesetMutex.RLock()
	defer fake.listFilesetMutex.RUnlock()
	return fake.listFilesetArgsForCall[i].filesystem, fake.listFilesetArgsForCall[i].fileset
}

func (fake *FakeSpectrumClient) ListFilesetReturns(result1 *core.Fileset, result2 error) {
	fake.ListFilesetStub = nil
	fake.listFilesetReturns = struct {
		result1 *core.Fileset
		result2 error
	}{result1, result2}
}

var _ core.SpectrumClient = new(FakeSpectrumClient)
