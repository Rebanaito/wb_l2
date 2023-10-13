package main

import "fmt"

type Weapon interface {
	GetName() string
	GetDmg() uint
	GetMag() uint
	Shoot()
}

type gun struct {
	name  string
	sound string
	dmg   uint
	mag   uint
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) GetDmg() uint {
	return g.dmg
}

func (g *gun) GetMag() uint {
	return g.mag
}

func (g *gun) Shoot() {
	for i := 0; i < int(g.mag); i++ {
		fmt.Printf("%s ", g.sound)
	}
	fmt.Println("*reloading sounds*")
}

type pistol556 struct {
	gun
}

type magnum357 struct {
	gun
}

type huntingRevolver struct {
	gun
}

func make556() Weapon {
	return &pistol556{
		gun: gun{
			name:  "5.56mm pistol",
			dmg:   28,
			mag:   5,
			sound: "bam",
		},
	}
}

func make357() Weapon {
	return &magnum357{
		gun: gun{
			name:  ".357 Magnum",
			dmg:   26,
			mag:   6,
			sound: "bap",
		},
	}
}

func makeHuntingRevolver() Weapon {
	return &huntingRevolver{
		gun: gun{
			name:  "Hunting revolver (GRA)",
			dmg:   58,
			mag:   5,
			sound: "BOOM",
		},
	}
}

func GRAVendotron(gunName string) Weapon {
	switch gunName {
	case "5.56mm pistol":
		return make556()
	case ".357 Magnum":
		return make357()
	case "Hunting revolver":
		return makeHuntingRevolver()
	}
	return nil
}

// func main() {
// 	maria := GRAVendotron("Maria")
// 	if maria == nil {
// 		fmt.Println("Gun Runners Arsenal does not sell this weapon")
// 	}
// 	pistol := GRAVendotron("5.56mm pistol")
// 	magnum := GRAVendotron(".357 Magnum")
// 	huntingRevolver := GRAVendotron("Hunting revolver")

// 	fmt.Println("First pistol is called", pistol.GetName())
// 	fmt.Println("It deals", pistol.GetDmg(), "damage per shot")
// 	fmt.Println("It can shoot", pistol.GetMag(), "times before having to reload")
// 	pistol.Shoot()
// 	fmt.Println()

// 	fmt.Println("Second pistol is called", magnum.GetName())
// 	fmt.Println("It deals", magnum.GetDmg(), "damage per shot")
// 	fmt.Println("It can shoot", magnum.GetMag(), "times before having to reload")
// 	magnum.Shoot()
// 	fmt.Println()

// 	fmt.Println("Third pistol is called", huntingRevolver.GetName())
// 	fmt.Println("It deals", huntingRevolver.GetDmg(), "damage per shot")
// 	fmt.Println("It can shoot", huntingRevolver.GetMag(), "times before having to reload")
// 	huntingRevolver.Shoot()
// }
