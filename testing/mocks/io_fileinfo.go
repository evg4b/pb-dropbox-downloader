package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i io/fs.FileInfo -o ./testing/mocks\io_fileinfo.go -n FileInfo

import (
	mm_fs "io/fs"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// FileInfo implements fs.FileInfo
type FileInfo struct {
	t minimock.Tester

	funcIsDir          func() (b1 bool)
	inspectFuncIsDir   func()
	afterIsDirCounter  uint64
	beforeIsDirCounter uint64
	IsDirMock          mFileInfoIsDir

	funcModTime          func() (t1 time.Time)
	inspectFuncModTime   func()
	afterModTimeCounter  uint64
	beforeModTimeCounter uint64
	ModTimeMock          mFileInfoModTime

	funcMode          func() (f1 mm_fs.FileMode)
	inspectFuncMode   func()
	afterModeCounter  uint64
	beforeModeCounter uint64
	ModeMock          mFileInfoMode

	funcName          func() (s1 string)
	inspectFuncName   func()
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mFileInfoName

	funcSize          func() (i1 int64)
	inspectFuncSize   func()
	afterSizeCounter  uint64
	beforeSizeCounter uint64
	SizeMock          mFileInfoSize

	funcSys          func() (p1 interface{})
	inspectFuncSys   func()
	afterSysCounter  uint64
	beforeSysCounter uint64
	SysMock          mFileInfoSys
}

// NewFileInfo returns a mock for fs.FileInfo
func NewFileInfo(t minimock.Tester) *FileInfo {
	m := &FileInfo{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.IsDirMock = mFileInfoIsDir{mock: m}

	m.ModTimeMock = mFileInfoModTime{mock: m}

	m.ModeMock = mFileInfoMode{mock: m}

	m.NameMock = mFileInfoName{mock: m}

	m.SizeMock = mFileInfoSize{mock: m}

	m.SysMock = mFileInfoSys{mock: m}

	return m
}

type mFileInfoIsDir struct {
	mock               *FileInfo
	defaultExpectation *FileInfoIsDirExpectation
	expectations       []*FileInfoIsDirExpectation
}

// FileInfoIsDirExpectation specifies expectation struct of the FileInfo.IsDir
type FileInfoIsDirExpectation struct {
	mock *FileInfo

	results *FileInfoIsDirResults
	Counter uint64
}

// FileInfoIsDirResults contains results of the FileInfo.IsDir
type FileInfoIsDirResults struct {
	b1 bool
}

// Expect sets up expected params for FileInfo.IsDir
func (mmIsDir *mFileInfoIsDir) Expect() *mFileInfoIsDir {
	if mmIsDir.mock.funcIsDir != nil {
		mmIsDir.mock.t.Fatalf("FileInfo.IsDir mock is already set by Set")
	}

	if mmIsDir.defaultExpectation == nil {
		mmIsDir.defaultExpectation = &FileInfoIsDirExpectation{}
	}

	return mmIsDir
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.IsDir
func (mmIsDir *mFileInfoIsDir) Inspect(f func()) *mFileInfoIsDir {
	if mmIsDir.mock.inspectFuncIsDir != nil {
		mmIsDir.mock.t.Fatalf("Inspect function is already set for FileInfo.IsDir")
	}

	mmIsDir.mock.inspectFuncIsDir = f

	return mmIsDir
}

// Return sets up results that will be returned by FileInfo.IsDir
func (mmIsDir *mFileInfoIsDir) Return(b1 bool) *FileInfo {
	if mmIsDir.mock.funcIsDir != nil {
		mmIsDir.mock.t.Fatalf("FileInfo.IsDir mock is already set by Set")
	}

	if mmIsDir.defaultExpectation == nil {
		mmIsDir.defaultExpectation = &FileInfoIsDirExpectation{mock: mmIsDir.mock}
	}
	mmIsDir.defaultExpectation.results = &FileInfoIsDirResults{b1}
	return mmIsDir.mock
}

//Set uses given function f to mock the FileInfo.IsDir method
func (mmIsDir *mFileInfoIsDir) Set(f func() (b1 bool)) *FileInfo {
	if mmIsDir.defaultExpectation != nil {
		mmIsDir.mock.t.Fatalf("Default expectation is already set for the FileInfo.IsDir method")
	}

	if len(mmIsDir.expectations) > 0 {
		mmIsDir.mock.t.Fatalf("Some expectations are already set for the FileInfo.IsDir method")
	}

	mmIsDir.mock.funcIsDir = f
	return mmIsDir.mock
}

// IsDir implements fs.FileInfo
func (mmIsDir *FileInfo) IsDir() (b1 bool) {
	mm_atomic.AddUint64(&mmIsDir.beforeIsDirCounter, 1)
	defer mm_atomic.AddUint64(&mmIsDir.afterIsDirCounter, 1)

	if mmIsDir.inspectFuncIsDir != nil {
		mmIsDir.inspectFuncIsDir()
	}

	if mmIsDir.IsDirMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsDir.IsDirMock.defaultExpectation.Counter, 1)

		mm_results := mmIsDir.IsDirMock.defaultExpectation.results
		if mm_results == nil {
			mmIsDir.t.Fatal("No results are set for the FileInfo.IsDir")
		}
		return (*mm_results).b1
	}
	if mmIsDir.funcIsDir != nil {
		return mmIsDir.funcIsDir()
	}
	mmIsDir.t.Fatalf("Unexpected call to FileInfo.IsDir.")
	return
}

// IsDirAfterCounter returns a count of finished FileInfo.IsDir invocations
func (mmIsDir *FileInfo) IsDirAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsDir.afterIsDirCounter)
}

