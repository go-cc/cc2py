package main

import (
	"testing"
)

type testCase struct {
	got  string
	want string
}

func testPinyin(t *testing.T, d []testCase) {
	for _, tc := range d {
		if tc.got != tc.want {
			t.Errorf(`Expected "%s", got "%s"`, tc.want, tc.got)
		}
	}
}

func TestConvert(t *testing.T) {
	hans := "〖中国银行〗，很行。"
	Separator := " "
	testData := []testCase{
		{
			cc2py(hans, 0, 0, Separator, false, false),
			"〖zhong guo yin xing 〗，hen xing 。",
		},
		{
			cc2py(hans, 0, 1, Separator, false, false),
			"〖z g y x 〗，h x 。",
		},
		// {
		// 	cc2py(hans, 0, 2, Separator, false, false),
		// 	"〖zh g y x 〗，h x 。",
		// },
		// {
		// 	cc2py(hans, 0, 9, Separator, false, false),
		// 	"〖ong uo in ing 〗，en ing 。",
		// },
		{
			cc2py(hans, 0, 0, Separator, false, false),
			"〖zhong guo yin xing 〗，hen xing 。",
		},
	}

	testPinyin(t, testData)

}
