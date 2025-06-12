package main

import(
	"fmt"

)

func main() {
	var q1, q2, q3, q4, q5 string = "-", "-", "-", "-", "-"
	var count int
	for {
		var message string
		fmt.Scan(&message)
		if message == "конец"{
			fmt.Println("1.", q1)
			fmt.Println("2.", q2)
			fmt.Println("3.", q3)
			fmt.Println("4.", q4)
			fmt.Println("5.", q5)
			break
		} else if message == "очередь" {
			fmt.Println("1.", q1)
			fmt.Println("2.", q2)
			fmt.Println("3.", q3)
			fmt.Println("4.", q4)
			fmt.Println("5.", q5)
			continue
		} else if message == "количество" {
			fmt.Println("Осталось свободных мест:", 5-count)
			fmt.Println("Всего человек в очереди:", count)
			continue
		} 
		var num int
		fmt.Scan(&num)
		
		if num < 1 || num > 5{
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
		} else if q1 != "-" && q2 != "-" && q3 != "-" && q4 != "-" && q5 != "-"{
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
		} else if num == 1{
			if q1 == "-"{
				q1 = message
				count += 1
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			}
		} else if num == 2{
			if q2 == "-"{
				q2 = message
				count += 1
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			}
		} else if num == 3{
			if q3 == "-"{
				q3 = message
				count += 1
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			}
		} else if num == 4{
			if q4 == "-"{
				q4 = message
				count += 1
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			}
		} else if num == 5{
			if q5 == "-"{
				q5 = message
				count += 1
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			}
		}
	}
}