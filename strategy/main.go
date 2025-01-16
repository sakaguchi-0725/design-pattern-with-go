package main

import "fmt"

/*
Strategyパターンの概要:
Strategyパターンは、アルゴリズムや処理ロジックをカプセル化し、動的に切り替えられるようにするデザインパターンです。
これにより、条件分岐を減らし、コードの柔軟性や保守性を向上させることができます。

特徴:
- アルゴリズムの切り替えが容易:
  実行時に異なるアルゴリズムを選択可能。
- クライアントコードの単純化:
  アルゴリズムの実装をContextに委ねることで、クライアントコードが簡潔になる。
- 拡張性が高い:
  新しいアルゴリズムを追加しても、既存のコードに影響を与えない。

このプログラムでは:
1. `ShippingStrategy`インターフェースを定義し、料金計算ロジックを統一。
2. 各配送方法（通常配送、速達、海外配送）をConcreteStrategyとして実装。
3. `ShippingCostCalculator`（Context）が、動的に配送方法を切り替えて料金計算を実行。

利点:
1. コードの再利用性:
   各配送方法のロジックが個別のクラスとして実装されているため、再利用が容易。
2. 柔軟性の向上:
   実行時に配送方法を選択可能で、条件分岐の削減につながる。
3. 拡張性の向上:
   新しい配送方法を追加する場合でも、Strategyクラスを1つ追加するだけで対応可能。

使用例:
- 料金計算:
  配送方法やプラン（例: 通常、速達、国際配送）によって料金が変わる場合。
- 認証方式の切り替え:
  OAuth、JWT、セッション認証など、異なる認証方法を動的に選択。
- フォーマット変換:
  データをJSON、XML、CSVなど、異なる形式で出力。
*/

func main() {
	calc := &ShippingCostCalculator{}

	// 通常配送
	calc.SetStrategy(&RegularShipping{})
	fmt.Printf("Regular Shipping Cost: %.2f\n", calc.Calculate(10.0))

	// 速達
	calc.SetStrategy(&ExpressShipping{})
	fmt.Printf("Express Shipping Cost: %.2f\n", calc.Calculate(10.0))

	// 海外発送
	calc.SetStrategy(&InternationalShipping{})
	fmt.Printf("International Shipping Cost: %.2f\n", calc.Calculate(10.0))
}
