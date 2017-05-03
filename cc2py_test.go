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
		// 0: normal. mnemonic~ nothing
		{
			cc2py(hans, 0, 0, Separator, false, false),
			"〖zhong guo yin xing 〗，hen xing 。",
		},
		{
			cc2py(hans, 0, 1, Separator, false, false),
			"〖z g y x 〗，h x 。",
		},
		{
			cc2py(hans, 0, 2, Separator, false, false),
			"〖zh g y x 〗，h x 。",
		},
		{
			cc2py(hans, 0, 9, Separator, false, false),
			"〖ong uo in ing 〗，en ing 。",
		},
		// 1: tone at the end. mnemonic~ single sided
		{
			cc2py(hans, 1, 0, Separator, false, false),
			"〖zhong1 guo2 yin2 xing2 〗，hen3 xing2 。",
		},
		{
			cc2py(hans, 1, 1, Separator, false, false),
			"〖z g y x 〗，h x 。",
		},
		{
			cc2py(hans, 1, 2, Separator, false, false),
			"〖zh g y x 〗，h x 。",
		},
		{
			cc2py(hans, 1, 9, Separator, false, false),
			"〖ong1 uo2 in2 ing2 〗，en3 ing2 。",
		},
		// 2: tone after yunmu. mnemonic~ double sided
		{
			cc2py(hans, 2, 0, Separator, false, false),
			"〖zho1ng guo2 yi2n xi2ng 〗，he3n xi2ng 。",
		},
		{
			cc2py(hans, 2, 1, Separator, false, false),
			"〖z g y x 〗，h x 。",
		},
		{
			cc2py(hans, 2, 2, Separator, false, false),
			"〖zh g y x 〗，h x 。",
		},
		{
			cc2py(hans, 2, 9, Separator, false, false),
			"〖o1ng uo2 i2n i2ng 〗，e3n i2ng 。",
		},
		// 3: tone on yunmu. mnemonic~ fancy
		{
			cc2py(hans, 0, 0, Separator, false, false),
			"〖zhong guo yin xing 〗，hen xing 。",
		},
		{
			cc2py(hans, 0, 1, Separator, false, false),
			"〖z g y x 〗，h x 。",
		},
		{
			cc2py(hans, 0, 2, Separator, false, false),
			"〖zh g y x 〗，h x 。",
		},
		{
			cc2py(hans, 0, 9, Separator, false, false),
			"〖ong uo in ing 〗，en ing 。",
		},
	}

	testPinyin(t, testData)

}
