package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)
const accountBalFile = "balance.txt"



func main(){

	var choice int
	var accountBal,err = fileops.GetFloatFromfile(accountBalFile)
	if err != nil {
		fmt.Println("Error")
		// return
		panic("cant continue")
	}


fmt.Println("welcome to Frankie Financial Services")
fmt.Println("please reach us on: ", randomdata.PhoneNumber())

	// for i:=0; 1<2; i++ {
	for {
			
presentOptions()

	fmt.Println("Your choice: ")
	fmt.Scan(&choice)


	switch choice {
	case 1:
		fmt.Println("Your balance is: ", accountBal)
	case 2:
		fmt.Println("How much do you want to deposit ? ", )

			var depositAmount float64
			fmt.Scan(&depositAmount)
			
			if depositAmount<=0{
				fmt.Println("please select one of the choices")
				continue
			}else {
				
			}
			accountBal += depositAmount
			fmt.Println(accountBal)
			fileops.WriteFloatToFile(accountBal,accountBalFile)
	case 3:
					fmt.Println("How much do you want to withdraw ? ", )

			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <=0 {
				fmt.Println("please select one of the choices")
				continue
			}else {

			}

			if withdrawalAmount > accountBal {
			fmt.Println("Invalid Amount")
			return
			}

			accountBal -= + withdrawalAmount
			fmt.Println(accountBal)
			fileops.WriteFloatToFile(accountBal,accountBalFile)

	default:
			fmt.Println("Thank you for usint frankie financial services")
			return
	}



	// if choice == 1 {
	// 	fmt.Println("Your balance is: ", accountBal)
	// }else if choice ==2 {
	// 	fmt.Println("How much do you want to deposit ? ", )

	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)
			
	// 		if depositAmount<=0{
	// 			fmt.Println("please select one of the choices")
	// 			continue
	// 		}else {
				
	// 		}
	// 		accountBal += depositAmount
	// 		fmt.Println(accountBal)
	// 	}else if choice == 3 {
			
	// 		fmt.Println("How much do you want to withdraw ? ", )

	// 		var withdrawalAmount float64
	// 		fmt.Scan(&withdrawalAmount)

	// 		if withdrawalAmount <=0 {
	// 			fmt.Println("please select one of the choices")
	// 			continue
	// 		}else {

	// 		}

	// 		if withdrawalAmount > accountBal {
	// 		fmt.Println("Invalid Amount")
	// 		return
	// 		}


	// 		accountBal -= + withdrawalAmount
	// 		fmt.Println(accountBal)
	// } else {
	// 	fmt.Println("Thank you for usint frankie financial services")
	// 	break
	// }

	}

	


}


