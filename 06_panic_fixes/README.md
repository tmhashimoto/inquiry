# 06: panic の直し方

`05_panic_cases` と同じケース名で、落ちない版を実行します。

```bash
go run . nil
go run . index
go run . map
go run . assert
go run . div
go run . explicit
go run . json
go run . recover
```

| panic | 直し方の目安 |
|---|---|
| nil ポインタ | `if p == nil` |
| index | `0 <= i && i < len(s)`、または error |
| nil map 書き込み | `make(map[K]V)` |
| 型アサーション | `v, ok := x.(T)` |
| ゼロ除算 | 除数を確認して error |
| 明示 panic | 業務では error を返す |
| JSON + any | 数値は `float64`。できれば構造体へ Unmarshal |

`recover` は普段の `if` の代わりには使いません（境界の安全網向け）。
