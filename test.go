package main

func intToRoman(num int) string {
	switch num {
		case 4: return "IV"
		case 9: return "IX"
		case 40: return "XL"
		case 90: return "XC"
		case 400: return "CD"
		case 900: return "CM"
	}

	var RomanNum string
	if num >= 1000 {
		num/1000
	}
}

func main() {



}
