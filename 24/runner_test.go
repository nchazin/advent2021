package main

import (
	"fmt"
	"testing"
)

func Test_imp(t *testing.T) {
	var tests = []struct {
		a    byte
		want [4]int64
	}{
		{'w', [4]int64{9, 0, 0, 0}},
		{'x', [4]int64{0, 9, 0, 0}},
		{'y', [4]int64{0, 0, 9, 0}},
		{'z', [4]int64{0, 0, 0, 9}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%c-%v", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			alu := Alu{}
			alu.imp(tt.a, 9)
			val := alu.get_all()
			if val != tt.want {
				t.Errorf("got %v, want %v", val, tt.want)
			}
		})
	}
}

func Test_store(t *testing.T) {
	var tests = []struct {
		a    byte
		want [4]int64
	}{
		{'w', [4]int64{9, 0, 0, 0}},
		{'x', [4]int64{0, 9, 0, 0}},
		{'y', [4]int64{0, 0, 9, 0}},
		{'z', [4]int64{0, 0, 0, 9}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%c-%v", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			alu := Alu{}
			alu.store(tt.a, 9)
			val := alu.get_all()
			if val != tt.want {
				t.Errorf("got %v, want %v", val, tt.want)
			}
		})
	}
}

func Test_fetch(t *testing.T) {
	var tests = []struct {
		a    byte
		init [4]int64
		want int64
	}{
		{'w', [4]int64{-1, 0, 0, 0}, -1},
		{'x', [4]int64{0, 3, 0, 0}, 3},
		{'y', [4]int64{0, 0, 9, 0}, 9},
		{'z', [4]int64{0, 0, 0, -8}, -8},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%c-%d", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			alu := Alu{tt.init[0], tt.init[1], tt.init[2], tt.init[3]}
			val := alu.fetch(tt.a)
			if val != tt.want {
				t.Errorf("got %v, want %v", val, tt.want)
			}
		})
	}

}

func test_add(t *testing.T) {
	var tests = []struct {
		a    byte
		init [4]int64
		b    []byte
		want [4]int64
	}{
		{'w', [4]int64{-1, 0, 0, 0}, []byte("-99"), [4]int64{-100, 0, 0, 0}},
		{'x', [4]int64{7, 3, 4, -2}, []byte{'w'}, [4]int64{7, 10, 4, -2}},
		{'y', [4]int64{0, 0, 112, -7}, []byte{'z'}, [4]int64{0, 0, 105, -7}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%c-%s-%d", tt.a, string(tt.b[:]), tt.want)
		t.Run(testname, func(t *testing.T) {
			alu := Alu{tt.init[0], tt.init[1], tt.init[2], tt.init[3]}
			alu.add(tt.a, tt.b)
			val := alu.get_all()
			if val != tt.want {
				t.Errorf("got %v, want %v", val, tt.want)

			}
		})
	}

}

func test_mul(t *testing.T) {
	var tests = []struct {
		a    byte
		init [4]int64
		b    []byte
		want [4]int64
	}{
		{'w', [4]int64{-1, 0, 0, 0}, []byte("-99"), [4]int64{99, 0, 0, 0}},
		{'x', [4]int64{7, 3, 4, -2}, []byte{'w'}, [4]int64{7, 28, 4, -2}},
		{'y', [4]int64{0, 0, 112, -7}, []byte{'w'}, [4]int64{0, 0, 0, -7}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%c-%s-%d", tt.a, string(tt.b[:]), tt.want)
		t.Run(testname, func(t *testing.T) {
			alu := Alu{tt.init[0], tt.init[1], tt.init[2], tt.init[3]}
			alu.mul(tt.a, tt.b)
			val := alu.get_all()
			if val != tt.want {
				t.Errorf("got %v, want %v", val, tt.want)

			}
		})
	}

}
