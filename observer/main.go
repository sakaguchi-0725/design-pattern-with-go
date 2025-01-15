package main

/*
Observerパターンの概要:
Observerパターンは、1対多の依存関係を定義するデザインパターンです。
1つのオブジェクト（Subject）の状態が変化したときに、それに依存している複数のオブジェクト（Observer）に通知を送る仕組みを提供します。

特徴:
- 状態変化の通知: Subjectの状態が変わると、登録された全てのObserverに自動的に通知されます。
- 疎結合な設計: SubjectとObserverが直接的に依存せず、Observerの追加や削除が簡単に行えます。
- リアルタイム更新: 状態が変わるたびにObserverが即座に更新を受け取ります。

このプログラムでは:
1. `NewsService`（Subject）がニュースの追加やObserverの管理を行います。
2. `User`（Observer）がニュースを受け取り、更新通知を処理します。
3. クライアントコード（`main`関数）は`Observer`を動的に登録し、ニュースの通知を実行します。

利点:
1. 拡張性の向上: Observerを動的に追加・削除でき、Subject側の変更なしにシステムを拡張可能。
2. コードの整理: 状態変化の通知ロジックをSubjectに集約し、Observerの更新処理を分離。
3. 柔軟性の向上: Observerごとに異なる通知方法を実装可能（例: メール通知、LINE通知など）。

使用例:
- リアルタイム通知: ユーザーが登録したイベントの通知。
- イベント駆動システム: ボタン押下やフォーム入力時のリアルタイム処理。
- データ同期: 複数の画面やデバイスでデータをリアルタイムに更新。
*/

func main() {
	newsService := MakeNewsService()

	// ユーザーを作成
	user1 := &User{Name: "Alice"}
	user2 := &User{Name: "Bob"}

	// ユーザーを購読者として登録
	newsService.RegisterObserver(user1)
	newsService.RegisterObserver(user2)

	// ニュースを追加して通知
	newsService.AddNews("Breaking News: Observer Pattern in Go!")
}
