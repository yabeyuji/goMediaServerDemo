# スマートリモコン on raspi セットアップ

### 今回使用した機器
- Raspberry Pi4(8GB) + SSD(1TB)
- Raspberry Pi Zero WH +  <a href="https://bit-trade-one.co.jp/adrszirs/">ADRSZIRS</a> x 2台

スマホから動画のアップロードできる機能もあるのでSSDのセットアップも手順に含めています。   
SSDはお好みのサイズで。   
赤外線送信にはhatと呼ばれる上にのせるタイプを使用しています。   
<br><br>

### raspiをsdカードにインストール
メディアサーバでモニタに出力するので最初からデスクトップ環境が入っているイメージをダウンロードします。   
手順は割愛します。   
<br><br>

### OSセットアップ
- 言語は日本語を選択   
- [Raspberry Piの設定]でSSHを有効化
- [Raspberry Piの設定]でI2Cを有効化
- [Raspberry Piの設定]でGPUメモリを512MBに変更   
  ※指定可能なmax値は896MBですが起動しなくなります
- 以下コマンドの実行
```
sudo apt install -y git libvlc-dev
```
<br><br>


### IPアドレス固定化
常に同じIPでアクセスできる方が便利なので。
```
sudo nano /etc/dhcpcd.conf
以下を追加(XXX YYY ZZZはルータの設定にあわせる)

# static ip
interface wlan0
static ip_address=192.168.XXX.YYY/24
static routers=192.168.XXX.ZZZ
static domain_name_servers=192.168.XXX.ZZZ
```
<br>

## git clone
```
cd ~
mkdir workspace && cd workspace
git clone git@github.com:YujiYabe/goMediaServerDemo.git goMediaServer
```
<br>

## ファイルを外部ストレージに保存
外部ストレージを接続   
ここでは外部ストレージのフォーマットはext4に指定しています   
ご自身の環境にあわせてください   
``` bash
# fstabバックアップ
cd /etc
sudo cp fstab fstab.bak

# 外部ストレージの情報を取得
sudo blkid /dev/sda1

# fstabを編集しマウント設定を追加
sudo nano /etc/fstab

# 以下を追加
PARTUUID=XXXXXXXX-XX /home/pi/workspace/goMediaServer/public/file ext4 nofail 0 0
```
<br>

### スクリーンセーバー解除
```
sudo nano /etc/xdg/lxsession/LXDE/autostart
# 以下を追加
@xset s off
@xset s noblank
@xset -dpms
```

```
sudo nano /etc/lightdm/lightdm.conf
# 以下を追加
[SeatDefaults]
xserver-command=X -s 0 -dpms
```
<br>

### 赤外線情報更新
家電のメーカーやモデルによって赤外線が違うので修正します   
以下のファイルのマップ変数を修正   
backend/device/internal/1_infrastructure/deviceapp/ir_data.go
```
var irDataList = map[string]string{}
irDataList["lightPower"] = "xxxxxxxxxxxxxxxxxxxxxx"
```

新しい赤外線情報などを追加したい場合などは下記ファイルも修正します。   
ここではdata-valueが「lightPower」の上記マップのキーとなります
public/vue/device/yyyy.vue
```
<b-col
  cols="4"
  :class="[commonClass, methodLightClass('lightPower')]"
  data-object="light"
  data-value="lightPower"
  @click="methodSendCommand($event)"
>
  <icon-power
    data-object="light"
    data-value="lightPower"
  />
</b-col>
```
<br><br>

### IR送信サーバ情報更新
IR送信サーバのアドレスをご自身の環境に変更します
backend/device/internal/1_infrastructure/deviceapp/deviceapp.go
```
shared.DataRoomLiving: "192.168.8.210",
shared.DataRoomBed:    "192.168.8.211",
```
<br><br>

### アプリセットアップ
```
mkdir ./public/file
mkdir ./public/file/video
mkdir ./public/file/temp
mkdir ./public/file/anime

echo "[]" > ./public/file/db.json
```

```
# vlcの出力を物理モニターに固定
echo 'export DISPLAY=:0' >> ~/.bashrc
```

```
git clone https://github.com/syndbg/goenv.git ~/.goenv

echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bashrc
echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(goenv init -)"' >> ~/.bashrc
echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.bashrc

exec $SHELL

goenv install 1.14.15
goenv global 1.14.15
```

<br>

### アプリ起動テスト
ターミナルを4つ立ち上げて以下を実行（起動しっぱなしになるので）   
ctrl+cで終了   
※初回起動はライブラリのダウンロードに時間がかかるのでサービス登録前に実施推奨   
```
make runws
```
```
make runfile
```
```
make runmedia
```
```
make rundevice
```

### サービス登録
```
make addws
```
```
make addfile
```
```
make addmedia
```
```
make adddevice
```
<br>

### サービス削除
```
cd /home/pi/workspace/goMediaServer
make delws
make delfile
make delmedia
make deldevice
```

<br>

# リンク

## 共通設計
<a href="docs/common_structure.md">こちら参照</a>
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
