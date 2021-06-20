# はじめに
<a href="../../../readme.md">こちら</a>のメディアサーバの解説です。
<br><br><br>

# 概要
### メディアサーバで以下の機能を提供
- ウェブサーバからのメディア操作をvlcに反映
- vlcからの再生位置をウェブサーバに転送
- ファイルサーバからのファイル管理情報をvlcに反映
<br><br>


# サーバ内の各層の機能概要
### infrastructure
- grpc
  - 他の機能との共通コンテンツの受信や送信
- media
  - vlcとgoのvlcラッパーで構成
  - grpcから受信したメディア操作の送信はチャンネルを使う
<br><br>

### adapter
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>


### usecase
<a href="../../../docs/common_structure.md">こちら</a>参照
<br><br>

### domain
- プリミティブな型の変換
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
<a href="backend/file/docs/readme.md">ファイルサーバ</a>   
<a href="backend/device/docs/readme.md">デバイスサーバ</a>   
<br><br>

## ブラウザ
<a href="public/docs/readme.md">こちら参照</a>
<br><br>
