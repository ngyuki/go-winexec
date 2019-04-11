# go-winexec

`/etc/wsl.conf` で以下のように options で `umask=22,fmask=111` のようにして +x が付与されないようにしたところ・・

```ini
[automount]
enabled = true
options = "metadata,uid=1000,gid=1000,umask=22,fmask=111,case=off"
```

`*.exe` を WSL から直接実できなくなってしまい、`C:\Windows` や `c:\Program Files` とかだと `chmod +x` も効かないため、どうにかして `*.exe` を実行するためのラッパー。

例えば次のようにメモ帳を実行できます。

```sh
./winexec.exe notepad.exe
```

## 課題

以下のように `cmd.exe` から別コマンドを実行するとコンソールになにも表示されない。

```sh
./winexec.exe cmd.exe /C php.exe -v
#=>
```

なぜか `cmd.exe` のビルトインコマンドなら大丈夫。

```sh
./winexec.exe cmd.exe /C echo ok
#=> ok
```

バッチファイルでも同様で、下記のようなバッチファイルを実行すると `php -v` だけ出力されない。

```bat
@echo off
echo a
php -v
echo z
```

なぜか `more` すると大丈夫。

```bat
@echo off
more | more
echo a
php -v
echo z
```

いちど `more` するとそれ以降はずっと大丈夫。

```bat
@echo off
more | more
echo a
php -v
echo z
git --version
```
