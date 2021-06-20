# はじめに
<a href="../../../readme.md">こちら</a>のデバイスサーバの解説です。
<br><br><br>

# 概要
### デバイスサーバで以下の機能を提供
- 受信した家電操作情報(エアコンの気温など)をIR送信サーバに転送
- 共有設定されている家電情報をウェブサーバに送信
<br><br>


# サーバ内の各層の機能概要
### infrastructure
- grpc
  - 他の機能との共通コンテンツの受信や送信
- device
  - 家電操作情報と赤外線データの関連付け
  - 赤外線データを追加・修正する場合はここのマップ変数に対して行う
<br><br>

### adapter
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### usecase
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### domain
- json.Marshal、json.Unmarshalなど標準パッケージ
- インメモリのデバイス管理情報の変更
<br><br><br>


# IR送信サーバ
http:GETリクエストで下記URLで送信する。   
http://IR送信サーバ:4001/赤外線データ情報   
backend/device/internal/1_infrastructure/deviceapp/deviceapp.go   


パスパラメータをそのまま赤外線に変換できればIR送信サーバ側の言語はどの言語でも問題ない。  
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
<a href="backend/file/docs/readme.md">ファイルサーバ</a>   
<a href="backend/media/docs/readme.md">メディアサーバ</a>   
<br><br>

## ブラウザ
<a href="public/docs/readme.md">こちら参照</a>
<br><br>
