package main

import (
	"reflect"
	"testing"
)

// func Test_compress(t *testing.T) {

// 	var tests = []struct {
// 		name     string
// 		chars    []byte
// 		expected int
// 	}{
// 		{"six letters", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
// 		{"single letter", []byte{'a'}, 2},
// 		{"lots of letters", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result := compress(test.chars)
// 			if result != test.expected {
// 				t.Error(test.name, "returned ", result, ", expected ", test.expected)
// 			}
// 		})
// 	}

// 	// fmt.Println(tests)
// 	// result := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
// 	// if result != 6 {
// 	// 	t.Error("incorrect result, expected 6, got: ", result)
// 	// }
// }

func Test_compress(t *testing.T) {
	// type args struct {
	// 	chars []byte
	// }
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{
			name:  "Test single character",
			chars: []byte{'a'},
			// want:  "a",
			want: 1,
		},
		{
			name:  "Test multiple characters",
			chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			// want:  "a2b2c3",
			want: 6,
		},
		{
			name:  "Test lots of characters",
			chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			// want:  "ab12",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.chars); got != tt.want {
				t.Errorf("compress() = %v (%v), want %v (%v)", got, reflect.TypeOf(got), tt.want, reflect.TypeOf(tt.want))
			}
		})
	}
}
