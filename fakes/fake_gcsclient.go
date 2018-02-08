// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/frodenas/gcs-resource"
	storage "google.golang.org/api/storage/v1"
)

type FakeGCSClient struct {
	BucketObjectsStub        func(bucketName string, prefix string) ([]string, error)
	bucketObjectsMutex       sync.RWMutex
	bucketObjectsArgsForCall []struct {
		bucketName string
		prefix     string
	}
	bucketObjectsReturns struct {
		result1 []string
		result2 error
	}
	bucketObjectsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ObjectGenerationsStub        func(bucketName string, objectPath string) ([]int64, error)
	objectGenerationsMutex       sync.RWMutex
	objectGenerationsArgsForCall []struct {
		bucketName string
		objectPath string
	}
	objectGenerationsReturns struct {
		result1 []int64
		result2 error
	}
	objectGenerationsReturnsOnCall map[int]struct {
		result1 []int64
		result2 error
	}
	DownloadFileStub        func(bucketName string, objectPath string, generation int64, localPath string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
		localPath  string
	}
	downloadFileReturns struct {
		result1 error
	}
	downloadFileReturnsOnCall map[int]struct {
		result1 error
	}
	UploadFileStub        func(bucketName string, objectPath string, objectContentType string, localPath string, predefinedACL string) (int64, error)
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		bucketName        string
		objectPath        string
		objectContentType string
		localPath         string
		predefinedACL     string
	}
	uploadFileReturns struct {
		result1 int64
		result2 error
	}
	uploadFileReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	URLStub        func(bucketName string, objectPath string, generation int64) (string, error)
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
	}
	uRLReturns struct {
		result1 string
		result2 error
	}
	uRLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DeleteObjectStub        func(bucketName string, objectPath string, generation int64) error
	deleteObjectMutex       sync.RWMutex
	deleteObjectArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
	}
	deleteObjectReturns struct {
		result1 error
	}
	deleteObjectReturnsOnCall map[int]struct {
		result1 error
	}
	GetBucketObjectInfoStub        func(bucketName, objectPath string) (*storage.Object, error)
	getBucketObjectInfoMutex       sync.RWMutex
	getBucketObjectInfoArgsForCall []struct {
		bucketName string
		objectPath string
	}
	getBucketObjectInfoReturns struct {
		result1 *storage.Object
		result2 error
	}
	getBucketObjectInfoReturnsOnCall map[int]struct {
		result1 *storage.Object
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGCSClient) BucketObjects(bucketName string, prefix string) ([]string, error) {
	fake.bucketObjectsMutex.Lock()
	ret, specificReturn := fake.bucketObjectsReturnsOnCall[len(fake.bucketObjectsArgsForCall)]
	fake.bucketObjectsArgsForCall = append(fake.bucketObjectsArgsForCall, struct {
		bucketName string
		prefix     string
	}{bucketName, prefix})
	fake.recordInvocation("BucketObjects", []interface{}{bucketName, prefix})
	fake.bucketObjectsMutex.Unlock()
	if fake.BucketObjectsStub != nil {
		return fake.BucketObjectsStub(bucketName, prefix)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bucketObjectsReturns.result1, fake.bucketObjectsReturns.result2
}

func (fake *FakeGCSClient) BucketObjectsCallCount() int {
	fake.bucketObjectsMutex.RLock()
	defer fake.bucketObjectsMutex.RUnlock()
	return len(fake.bucketObjectsArgsForCall)
}

func (fake *FakeGCSClient) BucketObjectsArgsForCall(i int) (string, string) {
	fake.bucketObjectsMutex.RLock()
	defer fake.bucketObjectsMutex.RUnlock()
	return fake.bucketObjectsArgsForCall[i].bucketName, fake.bucketObjectsArgsForCall[i].prefix
}

func (fake *FakeGCSClient) BucketObjectsReturns(result1 []string, result2 error) {
	fake.BucketObjectsStub = nil
	fake.bucketObjectsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) BucketObjectsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.BucketObjectsStub = nil
	if fake.bucketObjectsReturnsOnCall == nil {
		fake.bucketObjectsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.bucketObjectsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) ObjectGenerations(bucketName string, objectPath string) ([]int64, error) {
	fake.objectGenerationsMutex.Lock()
	ret, specificReturn := fake.objectGenerationsReturnsOnCall[len(fake.objectGenerationsArgsForCall)]
	fake.objectGenerationsArgsForCall = append(fake.objectGenerationsArgsForCall, struct {
		bucketName string
		objectPath string
	}{bucketName, objectPath})
	fake.recordInvocation("ObjectGenerations", []interface{}{bucketName, objectPath})
	fake.objectGenerationsMutex.Unlock()
	if fake.ObjectGenerationsStub != nil {
		return fake.ObjectGenerationsStub(bucketName, objectPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.objectGenerationsReturns.result1, fake.objectGenerationsReturns.result2
}

func (fake *FakeGCSClient) ObjectGenerationsCallCount() int {
	fake.objectGenerationsMutex.RLock()
	defer fake.objectGenerationsMutex.RUnlock()
	return len(fake.objectGenerationsArgsForCall)
}

func (fake *FakeGCSClient) ObjectGenerationsArgsForCall(i int) (string, string) {
	fake.objectGenerationsMutex.RLock()
	defer fake.objectGenerationsMutex.RUnlock()
	return fake.objectGenerationsArgsForCall[i].bucketName, fake.objectGenerationsArgsForCall[i].objectPath
}

func (fake *FakeGCSClient) ObjectGenerationsReturns(result1 []int64, result2 error) {
	fake.ObjectGenerationsStub = nil
	fake.objectGenerationsReturns = struct {
		result1 []int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) ObjectGenerationsReturnsOnCall(i int, result1 []int64, result2 error) {
	fake.ObjectGenerationsStub = nil
	if fake.objectGenerationsReturnsOnCall == nil {
		fake.objectGenerationsReturnsOnCall = make(map[int]struct {
			result1 []int64
			result2 error
		})
	}
	fake.objectGenerationsReturnsOnCall[i] = struct {
		result1 []int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) DownloadFile(bucketName string, objectPath string, generation int64, localPath string) error {
	fake.downloadFileMutex.Lock()
	ret, specificReturn := fake.downloadFileReturnsOnCall[len(fake.downloadFileArgsForCall)]
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
		localPath  string
	}{bucketName, objectPath, generation, localPath})
	fake.recordInvocation("DownloadFile", []interface{}{bucketName, objectPath, generation, localPath})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(bucketName, objectPath, generation, localPath)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downloadFileReturns.result1
}

func (fake *FakeGCSClient) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeGCSClient) DownloadFileArgsForCall(i int) (string, string, int64, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.downloadFileArgsForCall[i].bucketName, fake.downloadFileArgsForCall[i].objectPath, fake.downloadFileArgsForCall[i].generation, fake.downloadFileArgsForCall[i].localPath
}

func (fake *FakeGCSClient) DownloadFileReturns(result1 error) {
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGCSClient) DownloadFileReturnsOnCall(i int, result1 error) {
	fake.DownloadFileStub = nil
	if fake.downloadFileReturnsOnCall == nil {
		fake.downloadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGCSClient) UploadFile(bucketName string, objectPath string, objectContentType string, localPath string, predefinedACL string) (int64, error) {
	fake.uploadFileMutex.Lock()
	ret, specificReturn := fake.uploadFileReturnsOnCall[len(fake.uploadFileArgsForCall)]
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		bucketName        string
		objectPath        string
		objectContentType string
		localPath         string
		predefinedACL     string
	}{bucketName, objectPath, objectContentType, localPath, predefinedACL})
	fake.recordInvocation("UploadFile", []interface{}{bucketName, objectPath, objectContentType, localPath, predefinedACL})
	fake.uploadFileMutex.Unlock()
	if fake.UploadFileStub != nil {
		return fake.UploadFileStub(bucketName, objectPath, objectContentType, localPath, predefinedACL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadFileReturns.result1, fake.uploadFileReturns.result2
}

func (fake *FakeGCSClient) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeGCSClient) UploadFileArgsForCall(i int) (string, string, string, string, string) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return fake.uploadFileArgsForCall[i].bucketName, fake.uploadFileArgsForCall[i].objectPath, fake.uploadFileArgsForCall[i].objectContentType, fake.uploadFileArgsForCall[i].localPath, fake.uploadFileArgsForCall[i].predefinedACL
}

func (fake *FakeGCSClient) UploadFileReturns(result1 int64, result2 error) {
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) UploadFileReturnsOnCall(i int, result1 int64, result2 error) {
	fake.UploadFileStub = nil
	if fake.uploadFileReturnsOnCall == nil {
		fake.uploadFileReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.uploadFileReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) URL(bucketName string, objectPath string, generation int64) (string, error) {
	fake.uRLMutex.Lock()
	ret, specificReturn := fake.uRLReturnsOnCall[len(fake.uRLArgsForCall)]
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
	}{bucketName, objectPath, generation})
	fake.recordInvocation("URL", []interface{}{bucketName, objectPath, generation})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub(bucketName, objectPath, generation)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uRLReturns.result1, fake.uRLReturns.result2
}

