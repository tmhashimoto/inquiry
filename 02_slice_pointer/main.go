package main

import "fmt"

// スライスポインタ *[]T
// 前提: 普通のポインタは既習
//
// たとえ:
//   スライス変数 = 「メモ」（データの場所 / len / cap）
//   実際の数字の並び = 「データの置き場」
//
// ポイント:
//   - 置き場の中身を書き換える → []T で足りる
//   - メモを append / nil で差し替える → *[]T または戻り値
//   - *[]T（メモへのポインタ）と []*T（ポインタのスライス）は別物

func overwriteFirst(nums []int) {
	nums[0] = 999
}

func appendLocal(nums []int) {
	nums = append(nums, 100)
}

func appendViaPointer(nums *[]int) {
	*nums = append(*nums, 100)
}

func clearViaPointer(nums *[]int) {
	*nums = nil
}

func main() {
	// 要素の書き換え: 同じ置き場を見ているので []T で変わる
	fmt.Println("--- 要素の書き換え ([]T) ---")
	a := []int{1, 2, 3}
	fmt.Println("変更前:", a)
	overwriteFirst(a)
	fmt.Println("変更後:", a)

	// append を []T だけ: 手元のメモだけ更新され、呼び出し側は変わらない
	fmt.Println("\n--- append ([]T だけ) ---")
	b := []int{1, 2, 3}
	fmt.Println("変更前:", b)
	appendLocal(b)
	fmt.Println("変更後:", b)

	// append を *[]T: 呼び出し側のメモを差し替える
	fmt.Println("\n--- append (*[]T) ---")
	c := []int{1, 2, 3}
	fmt.Println("変更前:", c)
	appendViaPointer(&c)
	fmt.Println("変更後:", c)

	// nil 代入も *[]T
	fmt.Println("\n--- nil 代入 (*[]T) ---")
	d := []int{1, 2, 3}
	fmt.Println("変更前:", d)
	clearViaPointer(&d)
	fmt.Println("変更後:", d, "nil?", d == nil)

	// 実務では戻り値で受け直すことが多い
	fmt.Println("\n--- 戻り値で受け直す ---")
	e := []int{1, 2, 3}
	fmt.Println("変更前:", e)
	e = append(e, 100)
	fmt.Println("変更後:", e)

	// []*T: 要素がポインタのスライス（*[]T とは別）
	fmt.Println("\n--- []*T（要素がポインタ） ---")
	type Item struct{ Name string }
	items := []*Item{{Name: "りんご"}, {Name: "みかん"}}
	fmt.Println("変更前:", items[0].Name, items[1].Name)
	for _, it := range items {
		it.Name += "!"
	}
	fmt.Println("変更後:", items[0].Name, items[1].Name)

	// 同じ置き場を共有している例
	// s と t は同じ置き場を指す。t の len だけ増えても s の len は増えない
	fmt.Println("\n--- 同じ置き場の共有 ---")
	s := make([]int, 0, 3)
	s = append(s, 1, 2)
	t := s
	fmt.Println("append/書換前: s=", s, "t=", t)
	t = append(t, 3)
	t[0] = 9
	fmt.Println("append/書換後: s=", s, "t=", t)
}
