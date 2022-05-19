// 学習教材: https://news.mynavi.jp/techplus/article/gogogo-5/
package main
import (
	"fmt"
	"log"
	"io/ioutil"
	"path/filepath"
)
func main(){
	// カレントディレクトリ以下のファイル一覧を得て出力 --- (*1)
	for _, f := range GetFiles("."){
		fmt.Println(f)
	}
}
// 再帰的にファイルの一覧を得る --- (*2)
func GetFiles(dir string) []string{
	var result []string
	files, err := ioutil.ReadDir(dir)
	if err != nil{
		log.Fatal(err)
	}
	for _,file := range files {
		/*
		filepathパッケージのJoin関数
		https://jablogs.com/detail/27609
		*/
		fpath := filepath.Join(dir, file.Name()) 

		if file.IsDir() { // サブディレクトリの一覧を得る --- (*3)
			/*
			append関数 --- スライスに要素を付け加える
			append(引用元の値, 追加値1, 追加値2, 追加値3, ... , 追加値n)
				    s := []int{0, 10, 20, 30}
				    fmt.Println(s) // => [0 10 20 30]
				    s = append(s, 1, 2, 3, 4)
				    fmt.Println(s) // => [0 10 20 30 1 2 3 4]
			*/
			result = append(result, GetFiles(fpath)...)
			continue //ループの先頭に戻る
		}
		result = append(result, fpath)
	}
	return result
}