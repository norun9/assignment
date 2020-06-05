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
	newSlice := strings.Split(b, ".")

	wordsSlice := [][]string{}
	//最後に空白が入るため-1をしている(要リファクタリング)
	for i:=0; i<len(newSlice)-1; i++{
		collectSameWords := strings.Split(newSlice[i], " ")
		if collectSameWords[0] == ""{
			collectSameWords = append(collectSameWords[:0], collectSameWords[1:]...)
		}
		wordsSlice = append(wordsSlice, collectSameWords)
	}

	//スライス構造体の初期化で要素数を可変長にする
	words := make([]Words, len(wordsSlice))

	for i, word := range wordsSlice{
		w := word
		l := len(word)
		words[i].Name = &w[0]
		words[i].Length = &l
	}

	sort.Slice(words, func(i, j int) bool {
		if *words[i].Length > *words[j].Length {
			return true
		} else if *words[i].Length == *words[j].Length {
			if *words[i].Name < *words[j].Name {
				return true
			}
		}
		return false
	})

	for i:=0; i<3; i++{
		fmt.Println(*words[i].Name)
	}
}