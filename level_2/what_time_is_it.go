package main

import (
	"errors"
	"strings"
	"time"
    "unicode/utf8"
)
func TimeNow() time.Time{
	return time.Now()
}

var (
	timeNow = TimeNow()
)

func currentDayOfTheWeek() string{
	timeNowWeekday := timeNow.Weekday().String()
	if timeNowWeekday == "Monday"{
		timeNowWeekday = "Понедельник"
	} else if timeNowWeekday == "Tuesday"{
		timeNowWeekday = "Вторник"
	} else if timeNowWeekday == "Wednesday"{
		timeNowWeekday = "Среда"
	} else if timeNowWeekday == "Thursday"{
		timeNowWeekday = "Четверг"
	} else if timeNowWeekday == "Friday"{
		timeNowWeekday = "Пятница"
	} else if timeNowWeekday == "Saturday"{
		timeNowWeekday = "Суббота"
	} else if timeNowWeekday == "Sunday"{
		timeNowWeekday = "Воскресенье"
	} 
	return timeNowWeekday
}

func dayOrNight() string{
	if timeNow.Hour() >= 10 && timeNow.Hour() <= 22{
		return "День"
	}
	return "Ночь"
}

func nextFriday() int{
	friday := time.Friday
	if timeNow.Weekday() == friday{
		return 0
	}
	return int((friday - timeNow.Weekday() + 7) % 7)
}

func CheckCurrentDayOfTheWeek(answer string) bool{
	answer = strings.ToLower(answer)
	timeNowWeekday := strings.ToLower(currentDayOfTheWeek())

	return answer == timeNowWeekday
}

func CheckNowDayOrNight(answer string) (bool, error){
	answer = strings.ToLower(answer)
	if utf8.RuneCountInString(answer) != 4{
		err := errors.New("исправь свой ответ, а лучше ложись поспать")
		return false, err
	}

	if answer == strings.ToLower(dayOrNight()){
		return true, nil
	}
	return false, nil
}