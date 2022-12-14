## 手順  
  
1. スキーマ定義
  - hoge.graphqlsを生成。型定義
  - queryとmutationを定義する
    - query.graphqls
    - mutation.graphqls
2. コード生成  
```$ go run github.com/99designs/gqlgen```
  - `generated.go`やら`mutation.resolvers.go`、`query.resolvers.go`が作成されているはず
3. 生成されたコードの調整
  - 外部キー系の項目はそのまま構造体で定義されている（models_gen.go）
  - DBはだいたいIDで結合するはずで構造体のままだとまずいので、IDに変換してやる必要がある
    - models_gen.goは勝手に生成されるので放置。models.goを作成
    - models_gen.goのtypeをmodels.goに移行し、外部キー部分をIDに変更する  
    ※`hoge: Hoge!` となっている部分を`hogeID: ID!`にする
4. コード生成  
```$ go run github.com/99designs/gqlgen```
  - `hoge.resolvers.go`などが出来ているはず
5. リゾルバに処理を書く  
  - 各種リゾルバに処理を書いていく  
外部キー絡まないものは元の`query.resolvers.go`に書けばいいし、  
外部キーのものは*4*で生成されたresolversに書く  

## エラー処理  
- server.goのgraphqlサーバを編集する。  
[コード](https://github.com/watariRyo/go-graphql/blob/cc0e0a8e090d5a32790958878d5af6a75ab9c449/server.go#L46-L54)参照

## N + 1問題
1. dataloaderのrun  
```go get github.com/vektah/dataloaden```
- dataloader実装  
[dataloader](https://github.com/watariRyo/go-graphql/tree/main/graph/dataloader)以下参照  
`gen.go`を作成しloaderを作成するスキーマに対して`go run`する  
```go run github.com/vektah/dataloaden HogeLoader string "*github.com/watariRyo/go-graphql/graph/model.Hoge"```
※第二引数のstringはキー型によって変更する

2. graph/resolver.go  
- `graph/resolver.go`でLoaderを使うように定義
- `server.go`でGraphQLサーバ起動時に`CompanyLoader`を生成して、渡す。[コード](https://github.com/watariRyo/go-graphql/blob/cc0e0a8e090d5a32790958878d5af6a75ab9c449/server.go#L27-L44)
- [呼び出し](https://github.com/watariRyo/go-graphql/blob/cc0e0a8e090d5a32790958878d5af6a75ab9c449/graph/employee.resolvers.go#L20-L23)
  - `server.go`のほうで書いた`Fetch`の処理が呼び出され、一定溜めた後に処理が実行される
