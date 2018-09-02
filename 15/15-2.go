package main

func main() {
	// リストを宣言
	var a [4]int

	// リストを生成
	b := [4]int{}

	// リストを生成（初期値付き）
	c := [4]int{0, 1, 2, 3}

	// リストを生成（初期値付き/要素数は自動設定）
	d := [...]int{0, 1, 2, 3}

	type Vector3D [3]float64

	type Color [4]uint8
}
