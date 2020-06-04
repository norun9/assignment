package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"sort"
	"strings"
)

type Words struct {
	Name *string
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

	//var length int = len(lists)
	//WordSlice := make([]Words, int(length))

	//構造体初期化の要素数を可変長にしたい
	//オブジェクト{}が一つだけだとout of indexのエラーが起きる
	//ここさえリファクタしたらOK!!!!
	//words := []Words{{},{},{},{}}

	//スライス構造体の初期化で要素数を可変長にする
	words := make([]Words, len(lists))

	for i, value := range lists{
		v := value
		l := len(value)
		words[i].Name = &v[0]
		words[i].Length = &l
	}

	sort.Slice(words, func(i, j int) bool {
		return *words[i].Length > *words[j].Length
	})

	//一番初めは0番目の要素のLengthに設定
	prev :=  words[0]
	for j:=0; j<len(words);j++ {
			current := words[j]
			//今と一つ前の要素数が同じなら二つを比較してアルファベット順にする
			if *prev.Length == *current.Length{

			}
			prev = current
		fmt.Print(*words[j].Name)
		fmt.Println("\n")
	}

	//アルファベット順にするから3回で限定してはいけない
	//for i:=0; i<3; i++{
	//	fmt.Println(*words[i].Name)
	//	fmt.Println(*words[i].Length)
	//}
}