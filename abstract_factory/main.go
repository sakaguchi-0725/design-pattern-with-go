package main

import "fmt"

/*
Abstract Factoryパターンの概要:
Abstract Factoryパターンは、**関連する一連のオブジェクト群を生成するためのインターフェース**を提供し、
具体的なクラスに依存せずにオブジェクトを生成できるようにするデザインパターンです。

特徴:
- 複数の関連するオブジェクト（例: ボタン、スクロールバー）を1つの「ファクトリ」によってまとめて生成します。
- クライアントコードは、生成される具体的なオブジェクトのクラスを意識せずに操作可能です。
- ファクトリ自体を切り替えることで、一連のオブジェクト群を変更可能です（例: Windows用GUIからmacOS用GUIへの切り替え）。

利点:
1. **柔軟性**: 新しいオブジェクト群（ファミリー）を追加する際、既存コードに影響を与えずに拡張できます。
2. **一貫性の確保**: 関連するオブジェクト群を同時に生成するため、組み合わせのミスを防止できます。
3. **保守性の向上**: 生成ロジックを抽象化しているため、クライアントコードがシンプルになります。

使用例:
- **クロスプラットフォームのGUIライブラリ**:
  OSごとに異なるボタンやウィジェットを提供する。
- **データベース接続ライブラリ**:
  MySQL、PostgreSQLなど、異なるデータベース用の接続オブジェクト群を生成する。
- **テーマ切り替え**:
  ダークテーマ、ライトテーマごとに異なるスタイルオブジェクトを生成する。
*/

func main() {
	adidasFactory, _ := GetSportsFactory(adidas)
	nikeFactory, _ := GetSportsFactory(nike)

	nikeShoe := nikeFactory.makeShoe(27)
	nikeShirt := nikeFactory.makeShirt(10)

	adidasShoe := adidasFactory.makeShoe(26)
	adidasShirt := adidasFactory.makeShirt(11)

	printShoeDetail(nikeShoe)
	printShirtDetail(nikeShirt)

	printShoeDetail(adidasShoe)
	printShirtDetail(adidasShirt)
}

func printShoeDetail(s IShoe) {
	fmt.Println("=== Shoe Detail ===")
	fmt.Println("Logo: ", s.getLogo())
	fmt.Println("Size: ", s.getSize())
}

func printShirtDetail(s IShirt) {
	fmt.Println("=== Shirt Detail ===")
	fmt.Println("Logo: ", s.getLogo())
	fmt.Println("Size: ", s.getSize())
}
