package main

func makeAmericano(size float32) {
	beansAmt := (size / 8.0) * 5.0
	americano := &CoffeeMachine{}
	americano.startCoffee(beansAmt, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()
}

func makeLatte(size float32, foam bool) {
	beansAmt := (size / 8.0) * 5.0
	milkAmt := (size / 8.0) * 2.0

	latte := &CoffeeMachine{}
	latte.startCoffee(beansAmt, 4)
	latte.grindBeans()
	latte.useHotWater(size)
	latte.useMilk(milkAmt)
	latte.doFoam(foam)
	latte.endCoffee()
}
