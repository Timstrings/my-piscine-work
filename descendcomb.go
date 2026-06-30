package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true

	for num1 := 99; num1 >= 0; num1-- {
		for num2 := num1 - 1; num2 >= 0; num2-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			z01.PrintRune(rune('0' + num1/10))
			z01.PrintRune(rune('0' + num1%10))

			z01.PrintRune(' ')

			z01.PrintRune(rune('0' + num2/10))
			z01.PrintRune(rune('0' + num2%10))
		}
	}
}
