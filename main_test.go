package main

import "testing"

type Test struct {
	in  string
	out []string
}

var tests2 = []Test{
	{"в1", []string{}},
	{"A1", []string{"B3", "C2"}},
}

// TestFunc2 - тестирование функции по п.2
func TestFunc2(t *testing.T) {
	for i, test := range tests2 {
		te := Func2(test.in)
		if te == nil {
			t.Errorf("#%d: Func2(\"%v\")=%s; want %s", i, test.in, te, test.out)
		}
	}
}

type TestLenStruct struct {
	in  Koloda
	out int
}

var test3 = []TestLenStruct{
	{Koloda{[]Card{{false, 10, 2}, {false, 3, 3}, {false, 5, 0}, {false, 1, 2}, {false, 1, 2}, {false, 1, 2}}, 2}, 6},
	{Koloda{[]Card{{false, 10, 2}, {false, 3, 3}, {false, 5, 0}, {false, 1, 2}}, 2}, 4},
	{Koloda{[]Card{{false, 10, 2}, {false, 3, 3}, {false, 5, 0}}, 2}, 3},
	{Koloda{[]Card{{false, 10, 2}, {false, 3, 3}}, 2}, 2},
}

// TestLen - тестирование функции по п.3, Количество карт в колоде
func TestLen(t *testing.T) {
	for i, test := range test3 {
		te := test.in.Len()
		if te != test.out {
			t.Errorf("#%d: Len(%v)=%v; want %v", i, test.in, te, test.out)
		}
	}
}

// TestSwap - тестирование функции по п.3, Перемена двух карт местами
func TestSwap(t *testing.T) {
	for i, test := range test3 {
		te0 := test.in.Cards[0]
		te1 := test.in.Cards[1]
		test.in.Swap(0, 1)
		if te0 != test.in.Cards[1] && te1 != test.in.Cards[0] {
			t.Errorf("#%d: TestSwap(%v)=%v;", i, test.in, test.out)
		}
	}
}

// TestLess - тестирование функции по п.3, Перемена двух карт местами
func TestLess(t *testing.T) {
	for i, test := range test3 {

		if (test.in.Cards[0].Amount > test.in.Cards[1].Amount) != test.in.Less(0, 1) {
			t.Errorf("#%d: TestLess(%d,%d)=%v; want %v", i, 0, 1, test.in.Less(0, 1), test.in.Cards[0].Amount > test.in.Cards[1].Amount)
		}
	}
}
