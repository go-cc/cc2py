// 汉语拼音转换工具.

//go:generate sh -v cc2pyCLIGen.sh

package main

import (
	"fmt"
	"os"

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
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	s := `名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／19969月30日／59.70【元】，《三国演义》〖明〗罗贯中。`
	_ = s
	hans := "中国人的〖中国银行〗，很.行.。"

	// 默认
	a := pinyin.NewPinyin()
	//a.Separator = "_"
	fmt.Println(a.Convert(hans))

	// 包含声调
	a.SetStyle(pinyin.Tone)
	fmt.Println(a.Convert(hans))

	// 声调用数字表示
	a.SetStyle(pinyin.Tone2)
	fmt.Println(a.Convert(hans))

	// 声调在拼音后用数字表示
	a.SetStyle(pinyin.Tone3)
	fmt.Println(a.Convert(hans))

	// 开启多音字模式
	a.Heteronym = true
	fmt.Println(a.Convert(hans))
	a.SetStyle(pinyin.Tone)
	fmt.Println(a.Convert(hans))

	return nil
}
