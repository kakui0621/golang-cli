# golang-cli


## ビルド方法

1.環境変数を設定:

PowerShellでGoの環境変数を設定してビルドするには、以下のようにコマンドを実行してください。

```
$env:GOOS="windows"
$env:GOARCH="amd64"
```

2. ビルドコマンドの実行:

```powershell
go build -o <filename>.exe .\<filename>.go
```


## jsonformatter

1. ビルドコマンドの実行:
`go build -o ./jsonformatter/jsonformatter.exe`

2. 使用方法

```
echo '{"name":"John", "age":30, "city":"New York"}' | ./jsonformatter.exe

# 標準入力からJSON文字列を受け取り、整形されたJSONを標準出力に書き出します。
# json.UnmarshalでJSONを解析し、json.MarshalIndentで整形されたJSONを出力しています。
# このツールを使用する際は、JSON文字列をパイプしてツールに渡すことで、整形されたJSONを取得することができます。
```

## du

1. ビルドコマンドの実行:
`go build -o ./jsonformatter/jsonformatter.exe`

2. 使用方法

```
# 特定の拡張子（例：.log）のファイルのサイズを集計する場合:
./du -ext .log /path/to/directory

# サイズで結果をソートして表示する場合:
./du -sort /path/to/directory

# 2つの機能を組み合わせる場合:
./du -ext .log -sort /path/to/directory
```

## loganalyzer

1. ビルドコマンドの実行:

```
cd loganalyzer
`go build -o /loganalyzer.exe
```
2. 使用方法

```
# 実行（例えば、"ERROR"というキーワードを検索する場合）:
./loganalyzer.exe path_to_log_file.log ERROR

```


## httpclient

1. ビルドコマンドの実行:

```
cd httpclient
`go build -o /httpclient.exe
```

2. 使用方法

```
# 例:Googleのホームページの内容を取得する場合:
./httpclient.exe https://www.google.com
```


## localserver

指定したディレクトリ内の静的なHTML/CSS/JSファイルをホスティングするシンプルなWebサーバーを起動するツールです。

1. ビルドコマンドの実行:

```
cd localserver
`go build -o /localserver.exe
```

2. 使用方法

```
例:
myWebsiteというディレクトリ内のファイルをホスティングする場合:
./localserver.exe ./myWebsite
このコマンドを実行すると、http://localhost:8080/ でファイルにアクセスできるようになります。
```


## filechecker

指定されたディレクトリ内のファイルの内容をハッシュ化し、重複したファイルを検出するツールです。


1. ビルドコマンドの実行:

```
cd filechecker
`go build -o /filechecker.exe
```

2. 使用方法

```
例:
myFilesというディレクトリ内の重複ファイルを検出する場合:
./filechecker.exe ./myFiles

```





