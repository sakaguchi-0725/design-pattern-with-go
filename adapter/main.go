package main

/*
Adapterパターンの概要:
Adapterパターンは、互換性のない異なるインターフェースを持つクラスを、
クライアントコードが期待するインターフェースに適合させるための設計パターンです。

特徴:
- 既存のクラス（Adaptee）のインターフェースを変更することなく、新しい仕様に統一できます。
- クライアントコードはターゲットインターフェースだけを意識すればよく、実装の詳細を隠蔽できます。
- 異なるサービスやクラスを簡単に統一的に扱えるようになります。

このプログラムでは、支払い処理に対してAdapterパターンを適用し、
異なるインターフェースを持つ外部サービス（PayPal）を共通のインターフェース（Payment）に適合させています。

利点:
1. クライアントコードが具体的な支払いサービス（PayPal, Stripeなど）に依存せず、
   統一インターフェース（Payment）を通じて操作できるようになります。
2. 既存のコード（PayPalやStripe）を変更せずに統合できるため、レガシーコードや外部ライブラリの統合に最適です。
3. 新しい支払いサービスを追加する際も、Adapterを追加するだけで対応可能で、拡張性が高い設計です。

使用例:
- 外部APIの統合: 異なるサービス間のデータ形式を変換。
- 支払いシステム: PayPal、Stripeなど複数のプロバイダーを統一。
- レガシーシステムのモダナイズ: 古いシステムのインターフェースを新しい仕様に適合。
*/

func main() {
	paypal := &PayPal{}
	adapter := NewPayPalAdapter(paypal)
	adapter.Pay(50.2)
}
