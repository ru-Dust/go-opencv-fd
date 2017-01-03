// Automatically generated by MockGen. DO NOT EDIT!
// Source: src/detector/object_detector.go

package detector

import (
	gomock "github.com/golang/mock/gomock"
	opencv "github.com/lazywei/go-opencv/opencv"
	unsafe "unsafe"
)

// Mock of IplImageInterface interface
type MockIplImageInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockIplImageInterfaceRecorder
}

// Recorder for MockIplImageInterface (not exported)
type _MockIplImageInterfaceRecorder struct {
	mock *MockIplImageInterface
}

func NewMockIplImageInterface(ctrl *gomock.Controller) *MockIplImageInterface {
	mock := &MockIplImageInterface{ctrl: ctrl}
	mock.recorder = &_MockIplImageInterfaceRecorder{mock}
	return mock
}

func (_m *MockIplImageInterface) EXPECT() *_MockIplImageInterfaceRecorder {
	return _m.recorder
}

func (_m *MockIplImageInterface) Channels() int {
	ret := _m.ctrl.Call(_m, "Channels")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) Channels() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Channels")
}

func (_m *MockIplImageInterface) Depth() int {
	ret := _m.ctrl.Call(_m, "Depth")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) Depth() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Depth")
}

func (_m *MockIplImageInterface) Origin() int {
	ret := _m.ctrl.Call(_m, "Origin")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) Origin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Origin")
}

func (_m *MockIplImageInterface) Width() int {
	ret := _m.ctrl.Call(_m, "Width")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) Width() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Width")
}

func (_m *MockIplImageInterface) Height() int {
	ret := _m.ctrl.Call(_m, "Height")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) Height() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Height")
}

func (_m *MockIplImageInterface) WidthStep() int {
	ret := _m.ctrl.Call(_m, "WidthStep")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) WidthStep() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WidthStep")
}

func (_m *MockIplImageInterface) ImageSize() int {
	ret := _m.ctrl.Call(_m, "ImageSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) ImageSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageSize")
}

func (_m *MockIplImageInterface) ImageData() unsafe.Pointer {
	ret := _m.ctrl.Call(_m, "ImageData")
	ret0, _ := ret[0].(unsafe.Pointer)
	return ret0
}

func (_mr *_MockIplImageInterfaceRecorder) ImageData() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageData")
}

// Mock of HaarCascadeInterface interface
type MockHaarCascadeInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockHaarCascadeInterfaceRecorder
}

// Recorder for MockHaarCascadeInterface (not exported)
type _MockHaarCascadeInterfaceRecorder struct {
	mock *MockHaarCascadeInterface
}

func NewMockHaarCascadeInterface(ctrl *gomock.Controller) *MockHaarCascadeInterface {
	mock := &MockHaarCascadeInterface{ctrl: ctrl}
	mock.recorder = &_MockHaarCascadeInterfaceRecorder{mock}
	return mock
}

func (_m *MockHaarCascadeInterface) EXPECT() *_MockHaarCascadeInterfaceRecorder {
	return _m.recorder
}

func (_m *MockHaarCascadeInterface) DetectObjects(img *opencv.IplImage) []*opencv.Rect {
	ret := _m.ctrl.Call(_m, "DetectObjects", img)
	ret0, _ := ret[0].([]*opencv.Rect)
	return ret0
}

func (_mr *_MockHaarCascadeInterfaceRecorder) DetectObjects(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DetectObjects", arg0)
}

func (_m *MockHaarCascadeInterface) Release() {
	_m.ctrl.Call(_m, "Release")
}

func (_mr *_MockHaarCascadeInterfaceRecorder) Release() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Release")
}

// Mock of ObjectDetectorInterface interface
type MockObjectDetectorInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockObjectDetectorInterfaceRecorder
}

// Recorder for MockObjectDetectorInterface (not exported)
type _MockObjectDetectorInterfaceRecorder struct {
	mock *MockObjectDetectorInterface
}

func NewMockObjectDetectorInterface(ctrl *gomock.Controller) *MockObjectDetectorInterface {
	mock := &MockObjectDetectorInterface{ctrl: ctrl}
	mock.recorder = &_MockObjectDetectorInterfaceRecorder{mock}
	return mock
}

func (_m *MockObjectDetectorInterface) EXPECT() *_MockObjectDetectorInterfaceRecorder {
	return _m.recorder
}

func (_m *MockObjectDetectorInterface) Detect(img IplImageInterface) {
	_m.ctrl.Call(_m, "Detect", img)
}

func (_mr *_MockObjectDetectorInterfaceRecorder) Detect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Detect", arg0)
}

func (_m *MockObjectDetectorInterface) Draw(img IplImageInterface, value *opencv.Rect) {
	_m.ctrl.Call(_m, "Draw", img, value)
}

func (_mr *_MockObjectDetectorInterfaceRecorder) Draw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Draw", arg0, arg1)
}
