// 汉语拼音转换工具.
package main

////////////////////////////////////////////////////////////////////////////
// Porgram: pinyin
// Purpose: pinyin conversion Go library
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits:
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v cc2pyCLIGen.sh

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pinyin "github.com/go-cc/cc-pinyin"
	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "cc2py"
	VERSION   = "0.2.2"
	buildTime = "2021-12-18"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// default writer is os.Stdout
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

// Validate implements cli.Validator interface
func (argv *rootT) Validate(ctx *cli.Context) error {
	if argv.Tone < pinyin.Normal || argv.Tone > pinyin.Tone3 {
		return fmt.Errorf("tone %d out of range. run `%s -h` to get help.",
			argv.Tone, progname)
	}
	if argv.Truncate < pinyin.Normal ||
		argv.Truncate > pinyin.Initials && argv.Truncate < pinyin.ZeroConsonant ||
		argv.Truncate > pinyin.Finals {
		return fmt.Errorf("truncate %d out of range. run `%s -h` to get help.",
			argv.Truncate, progname)
	}
	return nil
}

func cc2pyC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()
	// fmt.Println(ctx.Args())
	argv := ctx.Argv().(*rootT)

	// input data
	var dd string
	if ctx.IsSet("--in") { // -i,--in option is specified
		data, err := ioutil.ReadAll(argv.Filei)
		abortOn("Input", err)
		argv.Filei.Close()
		dd = string(data)
	} else {
		dd = strings.Join(ctx.Args(), " ")
		if dd == "" {
			ctx.WriteUsage()
			os.Exit(0)
		}
	}

	fmt.Println(cc2py(dd, argv.Tone, argv.Truncate,
		argv.Separator, argv.Polyphone, argv.Capital))

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
