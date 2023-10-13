package main

import "fmt"

type maneuver interface {
	attack()
}

type defensive struct {
}

func (d *defensive) attack() {
	fmt.Println("The regiment is now holding a defensive position")
}

type envelopment struct {
}

func (e *envelopment) attack() {
	fmt.Println("The regiment has split up to flank the enemy")
}

type feint struct {
}

func (f *feint) attack() {
	fmt.Println("The regiment is feigning a retreat to lure the enemy")
}

type center struct {
}

func (c *center) attack() {
	fmt.Println("The regiment is exploiting a weakness in the enemy line to cut through it")
}

type regiment struct {
	personel   []struct{}
	atvs       []struct{}
	airSupport []struct{}
	tactic     maneuver
}

func (r *regiment) attack() {
	r.tactic.attack()
}

func (r *regiment) chooseTactic(m maneuver) {
	r.tactic = m
}

// func main() {
// 	def := &defensive{}
// 	cent := &center{}
// 	env := &envelopment{}
// 	feint := &feint{}
// 	regiment := &regiment{}
// 	regiment.chooseTactic(env)
// 	regiment.attack()
// 	regiment.chooseTactic(def)
// 	regiment.attack()
// 	regiment.chooseTactic(feint)
// 	regiment.attack()
// 	regiment.chooseTactic(cent)
// 	regiment.attack()
// }
