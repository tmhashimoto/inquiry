# 05: panic をわざと起こす

```bash
go run .          # 一覧
go run . nil      # 例
go run . json     # JSON の型まちがい
```

スタックトレースは次の順で読むとよいです。

1. 先頭の `panic: ...`（何が起きたか）
2. 自分の `main.go` の行（どこを直すか）

直し方は `06_panic_fixes` の同じケース名を実行してください。
