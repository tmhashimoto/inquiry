package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// panic の直し方サンプル（05_panic_cases と同じケース名）
//
//   go run . nil
//   go run . json
//   go run . recover

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
		fixNilPointer()
	case "index":
		fixIndexOutOfRange()
	case "map":
		fixNilMapWrite()
	case "assert":
		fixBadTypeAssert()
	case "div":
		fixDivideByZero()
	case "explicit":
		fixExplicitPanic()
	case "json":
		fixJSONTypeAssert()
	case "recover":
		demoRecover()
	default:
		fmt.Println("不明なケース:", args[0])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`使い方: go run . <ケース>

ケース一覧（05 と同じ名前 + recover）:
  nil       nil チェック
  index     範囲チェック / エラー返し
  map       make で初期化
  assert    comma-ok 形式
  div       除数が 0 でないか確認
  explicit  前提チェックして error を返す（panic しない）
  json      JSON: comma-ok / 構造体へ Unmarshal
  recover   panic を捕まえてログし、処理を続ける例

例: go run . nil`)
}

type User struct {
	Name string
}

func fixNilPointer() {
	fmt.Println("=== 修正: nil ポインタ ===")
	var u *User

	if u == nil {
		fmt.Println("u が nil なのでデフォルトを使う")
		u = &User{Name: "(未設定)"}
	}
	fmt.Println("名前:", u.Name)

	// ポインタを受け取る関数でも、入口で守る
	printUserName(nil)
	printUserName(&User{Name: "太郎"})
}

func printUserName(u *User) {
	if u == nil {
		fmt.Println("printUserName: nil なのでスキップ")
		return
	}
	fmt.Println("printUserName:", u.Name)
}

func fixIndexOutOfRange() {
	fmt.Println("=== 修正: 範囲チェック ===")
	scores := []int{60, 70, 80}

	for _, i := range []int{1, 3, -1} {
		v, err := at(scores, i)
		if err != nil {
			fmt.Printf("scores[%d] → エラー: %v\n", i, err)
			continue
		}
		fmt.Printf("scores[%d] = %d\n", i, v)
	}
}

func at(nums []int, i int) (int, error) {
	if i < 0 || i >= len(nums) {
		return 0, fmt.Errorf("index %d out of range (len=%d)", i, len(nums))
	}
	return nums[i], nil
}

func fixNilMapWrite() {
	fmt.Println("=== 修正: map を make する ===")
	var scores map[string]int

	if scores == nil {
		scores = make(map[string]int)
	}
	scores["太郎"] = 90
	fmt.Println(scores)

	// 構造体フィールドが map のときも同様
	room := &Classroom{}
	ensureScores(room)
	room.Scores["花子"] = 85
	fmt.Println(room.Scores)
}

type Classroom struct {
	Scores map[string]int
}

func ensureScores(c *Classroom) {
	if c.Scores == nil {
		c.Scores = make(map[string]int)
	}
}

func fixBadTypeAssert() {
	fmt.Println("=== 修正: comma-ok アサーション ===")
	var v any = "hello"

	n, ok := v.(int)
	if !ok {
		fmt.Printf("int ではない（実際の型: %T, 値: %v）\n", v, v)
	} else {
		fmt.Println("int として取れた:", n)
	}

	s, ok := v.(string)
	if ok {
		fmt.Println("string として取れた:", s)
	}
}

func fixDivideByZero() {
	fmt.Println("=== 修正: 除数チェック ===")
	a, b := 10, 0

	result, err := safeDiv(a, b)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println("結果:", result)
}

func safeDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func fixExplicitPanic() {
	fmt.Println("=== 修正: 前提を error で返す ===")
	path := ""
	cfg, err := loadConfig(path)
	if err != nil {
		fmt.Println("呼び出し側で扱う:", err)
		return
	}
	fmt.Println("設定:", cfg)
}

func loadConfig(path string) (string, error) {
	if path == "" {
		// ライブラリや普通の業務コードでは panic より error
		return "", errors.New("config path is required")
	}
	return "ok:" + path, nil
}

func fixJSONTypeAssert() {
	fmt.Println("=== 修正: JSON の型を正しく扱う ===")
	raw := []byte(`{"name":"太郎","age":20,"profile":{"city":"東京"}}`)

	// 方法1（おすすめ）: 構造体に直接入れる。型が合わなければ error になる（panic しない）
	type Profile struct {
		City string `json:"city"`
	}
	type Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Profile Profile `json:"profile"`
	}
	var p Person
	if err := json.Unmarshal(raw, &p); err != nil {
		fmt.Println("構造体 Unmarshal エラー:", err)
		return
	}
	fmt.Println("OK 構造体:", p)

	// 方法2: どうしても any を使うなら comma-ok + 実際の型（数値は float64）
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("変更前(any):", m)
	fmt.Printf("age の実際の型: %T\n", m["age"])

	ageFloat, ok := m["age"].(float64)
	if !ok {
		fmt.Println("age が数値ではない")
		return
	}
	age := int(ageFloat)

	name, ok := m["name"].(string)
	if !ok {
		fmt.Println("name が string ではない（キー無し・型違い含む）")
		return
	}

	profile, ok := m["profile"].(map[string]any)
	if !ok {
		fmt.Println("profile が object ではない")
		return
	}
	city, _ := profile["city"].(string)

	fmt.Println("OK any+comma-ok: name=", name, "age=", age, "city=", city)
}

func demoRecover() {
	fmt.Println("=== recover: パニックを捕まえる ===")
	fmt.Println("（普段の if の代わりには使わない。境界での安全網）")

	safeRun("良い処理", func() {
		fmt.Println("  普通に終了")
	})

	safeRun("悪い処理", func() {
		var u *User
		_ = u.Name // ここで panic
	})

	fmt.Println("recover したので main は生きている")
}

// safeRun は演習・サーバのリクエスト境界などで使うイメージ
func safeRun(name string, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  [%s] panic を捕捉: %v\n", name, r)
		}
	}()
	fmt.Printf("実行: %s\n", name)
	fn()
}
