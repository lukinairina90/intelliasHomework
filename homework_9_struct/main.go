package main

import "fmt"

//a) метод для типу Rectangle, який намалює в консолі прямокутник літерами “O” (як на прикладі в кутку),
//	 метод повинен давати вибір - малювати вертикально, або горизонтально.
//b) метод, що збільшить розмір Прямокутника в задану кількість разів
//c) на Прямокутнику зробити метод, що порівнює його площу із площею іншого Прямокутника
//d) створити тип Квадрат, зробити на Прямокутнику метод, що приймає Квадрат і відповідає скільки таких Квадратів можна
//	 вписати у цей Прямокутник
//e) Створити слайс Користувачів (структура як із лекції), вивести повне імʼя (fullName) і нік користувачів зі слайса,
//	 які старші за 20 років, і імʼя (firstName) яких “John”.

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	nick      string
	age       int
}

func NewUser(id string, email string, firstName string, lastName string, nick string, age int) User {
	return User{id: id, email: email, firstName: firstName, lastName: lastName, nick: nick, age: age}
}

func (u User) FullName() string {
	return u.firstName + " " + u.lastName
}

type Rectangle struct {
	Width  int
	Height int
}

func NewRectangle(width int, height int) Rectangle {
	return Rectangle{Width: width, Height: height}
}

func (r Rectangle) Draw(vertical bool) {
	if vertical {
		r.Width, r.Height = r.Height, r.Width
	}

	for i := 0; i < r.Height; i++ {
		for j := 0; j < r.Width; j++ {
			fmt.Print("0 ")
		}
		fmt.Println()
	}
}

func (r *Rectangle) Resize(size int) {
	r.Height, r.Width = r.Height*size, r.Width*size
}

func (r Rectangle) IsBiggerThan(rect Rectangle) bool {
	return r.Width*r.Height > rect.Width*rect.Height
}

func (r Rectangle) FitSquares(square Square) int {
	return (r.Width * r.Height) / (square.Side * 2)
}

type Square struct {
	Side int
}

func NewSquare(side int) Square {
	return Square{Side: side}
}

func main() {
	fmt.Println("----------------------------------------------Task a---------------------------------------------------")
	fmt.Println()
	rect1 := NewRectangle(9, 8)
	rect1.Draw(false)

	fmt.Println("----------------------------------------------Task b---------------------------------------------------")
	fmt.Println()
	rect1.Resize(2)
	rect1.Draw(false)

	fmt.Println("----------------------------------------------Task c---------------------------------------------------")
	fmt.Println()
	rect2 := NewRectangle(3, 5)
	res1 := rect2.IsBiggerThan(rect1)
	fmt.Println("Task c answer = ", res1)

	fmt.Println("----------------------------------------------Task d---------------------------------------------------")
	fmt.Println()
	square := NewSquare(2)
	res2 := rect1.FitSquares(square)
	fmt.Println("Task d answer = ", res2)

	fmt.Println("----------------------------------------------Task e---------------------------------------------------")
	fmt.Println()

	usersArr := []User{
		NewUser("f7d8sjnw1", "example1@gmail.com", "Karl", "Dekert", "KarlD", 23),
		NewUser("f7d8sjnw2", "example2@gmail.com", "Jan", "Vslasov", "JanV", 19),
		NewUser("f7d8sjnw3", "example3@gmail.com", "Monica", "Ruzh", "Mora", 25),
		NewUser("f7d8sjnw4", "example4@gmail.com", "John", "Dou", "Jou", 60),
		NewUser("f7d8sjnw5", "example5@gmail.com", "Mark", "Twen", "banana", 25),
	}

	fullName, nick := OldestTwenty(usersArr)
	fmt.Printf("fullName: %s\nnick:  %s", fullName, nick)
}

func OldestTwenty(users []User) (fullName, nick string) {
	for _, user := range users {
		if user.age > 20 && user.firstName == "John" {
			fullName = user.FullName()
			nick = user.nick
		}
	}
	return
}
