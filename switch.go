 package main

import ("fmt"
 		"time"
		)
func main(){


			// 	switch i:=1;i{

			// 	case 1:

			// 			fmt.Println("u r in case 1")
					
			// 	case 2:

			// 			fmt.Println("u r in case 2")
					

			// 	case 3:

			// 			fmt.Println("u r in case 3")
					

			// 		default :

			// 			fmt.Println("u r in default")
					
			// }

			switch time.Now().Weekday() {
			
			case time.Saturday, time.Sunday:
				fmt.Println("It's the weekend")
			
			default:
				fmt.Println("It's a weekday")
			}

			// t := time.Now()
			// fmt.Println(t.Hour())
			// switch t.Hour(){
			// case t.Hour():
			// 	fmt.Println("It's before noon")
			// default:
			// 	fmt.Println("It's after noon")
			// }

}