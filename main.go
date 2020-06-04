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

func (word Words) String() *string {
	return word.Name
}

type perLength []Words

func (length perLength) Len() int {
	return len()
}

func (a ByMass) Len() int           { return len(a) }
func (a ByMass) Less(i, j int) bool { return a[i].Mass < a[j].Mass }
func (a ByMass) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }




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

	var length int = len(lists)
	WordSlice := make([]Words, int(length))
	for i, value := range lists{
		v := value
		l := len(value)
		WordSlice[i].Name = &v[0]
		WordSlice[i].Length = &l
	}

	for _, w := range WordSlice{
		fmt.Print(*w.Name)
		fmt.Print(*w.Length)
		fmt.Println("\n")
	}
}
