package main

import "fmt"

type dog struct {
	name   string
	group  string
	height int
}
type bird struct {
	name   string
	group  string
	height int
}

func (d dog) describe() string {
	dscr := "私は" + d.group
	dscr += "、名は" + d.name
	return dscr
}

func (b bird) describe() string {
	dscr := "私は" + b.name
	dscr += "、" + b.group + "の仲間"
	return dscr
}

func main() {
	pome := dog{"ポメ", "ポメラニアン", 25}
	meji := dog{"メジロ", "スズメ", 12}

	fmt.Println(pome.describe())
	fmt.Println(meji.describe())
}
