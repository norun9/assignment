package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"sort"
	"strings"
)

type Words struct {
	Name   *string
	Length *int
}

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	b := string(body)
	//"."で区切って同じ文字の配列を作成
	slice := strings.Split(b, ".")

	lists := [][]string{}
	//最後に空白が入るため-1をしている(要リファクタリング)
	for i := 0; i<len(slice)-1; i++{
		sameStrSlice := strings.Split(slice[i], " ")
		if sameStrSlice[0] == ""{
			sameStrSlice = append(sameStrSlice[:0], sameStrSlice[1:]...)
		}
		lists = append(lists, sameStrSlice)
	}

	//スライス構造体の初期化で要素数を可変長にする
	words := make([]Words, len(lists))

	for i, value := range lists{
		v := value
		l := len(value)
		words[i].Name = &v[0]
		words[i].Length = &l
	}

	sort.Slice(words, func(i, j int) bool {
		if *words[i].Length > *words[j].Length {
			return true
		} else if *words[i].Length == *words[j].Length {
			if *words[j].Name > *words[i].Name {
				return true
			}
		}
		return false
	})

	for i:=0; i<3; i++{
		fmt.Println(*words[i].Name)
	}
}