// IsDirBeforeCounter returns a count of FileInfo.IsDir invocations
func (mmIsDir *FileInfo) IsDirBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsDir.beforeIsDirCounter)
}

// MinimockIsDirDone returns true if the count of the IsDir invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockIsDirDone() bool {
	for _, e := range m.IsDirMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsDirMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsDirCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsDir != nil && mm_atomic.LoadUint64(&m.afterIsDirCounter) < 1 {
		return false
	}
	return true
}

// MinimockIsDirInspect logs each unmet expectation
func (m *FileInfo) MinimockIsDirInspect() {
	for _, e := range m.IsDirMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.IsDir")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IsDirMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIsDirCounter) < 1 {
		m.t.Error("Expected call to FileInfo.IsDir")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsDir != nil && mm_atomic.LoadUint64(&m.afterIsDirCounter) < 1 {
		m.t.Error("Expected call to FileInfo.IsDir")
	}
}

type mFileInfoModTime struct {
	mock               *FileInfo
	defaultExpectation *FileInfoModTimeExpectation
	expectations       []*FileInfoModTimeExpectation
}

// FileInfoModTimeExpectation specifies expectation struct of the FileInfo.ModTime
type FileInfoModTimeExpectation struct {
	mock *FileInfo

	results *FileInfoModTimeResults
	Counter uint64
}

// FileInfoModTimeResults contains results of the FileInfo.ModTime
type FileInfoModTimeResults struct {
	t1 time.Time
}

// Expect sets up expected params for FileInfo.ModTime
func (mmModTime *mFileInfoModTime) Expect() *mFileInfoModTime {
	if mmModTime.mock.funcModTime != nil {
		mmModTime.mock.t.Fatalf("FileInfo.ModTime mock is already set by Set")
	}

	if mmModTime.defaultExpectation == nil {
		mmModTime.defaultExpectation = &FileInfoModTimeExpectation{}
	}

	return mmModTime
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.ModTime
func (mmModTime *mFileInfoModTime) Inspect(f func()) *mFileInfoModTime {
	if mmModTime.mock.inspectFuncModTime != nil {
		mmModTime.mock.t.Fatalf("Inspect function is already set for FileInfo.ModTime")
	}

	mmModTime.mock.inspectFuncModTime = f

	return mmModTime
}

// Return sets up results that will be returned by FileInfo.ModTime
func (mmModTime *mFileInfoModTime) Return(t1 time.Time) *FileInfo {
	if mmModTime.mock.funcModTime != nil {
		mmModTime.mock.t.Fatalf("FileInfo.ModTime mock is already set by Set")
	}

	if mmModTime.defaultExpectation == nil {
		mmModTime.defaultExpectation = &FileInfoModTimeExpectation{mock: mmModTime.mock}
	}
	mmModTime.defaultExpectation.results = &FileInfoModTimeResults{t1}
	return mmModTime.mock
}

//Set uses given function f to mock the FileInfo.ModTime method
func (mmModTime *mFileInfoModTime) Set(f func() (t1 time.Time)) *FileInfo {
	if mmModTime.defaultExpectation != nil {
		mmModTime.mock.t.Fatalf("Default expectation is already set for the FileInfo.ModTime method")
	}

	if len(mmModTime.expectations) > 0 {
		mmModTime.mock.t.Fatalf("Some expectations are already set for the FileInfo.ModTime method")
	}

	mmModTime.mock.funcModTime = f
	return mmModTime.mock
}

// ModTime implements fs.FileInfo
func (mmModTime *FileInfo) ModTime() (t1 time.Time) {
	mm_atomic.AddUint64(&mmModTime.beforeModTimeCounter, 1)
	defer mm_atomic.AddUint64(&mmModTime.afterModTimeCounter, 1)

	if mmModTime.inspectFuncModTime != nil {
		mmModTime.inspectFuncModTime()
	}

	if mmModTime.ModTimeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmModTime.ModTimeMock.defaultExpectation.Counter, 1)

		mm_results := mmModTime.ModTimeMock.defaultExpectation.results
		if mm_results == nil {
			mmModTime.t.Fatal("No results are set for the FileInfo.ModTime")
		}
		return (*mm_results).t1
	}
	if mmModTime.funcModTime != nil {
		return mmModTime.funcModTime()
	}
	mmModTime.t.Fatalf("Unexpected call to FileInfo.ModTime.")
	return
}

