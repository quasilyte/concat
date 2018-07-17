package concat

import (
	"strings"
	"testing"
	"unsafe"
)

func TestStringSize(t *testing.T) {
	have := unsafe.Sizeof(stringStruct{})
	want := unsafe.Sizeof("")
	if have != want {
		t.Errorf("string struct size mismatch: have %d, want %d", have, want)
	}
}

func TestConcat(t *testing.T) {
	const stringConst = "const"
	var stringVar = "var"

	inputs := [][]string{
		// concat2:
		{"", ""},
		{"x", ""},
		{"", "x"},
		{stringVar, stringConst},
		{"abc", stringVar},
		{stringConst, "abc"},
		{stringVar + stringVar, stringVar},

		// concat3:
		{"", "", ""},
		{"x", "", ""},
		{"", "x", ""},
		{stringVar, stringConst, "x"},
		{stringConst, "", ""},
		{stringVar, "x", stringVar + stringVar},
	}

	reverseStrings := func(xs []string) []string {
		ys := make([]string, len(xs))
		for i := range xs {
			ys[i] = xs[len(xs)-i-1]
		}
		return ys
	}

	type testCase struct {
		fn       func(args []string) string
		goldenFn func(args []string) string
		args     []string
	}
	var tests []testCase
	for _, xs := range inputs {
		var fn func([]string) string
		var goldenFn func([]string) string
		switch len(xs) {
		case 2:
			fn = func(xs []string) string {
				return Strings(xs[0], xs[1])
			}
			goldenFn = func(xs []string) string {
				return xs[0] + xs[1]
			}
		case 3:
			fn = func(xs []string) string {
				return Strings3(xs[0], xs[1], xs[2])
			}
			goldenFn = func(xs []string) string {
				return xs[0] + xs[1] + xs[2]
			}
		default:
			panic("invalid arguments count")
		}

		tests = append(tests,
			testCase{fn: fn, goldenFn: goldenFn, args: xs},
			testCase{fn: fn, goldenFn: goldenFn, args: reverseStrings(xs)})
	}

	for _, test := range tests {
		have := test.fn(test.args)
		want := test.goldenFn(test.args)
		if have != want {
			t.Errorf("concat(%v) result mismatch:\nhave: %q\nwant: %q",
				test.args, have, want)
		}
	}
}

var shortStrings = []string{"lorem ", "ipsum ", "dolor sit amet"}
var longerStrings = []string{
	strings.Repeat(shortStrings[0], 16),
	strings.Repeat(shortStrings[1], 16),
	strings.Repeat(shortStrings[2], 16),
}

func benchmarkConcat(b *testing.B, fn func([]string) string) {
	b.Run("short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fn(shortStrings)
		}
	})
	b.Run("longer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fn(longerStrings)
		}
	})
}

func BenchmarkConcat2Operator(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string { return xs[0] + xs[1] })
}

func BenchmarkConcat2Builder(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string {
		var builder strings.Builder
		builder.Grow(len(xs[0]) + len(xs[1]))
		builder.WriteString(xs[0])
		builder.WriteString(xs[1])
		return builder.String()
	})
}

func BenchmarkConcat2(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string {
		return Strings(xs[0], xs[1])
	})
}

func BenchmarkConcat3Operator(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string {
		return xs[0] + xs[1] + xs[2]
	})
}

func BenchmarkConcat3Builder(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string {
		var builder strings.Builder
		builder.Grow(len(xs[0]) + len(xs[1]) + len(xs[2]))
		builder.WriteString(xs[0])
		builder.WriteString(xs[1])
		builder.WriteString(xs[2])
		return builder.String()
	})
}

func BenchmarkConcat3(b *testing.B) {
	benchmarkConcat(b, func(xs []string) string {
		return Strings3(xs[0], xs[1], xs[2])
	})
}
