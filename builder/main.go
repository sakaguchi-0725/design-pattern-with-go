package main

import "fmt"

/*
Builderパターンの概要:
Builderパターンは、複雑なオブジェクトを段階的に作成する方法を提供するデザインパターンです。
特に、設定項目が多かったり、生成手順が複雑だったりするオブジェクトを扱うときに役立ちます。

特徴:
- 作成するオブジェクトの生成手順を1つずつ定義できます。
- 最後に「完成したオブジェクト」を受け取れる仕組みです。
- 複数の生成手順を組み合わせることで、異なるバリエーションのオブジェクトを作成できます。

利点:
1. 複雑なオブジェクトを作る手順を分かりやすく整理できます。
2. オブジェクトの生成過程を柔軟に変更できます。
3. 作成手順を隠蔽できるため、クライアントコードがシンプルになります。

このプログラムでは、Builderパターンを利用して、
複雑なWeb APIリクエスト（HTTPメソッド、URL、ヘッダー、ボディなど）を簡単に構築しています。

使用例:
- **Web APIリクエスト**:
  URL、メソッド（GET/POST）、ヘッダー、ボディなどの設定が必要なリクエストを組み立てる。
- **レポート生成**:
  PDFやExcelのような複雑なレポートを段階的に作成する。
- **ゲームキャラクター作成**:
  名前、見た目、能力値などを一つずつ設定して完成させる。
*/

func main() {
	builder := NewRequestBuilder()
	req, err := builder.
		SetBody("GET").
		SetURL("https://example.com/api").
		AddHeader("Authorization", "Bearer token").
		Build()

	if err != nil {
		fmt.Println("リクエストの構築に失敗しました: ", err)
		return
	}

	fmt.Println("リクエストの構築が成功しました")
	fmt.Println("Method: ", req.Method)
	fmt.Println("URL: ", req.URL)
	fmt.Println("Headers: ", req.Header)
}
