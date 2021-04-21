# Day 1 Katas

The following katas are designed to help familiarise you with Go's static typing, testing framework, and [standard library](https://golang.org/pkg/) (hints to certain packages are given, but these are not your only opportunities to make use of the tools at hand).

The katas are presented (roughly) in order of increasing difficulty, and are focussed around strings, numbers, and booleans. They can **all** be solved using the techniques covered in the lecture today (with no need to call upon complex data types!).

## Section A

### getLength(str string) string

- returns length of string
- can assume that this function will not accept non-English characters

```go
getLength("hello") // returns 5
getLength("I hear the drums echoing tonight") // returns 32
```

### modulo(n, d int) int

- takes a number and divisor and returns the remainder when dividing n by d

```go
modulo(6, 3) // returns 0
modulo(5, 2) // returns 1
```

### squareRoot(n int) float64

- returns square root of passed integer

```go
squareRoot(9) // returns 3.0
squareRoot(2) // returns 1.414...
```

### shoutString(str string) string

- returns capitalised string

```go
shoutString("let's go!") // returns LET'S GO!
```

### joinStrings(str1, str2 string) string

- takes two string arguments
- returns a single string of the two concatenated

```go
joinStrings("go", "lang") // returns "golang"
```

### raiseToThePower(n, p int) int

- return the result of raising m to the nth power

```go
raiseToThePower(9,2) // returns 81
raiseToThePower(12, 0) // returns 1
```

_HINT_: Whilst this function accepts integer arguments, you may want to use floats in your solution. Think carefully about how to make sure you're returning the correct integer.

### isBiggerThan42()

- returns true if the passed number is greater than 42
- otherwise returns false

```go
isBiggerThan42(41) // returns false
isBiggerThan42(42) // returns false
isBiggerThan42(99) // returns true
```

### celsiusToFahrenheit()

- takes a number representing temperature in degrees Celsius
- returns the equivalent temperature in degrees Fahrenheit

```go
celsiusToFahrenheit(12.1) // returns 53.78
celsiusToFahrenheit(31.0) // returns 87.8
```

_HINT:_ There is no "ideal" way of comparing floating point numbers. Consider utilising the `math` package to round your answers to a suitable degree of accuracy before comparing them.

### isMillenial()

- returns true if person's date of birth falls within the years 1981 and 1996 inclusively
- otherwise returns false

```go
isMillenial(1980) // returns false
isMillenial(1992) // returns true
isMillenial(1996) // returns true
```

## Section B

### readTrafficLight()

- check if the "traffic light" is red, green or amber and return a corresponding message
- it should be case insensitve

```go
titleCase("red") // returns "STOP!"
titleCase("amber") // returns "GET READY..."
titleCase("green") // returns "GO!"
titleCase("GREEN") // returns "GO!"
```

### passOrFail()

- calculate whether a student passes or fails an exam
- should take the number of correct answers and total number of questions as arguments
- returns "Top marks!" if they achieve 100%
- returns "Pass" if marks are 60% or higher
- otherwise return "Fail"

```go
passOrFail(10, 20) // returns "Fail"
passOrFail(60, 100) // returns "Pass"
passOrFail(100, 100) // returns "Top marks!"
```

### titleCaseString()

- capitalises the first character of each word of a string

```go
titleCase("gopher") // returns "Gopher"
titleCase("thunderbirds are go") // returns "Thunderbirds Are Go"
```

### sumEvensTo()

- sum all the even numbers up to and including the passed number

```go
sumEvensTo(1) // returns 0
sumEvensTo(14) // returns 56
sumEvensTo(6) // returns 12
```

### checkInfinitive()

- checks if a french word is an infinitive french verb
- Can assume a french infinitive verb is a word that ends with either "re", "ir" or "er"

```go
checkInfinitive("manger") // returns true
checkInfinitive("être") // returns true
checkInfinitive("aller") // returns true
checkInfinitive("dormir") // returns true

checkInfinitive("est") // returns false
checkInfinitive("suis") // returns false
checkInfinitive("allez") // returns false
```