func (fake *FakeGCSClient) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeGCSClient) URLArgsForCall(i int) (string, string, int64) {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return fake.uRLArgsForCall[i].bucketName, fake.uRLArgsForCall[i].objectPath, fake.uRLArgsForCall[i].generation
}

func (fake *FakeGCSClient) URLReturns(result1 string, result2 error) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) URLReturnsOnCall(i int, result1 string, result2 error) {
	fake.URLStub = nil
	if fake.uRLReturnsOnCall == nil {
		fake.uRLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.uRLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) DeleteObject(bucketName string, objectPath string, generation int64) error {
	fake.deleteObjectMutex.Lock()
	ret, specificReturn := fake.deleteObjectReturnsOnCall[len(fake.deleteObjectArgsForCall)]
	fake.deleteObjectArgsForCall = append(fake.deleteObjectArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
	}{bucketName, objectPath, generation})
	fake.recordInvocation("DeleteObject", []interface{}{bucketName, objectPath, generation})
	fake.deleteObjectMutex.Unlock()
	if fake.DeleteObjectStub != nil {
		return fake.DeleteObjectStub(bucketName, objectPath, generation)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteObjectReturns.result1
}

func (fake *FakeGCSClient) DeleteObjectCallCount() int {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return len(fake.deleteObjectArgsForCall)
}

func (fake *FakeGCSClient) DeleteObjectArgsForCall(i int) (string, string, int64) {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return fake.deleteObjectArgsForCall[i].bucketName, fake.deleteObjectArgsForCall[i].objectPath, fake.deleteObjectArgsForCall[i].generation
}

func (fake *FakeGCSClient) DeleteObjectReturns(result1 error) {
	fake.DeleteObjectStub = nil
	fake.deleteObjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGCSClient) DeleteObjectReturnsOnCall(i int, result1 error) {
	fake.DeleteObjectStub = nil
	if fake.deleteObjectReturnsOnCall == nil {
		fake.deleteObjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteObjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGCSClient) GetBucketObjectInfo(bucketName string, objectPath string) (*storage.Object, error) {
	fake.getBucketObjectInfoMutex.Lock()
	ret, specificReturn := fake.getBucketObjectInfoReturnsOnCall[len(fake.getBucketObjectInfoArgsForCall)]
	fake.getBucketObjectInfoArgsForCall = append(fake.getBucketObjectInfoArgsForCall, struct {
		bucketName string
		objectPath string
	}{bucketName, objectPath})
	fake.recordInvocation("GetBucketObjectInfo", []interface{}{bucketName, objectPath})
	fake.getBucketObjectInfoMutex.Unlock()
	if fake.GetBucketObjectInfoStub != nil {
		return fake.GetBucketObjectInfoStub(bucketName, objectPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getBucketObjectInfoReturns.result1, fake.getBucketObjectInfoReturns.result2
}

func (fake *FakeGCSClient) GetBucketObjectInfoCallCount() int {
	fake.getBucketObjectInfoMutex.RLock()
	defer fake.getBucketObjectInfoMutex.RUnlock()
	return len(fake.getBucketObjectInfoArgsForCall)
}

func (fake *FakeGCSClient) GetBucketObjectInfoArgsForCall(i int) (string, string) {
	fake.getBucketObjectInfoMutex.RLock()
	defer fake.getBucketObjectInfoMutex.RUnlock()
	return fake.getBucketObjectInfoArgsForCall[i].bucketName, fake.getBucketObjectInfoArgsForCall[i].objectPath
}

func (fake *FakeGCSClient) GetBucketObjectInfoReturns(result1 *storage.Object, result2 error) {
	fake.GetBucketObjectInfoStub = nil
	fake.getBucketObjectInfoReturns = struct {
		result1 *storage.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) GetBucketObjectInfoReturnsOnCall(i int, result1 *storage.Object, result2 error) {
	fake.GetBucketObjectInfoStub = nil
	if fake.getBucketObjectInfoReturnsOnCall == nil {
		fake.getBucketObjectInfoReturnsOnCall = make(map[int]struct {
			result1 *storage.Object
			result2 error
		})
	}
	fake.getBucketObjectInfoReturnsOnCall[i] = struct {
		result1 *storage.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bucketObjectsMutex.RLock()
	defer fake.bucketObjectsMutex.RUnlock()
	fake.objectGenerationsMutex.RLock()
	defer fake.objectGenerationsMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	fake.getBucketObjectInfoMutex.RLock()
	defer fake.getBucketObjectInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGCSClient) recordInvocation(key string, args []interface{}) {
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

var _ gcsresource.GCSClient = new(FakeGCSClient)
