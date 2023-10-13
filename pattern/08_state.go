package main

import "fmt"

type pizza struct {
	noPizza   state
	coldStove state
	coldPizza state
	ready     state

	currentState state
	timeCooked   uint
}

func (p *pizza) preparePizza() error {
	return p.currentState.preparePizza()
}

func (p *pizza) heatUpStove() error {
	return p.currentState.heatUpStove()
}

func (p *pizza) cook() error {
	return p.currentState.cook()
}

func (p *pizza) eat() error {
	return p.currentState.eat()
}

type state interface {
	preparePizza() error
	heatUpStove() error
	cook() error
	eat() error
}

type noPizza struct {
	pizza *pizza
}

func (no *noPizza) preparePizza() error {
	fmt.Println("Preparing the dough... adding ingredients... Done, time to prepare the stove")
	no.pizza.currentState = no.pizza.coldStove
	return nil
}

func (no *noPizza) heatUpStove() error {
	return fmt.Errorf("Too early, we don't have a pizza yet")
}

func (no *noPizza) cook() error {
	return fmt.Errorf("Too early, we don't have a pizza yet")
}

func (no *noPizza) eat() error {
	return fmt.Errorf("Too early, we don't have a pizza yet")
}

type coldStove struct {
	pizza *pizza
}

func (cold *coldStove) preparePizza() error {
	return fmt.Errorf("The pizza is already prepared")
}

func (cold *coldStove) heatUpStove() error {
	fmt.Println("Heating up the stove... it's ready to cook")
	cold.pizza.currentState = cold.pizza.coldPizza
	return nil
}

func (cold *coldStove) cook() error {
	return fmt.Errorf("Too early, the stove is still cold and so is the pizza")
}

func (cold *coldStove) eat() error {
	return fmt.Errorf("Too early, we haven't cooked the pizza yet because the stove is still cold")
}

type coldPizza struct {
	pizza *pizza
}

func (cold *coldPizza) preparePizza() error {
	return fmt.Errorf("The pizza is already prepared")
}

func (cold *coldPizza) heatUpStove() error {
	return fmt.Errorf("The stove is already hot, time to put pizza int")
}

func (cold *coldPizza) cook() error {
	fmt.Println("Cooking the pizza for 5 minutes")
	cold.pizza.timeCooked += 5
	if cold.pizza.timeCooked >= 15 {
		fmt.Println("The pizza is done, time to take it out")
		cold.pizza.currentState = cold.pizza.ready
	}
	return nil
}

func (cold *coldPizza) eat() error {
	return fmt.Errorf("Too early, the pizza is still cooking")
}

type ready struct {
	pizza *pizza
}

func (pizza *ready) preparePizza() error {
	return fmt.Errorf("You still have a pizza to eat, get at it while it's hot!")
}

func (pizza *ready) heatUpStove() error {
	return fmt.Errorf("Just eat the pizza you just cooked")
}

func (pizza *ready) cook() error {
	return fmt.Errorf("If you cook this pizza any longer it will burn")
}

func (pizza *ready) eat() error {
	pizza.pizza.currentState = pizza.pizza.noPizza
	fmt.Println("*CHEWING NOISES*")
	return nil
}

func main() {
	margarita := &pizza{}
	noPizza := &noPizza{pizza: margarita}
	coldStove := &coldStove{pizza: margarita}
	coldPizza := &coldPizza{pizza: margarita}
	ready := &ready{pizza: margarita}
	margarita.noPizza = noPizza
	margarita.coldStove = coldStove
	margarita.coldPizza = coldPizza
	margarita.ready = ready
	margarita.currentState = margarita.noPizza
	fmt.Println(margarita.eat())
	margarita.preparePizza()
	margarita.heatUpStove()
	fmt.Println(margarita.eat())
	margarita.cook()
	fmt.Println(margarita.eat())
	margarita.cook()
	margarita.cook()
	margarita.eat()
}
