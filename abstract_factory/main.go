package main

import "fmt"

/*
Abstract Factoryパターンの主なメリットは、
関連する一連のオブジェクトを生成するためのインターフェースを提供し、
具体的なクラスに依存せずにオブジェクトを生成できることです。

これにより、クライアントコードは生成される具体的なオブジェクトのクラスを知らなくても、
オブジェクトを生成して使用することができます。

これにより、コードの柔軟性と拡張性が向上します。

また、オブジェクトの生成と使用を分離することで、
コードの保守性も向上します。

使用例としては、データベースアクセスライブラリなどで
さまざまなデータベースに対応するために、
データベースアクセスの抽象化を行う場合などに利用されます。
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
