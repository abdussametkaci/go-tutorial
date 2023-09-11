package main

import "fmt"

var name, desc string
var radius int32
var mass float64
var active bool
var satellites []string

/*
var name, desc string = "Earth", "Planet"
var radius int32 = 6378
var mass float64 = 5.972E+24
var active bool = true
var satellites = []string{
  "Moon",
}
*/

/*
var name, desc = "Mars", "Planet"
var radius = 6755
var mass = 641693000000000.0
var active = true
var satellites = []string{
  "Phobos",
  "Deimos",
}
*/

/*
name := "Neptune"
    desc := "Planet"
    radius := 24764
    mass := 1.024e26
    active := true
    satellites := []string{
         "Naiad", "Thalassa", "Despina", "Galatea", "Larissa",
     "S/2004 N 1", "Proteus", "Triton", "Nereid", "Halimede",
         "Sao", "Laomedeia", "Neso", "Psamathe",
    }
*/

/*
var (
	name       string  = "Earth"
	desc       string  = "Planet"
	radius     int32   = 6378
	mass       float64 = 5.972e+24
	active     bool    = true
	satellites []string
)
*/

const (
	StarHyperGiant = iota
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
	StarDwarf
	StarSubDwarf
	StarWhiteDwarf
	StarRedDwarf
	StarBrownDwarf
)

const (
	StarDwarf2 = iota
	StarSubDwarf2
	StarWhiteDwarf2
	StarRedDwarf2
	StarBrownDwarf2
)

/*
const (
  StarDwarf byte = iota
  StarSubDwarf
  StarWhiteDwarf
  StarRedDwarf
  StarBrownDwarf
)
*/

/*
const (
  StarHyperGiant = 2.0*iota
  StarSuperGiant
  StarBrightGiant
  StarGiant
  StarSubGiant
)

*/

/*
_              = iota    // value 0
StarHyperGiant = 1 << iota
StarSuperGiant
StarBrightGiant
StarGiant
StarSubGiant
_          // value 64
StarDwarf
StarSubDwarf
StarWhiteDwarf
StarRedDwarf
StarBrownDwarf

---

StarHyperGiant = 2
StarSuperGiant = 4
StarBrightGiant = 8
StarGiant = 16
StarSubGiant = 32
StarDwarf = 128
StarSubDwarf = 256
StarWhiteDwarf = 512
StarRedDwarf = 1024
StarBrownDwarf = 2048
*/

func main() {
	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989e+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)

	fmt.Println("StarGiant", StarGiant)
	fmt.Println("StarWhiteDwarf2", StarWhiteDwarf2)

	fmt.Println("---")
	reverse("apple")
}

func reverse(s string) {
	for i := len(s) - 1; i >= 0; {
		fmt.Print(string(s[i]))
		i--
	}

	fmt.Println("")
}
