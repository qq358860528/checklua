// checklua lua语法检查工具
// xiaoyu chen
// 358860528@qq.com
// lua文件支持utf8 但是不支持utf8bom 格式 会报错1行1列错误
// 原生lua不支持utf8bom 这里不做支持
// 支持递归检查子目录
// checklua.exe 检查目录【支持绝对路径和相对路径】

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/gopher-lua"
)

func main() {

	//flag.Parse()
	checkdir := "./"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkdir = dir
	if len(os.Args) >= 2 {
		checkdir = os.Args[1]
	}

	L1 := lua.NewState()
	defer L1.Close()
	L1.OpenLibs()
	fmt.Println("开始检查lua 配置表!", checkdir)
	errcount := 0

	errarr := []string{}
	err = filepath.Walk(checkdir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".lua" {
			fmt.Println(path)
			//判断是否lua文件
			_, err := L1.LoadFile(path)
			//err := L1.DoFile("hello.lua")
			if err != nil {
				fmt.Println(err)
				errarr = append(errarr, err.Error())
				errcount++
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if errcount > 0 {
		fmt.Println("一共", errcount, "个配置表出错!")
		for _, e := range errarr {
			fmt.Println(e)
		}
	}
	fmt.Println("检查lua 配置表完毕!")
	//time.Sleep(2 * time.Second)
}
