package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// panic をわざと起こすサンプル
//
//   go run .          # 一覧
//   go run . nil      # 1ケース実行（まとめて実行すると最初で止まる）
//   go run . json
//
// 出たメッセージを読んだら、06_panic_fixes の同じケース名を実行する

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) < 1 {
		printUsage()
		return
	}

	switch args[0] {
	case "nil":
		caseNilPointer()
	case "index":
		caseIndexOutOfRange()
	case "map":
		caseNilMapWrite()
	case "assert":
		caseBadTypeAssert()
	case "div":
		caseDivideByZero()
	case "explicit":
		caseExplicitPanic()
	case "json":
		caseJSONTypeAssert()
	default:
		fmt.Println("不明なケース:", args[0])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`使い方: go run . <ケース>

ケース一覧:
  nil       nil ポインタ参照
  index     スライスの範囲外アクセス
  map       nil map への書き込み
  assert    型アサーション失敗
  div       整数のゼロ除算
  explicit  panic("...") を自分で呼ぶ
  json      JSON を any で受けて型を間違える

例: go run . nil`)
}

// --- ケース実装（どれも途中で panic する） ---

type User struct {
	Name string
}

func caseNilPointer() {
	fmt.Println("ケース: nil ポインタ参照")
	var u *User // nil
	fmt.Println("名前:", u.Name)
}

func caseIndexOutOfRange() {
	fmt.Println("ケース: index out of range")
	scores := []int{60, 70, 80}
	i := 3
	fmt.Println("scores[i] =", scores[i])
}

func caseNilMapWrite() {
	fmt.Println("ケース: nil map への書き込み")
	var scores map[string]int // nil（読み取りはできるが書き込みは panic）
	fmt.Println("読むだけなら OK:", scores["太郎"])
	scores["太郎"] = 90
}

func caseBadTypeAssert() {
	fmt.Println("ケース: 型アサーション失敗")
	var v any = "hello"
	n := v.(int)
	fmt.Println(n)
}

func caseDivideByZero() {
	fmt.Println("ケース: 整数のゼロ除算")
	a, b := 10, 0
	// 浮動小数の 0 除算は +Inf などになるが、整数は panic
	fmt.Println(a / b)
}

func caseExplicitPanic() {
	fmt.Println("ケース: 明示的な panic")
	// 到達してはいけない状態を表すときに使う（安易な多用は避ける）
	panic("これは意図した panic です: 設定ファイルが必須なのに無い")
}

func caseJSONTypeAssert() {
	fmt.Println("ケース: JSON を map[string]any で受けて型アサーションを間違える")

	// JSON の数値は、any に入れると float64 になる（int ではない）
	raw := []byte(`{"name":"太郎","age":20}`)
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		panic(err)
	}

	fmt.Println("Unmarshal 後:", m)
	fmt.Printf("age の実際の型: %T  値: %v\n", m["age"], m["age"])

	// よくある勘違い: JSON の 20 を int だと思ってキャストする → panic
	age := m["age"].(int)
	fmt.Println("age:", age)

	// ほかにも落ちやすい例（ここでは実行しない）:
	//   m["missing"].(string)           → キー無しは nil → panic
	//   m["name"].(map[string]any)      → string なのに map だと思う → panic
}
