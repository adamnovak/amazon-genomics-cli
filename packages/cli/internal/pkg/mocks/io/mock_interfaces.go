// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/mocks/io/interfaces.go

// Package iomocks is a generated GoMock package.
package iomocks

import (
	fs "io/fs"
	reflect "reflect"

	spec "github.com/aws/amazon-genomics-cli/internal/pkg/cli/spec"
	gomock "github.com/golang/mock/gomock"
	zerolog "github.com/rs/zerolog"
)

// MockOS is a mock of OS interface.
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *MockOSMockRecorder
}

// MockOSMockRecorder is the mock recorder for MockOS.
type MockOSMockRecorder struct {
	mock *MockOS
}

// NewMockOS creates a new mock instance.
func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &MockOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOS) EXPECT() *MockOSMockRecorder {
	return m.recorder
}

// Chdir mocks base method.
func (m *MockOS) Chdir(dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chdir", dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chdir indicates an expected call of Chdir.
func (mr *MockOSMockRecorder) Chdir(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chdir", reflect.TypeOf((*MockOS)(nil).Chdir), dir)
}

// MkdirTemp mocks base method.
func (m *MockOS) MkdirTemp(dir, pattern string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirTemp", dir, pattern)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MkdirTemp indicates an expected call of MkdirTemp.
func (mr *MockOSMockRecorder) MkdirTemp(dir, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirTemp", reflect.TypeOf((*MockOS)(nil).MkdirTemp), dir, pattern)
}

// Remove mocks base method.
func (m *MockOS) Remove(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockOSMockRecorder) Remove(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockOS)(nil).Remove), name)
}

// RemoveAll mocks base method.
func (m *MockOS) RemoveAll(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockOSMockRecorder) RemoveAll(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockOS)(nil).RemoveAll), path)
}

// Stat mocks base method.
func (m *MockOS) Stat(name string) (fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", name)
	ret0, _ := ret[0].(fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockOSMockRecorder) Stat(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockOS)(nil).Stat), name)
}

// UserHomeDir mocks base method.
func (m *MockOS) UserHomeDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHomeDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHomeDir indicates an expected call of UserHomeDir.
func (mr *MockOSMockRecorder) UserHomeDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHomeDir", reflect.TypeOf((*MockOS)(nil).UserHomeDir))
}

// MockZip is a mock of Zip interface.
type MockZip struct {
	ctrl     *gomock.Controller
	recorder *MockZipMockRecorder
}

// MockZipMockRecorder is the mock recorder for MockZip.
type MockZipMockRecorder struct {
	mock *MockZip
}

// NewMockZip creates a new mock instance.
func NewMockZip(ctrl *gomock.Controller) *MockZip {
	mock := &MockZip{ctrl: ctrl}
	mock.recorder = &MockZipMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZip) EXPECT() *MockZipMockRecorder {
	return m.recorder
}

// CompressToTmp mocks base method.
func (m *MockZip) CompressToTmp(srcPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompressToTmp", srcPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompressToTmp indicates an expected call of CompressToTmp.
func (mr *MockZipMockRecorder) CompressToTmp(srcPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompressToTmp", reflect.TypeOf((*MockZip)(nil).CompressToTmp), srcPath)
}

// MockTmp is a mock of Tmp interface.
type MockTmp struct {
	ctrl     *gomock.Controller
	recorder *MockTmpMockRecorder
}

// MockTmpMockRecorder is the mock recorder for MockTmp.
type MockTmpMockRecorder struct {
	mock *MockTmp
}

// NewMockTmp creates a new mock instance.
func NewMockTmp(ctrl *gomock.Controller) *MockTmp {
	mock := &MockTmp{ctrl: ctrl}
	mock.recorder = &MockTmpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTmp) EXPECT() *MockTmpMockRecorder {
	return m.recorder
}

// TempDir mocks base method.
func (m *MockTmp) TempDir(dir, pattern string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TempDir", dir, pattern)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TempDir indicates an expected call of TempDir.
func (mr *MockTmpMockRecorder) TempDir(dir, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TempDir", reflect.TypeOf((*MockTmp)(nil).TempDir), dir, pattern)
}

// Write mocks base method.
func (m *MockTmp) Write(namePattern, content string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", namePattern, content)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockTmpMockRecorder) Write(namePattern, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTmp)(nil).Write), namePattern, content)
}

// MockFileReader is a mock of FileReader interface.
type MockFileReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileReaderMockRecorder
}

// MockFileReaderMockRecorder is the mock recorder for MockFileReader.
type MockFileReaderMockRecorder struct {
	mock *MockFileReader
}

// NewMockFileReader creates a new mock instance.
func NewMockFileReader(ctrl *gomock.Controller) *MockFileReader {
	mock := &MockFileReader{ctrl: ctrl}
	mock.recorder = &MockFileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileReader) EXPECT() *MockFileReaderMockRecorder {
	return m.recorder
}

// ReadFile mocks base method.
func (m *MockFileReader) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockFileReaderMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileReader)(nil).ReadFile), arg0)
}

