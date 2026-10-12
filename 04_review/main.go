package main

import "fmt"

// 01〜03 の短い復習（変更前 / 変更後を表示）

type Student struct {
	Name  string
	Score int
}

func (s *Student) Raise(delta int) { s.Score += delta }

func main() {
	fmt.Println("--- 構造体 *T ---")
	st := Student{Name: "太郎", Score: 60}
	fmt.Println("変更前:", st)
	st.Raise(10)
	fmt.Println("変更後:", st)

	fmt.Println("\n--- スライス *[]T ---")
	nums := []int{1, 2, 3}
	fmt.Println("変更前:", nums)
	p := &nums
	*p = append(*p, 100)
	fmt.Println("変更後:", nums)

	fmt.Println("\n--- map *map ---")
	m := map[string]int{"太郎": 60}
	fmt.Println("中身 変更前:", m)
	m["太郎"] = 95
	fmt.Println("中身 変更後:", m)
	fmt.Println("差し替え 変更前:", m)
	pm := &m
	*pm = map[string]int{"リセット": 0}
	fmt.Println("差し替え 変更後:", m)
}
