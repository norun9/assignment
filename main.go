package main

import (
	"io/ioutil"
	"os"
	"fmt"
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
		//fmt.Println(len(sameStrSlice))
		//fmt.Println(sameStrSlice)
		lists = append(lists, sameStrSlice)
	}
	fmt.Println(len(lists))
	var length int = len(lists)
	hoges := make([]Words, int(length))
	for i, value := range lists{
		v := value
		l := len(value)
		hoges[i].Name = &v[0]
		hoges[i].Length = &l
	}

	for _, hoge := range hoges{
		fmt.Print(*hoge.Name)
		fmt.Print(*hoge.Length)
		fmt.Print(" ")
	}
}
