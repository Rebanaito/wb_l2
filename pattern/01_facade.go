package main

type chef struct {
	// A human being with his multitude of personal qualities
	// that might affect his work. There might be a methhod
	// for our customer to interact with the chef somehow,
	// but generally we want this guy in the kitchen and that's it
}

func (c chef) prepareOrder(order bool) *ingredients {
	return &ingredients{}
}

type ingredients struct {
	// Specific details about ingredient proportions and
	// instructions on how to prepare them.
	// Our customer might have some general idea, but they don't
	// need to know the details
}

type Order struct {
	// The only object our customer will actually see and feel
}

type oven struct {
	// Oven internals: temperature modes, mechanical positions, etc.
}

func (o oven) cook(ingredients *ingredients) Order {
	return Order{}
}

type kitchen struct {
	chef chef
	oven oven
}

// The end customers doesn't need to know the intricacies of
// how their food is prepared, this whole thing happens in the kitchen
func (k kitchen) processOrder(order bool) Order {
	ingredients := k.chef.prepareOrder(order)
	food := k.oven.cook(ingredients)
	return food
}

type waiters struct {
}

func (w waiters) giveAdvice() bool {
	return true
}

func (w waiters) passOrderToKitchen() bool {
	return true
}

// A restaraunt has many moving parts within it
// However, as a customer, we don't want to see the behind-the-scenes
type Restaraunt struct {
	kitchen kitchen
	waiters waiters
}

func (r Restaraunt) AskAboutMenu() bool {
	return r.waiters.giveAdvice()
}

func (r Restaraunt) MakeOrder() Order {
	order := r.waiters.passOrderToKitchen()
	return r.kitchen.processOrder(order)
}
