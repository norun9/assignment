package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Word struct {
	Name   *string
	Length *int
}

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	// 標準入力で受け取ったバイト型を文字列に変換
	b := string(body)
	//"."を基準に同じ文字のスライスを作成
	newSlice := strings.Split(b, ".")

	//newSliceの末尾要素が空白になるので末尾要素以外を再代入
	newSlice = newSlice[:len(newSlice)-1]

	//同じ文字同士のスライスを作成する為に、空の多次元スライスを変数wordsSliceで定義
	wordsSlice := [][]string{}

	for i := 0; i < len(newSlice); i++ {

		//空白を基準に同じ文字同士のスライスを作成
		collectSameWords := strings.Split(newSlice[i], " ")

		//collectSameWordsの時点で先頭要素が空のスライスが存在するので、先頭要素""を削除した残りの要素を追加
		if collectSameWords[0] == "" {
			collectSameWords = append(collectSameWords[:0], collectSameWords[1:]...)
		}
		//上記の条件以外はwordsSliceに対してcollectSameWordsスライスを追加
		wordsSlice = append(wordsSlice, collectSameWords)
	}

	//スライス構造体の初期化で要素数を可変長にする
	//wordsSliceの要素数に応じた数のスライスが作成される
	words := make([]Word, len(wordsSlice))

	for i, word := range wordsSlice {
		//変数wには要素の値を、変数lには各スライス内の要素数を定義
		w := word
		l := len(word)
		//構造体Wordの各フィールド名
		//Nameフィールドには要素の値を一つだけ代入(表示させる要素は一つだけで良い為)
		//Lengthフィールドには各スライス内の要素数を代入
		words[i].Name = &w[0]
		words[i].Length = &l
	}

	//隣の要素と順に比較していき要素の大きさ順と、もし要素数が比較対象と同じであればアルファベット順でソートしていく
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

	//TOP3表示
	//要素の大きさ順かつアルファベット順でソートしたスライスを3つまで順々に表示させる
	for i := 0; i < 3; i++ {
		fmt.Println(*words[i].Name)
	}
}
