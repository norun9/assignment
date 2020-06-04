package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

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
		//fmt.Println(len(sameStrSlice))
		//fmt.Println(sameStrSlice)
		lists = append(lists, sameStrSlice)
	}
	fmt.Println(lists)
}

type Words struct {
	Name string
	Length int
}

func (w Words) String() string{
	return w.Name
}

type ByLen []Words

func (l ByLen) Len() int {
	return len(l)
}
