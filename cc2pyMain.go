package main

import (
    "fmt"
    pybr "github.com/go-cc/cc-table/cc-pinyin-range"
)

func main() {
    s := `9月30日，金秋送爽，在祖国大地处处洋溢着和平幸福美好生活的今天，我们迎来了第三个法定“烈士纪念日”。雄伟壮丽的天安门广场上，鲜艳的五星红旗迎风飘扬，庄严肃穆的人民英雄纪念碑巍然矗立，指向云天`

    py := pybr.New()
    p, _ := py.Convert(s)

    fmt.Println(p)
 

    //设置分隔符
    //首字母是否大写
    py.Split = "-"
    py.Upper = false
    p, _ = py.Convert(s)
    fmt.Println(p)
}
