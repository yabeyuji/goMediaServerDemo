# はじめに
<a href="../../../readme.md">こちら</a>のファイルサーバの解説です。
<br><br><br>

# 概要
### ファイルサーバで以下の機能を提供
- ウェブサーバからのバイナリ（ビデオファイル）を受信しファイル保存
- ビデオファイル一覧をウェブサーバ・メディアサーバに送信
<br><br>


# サーバ内の各層の機能概要
### infrastructure
- grpc
  - 他の機能との共通コンテンツの受信や送信
  - バイナリ(wsでpostされたファイル)をstreamで受信
- file
  - 起動時にjsonファイルからファイル一覧を読み込みインメモリ化
  - インメモリで管理しているファイル一覧をjsonに保存
  - バイナリをmp4形式保存
  - mp4からanimeGIFを作成
<br><br>

### adapter
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### usecase
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### domain
- json.Marshal、json.Unmarshalなど標準パッケージ
- インメモリのファイル管理情報の変更
- 共通コンテンツの解析
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
<a href="backend/ws/docs/readme.md">ウェブサーバ</a>   
<a href="backend/media/docs/readme.md">メディアサーバ</a>   
<a href="backend/device/docs/readme.md">デバイスサーバ</a>   
<br><br>

## ブラウザ
<a href="public/docs/readme.md">こちら参照</a>
<br><br>
