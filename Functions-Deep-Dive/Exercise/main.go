package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var taxRates = [3]int {0,10,20}
	
	prices,err:=getPrices();
	if err!=nil{
		fmt.Println(err)
		return;
	}
	pricePointers:=make([]*int,0,len(prices));

	for index,_ := range prices{
		pricePointers = append(pricePointers, &prices[index])
	}
	var result = map[int][]int{};

	for index,val := range taxRates{
		result[val]=taxPrices(&taxRates[index],pricePointers...)
	}

	fmt.Print(result)
}


func getPrices() ([]int,error) {
	bytes,err:=os.ReadFile("prices.txt");
	if err!=nil{
		return nil,err;
	}

	stringValue:=string(bytes);

	var prices = []string{};

	prices=strings.Split(stringValue, ",")

	integerPrices :=[]int{};
	for _,val:=range prices{
		value,err := strconv.Atoi(val);
		if err!=nil{
			return nil,err
		}
		integerPrices = append(integerPrices,value)

	}
	return integerPrices,nil
}

func taxPrices(taxRate *int, prices ...*int) []int {
	result:=[]int{}
	for _,val := range prices {
		taxAmount:=(*taxRate*(*val))/100
		taxedPrice:=(*val)+taxAmount;

		result=append(result, taxedPrice)
	}
	return result; 
}