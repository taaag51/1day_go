package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}

func (d dog) describe() string {
	dscr := "わたしは" + d.group
	dscr += "、名は" + d.name
	return dscr
}

func (d dog) bigger(d2 dog) string {
	dscr := d.describe()
	dscr += "、私は" + d2.name

	diff := d.height - d2.height
	switch {
	case diff > 5:
		dscr += "より大きい"
	case diff < -5:
		dscr += "より小さい"
	default:
		dscr += "と同じぐらいである"
	}
	return dscr
}

// func biggest(dogs []dog) dog {
// 	biggest := dogs[0]

// 	for i := 1; i < len(dogs); i++ {
// 		if dogs[i].height > biggest.height {
// 			biggest = dogs[i]
// 		}
// 	}
// 	return biggest
// }

func main() {
	maru := dog{"マル", "マルチーズ", 22}
	shiba := dog{"シバ", "芝犬", 40}

	fmt.Println(maru.bigger(shiba))

	// dogs := []dog{pome, maru, shiba}

	// fmt.Printf("%sさんが最大です\n", biggest(dogs).name)
}
