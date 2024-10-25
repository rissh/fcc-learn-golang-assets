package main

func fizzbuzz() string{
	// ?
	var res string
	for i := 1; i < 101; i++ {
		if i % 3 == 0{
			res += "Fizz"
		} else if i % 5 == 0{
			res += "Buzz"
		} else {
			res += string(float64(i))
		}
	}
	return res
}

// don't touch below this line

func main() {
	fizzbuzz()
}
