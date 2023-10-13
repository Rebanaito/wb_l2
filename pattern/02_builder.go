package main

import (
	"fmt"
	"strings"
)

// The end product our client will get to enjoy
// They get a finished product, a combined result of properly
// mixing the cooking the ingredients. Our client must only
// have one way of interacting with this product - eating it!
// So how do we create this complex dish with many possible options
// while still having it be the same kind of object?
type Pizza struct {
	name   string
	crust  int8
	spicy  bool
	addons addons
}

type addons struct {
	count        int
	anchovy      bool
	pineapple    bool
	cheezeBorder bool
	meatballs    bool
}

// We use a builder. This structure mimicks the end product (Pizza),
// we cannot eat it but we can modify it's contents since it's
// not a pizza yet
type PizzaBuilder struct {
	name   string
	crust  int8
	spicy  bool
	addons addons
}

// Below are the methods we can use on the builder to
// specify what our pizza should be like
func (b PizzaBuilder) ChoosePizza(name string) PizzaBuilder {
	b.name = name
	return b
}

func (b PizzaBuilder) CrustThin() PizzaBuilder {
	b.crust = 1
	return b
}

func (b PizzaBuilder) CrustThick() PizzaBuilder {
	b.crust = 3
	return b
}

func (b PizzaBuilder) Spicy() PizzaBuilder {
	b.spicy = true
	return b
}

func (b PizzaBuilder) Anchovy() PizzaBuilder {
	b.addons.anchovy = true
	return b
}

func (b PizzaBuilder) CheeseBorder() PizzaBuilder {
	b.addons.cheezeBorder = true
	return b
}

func (b PizzaBuilder) Pineapple() PizzaBuilder {
	b.addons.pineapple = true
	return b
}

func (b PizzaBuilder) Meatballs() PizzaBuilder {
	b.addons.meatballs = true
	return b
}

// Function that allows us to make a new builder
func NewPizzaBuilder() PizzaBuilder {
	var builder PizzaBuilder
	builder.name = "Pepperoni"
	return builder
}

// The function that we'll use to make the pizza
func (b PizzaBuilder) MakePizza() Pizza {
	var pizza Pizza
	pizza.name = b.name
	pizza.crust = b.crust
	pizza.spicy = b.spicy
	pizza.addons = b.addons
	return pizza
}

// Just some output formatting methods to support our Eat() method
func (p Pizza) spice() string {
	if p.spicy {
		return "spicy"
	}
	return ""
}

func (p Pizza) crustType() string {
	switch p.crust {
	case 1:
		return "thin"
	case 3:
		return "thick"
	}
	return "regular"
}

func appendAddon(addon bool, name string, str strings.Builder, count *int) {
	if !addon {
		return
	}
	if *count > 1 {
		str.WriteString(", ")
	} else {
		str.WriteString(" and ")
	}
	str.WriteString(name)
	*count--
}

func (p Pizza) chosenAddons() string {
	count := p.addons.count
	if count <= 0 {
		return ""
	}
	var str strings.Builder
	appendAddon(p.addons.anchovy, "anchovy", str, &count)
	appendAddon(p.addons.cheezeBorder, "cheeze border", str, &count)
	appendAddon(p.addons.meatballs, "meatballs", str, &count)
	appendAddon(p.addons.pineapple, "pineapple", str, &count)
	return str.String()
}

// The only method our client will ever need, leaving them
// completely out of the process of making a pizza
func (p Pizza) Eat() {
	fmt.Printf("I am gonna destroy this %s %s with %s crust%s! *CHEWING NOISES*",
		p.spice(), p.name, p.crustType(), p.chosenAddons())
}

// func main() {
// 	pizza := NewPizzaBuilder().ChoosePizza("Margarita").CrustThin().Spicy().MakePizza()
// 	pizza.Eat()
// }
