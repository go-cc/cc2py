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


# Download binaries

- The latest binary executables are available under  
https://bintray.com/suntong/bin/{{.Name}}#files/{{.Name}}  
as the result of the Continuous-Integration process.

- I.e., they are built right from the source code during every git commit automatically by [travis-ci](https://travis-ci.org/), thus are always the latest.

- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `{{.Name}}-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.

- You may want to rename it to a shorter name, e.g., `{{.Name}}`, instead after downloading it. To do the downloading and renaming programatically, use the plain-downloading url  
https://dl.bintray.com/suntong/bin/{{.Name}}.


# Debian package

Available at 


# Install Source

To install the source code instead:

```
go get github.com/go-cc/cc2py
```

# Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome.
