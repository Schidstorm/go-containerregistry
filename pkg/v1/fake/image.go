// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	v1 "github.com/schidstorm/go-containerregistry/pkg/v1"
	"github.com/schidstorm/go-containerregistry/pkg/v1/types"
)

type FakeImage struct {
	ConfigFileStub        func() (*v1.ConfigFile, error)
	configFileMutex       sync.RWMutex
	configFileArgsForCall []struct {
	}
	configFileReturns struct {
		result1 *v1.ConfigFile
		result2 error
	}
	configFileReturnsOnCall map[int]struct {
		result1 *v1.ConfigFile
		result2 error
	}
	ConfigNameStub        func() (v1.Hash, error)
	configNameMutex       sync.RWMutex
	configNameArgsForCall []struct {
	}
	configNameReturns struct {
		result1 v1.Hash
		result2 error
	}
	configNameReturnsOnCall map[int]struct {
		result1 v1.Hash
		result2 error
	}
	DigestStub        func() (v1.Hash, error)
	digestMutex       sync.RWMutex
	digestArgsForCall []struct {
	}
	digestReturns struct {
		result1 v1.Hash
		result2 error
	}
	digestReturnsOnCall map[int]struct {
		result1 v1.Hash
		result2 error
	}
	LayerByDiffIDStub        func(v1.Hash) (v1.Layer, error)
	layerByDiffIDMutex       sync.RWMutex
	layerByDiffIDArgsForCall []struct {
		arg1 v1.Hash
	}
	layerByDiffIDReturns struct {
		result1 v1.Layer
		result2 error
	}
	layerByDiffIDReturnsOnCall map[int]struct {
		result1 v1.Layer
		result2 error
	}
	LayerByDigestStub        func(v1.Hash) (v1.Layer, error)
	layerByDigestMutex       sync.RWMutex
	layerByDigestArgsForCall []struct {
		arg1 v1.Hash
	}
	layerByDigestReturns struct {
		result1 v1.Layer
		result2 error
	}
	layerByDigestReturnsOnCall map[int]struct {
		result1 v1.Layer
		result2 error
	}
	LayersStub        func() ([]v1.Layer, error)
	layersMutex       sync.RWMutex
	layersArgsForCall []struct {
	}
	layersReturns struct {
		result1 []v1.Layer
		result2 error
	}
	layersReturnsOnCall map[int]struct {
		result1 []v1.Layer
		result2 error
	}
	ManifestStub        func() (*v1.Manifest, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
	}
	manifestReturns struct {
		result1 *v1.Manifest
		result2 error
	}
	manifestReturnsOnCall map[int]struct {
		result1 *v1.Manifest
		result2 error
	}
	MediaTypeStub        func() (types.MediaType, error)
	mediaTypeMutex       sync.RWMutex
	mediaTypeArgsForCall []struct {
	}
	mediaTypeReturns struct {
		result1 types.MediaType
		result2 error
	}
	mediaTypeReturnsOnCall map[int]struct {
		result1 types.MediaType
		result2 error
	}
	RawConfigFileStub        func() ([]byte, error)
	rawConfigFileMutex       sync.RWMutex
	rawConfigFileArgsForCall []struct {
	}
	rawConfigFileReturns struct {
		result1 []byte
		result2 error
	}
	rawConfigFileReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	RawManifestStub        func() ([]byte, error)
	rawManifestMutex       sync.RWMutex
	rawManifestArgsForCall []struct {
	}
	rawManifestReturns struct {
		result1 []byte
		result2 error
	}
	rawManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SizeStub        func() (int64, error)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct {
	}
	sizeReturns struct {
		result1 int64
		result2 error
	}
	sizeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImage) ConfigFile() (*v1.ConfigFile, error) {
	fake.configFileMutex.Lock()
	ret, specificReturn := fake.configFileReturnsOnCall[len(fake.configFileArgsForCall)]
	fake.configFileArgsForCall = append(fake.configFileArgsForCall, struct {
	}{})
	stub := fake.ConfigFileStub
	fakeReturns := fake.configFileReturns
	fake.recordInvocation("ConfigFile", []interface{}{})
	fake.configFileMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) ConfigFileCallCount() int {
	fake.configFileMutex.RLock()
	defer fake.configFileMutex.RUnlock()
	return len(fake.configFileArgsForCall)
}

