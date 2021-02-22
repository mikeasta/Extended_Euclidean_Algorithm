package main

import (
	"fmt";
	"strconv";
	"math";
)


func settingUpSlices(number1 , number2 *int,
	remainders, factors, x, y *[]int) (err bool, msg string) {
	err    = false
	msg    = ""
	fmt.Printf("Enter first number: ")
	fmt.Scanf("%d", number1)
	fmt.Printf("Enter second number: ")
	fmt.Scanf("%d", number2)

	// * Checking data
	if (*number1 == 0 || *number2 == 0) {
		err    = true
		msg	   = "Invalid data: input need to be non-zero"
	}

	// * Setting up X vector
	if (*number1 < 0) {
		*x = append(*x, -1)
		*number1 *= (-1)
	} else {
		*x = append(*x, 1)
	}

	*x = append(*x, 0)

	// * Setting up Y vector
	*y = append(*y, 0)

	if (*number2 < 0) {
		*y = append(*y, -1)
		*number2 *= (-1)
	} else {
		*y = append(*y, 1)
	}

	// *Setting up remainders
	if (*number1 >= *number2) {
		*remainders = append(*remainders, *number1)
		*remainders = append(*remainders, *number2) 
	} else {
		*remainders = append(*remainders, *number1)
		*remainders = append(*remainders, *number2)
	}

	// * Setting up factors
	*factors = append(*factors, 0, 0)
	return
}


func euclideanAlgorithm(remainders, factors *[]int) {
	var processRemainder, processFactor int
	remaindersStatic := *remainders
	for i := 0;; i++ {
		processFactor     = int(math.Floor(float64(remaindersStatic[i]) / float64(remaindersStatic[i + 1])))
		processRemainder  = remaindersStatic[i] - remaindersStatic[i + 1] * processFactor
		*remainders 	  = append(*remainders, processRemainder)
		remaindersStatic  = append(remaindersStatic, processRemainder)
		*factors   		  = append(*factors, processFactor)
		if (processRemainder == 0) { break }
	}
}

func calculateRatios(factors []int, x, y *[]int) {
	var currentX, currentY int
	xStatic := *x
	yStatic := *y
	for i := 0; i < len(factors) - 2; i++ {
		currentX = xStatic[i] - factors[i + 2] * xStatic[i + 1]
		currentY = yStatic[i] - factors[i + 2] * yStatic[i + 1]
		xStatic  = append(xStatic, currentX)
		yStatic  = append(yStatic, currentY) 
		*x  	 = append(*x, currentX)
		*y		 = append(*y, currentY)

		fmt.Printf(strconv.Itoa(i))
	}	
}

func logData(remainders, factors, x, y []int) {
	// * Find the longest num
	var maxLength int
	
	fmt.Printf("\n")
	for _, value := range *remainders {
		fmt.Printf(strconv.Itoa(value) + " ")
	}

	fmt.Printf("\n")
	for _, value := range *factors {
		fmt.Printf(strconv.Itoa(value) + " ")
	}

	fmt.Printf("\n")
	for _, value := range *x {
		fmt.Printf(strconv.Itoa(value) + " ")
	}

	fmt.Printf("\n")
	for _, value := range *y {
		fmt.Printf(strconv.Itoa(value) + " ")
	}
}

func main() {
	var number1, number2 int
	var remainders, factors, x, y  []int

	// * Step 1. Initialize data
	err, msg := settingUpSlices(&number1, &number2, &remainders, &factors, &x, &y);

	// ! Validation for Step 1.
	if (err) {
		fmt.Printf(msg)
	}

	// * Step 2. Find last remainder (Euclidean Algorithm) 
	euclideanAlgorithm(&remainders, &factors)


	// * Step 3. Calculate ratios
	calculateRatios(factors, &x, &y)

	// * Step 4. Log the result 
	logData(remainders, factors, x, y)

}