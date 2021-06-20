# はじめに
<a href="../readme.md">こちら</a>の共通設計です。
<br><br><br>


# 概要
ウェブサーバ上で以下4つの機能を独立してgrpcで通信しているが、一定のルールを設けて内部設計を同じにしている。   
また、共通の型、共通の処理は別パッケージに抜き出してそれぞれが参照できるようにシンボリックリンクで共通利用できるように設計
- ウェブサーバ
- ファイルサーバ
- メディアサーバ
- デバイスサーバ
<br><br><br>


# ディレクトリ設計
大枠の設計に<a href="https://github.com/golang-standards/project-layout/blob/master/README_ja.md">Standard Go Project Layout</a>を適用   
internal内のディレクトリ設計に<a href="https://blog.tai2.net/the_clean_architecture.html">clean Architecture</a>を適用
<br><br><br>



# internal設計

- プログラムの煩雑化を抑えるために以下ルールを適用
  - エラーの発生する処理はinfrastructureとdomainにしか書かない。
  - adapterとuseCaseはinfrastructureとdomainの呼び出ししかしない。


<br>

|infrastructure|adapter|useCase|domain|
|--------------|-------|-------|------|
| ①※1        |       |       |      |
|              | ②    |       |      |
| ⑦※2        |       | ③    | ④※1|
|              | ⑤    |       |      |
| ⑥※1        |       |       |      |
<br>

※1 エラーの起点箇所(ここ以外は呼び出ししかしないのでエラーの起点とはならない)   
※2 ⑥から①へのチャネルを介したデータ送信
### infrastructure
- infrastructure依存の型を使用できる   
- WAFのコンテキストやパラメータ、アップロードファイル、データベースの構造体に合わせた型などはここでしか使用できない。
<br><br>

### ① 外部からの受付処理やデーモンなど(grpc,webSocket,vlc)
- エラーの発生起点
- 一つ以上のcontrollerを扱う
- infrastructure依存の型や処理をフィルタ※してcontrollerに渡す   
※httpリクエスト構造体からdomain構造体への変換など
- バリデーションなどもここで。
<br><br>

### ② infrastructureから受け取った処理をuseCaseに渡す
- useCaseを組み合わせてresponseを返却する
- 実際の処理は書かずにuseCaseの呼び出ししかしない
- infrastructureで受け取った引数で処理分岐が必要になる場合はここで実施する
<br><br>

### ③ controllerから受け取った処理をdomainかserviceに渡す
- serviceとdomainを組み合わせてresponseを返却する
- 実際の処理は書かずにserviceとdomainの呼び出ししかしない
- 複数のservice(grpcとvlc等)の処理をまぜることが可能
- serviceから取得したデータの加工の必要であればdomainで実施する
- 処理単位はなるべく細かくする   
- 引数による内部で処理分岐はできない   
<br><br>

### ④ useCaseから呼び出させるdomainデータ型とデータの処理
- エラーの発生起点
- domainで扱うべきデータ型とその振る舞いについて実施
- JSONMarshalやJSONUnmarshalなどエラーが発生する標準関数もここで実施
<br><br>

### ⑤ useCaseから呼び出されinfrastructureに渡す処理
- 複数の同一のinfrastructureを含めることが可能(grpcとwebSocketなど他のinfrastructureの処理をまぜることはできない)
- infrastructureの結果によって同じinfrastructureの問い合わせは可能(副問合せなどでできない処理など)※   
  ※外部引数による処理分岐はできるが推奨されない。infrastructureno
  結果によって処理分岐は可能
- 取得したデータの加工はせずにそのまま返す
<br><br>

### ⑥ infrastructureの処理
- 基本的に1処理しかしない
- エラーの発生起点
- infrastructure依存の処理(sqlテーブル構造体からdomain構造体への相互変換など)
- 取得したデータの加工を行うことはできる(idのみ取り出しなど)
- デーモン(webSocket本体やvlc本体)に対しての処理はここからチャンネルを使って送信する
<br><br>


### ⑦ 内部通信
- 以下の処理でチャネルを使って情報をメインプロセスに送信している
  - ウェブサーバ・・grpcで受信した情報をブラウザに送信するため、コントローラ→ユースケース→サービスを介してインフラストラクチャのチャネルにデータを送信。
  - メディアサーバ・・grpcで受信した情報をvlcに送信するため、コントローラ→ユースケース→サービスを介してインフラストラクチャのチャネルにデータを送信。
  
  
# リンク
## セットアップ
<a href="docs/setup.md">こちら参照</a>
<br><br>

## サーバ
以下参照   
<a href="backend/ws/docs/readme.md">ウェブサーバ</a>   
<a href="backend/file/docs/readme.md">ファイルサーバ</a>   
<a href="backend/media/docs/readme.md">メディアサーバ</a>   
<a href="backend/device/docs/readme.md">デバイスサーバ</a>   
<br><br>

## ブラウザ
<a href="public/docs/readme.md">こちら参照</a>
<br><br>