// ModTimeAfterCounter returns a count of finished FileInfo.ModTime invocations
func (mmModTime *FileInfo) ModTimeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmModTime.afterModTimeCounter)
}

// ModTimeBeforeCounter returns a count of FileInfo.ModTime invocations
func (mmModTime *FileInfo) ModTimeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmModTime.beforeModTimeCounter)
}

// MinimockModTimeDone returns true if the count of the ModTime invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockModTimeDone() bool {
	for _, e := range m.ModTimeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ModTimeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterModTimeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcModTime != nil && mm_atomic.LoadUint64(&m.afterModTimeCounter) < 1 {
		return false
	}
	return true
}

// MinimockModTimeInspect logs each unmet expectation
func (m *FileInfo) MinimockModTimeInspect() {
	for _, e := range m.ModTimeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.ModTime")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ModTimeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterModTimeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.ModTime")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcModTime != nil && mm_atomic.LoadUint64(&m.afterModTimeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.ModTime")
	}
}

type mFileInfoMode struct {
	mock               *FileInfo
	defaultExpectation *FileInfoModeExpectation
	expectations       []*FileInfoModeExpectation
}

// FileInfoModeExpectation specifies expectation struct of the FileInfo.Mode
type FileInfoModeExpectation struct {
	mock *FileInfo

	results *FileInfoModeResults
	Counter uint64
}

// FileInfoModeResults contains results of the FileInfo.Mode
type FileInfoModeResults struct {
	f1 mm_fs.FileMode
}

// Expect sets up expected params for FileInfo.Mode
func (mmMode *mFileInfoMode) Expect() *mFileInfoMode {
	if mmMode.mock.funcMode != nil {
		mmMode.mock.t.Fatalf("FileInfo.Mode mock is already set by Set")
	}

	if mmMode.defaultExpectation == nil {
		mmMode.defaultExpectation = &FileInfoModeExpectation{}
	}

	return mmMode
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.Mode
func (mmMode *mFileInfoMode) Inspect(f func()) *mFileInfoMode {
	if mmMode.mock.inspectFuncMode != nil {
		mmMode.mock.t.Fatalf("Inspect function is already set for FileInfo.Mode")
	}

	mmMode.mock.inspectFuncMode = f

	return mmMode
}

// Return sets up results that will be returned by FileInfo.Mode
func (mmMode *mFileInfoMode) Return(f1 mm_fs.FileMode) *FileInfo {
	if mmMode.mock.funcMode != nil {
		mmMode.mock.t.Fatalf("FileInfo.Mode mock is already set by Set")
	}

	if mmMode.defaultExpectation == nil {
		mmMode.defaultExpectation = &FileInfoModeExpectation{mock: mmMode.mock}
	}
	mmMode.defaultExpectation.results = &FileInfoModeResults{f1}
	return mmMode.mock
}

//Set uses given function f to mock the FileInfo.Mode method
func (mmMode *mFileInfoMode) Set(f func() (f1 mm_fs.FileMode)) *FileInfo {
	if mmMode.defaultExpectation != nil {
		mmMode.mock.t.Fatalf("Default expectation is already set for the FileInfo.Mode method")
	}

	if len(mmMode.expectations) > 0 {
		mmMode.mock.t.Fatalf("Some expectations are already set for the FileInfo.Mode method")
	}

	mmMode.mock.funcMode = f
	return mmMode.mock
}

// Mode implements fs.FileInfo
func (mmMode *FileInfo) Mode() (f1 mm_fs.FileMode) {
	mm_atomic.AddUint64(&mmMode.beforeModeCounter, 1)
	defer mm_atomic.AddUint64(&mmMode.afterModeCounter, 1)

	if mmMode.inspectFuncMode != nil {
		mmMode.inspectFuncMode()
	}

	if mmMode.ModeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMode.ModeMock.defaultExpectation.Counter, 1)

		mm_results := mmMode.ModeMock.defaultExpectation.results
		if mm_results == nil {
			mmMode.t.Fatal("No results are set for the FileInfo.Mode")
		}
		return (*mm_results).f1
	}
	if mmMode.funcMode != nil {
		return mmMode.funcMode()
	}
	mmMode.t.Fatalf("Unexpected call to FileInfo.Mode.")
	return
}

