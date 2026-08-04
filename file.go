package gophers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"github.com/hinst/go-gophers/file_mode"
)

func ReadJsonFile[T any](filePath string, receiver T) T {
	var file = AssertResultError(os.Open(filePath))
	defer IoCloseSilently(file)
	AssertError(json.NewDecoder(file).Decode(receiver))
	return receiver
}

func ReadXmlFile[T any](filePath string, receiver T) T {
	var file = AssertResultError(os.Open(filePath))
	defer IoCloseSilently(file)
	AssertError(xml.NewDecoder(file).Decode(receiver))
	return receiver
}

func ReadTextFile(filePath string) string {
	return string(ReadBytesFile(filePath))
}

func ReadBytesFile(filePath string) []byte {
	var fileContent = AssertResultError(os.ReadFile(filePath))
	return fileContent
}

func WriteBytesFile(filePath string, data []byte) {
	AssertError(os.WriteFile(filePath, data, file_mode.USER_RW))
}

func WriteJsonFile[T any](filePath string, data T) {
	var jsonBytes = AssertResultError(json.Marshal(data))
	AssertError(os.WriteFile(filePath, jsonBytes, file_mode.USER_RW))
}

func WriteTextFile(filePath string, text string) {
	AssertError(os.WriteFile(filePath, []byte(text), file_mode.USER_RW))
}

func CheckFileExists(filePath string) bool {
	var info, err = os.Stat(filePath)
	return err == nil && !info.IsDir()
}

func CheckDirectoryExists(directoryPath string) bool {
	var info, err = os.Stat(directoryPath)
	return err == nil && info.IsDir()
}

func CopyFile(destinationPath string, sourcePath string) (size int64) {
	var sourceFile = AssertResultError(os.Open(sourcePath))
	defer IoCloseSilently(sourceFile)
	var destinationFile = AssertResultError(os.Create(destinationPath))
	defer IoClose(destinationFile)
	return AssertResultError(io.Copy(destinationFile, sourceFile))
}

// https://stackoverflow.com/questions/29505089/how-can-i-compare-two-files-in-golang
func CheckFilesEqual(file1, file2 string) bool {
	const chunkSize = 64000
	f1, err := os.Open(file1)
	if err != nil {
		return false
	}
	defer IoCloseSilently(f1)

	f2, err := os.Open(file2)
	if err != nil {
		return false
	}
	defer IoCloseSilently(f2)

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				panic("File comparison error")
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func CheckTextFilesEqual(file1, file2 string) bool {
	var text1 = ""
	if CheckFileExists(file1) {
		text1 = ReadTextFile(file1)
	}
	var text2 = ""
	if CheckFileExists(file2) {
		text2 = ReadTextFile(file2)
	}
	// ignore line breaks
	text1 = strings.ReplaceAll(text1, "\r\n", "\n")
	text2 = strings.ReplaceAll(text2, "\r\n", "\n")
	// trim whitespace
	text1 = strings.TrimSpace(text1)
	text2 = strings.TrimSpace(text2)
	return text1 == text2
}
