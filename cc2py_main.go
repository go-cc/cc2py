// 汉语拼音转换工具.
package main

////////////////////////////////////////////////////////////////////////////
// Program: cc2py
// Purpose: Chinese-Character to Pinyin converter
// Authors: Tong Sun (c) 2017-2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cc2py_cliGen.sh

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pinyin "github.com/go-cc/cc-pinyin"
	"github.com/go-easygen/go-flags"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Filei     string `short:"i" long:"in" description:"the Chinese text file to read from (or \"-\" for stdin)"`
	Tone      int    `short:"t" long:"tone" env:"CC2PY_TONE" description:"tone selection\n\t\t\t  0: normal. mnemonic~ nothing\n\t\t\t  1: tone at the end. mnemonic~ single sided\n\t\t\t  2: tone after yunmu. mnemonic~ double sided\n\t\t\t  3: tone on yunmu. mnemonic~ fancy"`
	Truncate  int    `short:"l" long:"truncate" env:"CC2PY_TRUNCATE" description:"select only part of the pinyin\n\t\t\t  0: normal. mnemonic~ nothing truncated\n\t\t\t  1: leave first char. mnemonic~ one\n\t\t\t  2: leave first shengmu. mnemonic~ might be two\n\t\t\t  9: leave only yunmu. mnemonic~ last"`
	Separator string `short:"s" long:"separator" env:"CC2PY_SEPARATOR" description:"separator string between each pinyin" default:" "`
	Polyphone bool   `short:"p" long:"polyphone" env:"CC2PY_POLYPHONE" description:"polyphone support, output each polyphone pinyin available"`
	Capital   bool   `short:"c" long:"capitalized" env:"CC2PY_CAPITAL" description:"capitalized each pinyin word"`

	// positional arguments
	Args struct {
		CCStrs []string
	} `positional-args:"yes" required:"yes"`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "cc2py"
	version   = "0.2.2"
	date = "2022-01-18"

	// Opts store all the configurable options
	Opts OptsT
)

var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println("")
	err := Validate()
	abortOn("cli options", err)
	cc2pyC()
}

// Validate cli values sanity
func Validate() error {
	if Opts.Tone < pinyin.Normal || Opts.Tone > pinyin.Tone3 {
		return fmt.Errorf("tone %d out of range. run `%s -h` to get help.",
			Opts.Tone, progname)
	}
	if Opts.Truncate < pinyin.Normal ||
		Opts.Truncate > pinyin.Initials && Opts.Truncate < pinyin.ZeroConsonant ||
		Opts.Truncate > pinyin.Finals {
		return fmt.Errorf("truncate %d out of range. run `%s -h` to get help.",
			Opts.Truncate, progname)
	}
	return nil
}

func cc2pyC() error {
	// input data
	var dd string
	if len(Opts.Filei) != 0 { // -i,--in option is specified
		var data []byte
		var err error
		if Opts.Filei == "-" {
			data, err = ioutil.ReadAll(os.Stdin)
		} else {
			data, err = ioutil.ReadFile(Opts.Filei)
		}
		abortOn("Input", err)
		dd = string(data)
	} else {
		dd = strings.Join(Opts.Args.CCStrs, " ")
		if dd == "" {
			parser.WriteHelp(os.Stdout)
			os.Exit(0)
		}
	}

	fmt.Println(cc2py(dd, Opts.Tone, Opts.Truncate,
		Opts.Separator, Opts.Polyphone, Opts.Capital))

	return nil
}

func cc2py(dd string, tone, truncate int, separator string, polyphone, capitalized bool) string {
	py := pinyin.NewPinyin(tone, truncate, separator, polyphone, capitalized)
	return py.Convert(dd)
}

//==========================================================================
// support functions

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