// ModeAfterCounter returns a count of finished FileInfo.Mode invocations
func (mmMode *FileInfo) ModeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMode.afterModeCounter)
}

// ModeBeforeCounter returns a count of FileInfo.Mode invocations
func (mmMode *FileInfo) ModeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMode.beforeModeCounter)
}

// MinimockModeDone returns true if the count of the Mode invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockModeDone() bool {
	for _, e := range m.ModeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ModeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterModeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMode != nil && mm_atomic.LoadUint64(&m.afterModeCounter) < 1 {
		return false
	}
	return true
}

// MinimockModeInspect logs each unmet expectation
func (m *FileInfo) MinimockModeInspect() {
	for _, e := range m.ModeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.Mode")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ModeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterModeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Mode")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMode != nil && mm_atomic.LoadUint64(&m.afterModeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Mode")
	}
}

type mFileInfoName struct {
	mock               *FileInfo
	defaultExpectation *FileInfoNameExpectation
	expectations       []*FileInfoNameExpectation
}

// FileInfoNameExpectation specifies expectation struct of the FileInfo.Name
type FileInfoNameExpectation struct {
	mock *FileInfo

	results *FileInfoNameResults
	Counter uint64
}

// FileInfoNameResults contains results of the FileInfo.Name
type FileInfoNameResults struct {
	s1 string
}

// Expect sets up expected params for FileInfo.Name
func (mmName *mFileInfoName) Expect() *mFileInfoName {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("FileInfo.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &FileInfoNameExpectation{}
	}

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.Name
func (mmName *mFileInfoName) Inspect(f func()) *mFileInfoName {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for FileInfo.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by FileInfo.Name
func (mmName *mFileInfoName) Return(s1 string) *FileInfo {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("FileInfo.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &FileInfoNameExpectation{mock: mmName.mock}
	}
	mmName.defaultExpectation.results = &FileInfoNameResults{s1}
	return mmName.mock
}

//Set uses given function f to mock the FileInfo.Name method
func (mmName *mFileInfoName) Set(f func() (s1 string)) *FileInfo {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the FileInfo.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the FileInfo.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Name implements fs.FileInfo
func (mmName *FileInfo) Name() (s1 string) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName()
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)

		mm_results := mmName.NameMock.defaultExpectation.results
		if mm_results == nil {
			mmName.t.Fatal("No results are set for the FileInfo.Name")
		}
		return (*mm_results).s1
	}
	if mmName.funcName != nil {
		return mmName.funcName()
	}
	mmName.t.Fatalf("Unexpected call to FileInfo.Name.")
	return
}

// NameAfterCounter returns a count of finished FileInfo.Name invocations
func (mmName *FileInfo) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of FileInfo.Name invocations
func (mmName *FileInfo) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockNameInspect logs each unmet expectation
func (m *FileInfo) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.Name")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Name")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Name")
	}
}

type mFileInfoSize struct {
	mock               *FileInfo
	defaultExpectation *FileInfoSizeExpectation
	expectations       []*FileInfoSizeExpectation
}

