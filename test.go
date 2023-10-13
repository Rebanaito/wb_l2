package main

func main() {
	pizza := pattern().NewPizzaBuilder().ChoosePizza("Margarita").CrustThin().Spicy().MakePizza()
	pizza.Eat()
}
