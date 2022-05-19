package main
import(
	"fmt"
	"log"
	"io/ioutil"
)
func main(){
	// カレントディレクトリのファイル一覧を得る
	/*ioutil.ReadDir関数 --- 引数にディレクトリを与えるとstring型スライスを返す*/
	files, err := ioutil.ReadDir(".")
	/* 
	Go言語ではエラーのない場合は結果がnilではないことを前提にしている
	*/
	if err != nil { // nilは「無」を表す。nullと同じイメージ
		log.Fatal(err) // ★Fatal(=強制終了)は極力使わない.
	}
	// 取得した一覧を表示
	for _, file := range files {
		fmt.Println(file.Name())
	}
}