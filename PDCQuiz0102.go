package pdcquiz0102

import (
	"fmt"
)

func RunMenu() {

	end := false

	for !end {

		fmt.Println(`Please select an option:
1 - Print Covid19 cases in Pakistan 
2 - Print Covid19 cases in SouthKorea
3 - Print Covid19 cases in France
4 - Print my personalized message to Coronavirus
0 - Exit`)

		var selected int
		fmt.Scanf("%d", &selected)

		if selected < 0 || selected > 4 {
			fmt.Println("Please enter a valid option..")
			continue
		}

		switch selected {
		case 1:
			fmt.Println("3/29/2020:	1,526 cases in Pakistan")
		case 2:
			fmt.Println("3/29/2020: 9,583 cases in South Korea")
		case 3:
			fmt.Println("3/29/2020: 37,575 cases in France")
		case 4:
			fmt.Println("	='(		")
		case 0:
			fmt.Println("Exiting menu")
			end = true
		}
	}

}
