# golangで色々試してみるやつ

## ent

[github](https://github.com/ent/ent)
gormはinterface{}とリフレクションで機能を実現(してるらしい)->想定外の不具合が発生する
entは完全な型指定と明示的なAPI(コード自動生成による)->静的型付けと相性がいい、GraphQLとも相性がいい

- スキーマの雛形ファイルの生成

```shell
go run entgo.io/ent/cmd/ent init User
```

- スキーマファイルを元に、コードを自動生成

```shell
go generate ./ent
```
