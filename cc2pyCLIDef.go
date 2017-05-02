////////////////////////////////////////////////////////////////////////////
// Program: cc2pyC
// Purpose: Chinese-Character to Pinyin converter
// Authors: Tong Sun (c) 2016-2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// cc2pyC

type rootT struct {
	cli.Helper
	Filei     *clix.Reader `cli:"i,in" usage:"the Chinese text file to read from (or stdin)\n\t\t\tif not specified, read from command line instead"`
	Tone      int          `cli:"t,tone" usage:"tone selection\n\t\t\t  0: normal. mnemonic~ nothing\n\t\t\t  1: tone at the end. mnemonic~ single sided\n\t\t\t  2: tone after yunmu. mnemonic~ double sided\n\t\t\t  3: tone on yunmu. mnemonic~ fancy"`
	Truncate  int          `cli:"l,truncate" usage:"select only part of the pinyin\n\t\t\t  0: normal. mnemonic~ nothing truncated\n\t\t\t  1: leave first char. mnemonic~ one\n\t\t\t  2: leave first shengmu. mnemonic~ might be two\n\t\t\t  9: leave only yunmu. mnemonic~ last"`
	Separator string       `cli:"s,separator" usage:"separator string between each pinyin" dft:" "`
	Polyphone bool         `cli:"p,polyphone" usage:"polyphone support, output each polyphone pinyin available"`
	Capital   bool         `cli:"c,capitalized" usage:"capitalized each pinyin word"`
}

var root = &cli.Command{
	Name: "cc2pyC",
	Desc: "Chinese-Character to Pinyin converter\nbuilt on " + buildTime,
	Text: "Converter Chinese to pinyin in different ways",
	Argv: func() interface{} { return new(rootT) },
	Fn:   cc2pyC,

	CanSubRoute: true,
}

// func main() {
// 	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
// 	//NOTE: You can set any writer implements io.Writer
// 	// default writer is os.Stdout
// 	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	fmt.Println("")
// }

// func cc2pyC(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }
