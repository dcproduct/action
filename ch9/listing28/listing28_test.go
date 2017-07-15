// Sample benchmarks to test which function is better for converting
// an integer into a string. First using the fmt.Sprintf function,
// then the strconv.FormatInt function and then strconv.Itoa.

package listing28_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf provides performanace numbers for the
// fmt.Sprint function.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt funcation.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa funcation.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
