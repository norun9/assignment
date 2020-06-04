package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println(len(lists)+1)
	rank := [][]string{}
	fmt.Printf("配列listsの要素数：%d\n", len(lists))
	var prev int =  len(lists[0])

	for j := 0; j < len(lists); j++{
		//fmt.Printf("配列listsの要素：%v\n", lists[j])
		//fmt.Printf("配列内の要素数：%v\n",len(lists[j]))
		fmt.Printf("now word：%v\n",lists[j][0])

		current := len(lists[j])
		fmt.Printf("前の要素数：%v\n",prev)
		fmt.Printf("今の要素数：%v\n",current)
		if current > prev {
			rank = append([][]string{lists[j]}, rank...)
		//}else{
		//
		}
		prev = current
	}
	fmt.Println(rank)


}