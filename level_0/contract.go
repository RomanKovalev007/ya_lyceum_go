package main

import(
	"fmt"
	"time"
	"math"
)

func main() {
	var input_date string
	var name string
	var surname string
	var fathername string
	var pay1 float64
	var pay2 float64
	var pay3 float64

	fmt.Scan(&input_date)
	date, err := time.Parse("02.01.2006", input_date)
	if err != nil {
		println("Ошибка при парсинге", err)
	}
	fmt.Scan(&name)
	fmt.Scan(&surname)
	fmt.Scan(&fathername)
	fmt.Scan(&pay1)
	fmt.Scan(&pay2)
	fmt.Scan(&pay3)

	new_date := date.AddDate(0, 0, 15)
	pay_sum := pay1 + pay2 + pay3
	pay_sum_r := int(pay_sum)
	pay_sum_kop := int(math.Round((pay_sum - float64(pay_sum_r))*100))
	
	message := fmt.Sprintf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.
Общая сумма выплат составит %d руб. %d коп.
	
С уважением, 
Гл. бух. Иванов А.Е.`, 
			surname, name, fathername,
			new_date.Format("02.01.2006"),
			pay_sum_r, pay_sum_kop)
	
		fmt.Println(message)
}