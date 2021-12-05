func plusOne(digits []int) []int {
    flag := false
    for _, v := range digits {
        if v != 9 {
            flag = true
            break
        }
    }
    var ans []int
    if flag {
        ans = make([]int, len(digits))
    } else {
        ans = make([]int, len(digits) + 1)
    }
    
    flag = true
    b := 1
    for i := len(digits) - 1; i >= 0; i-- {
        if !flag {
            ans[i] = digits[i]
            continue
        }
        
        if digits[i] + b <= 9 {
            ans[i] = digits[i] + b
            flag = !flag
        } else {
            ans[i] = 0
        }
    }
    if flag {
        ans[0] = 1
    }
    return ans
}
