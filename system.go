package gophers

import (
	"os"
	"path/filepath"
	"reflect"
	"unsafe"
)

var SizeOfInt = unsafe.Sizeof(int(0))
var BitSizeOfInt = SizeOfInt * 8

var SizeOfUint = unsafe.Sizeof(uint(0))
var BitSizeOfUInt = SizeOfUint * 8

var ExecutableFilePath string
var ExecutableFileDirectory string

func init() {
	var executableFilePath, e = os.Executable()
	AssertError(e)
	ExecutableFilePath, e = filepath.Abs(executableFilePath)
	AssertError(e)
	ExecutableFileDirectory = filepath.Dir(ExecutableFilePath)
}

func IsNil(value interface{}) bool {
	return value == nil || reflect.ValueOf(value).IsNil()
}

// This function exists for better readability: to avoid !isNil { ugly }
func IsThere(value interface{}) bool {
	return !IsNil(value)
}
