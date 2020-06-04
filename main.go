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

func (word Words) String() *string {
	return word.Name
}

type perLength []Words

func (l perLength) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
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
	words := []Words{{},{},{},{}}
	//words := []Words{}
	fmt.Println(words[0])
	fmt.Println(lists[0][0])

	for i, value := range lists{
		v := value
		l := len(value)
		words[i].Name = &v[0]
		words[i].Length = &l
	}
	fmt.Println(*words[0].Name)


	for _, w := range words{
		fmt.Print(*w.Name)
		fmt.Print(*w.Length)
		fmt.Println("\n")
	}

	sort.Slice()







}