// MockFileWriter is a mock of FileWriter interface.
type MockFileWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileWriterMockRecorder
}

// MockFileWriterMockRecorder is the mock recorder for MockFileWriter.
type MockFileWriterMockRecorder struct {
	mock *MockFileWriter
}

// NewMockFileWriter creates a new mock instance.
func NewMockFileWriter(ctrl *gomock.Controller) *MockFileWriter {
	mock := &MockFileWriter{ctrl: ctrl}
	mock.recorder = &MockFileWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileWriter) EXPECT() *MockFileWriterMockRecorder {
	return m.recorder
}

// WriteFile mocks base method.
func (m *MockFileWriter) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockFileWriterMockRecorder) WriteFile(filename, data, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFileWriter)(nil).WriteFile), filename, data, perm)
}

// MockFormat is a mock of Format interface.
type MockFormat struct {
	ctrl     *gomock.Controller
	recorder *MockFormatMockRecorder
}

// MockFormatMockRecorder is the mock recorder for MockFormat.
type MockFormatMockRecorder struct {
	mock *MockFormat
}

// NewMockFormat creates a new mock instance.
func NewMockFormat(ctrl *gomock.Controller) *MockFormat {
	mock := &MockFormat{ctrl: ctrl}
	mock.recorder = &MockFormatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFormat) EXPECT() *MockFormatMockRecorder {
	return m.recorder
}

// LogsPrintLn mocks base method.
func (m *MockFormat) LogsPrintLn(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogsPrintLn", varargs...)
}

// LogsPrintLn indicates an expected call of LogsPrintLn.
func (mr *MockFormatMockRecorder) LogsPrintLn(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsPrintLn", reflect.TypeOf((*MockFormat)(nil).LogsPrintLn), args...)
}

// MockLog is a mock of Log interface.
type MockLog struct {
	ctrl     *gomock.Controller
	recorder *MockLogMockRecorder
}

// MockLogMockRecorder is the mock recorder for MockLog.
type MockLogMockRecorder struct {
	mock *MockLog
}

// NewMockLog creates a new mock instance.
func NewMockLog(ctrl *gomock.Controller) *MockLog {
	mock := &MockLog{ctrl: ctrl}
	mock.recorder = &MockLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLog) EXPECT() *MockLogMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockLog) Info() *zerolog.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(*zerolog.Event)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockLogMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLog)(nil).Info))
}

// MockSpec is a mock of Spec interface.
type MockSpec struct {
	ctrl     *gomock.Controller
	recorder *MockSpecMockRecorder
}

// MockSpecMockRecorder is the mock recorder for MockSpec.
type MockSpecMockRecorder struct {
	mock *MockSpec
}

// NewMockSpec creates a new mock instance.
func NewMockSpec(ctrl *gomock.Controller) *MockSpec {
	mock := &MockSpec{ctrl: ctrl}
	mock.recorder = &MockSpecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpec) EXPECT() *MockSpecMockRecorder {
	return m.recorder
}

// FromJson mocks base method.
func (m *MockSpec) FromJson(manifestFilePath string) (spec.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromJson", manifestFilePath)
	ret0, _ := ret[0].(spec.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromJson indicates an expected call of FromJson.
func (mr *MockSpecMockRecorder) FromJson(manifestFilePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromJson", reflect.TypeOf((*MockSpec)(nil).FromJson), manifestFilePath)
}

// MockJson is a mock of Json interface.
type MockJson struct {
	ctrl     *gomock.Controller
	recorder *MockJsonMockRecorder
}

// MockJsonMockRecorder is the mock recorder for MockJson.
type MockJsonMockRecorder struct {
	mock *MockJson
}

// NewMockJson creates a new mock instance.
func NewMockJson(ctrl *gomock.Controller) *MockJson {
	mock := &MockJson{ctrl: ctrl}
	mock.recorder = &MockJsonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJson) EXPECT() *MockJsonMockRecorder {
	return m.recorder
}

// Marshal mocks base method.
func (m *MockJson) Marshal(v interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal", v)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal.
func (mr *MockJsonMockRecorder) Marshal(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockJson)(nil).Marshal), v)
}

// Unmarshal mocks base method.
func (m *MockJson) Unmarshal(data []byte, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockJsonMockRecorder) Unmarshal(data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockJson)(nil).Unmarshal), data, v)
}
