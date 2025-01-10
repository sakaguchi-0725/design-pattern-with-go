package main

/*
Factory Methodパターンの概要:
Factory Methodパターンは、オブジェクト生成の責任をクライアントから分離し、
生成方法をサブクラスや関数に委ねることで柔軟性を高めるデザインパターンです。

特徴:
- クライアントコード（呼び出し側）は、具体的なクラスではなく抽象インターフェースに依存します。
- 実際のオブジェクト生成はサブクラスや特定の関数が行うため、生成ロジックの変更や新しい型の追加が容易です。

利点:
1. クライアントコードが生成ロジックに依存しないため、拡張性が高まります。
2. 新しい型のオブジェクトを追加する際、既存コードへの影響を最小限に抑えられます。

このプログラムでは、Factory Methodパターンを利用して、
異なる通知方法（メール通知/SMS通知）を動的に生成しています。

使用例:
- **通知システム**:
  メール通知、SMS通知など、異なる通知方法に応じた通知オブジェクトを生成する。
- **ドキュメントフォーマット**:
  PDF、Wordなど、ファイル形式ごとに異なるエクスポートオブジェクトを生成する。
- **ログ出力**:
  ファイル、コンソール、リモートサーバーなど、出力先に応じて異なるLoggerオブジェクトを生成する。
*/

func main() {
	var factory NotifierFactory

	factory = &EmailNotifierFactory{}
	notifier := factory.NewNotifier()
	notifier.Notify("メールでのお知らせです")

	factory = &SMSNotifierFactory{}
	notifier = factory.NewNotifier()
	notifier.Notify("SMSでのお知らせです")
}
