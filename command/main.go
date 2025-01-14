package main

/*
Commandパターンの概要:
Commandパターンは、操作（リクエストやタスク）をオブジェクトとしてカプセル化し、
操作の発行者（呼び出し元）と実行者（受け取り手）を分離するデザインパターンです。

特徴:
- リクエストをオブジェクト化することで、操作の履歴記録、再実行（Redo）、取り消し（Undo）を簡単に実現可能。
- 発行者と実行者を分離することで、操作の柔軟性と拡張性を向上。
- 操作をキューに追加して後から実行する非同期タスク管理や、操作をスケジュール化する用途にも適用可能。

このプログラムの目的:
このプログラムでは、Commandパターンを用いて異なる種類の通知タスク（メール送信、LINE送信）を統一的に管理し、順次実行します。
`Command`インターフェースを基準とすることで、新しい通知方法（SMSやプッシュ通知）を簡単に追加できるよう設計されています。

利点:
1. **操作の抽象化**:
   - メール送信、LINE送信など異なるタスクを統一された方法で管理可能。
2. **拡張性の向上**:
   - 新しい通知方法を追加する際、`Command`インターフェースを実装するだけで対応可能。
3. **柔軟性の向上**:
   - タスクをキューに追加して後から実行したり、スケジュール管理にも対応。

使用例:
- **非同期通知システム**:
  メールやLINE、SMSの送信タスクを管理。
- **バッチ処理**:
  一括タスクの管理や順次実行。
- **操作履歴の管理**:
  取り消し（Undo）や再実行（Redo）を実現。
*/

func main() {
	queue := &TaskQueue{}

	// コマンドを追加
	queue.Add(&SendEmailCommand{
		Recipient: "alice@example.com",
		Message:   "Hello Alice!",
	})
	queue.Add(&SendEmailCommand{
		Recipient: "bob@example.com",
		Message:   "Hello Bob!",
	})
	queue.Add(&SendLineCommand{
		LineID:  "john-XDG0-i1G2",
		Message: "Hello John!",
	})

	queue.Process()
}
