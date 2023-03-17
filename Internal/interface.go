package main
import ("fmt");

func main() {
	var x = 7

	switch x {
	case 10:
		fmt.Println("Perfect")
	case 8:
		fmt.Println("Good")
	default:
		fmt.Println("Not bad")
	}



	var car map[string]int

	car = map[string]int{}

	car["TYT"] = 10
	car["Honda"] = 20

	fmt.Print("TYT", car["TYT"])
	fmt.Print("Lambo", car["Honda"])
}