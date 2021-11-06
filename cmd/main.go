package main

//1-1(hello,world!)

// func main() { 
// 	println("Hello, Go!")
// }


//----------------------------------------------------------------->


//1-2 (converter)

// func main() {
// 	amount := 100.0
// 	sellRate:=0.1470
	
// 	result := amount/sellRate
// 	println(result)
// }

//-------------------------------------------------------------------->


//2-1 (deposit)

// func main(){
// 	amount:=999_999_00
// 	minPersent:=4
// 	maxPersent:=6
// 	minIncome:=amount/100*minPersent
// 	maxIncome:=amount/100*maxPersent
// 	println(minIncome)
// 	println(maxIncome)
// }

//----------------------------------------------------------------------->


//2-2 (megafon)

// func main() {
// 	amount:=666
// 	welcomeBonus:=5
// 	point:=amount/5
// 	bonus:=welcomeBonus+point
// 	println(bonus)
// }

//----------------------------------------------------------------------->


//3-1(welcome)

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
	
// )

// func main() {
// 	flag.Parse()
// 	name:=flag.Arg(0)
// 	code:=flag.Arg(1)
// 	template1:="Добро пожаловать, {username}!"
// 	message1:=strings.ReplaceAll(template1,"{username}", name)
// 	template2:="Ваш код доступа: {code}."
// 	message2:=strings.ReplaceAll(template2,"{code}", code)
// 	fmt.Println(message1,message2)
// }


//------------------------------------------------------------------------------>


//3-2(phone)

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
	
// )

// func main() {
// 	flag.Parse()
// 	number:=flag.Arg(0) 
// 	template :=strings.ReplaceAll(number, ")","")
// 	template =strings.ReplaceAll(template, "(", "" )
// 	template = strings.ReplaceAll(template, "-", "" )
// 	template =strings.ReplaceAll(template, " ", "" )
// 	fmt.Println(template)
// }


//------------------------------------------------------------------>

//4-1(mobi)

// import (
// 	"fmt"
//  	"mobi/pkg/commission"   //+ pkg/commission (bonus.go)
// )
// func main() {
	
// 	result := commission.Calculate(9_000_00)
// 	fmt.Println(result)
// }


//------------------------------------------------------------------>

//4-2 (megafon ) 

// import (
// "fmt"
// "megafon/pkg/billing"  //+pkg/billing (billing.go)
// )
// func main() {
// 	result:=billing.Calculate1000(100)
// 	fmt.Println(result)

// 	result1:=billing.Calculate2000(3000)
// 	fmt.Println(result1)

// 	result2:=billing.Calculate3000(4000)
// 	fmt.Println(result2)

// 	result3:=billing.Calculate5000(4000)
// 	fmt.Println(result3)

// 	result4:=billing.Calculate10000(10000)
// 	fmt.Println(result4)
// }


//--------------------------------------------------------------------->

//5-1(bank)

// import (
// 	"bank/pkg/bank/deposit"
// 	"fmt"
// )

// func main() {
// 	{
// 		fmt.Println(deposit.Calculate(0, "TJS"))
// 	}
// 	{
// 		fmt.Println(deposit.Calculate(0, "USD"))
// 	}
// 	{
// 		fmt.Println(deposit.Calculate(500_000_00, "TJS"))
// 	}
// 	{	
// 		fmt.Println( deposit.Calculate(500_000_00, "USD"))
// 	}
// 	{
// 		fmt.Println(deposit.Calculate(1_000_000_00, "TJS"))
// 	}
// 	{
// 	fmt.Println(deposit.Calculate(1_000_000_00, "USD"))
// 	}
// }


//------------------------------------------------------------------------------>

//5-2(bank)
// package main

// import ( "fmt"
// "bank/pkg/bank/transfer")

// func main() {
// 	fmt.Println(transfer.Total(0))
// 	fmt.Println(transfer.Total(5_000_00))
// 	fmt.Println(transfer.Total(10_000_00))
// }


//----------------------------------------------------------------------->

//6-1(bank)
// package bank

// import (
// 	"bank/pkg/bank/types"
// 	"fmt"
// )
// func main() {
// 	card := types.Card{Balance : 0 ,ID :1000, Active: true, }
// 	fmt.Printf("%+v",card)
// }