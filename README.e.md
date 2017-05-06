---
ex1: cc2py -t 3 "中国人的〖中国银行〗，很.行.。"
ex2: echo "中国人的〖中国银行〗，很.行.。" | tee /tmp/pytest.txt | cc2py -t 1 -i
ex3: cc2py -i /tmp/pytest.txt
---

# {{.Name}}

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

{{pkgdoc}}

## {{toc 5}}

# {{.Name}} - Chinese-Character to Pinyin converter

`{{.Name}}` will convert Chinese-Character to pinyin.

# Usage

#### $ {{exec "cc2py" | color "sh"}}

## Examples

```sh
$ {{shell .ex1}}

$ {{shell .ex2}}

$ {{shell .ex3}}

$ {{shell "cc2py -i /tmp/pytest.txt -p -t 3"}}

$ {{shell "cc2py -i /tmp/pytest.txt -l 2 -c "}}

$ {{shell "cc2py -i /tmp/pytest.txt -l 1 -s '' -c "}}
```


All patches welcome. 
