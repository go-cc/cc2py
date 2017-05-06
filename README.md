
# cc2py

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-cc/cc2py?status.svg)](http://godoc.org/github.com/go-cc/cc2py)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cc/cc2py)](https://goreportcard.com/report/github.com/go-cc/cc2py)
[![travis Status](https://travis-ci.org//go-cc/cc2py.svg?branch=master)](https://travis-ci.org//go-cc/cc2py)

汉语拼音转换工具.


## TOC
- [cc2py - Chinese-Character to Pinyin converter](#cc2py---chinese-character-to-pinyin-converter)
- [Usage](#usage)
  - [$ cc2py](#-cc2py)
- [Examples](#examples)

# cc2py - Chinese-Character to Pinyin converter

`cc2py` will convert Chinese-Character to pinyin.

# Usage

#### $ cc2py
```sh
Chinese-Character to Pinyin converter
built on 2017-05-05

Converter Chinese to pinyin in different ways

Options:

  -h, --help            display help information
  -i, --in              the Chinese text file to read from (or stdin)
			if not specified, read from command line instead
  -t, --tone            tone selection
			  0: normal. mnemonic~ nothing
			  1: tone at the end. mnemonic~ single sided
			  2: tone after yunmu. mnemonic~ double sided
			  3: tone on yunmu. mnemonic~ fancy
  -l, --truncate        select only part of the pinyin
			  0: normal. mnemonic~ nothing truncated
			  1: leave first char. mnemonic~ one
			  2: leave first shengmu. mnemonic~ might be two
			  9: leave only yunmu. mnemonic~ last
  -s, --separator[= ]   separator string between each pinyin
  -p, --polyphone       polyphone support, output each polyphone pinyin available
  -c, --capitalized     capitalized each pinyin word
```

## Examples

```sh
$ cc2py -t 3 "中国人的〖中国银行〗，很.行.。"
zhōng guó rén de 〖zhōng guó yín xíng 〗，hěn .xíng .。

$ echo "中国人的〖中国银行〗，很.行.。" | tee /tmp/pytest.txt | cc2py -t 1 -i
zhong1 guo2 ren2 de 〖zhong1 guo2 yin2 xing2 〗，hen3 .xing2 .。

$ cc2py -i /tmp/pytest.txt
zhong guo ren de 〖zhong guo yin xing 〗，hen .xing .。

$ cc2py -i /tmp/pytest.txt -p -t 3
zhōng/zhòng guó rén de/dì/dí 〖zhōng/zhòng guó yín xíng/háng/xìng/hàng/héng 〗，hěn .xíng/háng/xìng/hàng/héng .。

$ cc2py -i /tmp/pytest.txt -l 2 -c 
Zh G R D 〖Zh G Y X 〗，H .X .。

$ cc2py -i /tmp/pytest.txt -l 1 -s '' -c 
ZGRD〖ZGYX〗，H.X.。
```


All patches welcome. 
