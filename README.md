# 初めに
このリポジトリは静的解析・skeltonの勉強用のため、コード内のコメントは全て自分向けに作成しています。  
また、コミットの積み方も適当です。

# 使い方
`go build` で実行可能ファイルを作成し、PATHを通します。  
その後、カレントディレクトリを静的解析を行いたいコードがあるディレクトリにし、コマンドを実行します。  
```
go vet -vettool=`which checkvarname` ./...
```

このコマンドによって、checkvarnameの静的解析がカレントディレクトリ以下の全てのgoコードに行なわれます。  
私の過去のポートフォリオに対して上記のコマンドを行ったところ、下記の結果が出ました。

```
% go vet -vettool=`which checkvarname` ./...
# github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol
dbcontrol/comments.go:85:2: _ or - is used
dbcontrol/comments.go:102:2: _ or - is used
dbcontrol/initdb.go:9:2: _ or - is used
# github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol_test
dbcontrol/comments_test.go:1:9: _ or - is used
dbcontrol/comments_test.go:35:6: _ or - is used
dbcontrol/comments_test.go:92:6: _ or - is used
dbcontrol/comments_test.go:137:6: _ or - is used
dbcontrol/episodes_test.go:1:9: _ or - is used
dbcontrol/episodes_test.go:41:6: _ or - is used
dbcontrol/programs_test.go:1:9: _ or - is used
dbcontrol/programs_test.go:35:6: _ or - is used
以下略
```
