# Go ポインタ実習（構造体 / スライス / map）+ panic

普通のポインタ（`*int` など）を学んだあとに進む教材です。  
コメントで説明し、実行すると **変更前 / 変更後** が印字されます。

## 準備

```bash
git clone https://github.com/tmhashimoto/inquiry.git
cd inquiry
```

Go 1.22 以降を想定しています。

## ポインタ編

```bash
go run ./01_struct_pointer   # 構造体のポインタ *T
go run ./02_slice_pointer    # スライスのポインタ *[]T
go run ./03_map_pointer      # map のポインタ *map
go run ./04_review           # 短い復習
```

覚えるときの軸:

> 中身を触るのか、変数そのものを差し替えるのか？

| 型 | いつポインタが要るか |
|---|---|
| `*Student` | フィールドを書き換えたいとき |
| `*[]T` | `append` / `nil` でスライス変数自体を差し替えるとき |
| `*map[K]V` | map 変数自体を差し替えるとき |

たとえ:

- スライス変数＝メモ、実際のデータ＝置き場
- map 変数＝メモ、実際のキーと値＝表

## panic 編

まずわざと落としてメッセージを読み、同じ名前のケースで直し方を見ます。

```bash
# 落とす
go run ./05_panic_cases nil
go run ./05_panic_cases json

# 直す（ケース名は上と同じ）
go run ./06_panic_fixes nil
go run ./06_panic_fixes json
go run ./06_panic_fixes recover
```

引数なしで実行すると、使えるケース一覧が出ます。

| ケース | 内容 |
|---|---|
| `nil` | nil ポインタ |
| `index` | スライスの範囲外 |
| `map` | nil map への書き込み |
| `assert` | 型アサーション失敗 |
| `div` | ゼロ除算 |
| `explicit` | 明示的な `panic` |
| `json` | JSON を `any` で受けて型を間違える（数値は `float64`） |

## ディレクトリ

```
01_struct_pointer/   構造体 *T
02_slice_pointer/    スライス *[]T
03_map_pointer/      map *map
04_review/           復習
05_panic_cases/      panic を起こす
06_panic_fixes/      直し方
```
