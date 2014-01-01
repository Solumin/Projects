package main

import (
	"fmt"
)

func formatTriplet(trip int) string {
	if trip < 100 {
		return formatTens(trip)
	}

	if trip%100 == 0 {
		return fmt.Sprintf("%s hundred", formatDigit(trip/100))
	}

	return fmt.Sprintf("%s hundred %s", formatDigit(trip/100),
		formatTens(trip%100))
}

func formatDigit(digit int) string {
	switch digit {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "null digit"
	}
}

func formatTens(tens int) string {
	if tens < 10 {
		return formatDigit(tens)
	}

	if tens < 20 {
		switch tens {
		case 11:
			return "eleven"
		case 12:
			return "twelve"
		case 13:
			return "thirteen"
		// 14
		case 15:
			return "fifteen"
		// 16, 17
		case 18:
			return "eighteen"
		// 19
		default:
			return fmt.Sprintf("%steen", formatDigit(tens%10))
		}
	}

	switch tens / 10 {
	case 2:
		return fmt.Sprintf("twenty-%s", formatDigit(tens%10))
	case 3:
		return fmt.Sprintf("thirty-%s", formatDigit(tens%10))
	case 4:
		return fmt.Sprintf("fourty-%s", formatDigit(tens%10))
	case 5:
		return fmt.Sprintf("fifty-%s", formatDigit(tens%10))
	case 6:
		return fmt.Sprintf("sixty-%s", formatDigit(tens%10))
	case 7:
		return fmt.Sprintf("seventy-%s", formatDigit(tens%10))
	case 8:
		return fmt.Sprintf("eighty-%s", formatDigit(tens%10))
	case 9:
		return fmt.Sprintf("ninety-%s", formatDigit(tens%10))
	}

	return ""
}

func main() {
	fmt.Printf(" 12 = %s\n", formatTriplet(12))
	fmt.Printf(" 14 = %s\n", formatTriplet(14))
	fmt.Printf(" 99 = %s\n", formatTriplet(99))
	fmt.Printf("200 = %s\n", formatTriplet(200))
	fmt.Printf("314 = %s\n", formatTriplet(314))
}
