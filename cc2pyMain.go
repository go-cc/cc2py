// 汉语拼音转换工具.

//go:generate sh -v cc2pyCLIGen.sh

package main

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

var progname = "cc2py"
var buildTime = "2017-04-28"

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

func cc2pyC(ctx *cli.Context) error {
	// ctx.JSON(ctx.RootArgv())
	// ctx.JSON(ctx.Argv())
	// fmt.Println()
	// fmt.Println(ctx.Args())
	argv := ctx.Argv().(*rootT)

	// input data
	var dd string
	if ctx.IsSet("--Filei") { // --Filei option is specified
		data, err := ioutil.ReadAll(argv.Filei)
		abortOn("Input", err)
		argv.Filei.Close()
		print(data)
		dd = string(data)
		print(dd)
	} else {
		dd = strings.Join(ctx.Args(), " ")
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
