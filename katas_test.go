package katas_test

import (
	katas "go-katas"
	"math"
	"testing"
	assert "testing-utilities"
)

func TestGetLength(t *testing.T) {
	t.Run("returns 0 when passed empty string", func(t *testing.T) {
		got := katas.GetLength("")
		want := 0
		assert.Int(t, got, want)
	})
	t.Run("returns 5 when passed string word: hello", func(t *testing.T) {
		got := katas.GetLength("hello")
		want := 5
		assert.Int(t, got, want)
	})
	t.Run("returns 32 when passed string sentence: I hear the drums echoing tonight", func(t *testing.T) {
		got := katas.GetLength("I hear the drums echoing tonight")
		want := 32
		assert.Int(t, got, want)
	})
}

func TestModulo(t *testing.T) {
	t.Run("returns 0 when passed 6, 3", func(t *testing.T) {
		got := katas.Modulo(6, 3)
		want := 0
		assert.Int(t, got, want)
	})
	t.Run("returns 1 when passed 5, 2", func(t *testing.T) {
		got := katas.Modulo(5, 2)
		want := 1
		assert.Int(t, got, want)
	})
}

func TestSquareRoot(t *testing.T) {
	t.Run("returns 3 when passed 9", func(t *testing.T) {
		got := katas.SquareRoot(9)
		want := 3.0
		assert.Float(t, got, want)
	})
	t.Run("returns 5 when passed 25", func(t *testing.T) {
		got := katas.SquareRoot(25)
		want := 5.0
		assert.Float(t, got, want)
	})
	t.Run("returns 1.414... when passed 2", func(t *testing.T) {
		got := katas.SquareRoot(2)
		want := math.Sqrt(2.0)
		assert.Float(t, got, want)
	})
}

func TestShoutString(t *testing.T) {
	t.Run("returns A when passed a", func(t *testing.T) {
		got := katas.ShoutString("a")
		want := "A"
		assert.String(t, got, want)
	})
	t.Run("returns LET'S GO! when passed let's go!", func(t *testing.T) {
		got := katas.ShoutString("let's go!")
		want := "LET'S GO!"
		assert.String(t, got, want)
	})
}

func TestJoinStrings(t *testing.T) {
	t.Run("Joins strings a and b to return ab", func(t *testing.T) {
		got := katas.JoinStrings("a", "b")
		want := "ab"
		assert.String(t, got, want)
	})
	t.Run("Joins strings go and lang to return golang", func(t *testing.T) {
		got := katas.JoinStrings("go", "lang")
		want := "golang"
		assert.String(t, got, want)
	})
}

func TestRaiseToThePower(t *testing.T) {
	t.Run("Raises 9 to the power of 2 returning 81", func(t *testing.T) {
		got := katas.RaiseToThePower(9, 2)
		want := 81
		assert.Int(t, got, want)
	})
	t.Run("Raises 12 to the power of 0 returning 1", func(t *testing.T) {
		got := katas.RaiseToThePower(12, 0)
		want := 1
		assert.Int(t, got, want)
	})
	t.Run("Raises 10 to the power of 1 returning 10", func(t *testing.T) {
		got := katas.RaiseToThePower(10, 1)
		want := 10
		assert.Int(t, got, want)
	})
}

func TestIsBiggerThan42(t *testing.T) {
	t.Run("Returns true when passed 99 as it is bigger than 42", func(t *testing.T) {
		got := katas.IsBiggerThan42(99)
		want := true
		assert.Bool(t, got, want)
	})
	t.Run("Returns false when passed 41 as it is smaller than 42", func(t *testing.T) {
		got := katas.IsBiggerThan42(41)
		want := false
		assert.Bool(t, got, want)
	})
	t.Run("Returns false when passed 42 as it is the same as 42", func(t *testing.T) {
		got := katas.IsBiggerThan42(42)
		want := false
		assert.Bool(t, got, want)
	})
}

func TestCelsiusToFahrenheit(t *testing.T) {
	t.Run("takes 12.1 degrees Celsius and returns the equivalent temperature in degrees Fahrenheit, 53.78", func(t *testing.T) {
		got := katas.CelsiusToFahrenheit(12.1)
		want := 53.78
		assert.Float(t, got, want)
	})
	t.Run("takes 31.0 degrees Celsius and returns the equivalent temperature in degrees Fahrenheit, 87.8", func(t *testing.T) {
		got := katas.CelsiusToFahrenheit(31.0)
		want := 87.8
		assert.Float(t, got, want)
	})
}

func TestIsMillenial(t *testing.T) {
	t.Run("returns false when passed the year 1980 as it is not a year in which a Millenial is born", func(t *testing.T) {
		got := katas.IsMillenial(1980)
		want := false
		assert.Bool(t, got, want)
	})
	t.Run("returns true when passed the year 1992 as it is a year in which a Millenial is born", func(t *testing.T) {
		got := katas.IsMillenial(1992)
		want := true
		assert.Bool(t, got, want)
	})
}

