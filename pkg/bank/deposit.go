
package bank

// //package deposit //5-1(bank)

// //Calculate расситывает параметры вклада
// func Calculate(amount int , currency string) (min int, max int) {
// 	minPercent, maxPercent :=PercentFor(currency)

// 	min=amount*minPercent/100/12
// 	max=amount*maxPercent/100/12

// 	return min,max
// }

// //PercentFor возвращает процент по вкладу по валюте
// func PercentFor(currency string) (min int, max int) {
// 	switch currency{
// 	case "TJS":
// 		return 4,6
// 	case "USD":
// 		return 1,2
// 	default:
// 		return 0,0
// 	}
// }