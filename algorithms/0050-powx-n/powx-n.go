package powx_n

func myPow(x float64, n int) float64 {
    if n < 0 {
        x, n = 1/x, -n
    }
    return pow(x, n)
}

func pow(x float64, n int) float64 {
    if n <= 0 {
        return 1
    }
    temp, i := x, 1
    for 2*i <= n {
        x, i = x*x, 2*i
    }
    return x * pow(temp, n-i)
}