func TestReadTrafficLight(t *testing.T) {
	t.Run("returns the string STOP! when passed the string red", func(t *testing.T) {
		got := katas.ReadTrafficLight("red")
		want := "STOP!"
		assert.String(t, got, want)
	})

	t.Run("returns the string GET READY... when passed the string amber", func(t *testing.T) {
		got := katas.ReadTrafficLight("amber")
		want := "GET READY..."
		assert.String(t, got, want)
	})

	t.Run("returns the string GO! when passed the string green", func(t *testing.T) {
		got := katas.ReadTrafficLight("green")
		want := "GO!"
		assert.String(t, got, want)
	})

	t.Run("is case insensitive", func(t *testing.T) {
		got := katas.ReadTrafficLight("AMBER")
		want := "GET READY..."
		assert.String(t, got, want)
	})
}

func TestPassOrFail(t *testing.T) {
	t.Run("returns string Fail if student scores less than 60% on exam", func(t *testing.T) {
		got := katas.PassOrFail(10, 20)
		want := "Fail"
		assert.String(t, got, want)
	})
	t.Run("returns string Pass if student scores 60% or more on exam", func(t *testing.T) {
		got := katas.PassOrFail(60, 100)
		want := "Pass"
		assert.String(t, got, want)
	})
	t.Run("returns string Top marks! if student scores 100% on exam", func(t *testing.T) {
		got := katas.PassOrFail(100, 100)
		want := "Top marks!"
		assert.String(t, got, want)
	})
}

func TestTitleCaseString(t *testing.T) {
	t.Run("Capitalizes the first letter of a single word", func(t *testing.T) {
		got := katas.TitleCaseString("gopher")
		want := "Gopher"
		assert.String(t, got, want)
	})
	t.Run("Capitalizes the first letter of each word in a sentence", func(t *testing.T) {
		got := katas.TitleCaseString("thunderbirds are go")
		want := "Thunderbirds Are Go"
		assert.String(t, got, want)
	})
}

func TestSumEvensTo(t *testing.T) {
	t.Run("returns 0 when passed 1", func(t *testing.T) {
		got := katas.SumEvensTo(1)
		want := 0
		assert.Int(t, got, want)
	})
	t.Run("returns 12 when passed 6", func(t *testing.T) {
		got := katas.SumEvensTo(6)
		want := 12
		assert.Int(t, got, want)
	})
	t.Run("returns 56 when passed 14", func(t *testing.T) {
		got := katas.SumEvensTo(14)
		want := 56
		assert.Int(t, got, want)
	})
}

func TestCheckInfinitive(t *testing.T) {
	t.Run("returns false if word is not an infinitive french verb", func(t *testing.T) {
		t.Run("est", func(t *testing.T) {
			got := katas.CheckInfinitive("est")
			want := false
			assert.Bool(t, got, want)
		})
		t.Run("suis", func(t *testing.T) {
			got := katas.CheckInfinitive("suis")
			want := false
			assert.Bool(t, got, want)
		})
		t.Run("allez", func(t *testing.T) {
			got := katas.CheckInfinitive("allez")
			want := false
			assert.Bool(t, got, want)
		})
	})
	t.Run("returns true if word is an infinitive french verb", func(t *testing.T) {
		t.Run("manger", func(t *testing.T) {
			got := katas.CheckInfinitive("manger")
			want := true
			assert.Bool(t, got, want)
		})
		t.Run("être", func(t *testing.T) {
			got := katas.CheckInfinitive("être")
			want := true
			assert.Bool(t, got, want)
		})
		t.Run("aller", func(t *testing.T) {
			got := katas.CheckInfinitive("aller")
			want := true
			assert.Bool(t, got, want)
		})
		t.Run("dormir", func(t *testing.T) {
			got := katas.CheckInfinitive("dormir")
			want := true
			assert.Bool(t, got, want)
		})
	})
}

func TestHugsAndKisses(t *testing.T) {
	t.Run("passes in 1 and returns x", func(t *testing.T) {
		got := katas.HugsAndKisses(1)
		want := "x"
		assert.String(t, got, want)
	})
	t.Run("passes in 2 and returns xo", func(t *testing.T) {
		got := katas.HugsAndKisses(2)
		want := "xo"
		assert.String(t, got, want)
	})
	t.Run("passes in a number and returns the appropriate amount of hugs and kisses", func(t *testing.T) {
		got := katas.HugsAndKisses(5)
		want := "xoxox"
		assert.String(t, got, want)
	})
}

func TestFizzBuzz(t *testing.T) {
	t.Run("returns ... when passed a number which is not divisible by 3 nor 5", func(t *testing.T) {
		got := katas.FizzBuzz(1)
		want := "..."
		assert.String(t, got, want)
	})
	t.Run("returns fizz! when passed a number which is divisible by 3 ", func(t *testing.T) {
		got := katas.FizzBuzz(6)
		want := "fizz!"
		assert.String(t, got, want)
	})
	t.Run("returns buzz! when passed a number which is divisible by 5 ", func(t *testing.T) {
		got := katas.FizzBuzz(25)
		want := "buzz!"
		assert.String(t, got, want)
	})
	t.Run("returns fizzbuzz! when passed a number which is divisible by 5 and 3", func(t *testing.T) {
		got := katas.FizzBuzz(30)
		want := "fizzbuzz!"
		assert.String(t, got, want)
	})
}

func CountTheVowels(t *testing.T) {
	t.Run("returns 1 when passed a string with one vowel present", func(t *testing.T) {

	})
}
