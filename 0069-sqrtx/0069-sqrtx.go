func mySqrt(x int) int {
    l, r := 0, x
    res := 0
    for l <= r {
        m := (l+r) / 2
        if m*m > x {
            r = m -1 
        }else if m*m < x {
            l = m +1 
            res = m
        }else {
            return m
        }
    }
    return res
}