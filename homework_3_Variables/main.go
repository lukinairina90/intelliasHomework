// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання.
// Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

package main

import "fmt"

func main() {
	const applePrice float64 = 5.99
	const pearPrice float64 = 7
	const budget float64 = 23

	var applesAmount float64 = 9
	var pearsAmount float64 = 8

	priceFor9ApplesAnd8Pears := applePrice*applesAmount + pearPrice*pearsAmount
	howManyApplesCanWeBuy := int(budget / applesAmount)
	howManyPearsCanWeBuy := int(budget / pearsAmount)
	canWeBuy2PearsAnd2Apples := applePrice*2+pearPrice*2 <= budget

	fmt.Printf("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\t %.2fгрн\n", priceFor9ApplesAnd8Pears)
	fmt.Printf("Скільки яблок ми можемо купити?\t %dшт.\n", howManyApplesCanWeBuy)
	fmt.Printf("Скільки груш ми можемо купити?\t %dшт.\n", howManyPearsCanWeBuy)
	if canWeBuy2PearsAnd2Apples == false {
		fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?\tні")
	} else {
		fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?\tтак")
	}
}
