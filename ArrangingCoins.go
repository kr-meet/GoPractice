func arrangeCoins(n int) int {
    
    s, e := 0, n    
    for s <= e {
            
        mid := s + (e - s) / 2; 
        coin := mid * (mid + 1) / 2;
        if coin > n {
            e = mid - 1
        } else if coin < n {
            s = mid + 1
        } else {
            return mid
        }
    }
    return e
}