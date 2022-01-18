////////////////////////////////////////////////////////////////////////////
// Program: cc2pyC
// Purpose: Chinese-Character to Pinyin converter
// Authors: Tong Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
//  	"fmt"
//  	"os"

//	"github.com/go-easygen/go-flags"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Filei	string	`short:"i" long:"in" description:"the Chinese text file to read from (or stdin)\n\t\t\tif not specified, read from command line instead"`
//  	Tone	int	`short:"t" long:"tone" env:"CC2PY_TONE" description:"tone selection\n\t\t\t  0: normal. mnemonic~ nothing\n\t\t\t  1: tone at the end. mnemonic~ single sided\n\t\t\t  2: tone after yunmu. mnemonic~ double sided\n\t\t\t  3: tone on yunmu. mnemonic~ fancy"`
//  	Truncate	int	`short:"l" long:"truncate" env:"CC2PY_TRUNCATE" description:"select only part of the pinyin\n\t\t\t  0: normal. mnemonic~ nothing truncated\n\t\t\t  1: leave first char. mnemonic~ one\n\t\t\t  2: leave first shengmu. mnemonic~ might be two\n\t\t\t  9: leave only yunmu. mnemonic~ last"`
//  	Separator	string	`short:"s" long:"separator" env:"CC2PY_SEPARATOR" description:"separator string between each pinyin" default:" "`
//  	Polyphone	bool	`short:"p" long:"polyphone" env:"CC2PY_POLYPHONE" description:"polyphone support, output each polyphone pinyin available"`
//  	Capital	bool	`short:"c" long:"capitalized" env:"CC2PY_CAPITAL" description:"capitalized each pinyin word"`
//  }

// Template for main starts here

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "cc2pyC"
//          version   = "0.1.0"
//          date = "2022-01-17"

//  	// Opts store all the configurable options
//  	Opts OptsT
//  )
//
//  var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//
//  	if _, err := parser.Parse(); err != nil {
//  		switch flagsErr := err.(type) {
//  		case flags.ErrorType:
//  			if flagsErr == flags.ErrHelp {
//  				os.Exit(0)
//  			}
//  			os.Exit(1)
//  		default:
//  			fmt.Println()
//  			parser.WriteHelp(os.Stdout)
//  			os.Exit(1)
//  		}
//  	}
//  	fmt.Println("")
//  }
// Template for main ends here