func (fake *FakeImage) ConfigFileCalls(stub func() (*v1.ConfigFile, error)) {
	fake.configFileMutex.Lock()
	defer fake.configFileMutex.Unlock()
	fake.ConfigFileStub = stub
}

func (fake *FakeImage) ConfigFileReturns(result1 *v1.ConfigFile, result2 error) {
	fake.configFileMutex.Lock()
	defer fake.configFileMutex.Unlock()
	fake.ConfigFileStub = nil
	fake.configFileReturns = struct {
		result1 *v1.ConfigFile
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) ConfigFileReturnsOnCall(i int, result1 *v1.ConfigFile, result2 error) {
	fake.configFileMutex.Lock()
	defer fake.configFileMutex.Unlock()
	fake.ConfigFileStub = nil
	if fake.configFileReturnsOnCall == nil {
		fake.configFileReturnsOnCall = make(map[int]struct {
			result1 *v1.ConfigFile
			result2 error
		})
	}
	fake.configFileReturnsOnCall[i] = struct {
		result1 *v1.ConfigFile
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) ConfigName() (v1.Hash, error) {
	fake.configNameMutex.Lock()
	ret, specificReturn := fake.configNameReturnsOnCall[len(fake.configNameArgsForCall)]
	fake.configNameArgsForCall = append(fake.configNameArgsForCall, struct {
	}{})
	stub := fake.ConfigNameStub
	fakeReturns := fake.configNameReturns
	fake.recordInvocation("ConfigName", []interface{}{})
	fake.configNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) ConfigNameCallCount() int {
	fake.configNameMutex.RLock()
	defer fake.configNameMutex.RUnlock()
	return len(fake.configNameArgsForCall)
}

func (fake *FakeImage) ConfigNameCalls(stub func() (v1.Hash, error)) {
	fake.configNameMutex.Lock()
	defer fake.configNameMutex.Unlock()
	fake.ConfigNameStub = stub
}

func (fake *FakeImage) ConfigNameReturns(result1 v1.Hash, result2 error) {
	fake.configNameMutex.Lock()
	defer fake.configNameMutex.Unlock()
	fake.ConfigNameStub = nil
	fake.configNameReturns = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) ConfigNameReturnsOnCall(i int, result1 v1.Hash, result2 error) {
	fake.configNameMutex.Lock()
	defer fake.configNameMutex.Unlock()
	fake.ConfigNameStub = nil
	if fake.configNameReturnsOnCall == nil {
		fake.configNameReturnsOnCall = make(map[int]struct {
			result1 v1.Hash
			result2 error
		})
	}
	fake.configNameReturnsOnCall[i] = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Digest() (v1.Hash, error) {
	fake.digestMutex.Lock()
	ret, specificReturn := fake.digestReturnsOnCall[len(fake.digestArgsForCall)]
	fake.digestArgsForCall = append(fake.digestArgsForCall, struct {
	}{})
	stub := fake.DigestStub
	fakeReturns := fake.digestReturns
	fake.recordInvocation("Digest", []interface{}{})
	fake.digestMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) DigestCallCount() int {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	return len(fake.digestArgsForCall)
}

func (fake *FakeImage) DigestCalls(stub func() (v1.Hash, error)) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = stub
}

func (fake *FakeImage) DigestReturns(result1 v1.Hash, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	fake.digestReturns = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) DigestReturnsOnCall(i int, result1 v1.Hash, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	if fake.digestReturnsOnCall == nil {
		fake.digestReturnsOnCall = make(map[int]struct {
			result1 v1.Hash
			result2 error
		})
	}
	fake.digestReturnsOnCall[i] = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) LayerByDiffID(arg1 v1.Hash) (v1.Layer, error) {
	fake.layerByDiffIDMutex.Lock()
	ret, specificReturn := fake.layerByDiffIDReturnsOnCall[len(fake.layerByDiffIDArgsForCall)]
	fake.layerByDiffIDArgsForCall = append(fake.layerByDiffIDArgsForCall, struct {
		arg1 v1.Hash
	}{arg1})
	stub := fake.LayerByDiffIDStub
	fakeReturns := fake.layerByDiffIDReturns
	fake.recordInvocation("LayerByDiffID", []interface{}{arg1})
	fake.layerByDiffIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) LayerByDiffIDCallCount() int {
	fake.layerByDiffIDMutex.RLock()
	defer fake.layerByDiffIDMutex.RUnlock()
	return len(fake.layerByDiffIDArgsForCall)
}

func (fake *FakeImage) LayerByDiffIDCalls(stub func(v1.Hash) (v1.Layer, error)) {
	fake.layerByDiffIDMutex.Lock()
	defer fake.layerByDiffIDMutex.Unlock()
	fake.LayerByDiffIDStub = stub
}

func (fake *FakeImage) LayerByDiffIDArgsForCall(i int) v1.Hash {
	fake.layerByDiffIDMutex.RLock()
	defer fake.layerByDiffIDMutex.RUnlock()
	argsForCall := fake.layerByDiffIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImage) LayerByDiffIDReturns(result1 v1.Layer, result2 error) {
	fake.layerByDiffIDMutex.Lock()
	defer fake.layerByDiffIDMutex.Unlock()
	fake.LayerByDiffIDStub = nil
	fake.layerByDiffIDReturns = struct {
		result1 v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) LayerByDiffIDReturnsOnCall(i int, result1 v1.Layer, result2 error) {
	fake.layerByDiffIDMutex.Lock()
	defer fake.layerByDiffIDMutex.Unlock()
	fake.LayerByDiffIDStub = nil
	if fake.layerByDiffIDReturnsOnCall == nil {
		fake.layerByDiffIDReturnsOnCall = make(map[int]struct {
			result1 v1.Layer
			result2 error
		})
	}
	fake.layerByDiffIDReturnsOnCall[i] = struct {
		result1 v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) LayerByDigest(arg1 v1.Hash) (v1.Layer, error) {
	fake.layerByDigestMutex.Lock()
	ret, specificReturn := fake.layerByDigestReturnsOnCall[len(fake.layerByDigestArgsForCall)]
	fake.layerByDigestArgsForCall = append(fake.layerByDigestArgsForCall, struct {
		arg1 v1.Hash
	}{arg1})
	stub := fake.LayerByDigestStub
	fakeReturns := fake.layerByDigestReturns
	fake.recordInvocation("LayerByDigest", []interface{}{arg1})
	fake.layerByDigestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) LayerByDigestCallCount() int {
	fake.layerByDigestMutex.RLock()
	defer fake.layerByDigestMutex.RUnlock()
	return len(fake.layerByDigestArgsForCall)
}

func (fake *FakeImage) LayerByDigestCalls(stub func(v1.Hash) (v1.Layer, error)) {
	fake.layerByDigestMutex.Lock()
	defer fake.layerByDigestMutex.Unlock()
	fake.LayerByDigestStub = stub
}

func (fake *FakeImage) LayerByDigestArgsForCall(i int) v1.Hash {
	fake.layerByDigestMutex.RLock()
	defer fake.layerByDigestMutex.RUnlock()
	argsForCall := fake.layerByDigestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImage) LayerByDigestReturns(result1 v1.Layer, result2 error) {
	fake.layerByDigestMutex.Lock()
	defer fake.layerByDigestMutex.Unlock()
	fake.LayerByDigestStub = nil
	fake.layerByDigestReturns = struct {
		result1 v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) LayerByDigestReturnsOnCall(i int, result1 v1.Layer, result2 error) {
	fake.layerByDigestMutex.Lock()
	defer fake.layerByDigestMutex.Unlock()
	fake.LayerByDigestStub = nil
	if fake.layerByDigestReturnsOnCall == nil {
		fake.layerByDigestReturnsOnCall = make(map[int]struct {
			result1 v1.Layer
			result2 error
		})
	}
	fake.layerByDigestReturnsOnCall[i] = struct {
		result1 v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Layers() ([]v1.Layer, error) {
	fake.layersMutex.Lock()
	ret, specificReturn := fake.layersReturnsOnCall[len(fake.layersArgsForCall)]
	fake.layersArgsForCall = append(fake.layersArgsForCall, struct {
	}{})
	stub := fake.LayersStub
	fakeReturns := fake.layersReturns
	fake.recordInvocation("Layers", []interface{}{})
	fake.layersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) LayersCallCount() int {
	fake.layersMutex.RLock()
	defer fake.layersMutex.RUnlock()
	return len(fake.layersArgsForCall)
}

func (fake *FakeImage) LayersCalls(stub func() ([]v1.Layer, error)) {
	fake.layersMutex.Lock()
	defer fake.layersMutex.Unlock()
	fake.LayersStub = stub
}

func (fake *FakeImage) LayersReturns(result1 []v1.Layer, result2 error) {
	fake.layersMutex.Lock()
	defer fake.layersMutex.Unlock()
	fake.LayersStub = nil
	fake.layersReturns = struct {
		result1 []v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) LayersReturnsOnCall(i int, result1 []v1.Layer, result2 error) {
	fake.layersMutex.Lock()
	defer fake.layersMutex.Unlock()
	fake.LayersStub = nil
	if fake.layersReturnsOnCall == nil {
		fake.layersReturnsOnCall = make(map[int]struct {
			result1 []v1.Layer
			result2 error
		})
	}
	fake.layersReturnsOnCall[i] = struct {
		result1 []v1.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Manifest() (*v1.Manifest, error) {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
	}{})
	stub := fake.ManifestStub
	fakeReturns := fake.manifestReturns
	fake.recordInvocation("Manifest", []interface{}{})
	fake.manifestMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeImage) ManifestCalls(stub func() (*v1.Manifest, error)) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = stub
}

func (fake *FakeImage) ManifestReturns(result1 *v1.Manifest, result2 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 *v1.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) ManifestReturnsOnCall(i int, result1 *v1.Manifest, result2 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 *v1.Manifest
			result2 error
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 *v1.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) MediaType() (types.MediaType, error) {
	fake.mediaTypeMutex.Lock()
	ret, specificReturn := fake.mediaTypeReturnsOnCall[len(fake.mediaTypeArgsForCall)]
	fake.mediaTypeArgsForCall = append(fake.mediaTypeArgsForCall, struct {
	}{})
	stub := fake.MediaTypeStub
	fakeReturns := fake.mediaTypeReturns
	fake.recordInvocation("MediaType", []interface{}{})
	fake.mediaTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) MediaTypeCallCount() int {
	fake.mediaTypeMutex.RLock()
	defer fake.mediaTypeMutex.RUnlock()
	return len(fake.mediaTypeArgsForCall)
}

func (fake *FakeImage) MediaTypeCalls(stub func() (types.MediaType, error)) {
	fake.mediaTypeMutex.Lock()
	defer fake.mediaTypeMutex.Unlock()
	fake.MediaTypeStub = stub
}

func (fake *FakeImage) MediaTypeReturns(result1 types.MediaType, result2 error) {
	fake.mediaTypeMutex.Lock()
	defer fake.mediaTypeMutex.Unlock()
	fake.MediaTypeStub = nil
	fake.mediaTypeReturns = struct {
		result1 types.MediaType
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) MediaTypeReturnsOnCall(i int, result1 types.MediaType, result2 error) {
	fake.mediaTypeMutex.Lock()
	defer fake.mediaTypeMutex.Unlock()
	fake.MediaTypeStub = nil
	if fake.mediaTypeReturnsOnCall == nil {
		fake.mediaTypeReturnsOnCall = make(map[int]struct {
			result1 types.MediaType
			result2 error
		})
	}
	fake.mediaTypeReturnsOnCall[i] = struct {
		result1 types.MediaType
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) RawConfigFile() ([]byte, error) {
	fake.rawConfigFileMutex.Lock()
	ret, specificReturn := fake.rawConfigFileReturnsOnCall[len(fake.rawConfigFileArgsForCall)]
	fake.rawConfigFileArgsForCall = append(fake.rawConfigFileArgsForCall, struct {
	}{})
	stub := fake.RawConfigFileStub
	fakeReturns := fake.rawConfigFileReturns
	fake.recordInvocation("RawConfigFile", []interface{}{})
	fake.rawConfigFileMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) RawConfigFileCallCount() int {
	fake.rawConfigFileMutex.RLock()
	defer fake.rawConfigFileMutex.RUnlock()
	return len(fake.rawConfigFileArgsForCall)
}

func (fake *FakeImage) RawConfigFileCalls(stub func() ([]byte, error)) {
	fake.rawConfigFileMutex.Lock()
	defer fake.rawConfigFileMutex.Unlock()
	fake.RawConfigFileStub = stub
}

func (fake *FakeImage) RawConfigFileReturns(result1 []byte, result2 error) {
	fake.rawConfigFileMutex.Lock()
	defer fake.rawConfigFileMutex.Unlock()
	fake.RawConfigFileStub = nil
	fake.rawConfigFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) RawConfigFileReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.rawConfigFileMutex.Lock()
	defer fake.rawConfigFileMutex.Unlock()
	fake.RawConfigFileStub = nil
	if fake.rawConfigFileReturnsOnCall == nil {
		fake.rawConfigFileReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.rawConfigFileReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) RawManifest() ([]byte, error) {
	fake.rawManifestMutex.Lock()
	ret, specificReturn := fake.rawManifestReturnsOnCall[len(fake.rawManifestArgsForCall)]
	fake.rawManifestArgsForCall = append(fake.rawManifestArgsForCall, struct {
	}{})
	stub := fake.RawManifestStub
	fakeReturns := fake.rawManifestReturns
	fake.recordInvocation("RawManifest", []interface{}{})
	fake.rawManifestMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) RawManifestCallCount() int {
	fake.rawManifestMutex.RLock()
	defer fake.rawManifestMutex.RUnlock()
	return len(fake.rawManifestArgsForCall)
}

func (fake *FakeImage) RawManifestCalls(stub func() ([]byte, error)) {
	fake.rawManifestMutex.Lock()
	defer fake.rawManifestMutex.Unlock()
	fake.RawManifestStub = stub
}

func (fake *FakeImage) RawManifestReturns(result1 []byte, result2 error) {
	fake.rawManifestMutex.Lock()
	defer fake.rawManifestMutex.Unlock()
	fake.RawManifestStub = nil
	fake.rawManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) RawManifestReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.rawManifestMutex.Lock()
	defer fake.rawManifestMutex.Unlock()
	fake.RawManifestStub = nil
	if fake.rawManifestReturnsOnCall == nil {
		fake.rawManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.rawManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Size() (int64, error) {
	fake.sizeMutex.Lock()
	ret, specificReturn := fake.sizeReturnsOnCall[len(fake.sizeArgsForCall)]
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct {
	}{})
	stub := fake.SizeStub
	fakeReturns := fake.sizeReturns
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImage) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeImage) SizeCalls(stub func() (int64, error)) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = stub
}

func (fake *FakeImage) SizeReturns(result1 int64, result2 error) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) SizeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	if fake.sizeReturnsOnCall == nil {
		fake.sizeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.sizeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configFileMutex.RLock()
	defer fake.configFileMutex.RUnlock()
	fake.configNameMutex.RLock()
	defer fake.configNameMutex.RUnlock()
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	fake.layerByDiffIDMutex.RLock()
	defer fake.layerByDiffIDMutex.RUnlock()
	fake.layerByDigestMutex.RLock()
	defer fake.layerByDigestMutex.RUnlock()
	fake.layersMutex.RLock()
	defer fake.layersMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.mediaTypeMutex.RLock()
	defer fake.mediaTypeMutex.RUnlock()
	fake.rawConfigFileMutex.RLock()
	defer fake.rawConfigFileMutex.RUnlock()
	fake.rawManifestMutex.RLock()
	defer fake.rawManifestMutex.RUnlock()
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImage) recordInvocation(key string, args []interface{}) {
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

var _ v1.Image = new(FakeImage)
