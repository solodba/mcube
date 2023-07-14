# protoc命令生成go文件
```sh
protoc -I=. --go_out=. --go_opt=module="github.com/solodba/mcube" pb/meta/meta.proto
protoc -I=. --go_out=. --go_opt=module="github.com/solodba/mcube" pb/page/page.proto
```

# protoc-go-inject-tag给go文件打标签
```sh
protoc-go-inject-tag -input=pb/*/*.pb.go
```