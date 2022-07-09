package main

import (
	"fmt"
)

//Домашнє завдання по інтерфейсах:
//кожна тварина в залежності від свого типу споживає різку кількість кілограмів їжі на місяць
//собака - їсть 10кг корму на кожні 5кг власної ваги
//кішка - 7кг на кожен кілограм власної ваги
//корова - 25кг на кожен кілограм власної ваги
//на фермі може бути різна кількість собак, кішок та корів
//кажен тип тварини має сам розраховувати для себе вагу корму

//написати динамічну функцію, яка буде виводити в консоль для кожної тварини на фермі її назву, вагу,
//та скільки їжі треба для того щоб її прогодувати ця функція також моє повертати сумму кг корму на всю ферму
//для подальшого виводу у консоль

const (
	dogEating        = 10
	catEating        = 7
	cowEating        = 25
	weightFeedForDog = 5
)

type Dog struct {
	name       string
	weight     int
	kind       string
	animalType string
}

func (d Dog) CalcFeedWeight() int {
	return (dogEating / weightFeedForDog) * d.weight
}

func (d Dog) String() string {
	return printInfo(d.animalType, d.weight, d.CalcFeedWeight())
}

type Cat struct {
	name       string
	weight     int
	kind       string
	animalType string
}

func (c Cat) CalcFeedWeight() int {
	return c.weight * catEating
}

func (c Cat) String() string {
	return printInfo(c.animalType, c.weight, c.CalcFeedWeight())
}

type Cow struct {
	name       string
	weight     int
	kind       string
	animalType string
}

func (co Cow) CalcFeedWeight() int {
	return co.weight * cowEating
}

func (co Cow) String() string {
	return printInfo(co.animalType, co.weight, co.CalcFeedWeight())
}

func printInfo(animalType string, weight, totalAmount int) string {
	return fmt.Sprintf("%q -  weight: %dkg, food intake per month: %dkg\n", animalType, weight, totalAmount)
}

type AnimalInfoPrinter interface {
	fmt.Stringer
	CalcFeedWeight() int
}

func main() {
	animals := make([]AnimalInfoPrinter, 0)
	animals = append(animals,
		Dog{"Lacky", 10, "doberman", "dog"},
		Dog{"Bobick", 113, "husky", "dog"},
		Dog{"Lis", 14, "dachshund", "dog"},
		Dog{"Ben", 5, "labrador", "dog"},
		Cat{"Myrka", 10, "siam", "cat"},
		Cat{"Glasha", 10, "persian", "cat"},
		Cat{"Tomas", 10, "british", "cat"},
		Cat{"Murchik", 10, "bengal", "cat"},
		Cow{"Marta", 150, "black", "cow"},
		Cow{"Zorka", 100, "white", "cow"},
		Cow{"Breez", 300, "red", "cow"},
	)

	PrintFarmInfo(animals)
}

func PrintFarmInfo(f []AnimalInfoPrinter) {
	amount := 0
	for _, ftr := range f {
		fmt.Printf(ftr.String())
		amount += ftr.CalcFeedWeight()
	}

	fmt.Println()
	fmt.Printf("the total amount of feed per month required for your farm = %dkg\n", amount)
}
