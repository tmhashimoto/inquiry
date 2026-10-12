package main

import "fmt"

// map ポインタ *map[K]V
// 前提: 普通のポインタは既習 / 構造体ポインタは 01 で扱い済み
//
// たとえ:
//   map 変数 = 「どの表を見るか」というメモ
//   実際のキーと値 = 「表」
//
// ポイント:
//   - 表の中身を書き換える → map[K]V で足りる
//   - メモを別の表に差し替える → *map[K]V または戻り値

func setScore(scores map[string]int, name string, score int) {
	scores[name] = score
}

func replaceMap(scores map[string]int) {
	scores = map[string]int{"誰か": 0}
}

func replaceMapPtr(scores *map[string]int) {
	*scores = map[string]int{"誰か": 0}
}

func main() {
	// エントリ更新: 同じ表を見ているので *map は不要
	fmt.Println("--- エントリ更新 (map) ---")
	m := map[string]int{"太郎": 60}
	fmt.Println("変更前:", m)
	setScore(m, "太郎", 90)
	fmt.Println("変更後:", m)

	// 変数の差し替えを値渡し: 手元のメモだけ変わる
	fmt.Println("\n--- 差し替え (map だけ) ---")
	fmt.Println("変更前:", m)
	replaceMap(m)
	fmt.Println("変更後:", m)

	// 変数の差し替えを *map: 呼び出し側のメモが変わる
	fmt.Println("\n--- 差し替え (*map) ---")
	fmt.Println("変更前:", m)
	replaceMapPtr(&m)
	fmt.Println("変更後:", m)

	// 実務では戻り値で受け直すことが多い
	fmt.Println("\n--- 戻り値で受け直す ---")
	fmt.Println("変更前:", m)
	m = map[string]int{}
	fmt.Println("変更後:", m)
}
