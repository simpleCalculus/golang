func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } 
    X := 0
    Y := x
    for x > 0 {
        X = X * 10 + (x % 10)
        x = x / 10
    }
    return X == Y 
}
