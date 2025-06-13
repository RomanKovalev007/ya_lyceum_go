package main

import (
	"fmt"
	"math"
	"time"
)

// модуль расчета БАЗОВОЙ ЦЕНЫ
type TripParameters struct {
	Distance float64
	Duration float64
}

const (
	pricePerKm = 10.0
	pricePerMinute = 2.0
)

func CalculateBasePrice(tp TripParameters) float64{
	basePrice := tp.Distance * pricePerKm + tp.Duration * pricePerMinute
	return basePrice
}
// конец модуля

// модуль расчета цены в ЧАСЫ ПИК
func GetTimeMultiplier(t time.Time) float64 {
    hour := t.Hour()
    isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday // Проверка, что сегодня суббота или воскресенье (выходные)

    switch {
    case hour >= 0 && hour < 5:
        return 1.5 // Ночной тариф
    case hour >= 7 && hour < 10 && !isWeekend:
        return 1.3 // Утренний час пик
    case isWeekend:
        return 1.2 // Выходные
    default:
        return 1.0
    }
}
// конец модуля

// модуль расчета цены с учетом ПОГОДНЫХ УСЛОВИЙ
type WeatherCondition int

const (
    Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64{
	k := 1.0
	if weather.Condition == Rain{
		k += 0.125
	} else if weather.Condition == HeavyRain{
		k += 0.2
	} else if weather.Condition == Snow{
		k += 0.15
	}
	if weather.WindSpeed > 15 {
		k += 0.1
	}
	return k
}
// конец модуля

// модуль расчета стоимости с учетом ПРОБОК
type TrafficClient interface {
    GetTrafficLevel(lat, lng float64) int // 1–5
}

func GetTrafficMultiplier(trafficLevel int) float64 {
    return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
    TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
    return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}
// конец модуля

// модуль ограничивающий максимальную/минимальную стоимость
const(
	minPrice = 99.0
	maxPrice = 20000
)

func ApplyPriceLimits(price float64) float64{
	if price > maxPrice{
		return maxPrice
	}
	if price < minPrice{
		return minPrice
	}
	return price
}
// конец модуля

// общий расчет стоимости
func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
    base := CalculateBasePrice(trip)
    timeMult := GetTimeMultiplier(now)
    weatherMult := GetWeatherMultiplier(weather)
    trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

    finalPrice := base * timeMult * weatherMult * trafficMult

    return ApplyPriceLimits(math.Round(finalPrice))
}

func main() {
    calculator := PriceCalculator{
        TrafficClient: &RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
    }
    
    price := calculator.CalculatePrice(
        TripParameters{Distance: 8.5, Duration: 20},
        time.Now(),
        WeatherData{HeavyRain, 10},
        55.751244, 37.618423,
    )
    
    fmt.Printf("Ваша цена: %.0f руб.\n", price)
}