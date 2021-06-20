# はじめに
<a href="../../../readme.md">こちら</a>のウェブサーバの解説です。
<br><br><br>

# 概要
### ウェブサーバで以下の機能を提供
- スマホからのビデオファイルをファイルサーバに転送
- ファイルサーバからのビデオファイル一覧をスマホに転送
- スマホからのメディア操作をメディアサーバに転送
- メディアサーバからの再生位置をスマホに転送
- 家電操作情報(エアコンの気温など)をデバイスサーバに転送
- デバイスサーバからの家電操作情報(エアコンの気温など)をスマホに転送
<br><br>


# サーバ内の各層の機能概要
### infrastructure
- grpc
  - 他の機能との共通コンテンツの受信や送信
  - バイナリ(wsでpostされたファイル)をstreamでファイルサーバに送信
- ws
  - echoとwebSocketで構成
  - webSocketでブラウザへの送信はチャンネルを使う
  - GET:indexとPOST:fileはrestで受信、それ以外はwebSocket上で受信
  - POST:fileはinfrastructure内でバイナリ変換後controllerに渡す
- network
  - WiFiのlocalアドレスを取得
<br><br>

### adapter
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### usecase
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### domain
- webSocketで受け取ったデータを元に宛先サーバの判別
    - ファイルサーバ
    - メディアサーバ
    - デバイスサーバ
<br><br>

# リンク
## セットアップ
<a href="docs/setup.md">こちら参照</a>
<br><br>

## 共通設計
<a href="docs/common_structure.md">こちら参照</a>
<br><br>

## サーバ
以下参照   
<a href="backend/file/docs/readme.md">ファイルサーバ</a>   
<a href="backend/media/docs/readme.md">メディアサーバ</a>   
<a href="backend/device/docs/readme.md">デバイスサーバ</a>   
<br><br>

## ブラウザ
<a href="public/docs/readme.md">こちら参照</a>
<br><br>
