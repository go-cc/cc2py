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

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "cc2py"
	version  = "0.2.3"
	date     = "2022-01-22"

	// Opts store all the configurable options
	Opts OptsT
)

var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	Opts.Version = showVersion
	Opts.Verbflg = func() {
		Opts.Verbose++
	}

	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println("")
	DoCc2py()
}

// DoCc2py implements the business logic of command `cc2py`
func DoCc2py() error {
	err := Validate()
	abortOn("cli options", err)
	return cc2pyC()
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

func showVersion() {
	fmt.Fprintf(os.Stderr, "cc2py - Chinese-Character to Pinyin converter\n")
	fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Tong Sun\n\n")
	fmt.Fprintf(os.Stderr, "Converter Chinese to pinyin in different ways\n\nBuilt on %s\nVersion %s\n",
		date, version)
	os.Exit(0)
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
