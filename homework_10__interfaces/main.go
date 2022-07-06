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

type Dog struct {
	name       string
	weight     int
	kind       string
	animalType string
	amountFeed int
	weightFeed int
}

func (d Dog) CalcFeedWeight() int {
	return (d.amountFeed / d.weightFeed) * d.weight
}

func (d Dog) String() string {
	return fmt.Sprintf("%q -  weight: %dkg, food intake per month: %dkg", d.animalType, d.weight, d.CalcFeedWeight())
}

type Cat struct {
	name       string
	weight     int
	kind       string
	animalType string
	amountFeed int
	weightFeed int
}

func (c Cat) CalcFeedWeight() int {
	return c.weight * c.amountFeed
}

func (c Cat) String() string {
	return fmt.Sprintf("%q -  weight: %dkg, food intake per month: %dkg", c.animalType, c.weight, c.CalcFeedWeight())
}

type Cow struct {
	name       string
	weight     int
	kind       string
	animalType string
	amountFeed int
	weightFeed int
}

func (co Cow) CalcFeedWeight() int {
	return co.weight * co.amountFeed
}

func (co Cow) String() string {
	return fmt.Sprintf("%q -  weight: %dkg, food intake per month: %dkg", co.animalType, co.weight, co.CalcFeedWeight())
}

type FeetCalculator interface {
	CalcFeedWeight() int
}

type AnimalPrinter interface {
	fmt.Stringer
	FeetCalculator
}

func main() {
	dogs := []Dog{
		{"Lacky", 10, "doberman", "dog", 10, 5},
		{"Bobick", 113, "husky", "dog", 10, 5},
		{"Lis", 14, "dachshund", "dog", 10, 5},
		{"Ben", 5, "labrador", "dog", 10, 5},
	}

	cats := []Cat{
		{"Myrka", 10, "siam", "cat", 7, 1},
		{"Glasha", 10, "persian", "cat", 7, 1},
		{"Tomas", 10, "british", "cat", 7, 1},
		{"Murchik", 10, "bengal", "cat", 7, 1},
	}
	cows := []Cow{
		{"Marta", 150, "black", "cow", 25, 1},
		{"Zorka", 100, "white", "cow", 25, 1},
		{"Breez", 300, "red", "cow", 25, 1},
	}

	animals := make([]AnimalPrinter, 0, len(dogs))
	for _, dog := range dogs {
		animals = append(animals, dog)
	}

	for _, cat := range cats {
		animals = append(animals, cat)
	}

	for _, cow := range cows {
		animals = append(animals, cow)
	}

	PrintFarm(animals)
}

func PrintFarm(f []AnimalPrinter) {
	amount := 0
	for _, ftr := range f {
		fmt.Println(ftr.String())
		amount += ftr.CalcFeedWeight()
	}

	fmt.Println()
	fmt.Printf("the total amount of feed per month required for your farm = %dkg\n", amount)
}