// FileInfoSizeExpectation specifies expectation struct of the FileInfo.Size
type FileInfoSizeExpectation struct {
	mock *FileInfo

	results *FileInfoSizeResults
	Counter uint64
}

// FileInfoSizeResults contains results of the FileInfo.Size
type FileInfoSizeResults struct {
	i1 int64
}

// Expect sets up expected params for FileInfo.Size
func (mmSize *mFileInfoSize) Expect() *mFileInfoSize {
	if mmSize.mock.funcSize != nil {
		mmSize.mock.t.Fatalf("FileInfo.Size mock is already set by Set")
	}

	if mmSize.defaultExpectation == nil {
		mmSize.defaultExpectation = &FileInfoSizeExpectation{}
	}

	return mmSize
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.Size
func (mmSize *mFileInfoSize) Inspect(f func()) *mFileInfoSize {
	if mmSize.mock.inspectFuncSize != nil {
		mmSize.mock.t.Fatalf("Inspect function is already set for FileInfo.Size")
	}

	mmSize.mock.inspectFuncSize = f

	return mmSize
}

// Return sets up results that will be returned by FileInfo.Size
func (mmSize *mFileInfoSize) Return(i1 int64) *FileInfo {
	if mmSize.mock.funcSize != nil {
		mmSize.mock.t.Fatalf("FileInfo.Size mock is already set by Set")
	}

	if mmSize.defaultExpectation == nil {
		mmSize.defaultExpectation = &FileInfoSizeExpectation{mock: mmSize.mock}
	}
	mmSize.defaultExpectation.results = &FileInfoSizeResults{i1}
	return mmSize.mock
}

//Set uses given function f to mock the FileInfo.Size method
func (mmSize *mFileInfoSize) Set(f func() (i1 int64)) *FileInfo {
	if mmSize.defaultExpectation != nil {
		mmSize.mock.t.Fatalf("Default expectation is already set for the FileInfo.Size method")
	}

	if len(mmSize.expectations) > 0 {
		mmSize.mock.t.Fatalf("Some expectations are already set for the FileInfo.Size method")
	}

	mmSize.mock.funcSize = f
	return mmSize.mock
}

// Size implements fs.FileInfo
func (mmSize *FileInfo) Size() (i1 int64) {
	mm_atomic.AddUint64(&mmSize.beforeSizeCounter, 1)
	defer mm_atomic.AddUint64(&mmSize.afterSizeCounter, 1)

	if mmSize.inspectFuncSize != nil {
		mmSize.inspectFuncSize()
	}

	if mmSize.SizeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSize.SizeMock.defaultExpectation.Counter, 1)

		mm_results := mmSize.SizeMock.defaultExpectation.results
		if mm_results == nil {
			mmSize.t.Fatal("No results are set for the FileInfo.Size")
		}
		return (*mm_results).i1
	}
	if mmSize.funcSize != nil {
		return mmSize.funcSize()
	}
	mmSize.t.Fatalf("Unexpected call to FileInfo.Size.")
	return
}

// SizeAfterCounter returns a count of finished FileInfo.Size invocations
func (mmSize *FileInfo) SizeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSize.afterSizeCounter)
}

// SizeBeforeCounter returns a count of FileInfo.Size invocations
func (mmSize *FileInfo) SizeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSize.beforeSizeCounter)
}

// MinimockSizeDone returns true if the count of the Size invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockSizeDone() bool {
	for _, e := range m.SizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSizeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSize != nil && mm_atomic.LoadUint64(&m.afterSizeCounter) < 1 {
		return false
	}
	return true
}

// MinimockSizeInspect logs each unmet expectation
func (m *FileInfo) MinimockSizeInspect() {
	for _, e := range m.SizeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.Size")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SizeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSizeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Size")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSize != nil && mm_atomic.LoadUint64(&m.afterSizeCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Size")
	}
}

type mFileInfoSys struct {
	mock               *FileInfo
	defaultExpectation *FileInfoSysExpectation
	expectations       []*FileInfoSysExpectation
}

// FileInfoSysExpectation specifies expectation struct of the FileInfo.Sys
type FileInfoSysExpectation struct {
	mock *FileInfo

	results *FileInfoSysResults
	Counter uint64
}

// FileInfoSysResults contains results of the FileInfo.Sys
type FileInfoSysResults struct {
	p1 interface{}
}

