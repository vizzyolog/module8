package main

import (
	"app/elec"
	"fmt"
)

func main() {

	newIPhone := elec.NewApplePhone("iPhone")
	printCharacteristics(newIPhone)

	newAndroid := elec.NewAndroidPhone("Meizu", "M5", "Android v51")
	printCharacteristics(newAndroid)

	newRadio := elec.NewRadioPhone("Радиотехника", "42", 5)
	printCharacteristics(newRadio)

}

func printCharacteristics(phone elec.Phone) {

	fmt.Printf("Brand: %v  Model: %v Type: %v ", phone.Brand(), phone.Model(), phone.Type())

	smart, ok := phone.(elec.Smartphone)
	if ok {
		fmt.Printf("OS %s \n", smart.OS())
	}

	radio, ok := phone.(elec.StationPhone)
	if ok {
		fmt.Printf("buttons Amount %v\n", radio.ButtonsCount())
	}

	fmt.Println()

}
