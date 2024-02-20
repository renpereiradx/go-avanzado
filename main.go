package main

import "fmt"

// type Any struct{
// 	Name string
// 	LastName string
// 	Age int32
// }

func main() {
	rows := [][]any{
		{"John", "Smith", int32(36)},
		{"Jane", "Doe", int32(29)},
	}

	for i := range rows {
		for j := range rows[i] {
			fmt.Println(i, j)
			fmt.Println(rows[i][j])
		}
	}

}
