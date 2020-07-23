package testing

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func stringContains(message string, search string) bool {

	searchIndex := 0

	for i := 0; i < len(message); i++ {
		if i+len(search)-searchIndex > len(message) {
			return false
		}

		if message[i] == search[searchIndex] {
			searchIndex++

			if searchIndex == len(search) {
				return true
			}
		}
	}

	return false
}
func byteSliceContains(message []byte, search []byte) bool {

	searchIndex := 0

	for i := 0; i < len(message); i++ {
		if i+len(search)-searchIndex > len(message) {
			return false
		}

		if message[i] == search[searchIndex] {
			searchIndex++

			if searchIndex == len(search) {
				return true
			}
		}
	}

	return false
}

var content, _ = ioutil.ReadFile("test_files/big_file.html")

func BenchmarkByteSliceContains(b *testing.B) {
	fmt.Println(b.N)
	search := []byte("https://www.w3.org/2000/svg")
	for i := 0; i < b.N; i++ {
		byteSliceContains(content, search)
	}
}

func BenchmarkStringContains(b *testing.B) {
	fmt.Println(b.N)
	contentString := string(content)
	for i := 0; i < b.N; i++ {
		stringContains(contentString, "https://www.w3.org/2000/svg")
	}
}

func BenchmarkStringContainsFramework(b *testing.B) {
	fmt.Println(b.N)
	contentString := string(content)
	for i := 0; i < b.N; i++ {
		strings.Contains(contentString, "https://www.w3.org/2000/svg")
	}
}