### hugsAndKisses()

- takes a number of alternating Os and Xs to return

```go
hugsAndKisses(1) // returns "x"
hugsAndKisses(2) // returns "xo"
hugsAndKisses(5) // returns "xoxox"
```

### fizzbuzz()

- when passed a number this function should return a string
- "fizz!" if the number is divisible by 3
- "buzz!" if the number is divisible by 5
- "fizzbuzz!" if the number is divisible by 3 and 5
- "..." if the number is neither divisible by 3 nor 5

```go
fizzbuzz(1) // returns "..."
fizzbuzz(6) // returns "fizz!"
fizzbuzz(30) // returns "fizzbuzz!"
fizzbuzz(25) // returns "buzz!"
```

### countTheVowels()

- counts the number of vowels present in a string

```go
countTheVowels("go") // returns 1
countTheVowels("going") // returns 2
countTheVowels("going, going, gone!") // 6
```

## Section C

### getMiddle()

- returns for middle or middle two characters of a string
- you may assume there will be no spaces and no non-English characters

```go
countTheVowels("northcoders") // returns "c"
countTheVowels("golang") // returns "la"

// BONUS - non-English characters (HARD)
countTheVowels("être") // returns "tr"
```

_HINT_: The built-in `unicode/utf8` package can help us with strings where the `len()` function doesn't cut it.

### extractNumber()

- extracts out the digits found within string of characters
- the digits are always found within a pair of parentheses

```go
extractNumber("v(1)mdkupz") // returns 1
extractNumber("kajsdföû(38)kjsdkn") // returns 38
extractNumber("poqpewjasiournsanu(1729)po") // returns 1729
```

### isPythagoreanTriple()

- [Pythagoras' theorem](https://www.mathsisfun.com/pythagoras.html) states that the formula `a² + b² = c²` holds true for any right-angled triangle
- a "**Pythagorean triple**" is any 3 integers which satisfy that formula
- this function takes 3 numbers, representing the three sides of a triangle
- it returns `true` if the 3 numbers form a Pythagorean triple
- it returns `false` if they do not
- _NOTE:_ The numbers do not have to be given in the order they are written in the formula

```go
isPythagoreanTriple(3, 4, 5)    // returns true
isPythagoreanTriple(3, 4, 7)    // returns false
isPythagoreanTriple(13, 5, 12)  // returns true
```

### getCentury()

- returns the century of the given year

```go
getCentury(1) // returns "1st"
getCentury(1899) // returns "19th"
getCentury(2020) // returns "21st"
```

### sumOfDigitSquares()

- takes a number as its only argument
- return the sum of the squares of its individual digits

```go
sumOfDigitSquares(6)    // returns 6^2 = 36
sumOfDigitSquares(401)  // returns 4^2 + 0^2 + 1^2 = 17
sumOfDigitSquares(-128) // returns 1^2 + 2^2 + 8^2 = 69
```

_HINT:_ You might want to make use of the [`strconv` package](https://golang.org/pkg/strconv/) to convert between numbers and strings.

### pigLatin()

- takes a word and converts it into pig Latin
- to translate to pig Latin you must the take the first consonant (or consecutive consonants) of a word, move it to the end and append "ay"
- if the string begins with a vowel, just append "way"

```go
pigLatin("paper")   // returns "aperpay"
pigLatin("crane")   // returns "anecray"
pigLatin("origami") // returns "origamiway
```

### isJumpingNumber()

- a "**jumping number**" is one where all the adjacent digits differ only by one (see the examples below)
- this function takes a number as its only argument
- returns `true` if it is a jumping number and `false` if not

```go
isJumpingNumber(8)        // true - all single digits are jumping numbers
isJumpingNumber(34)       // true
isJumpingNumber(37)       // false
isJumpingNumber(5456787)  // true
isJumpingNumber(78901)    // false - 9 and 0 do NOT differ by one
```
