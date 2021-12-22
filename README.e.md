---
ex1: cc2py -t 3 "中国人的〖中国银行〗，很.行.。"
ex2: echo "中国人的〖中国银行〗，很.行.。" | tee /tmp/pytest.txt | cc2py -t 1 -i
ex3: cc2py -i /tmp/pytest.txt -t 2
---

## {{toc 5}}
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)

## {{.Name}} - Chinese-Character to Pinyin converter

`{{.Name}}` will convert Chinese-Character to pinyin.

## Usage

### $ {{exec "cc2py" | color "sh"}}

### Examples

```sh
$ {{shell .ex1}}

$ {{shell .ex2}}

$ {{shell .ex3}}

$ {{shell "cc2py -i /tmp/pytest.txt -p -t 3"}}

$ {{shell "cc2py -i /tmp/pytest.txt -l 1 -c "}}

$ {{shell "cc2py -i /tmp/pytest.txt -l 2 -s '' -c "}}
```
