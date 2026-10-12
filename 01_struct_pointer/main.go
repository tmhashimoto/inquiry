package main

import "fmt"

// 構造体ポインタ *T
// 前提: *int などの普通のポインタは既習
//
// ポイント:
//   - 値渡しはコピーを触る → 呼び出し側は変わらない
//   - *T で渡すと本体を触る → 呼び出し側も変わる

type Student struct {
	Name  string
	Score int
}

func renameByValue(s Student, name string) {
	s.Name = name
}

func renameByPointer(s *Student, name string) {
	s.Name = name
}

func (s Student) RaiseByValue(delta int) {
	s.Score += delta
}

func (s *Student) RaiseByPointer(delta int) {
	s.Score += delta
}

func main() {
	st := Student{Name: "太郎", Score: 60}
	fmt.Println("初期値:", st)

	// 値渡し: コピーだけ変わる
	fmt.Println("\n--- 値渡し ---")
	fmt.Println("変更前:", st)
	renameByValue(st, "次郎")
	fmt.Println("変更後:", st)

	// ポインタ渡し: 本体が変わる
	fmt.Println("\n--- ポインタ渡し (*Student) ---")
	fmt.Println("変更前:", st)
	renameByPointer(&st, "次郎")
	fmt.Println("変更後:", st)

	// メソッドも同じ
	fmt.Println("\n--- 値レシーバ ---")
	fmt.Println("変更前:", st)
	st.RaiseByValue(10)
	fmt.Println("変更後:", st)

	fmt.Println("\n--- ポインタレシーバ ---")
	fmt.Println("変更前:", st)
	st.RaiseByPointer(10)
	fmt.Println("変更後:", st)

	// 最初からポインタで持つ場合
	fmt.Println("\n--- ポインタ変数 ---")
	p := &Student{Name: "花子", Score: 80}
	fmt.Println("変更前:", *p)
	p.RaiseByPointer(5)
	fmt.Println("変更後:", *p)
}