// Expect sets up expected params for FileInfo.Sys
func (mmSys *mFileInfoSys) Expect() *mFileInfoSys {
	if mmSys.mock.funcSys != nil {
		mmSys.mock.t.Fatalf("FileInfo.Sys mock is already set by Set")
	}

	if mmSys.defaultExpectation == nil {
		mmSys.defaultExpectation = &FileInfoSysExpectation{}
	}

	return mmSys
}

// Inspect accepts an inspector function that has same arguments as the FileInfo.Sys
func (mmSys *mFileInfoSys) Inspect(f func()) *mFileInfoSys {
	if mmSys.mock.inspectFuncSys != nil {
		mmSys.mock.t.Fatalf("Inspect function is already set for FileInfo.Sys")
	}

	mmSys.mock.inspectFuncSys = f

	return mmSys
}

// Return sets up results that will be returned by FileInfo.Sys
func (mmSys *mFileInfoSys) Return(p1 interface{}) *FileInfo {
	if mmSys.mock.funcSys != nil {
		mmSys.mock.t.Fatalf("FileInfo.Sys mock is already set by Set")
	}

	if mmSys.defaultExpectation == nil {
		mmSys.defaultExpectation = &FileInfoSysExpectation{mock: mmSys.mock}
	}
	mmSys.defaultExpectation.results = &FileInfoSysResults{p1}
	return mmSys.mock
}

//Set uses given function f to mock the FileInfo.Sys method
func (mmSys *mFileInfoSys) Set(f func() (p1 interface{})) *FileInfo {
	if mmSys.defaultExpectation != nil {
		mmSys.mock.t.Fatalf("Default expectation is already set for the FileInfo.Sys method")
	}

	if len(mmSys.expectations) > 0 {
		mmSys.mock.t.Fatalf("Some expectations are already set for the FileInfo.Sys method")
	}

	mmSys.mock.funcSys = f
	return mmSys.mock
}

// Sys implements fs.FileInfo
func (mmSys *FileInfo) Sys() (p1 interface{}) {
	mm_atomic.AddUint64(&mmSys.beforeSysCounter, 1)
	defer mm_atomic.AddUint64(&mmSys.afterSysCounter, 1)

	if mmSys.inspectFuncSys != nil {
		mmSys.inspectFuncSys()
	}

	if mmSys.SysMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSys.SysMock.defaultExpectation.Counter, 1)

		mm_results := mmSys.SysMock.defaultExpectation.results
		if mm_results == nil {
			mmSys.t.Fatal("No results are set for the FileInfo.Sys")
		}
		return (*mm_results).p1
	}
	if mmSys.funcSys != nil {
		return mmSys.funcSys()
	}
	mmSys.t.Fatalf("Unexpected call to FileInfo.Sys.")
	return
}

// SysAfterCounter returns a count of finished FileInfo.Sys invocations
func (mmSys *FileInfo) SysAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSys.afterSysCounter)
}

// SysBeforeCounter returns a count of FileInfo.Sys invocations
func (mmSys *FileInfo) SysBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSys.beforeSysCounter)
}

// MinimockSysDone returns true if the count of the Sys invocations corresponds
// the number of defined expectations
func (m *FileInfo) MinimockSysDone() bool {
	for _, e := range m.SysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSysCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSys != nil && mm_atomic.LoadUint64(&m.afterSysCounter) < 1 {
		return false
	}
	return true
}

// MinimockSysInspect logs each unmet expectation
func (m *FileInfo) MinimockSysInspect() {
	for _, e := range m.SysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to FileInfo.Sys")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSysCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Sys")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSys != nil && mm_atomic.LoadUint64(&m.afterSysCounter) < 1 {
		m.t.Error("Expected call to FileInfo.Sys")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FileInfo) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockIsDirInspect()

		m.MinimockModTimeInspect()

		m.MinimockModeInspect()

		m.MinimockNameInspect()

		m.MinimockSizeInspect()

		m.MinimockSysInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FileInfo) MinimockWait(timeout mm_time.Duration) {
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

func (m *FileInfo) minimockDone() bool {
	done := true
	return done &&
		m.MinimockIsDirDone() &&
		m.MinimockModTimeDone() &&
		m.MinimockModeDone() &&
		m.MinimockNameDone() &&
		m.MinimockSizeDone() &&
		m.MinimockSysDone()
}
