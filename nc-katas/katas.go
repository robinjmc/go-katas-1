package katas

import (
	"math"
	"strings"
)

/* ------- Section A ------- */

func GetLength(str string) int {
	return len(str)
}

func Modulo(n, d int) int {
	return n % d
}

func SquareRoot(n int) float64 {
	return math.Sqrt(float64(n))
}

func ShoutString(str string) string {
	return strings.ToUpper(str)
}

func JoinStrings(str1, str2 string) string {
	return str1 + str2
}

func RaiseToThePower(n, p int) int {
	if p == 0 {
		return 1
	} else if p == 1 {
		return n
	} else {
		powerRes := 0
		for i := 1; i < p; i++ {
			powerRes += n * n
		}
		return powerRes
	}
}

func IsBiggerThan42(n int) bool {
	return n > 42
}

func CelsiusToFahrenheit(c float64) float64 {
	//  °F = (°C x 1.8) + 32
	f := (c * 1.8) + 32
	roundToTwo := math.Round(f*100) / 100
	return roundToTwo
}

func IsMillenial(y uint16) bool {
	var top uint16 = 1997
	var bottom uint16 = 1980
	return bottom < y && y < top
}

/* ------- Section B ------- */

func ReadTrafficLight(s string) string {
	tL := strings.ToLower(s)
	if tL == "red" {
		return "STOP!"
	} else if tL == "amber" {
		return "GET READY..."
	} else {
		return "GO!"
	}
}

func PassOrFail(c, t int) string {
	if c == t {
		return "Top marks!"
	}
	p := float64(c) / float64(t)
	if p < .6 {
		return "Fail"
	} else {
		return "Pass"
	}
}

func TitleCaseString(t string) string {
	return strings.Title(t)
}

func SumEvensTo(n int) int {
	s := 0
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			s += i
		}
	}
	return s
}

func CheckInfinitive(v string) bool {
	if strings.HasSuffix(v, "er") {
		return true
	} else if strings.HasSuffix(v, "re") {
		return true
	} else if strings.HasSuffix(v, "ir") {
		return true
	} else {
		return false
	}
}

func HugsAndKisses(n int) string {
	var hK string
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			hK += "x"
		} else {
			hK += "o"
		}
	}
	return hK
}

func FizzBuzz(n int) string {
	isFizz := n%3 == 0
	isBuzz := n%5 == 0
	if isFizz && isBuzz {
		return "fizzbuzz!"
	} else if isFizz {
		return "fizz!"
	} else if isBuzz {
		return "buzz!"
	} else {
		return "..."
	}
}

// func CountTheVowels() {}

/* ------- Section C ------- */

// func GetMiddle() {}

// func ExtractNumber() {}

// func IsPythagoreanTriple() {}

// func GetCentury() {}

// func SumOfDigitSquares() {}

// func PigLatin() {}

// func IsJumpingNumber() {}
