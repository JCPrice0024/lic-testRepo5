package main

import (
	"fmt"

	lictestrepo2 "github.com/JCPrice0024/lic-testRepo2"
	lictestrepo3 "github.com/JCPrice0024/lic-testRepo3"
	lictestrepo4 "github.com/JCPrice0024/lic-testRepo4"
)

func main() {

	fmt.Println(lictestrepo2.WelcomeUser("JCPrice0024"))
	fmt.Println(lictestrepo3.Subtraction(7127, 123))
	fmt.Println(lictestrepo4.Addition(7127, 123))
}
