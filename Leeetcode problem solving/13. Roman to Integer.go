func romanToInt(s string) int {
    ln := len(s)
    ans := 0
    for i := ln - 1; i >= 0; i-- {
        x := s[i]
        // V or X
        if x == 'V' {
            if i > 0 && s[i-1] == 'I' {
                ans = ans + 4
                i--
            } else {
                ans = ans + 5
            }
        } else if x == 'X' {
            if i > 0 && s[i-1] == 'I' {
                ans = ans + 9
                i--
            } else {
                ans = ans + 10
            }
        }
        // L or C
        if x == 'L' {
            if i > 0 && s[i-1] == 'X' {
                ans = ans + 40
                i--
            } else {
                ans = ans + 50
            }
        } else if x == 'C' {
            if i > 0 && s[i-1] == 'X' {
                ans = ans + 90
                i--
            } else {
                ans = ans + 100
            }
        }
        // D or M
        if x == 'D' {
            if i > 0 && s[i-1] == 'C' {
                ans = ans + 400
                i--
            } else {
                ans = ans + 500
            }
        } else if x =='M' {
            if i > 0 && s[i-1] == 'C' {
                ans = ans + 900
                i--
            } else {
                ans = ans + 1000
            }
        }
        // I
        if x == 'I' {
            ans++
        }
    }
    return ans
}
