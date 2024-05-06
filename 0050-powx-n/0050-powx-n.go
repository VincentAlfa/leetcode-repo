func myPow(x float64, n int) float64 {
	if x == 0{ return 0 } 
	if n == 0{ return 1 }
    
	tempN:= n
	absN := math.Abs(float64(n))
	result := myPow(x, int(absN)/2)

	if n % 2 == 0 {
		result = result * result
	}else{
		result = x * result * result
	}

	if tempN > 0 { return result }
	return 1/ result
}