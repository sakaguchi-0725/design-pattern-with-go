package main

import "fmt"

/*
Iteratorパターンの概要:
Iteratorパターンは、コレクション（リスト、配列、ツリーなど）の要素を
順番に操作する方法を提供するデザインパターンです。
クライアントコードは、コレクションの内部構造に依存せず、
共通のインターフェースを使ってコレクション内の要素を操作できます。

特徴:
- コレクションの構造を隠蔽し、簡潔なアクセス手段を提供します。
- 複数の探索アルゴリズムやアクセス方法を実現可能です。
- コレクションとその要素へのアクセスロジックを分離することで、柔軟性が向上します。

このプログラムでは:
- `UserCollection`がユーザーリストを保持し、専用の`UserIterator`を提供します。
- クライアントコード（`main`関数）は、`Iterator`インターフェースを通じてデータを順次処理します。

利点:
1. クライアントコードは、コレクションの内部構造（配列やリストなど）に依存しません。
2. 複数の探索方法を簡単に追加できるため、拡張性が高い設計です。
3. コレクションごとに異なるIteratorを持つことで、柔軟にアクセス方法をカスタマイズできます。

使用例:
- データベースの結果セットを順次処理する。
- REST APIから取得したリストデータを順に処理する。
- ページネーションやバッチ処理を実装する。
*/

func main() {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	collection := &UserCollection{users}
	iterator := collection.NewUserIterator()

	for iterator.HasNext() {
		user := iterator.Next()
		fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
	}
}
