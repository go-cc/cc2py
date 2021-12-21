
# cc2py

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-cc/cc2py?status.svg)](http://godoc.org/github.com/go-cc/cc2py)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cc/cc2py)](https://goreportcard.com/report/github.com/go-cc/cc2py)
[![Build Status](https://github.com/go-cc/cc2py/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/go-cc/cc2py/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)

汉语拼音转换工具.


## TOC
- [cc2py - Chinese-Character to Pinyin converter](#cc2py---chinese-character-to-pinyin-converter)
- [Usage](#usage)
  - [$ cc2py](#-cc2py)
  - [Examples](#examples)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)

## cc2py - Chinese-Character to Pinyin converter

`cc2py` will convert Chinese-Character to pinyin.

## Usage

### $ cc2py
```sh
Chinese-Character to Pinyin converter
built on 2021-12-18

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

### Examples

```sh
$ cc2py -t 3 "中国人的〖中国银行〗，很.行.。"
zhōng guó rén de 〖zhōng guó yín xíng 〗，hěn .xíng .。

$ echo "中国人的〖中国银行〗，很.行.。" | tee /tmp/pytest.txt | cc2py -t 1 -i
zhong1 guo2 ren2 de 〖zhong1 guo2 yin2 xing2 〗，hen3 .xing2 .。

$ cc2py -i /tmp/pytest.txt -t 2
zho1ng guo2 re2n de 〖zho1ng guo2 yi2n xi2ng 〗，he3n .xi2ng .。

$ cc2py -i /tmp/pytest.txt -p -t 3
zhōng/zhòng guó rén de/dī/dí/dì 〖zhōng/zhòng guó yín xíng/háng/héng/xìng/hàng 〗，hěn .xíng/háng/héng/xìng/hàng .。

$ cc2py -i /tmp/pytest.txt -l 1 -c 
Z G R D 〖Z G Y X 〗，H .X .。

$ cc2py -i /tmp/pytest.txt -l 2 -s '' -c 
ZhGRD〖ZhGYX〗，H.X.。
```

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/go-cc/cc2py/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `cc2py_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf cc2py_*_linux_amd64.tar.gz
sudo mv -v cc2py_*_linux_amd64/cc2py /usr/local/bin/
rmdir -v cc2py_*_linux_amd64
```


### Distro package

- Packages available for Linux distros are
  * [Alpine Linux](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-alpine)
  * [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb)
  * [RedHat](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-rpm)

The repo setup instruction url has been given above.
For example, for [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb) --

### Debian package


```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/suntong/repo/setup.deb.sh' \
  | sudo -E bash

# That's it. You then can do your normal operations, like

sudo apt-get update
apt-cache policy cc2py

sudo apt-get install -y cc2py
```

## Install Source

To install the source code instead:

```
go get -v -u github.com/go-cc/cc2py
```

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.

All patches welcome.
