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
	words := make([]Words, len(lists))

	for i, value := range lists{
		v := value
		l := len(value)
		words[i].Name = &v[0]
		words[i].Length = &l
	}

	//for _, w := range words{
	//	fmt.Print(*w.Name)
	//	fmt.Print(*w.Length)
	//	fmt.Println("\n")
	//}

	sort.Slice(words, func(i, j int) bool {
		return *words[i].Length > *words[j].Length
	})

	for i:=0; i<3; i++{
		fmt.Println(*words[i].Name)
	}
	//for _, w := range words {
	//	fmt.Println(*w.Name)
	//}
}
