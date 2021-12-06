func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a        
    }
    var sb strings.Builder
    var d byte
    for i, j := len(a) - 1, len(b) - 1; i >= 0; {
        var c byte
        if j >= 0 {
            c, d = helper(a[i], b[j], d)
        } else {
            c, d = helper(a[i], '0', d)
        }
        fmt.Println(string(c + '0'))
        sb.WriteString(string(c + '0'))   
        i--
        j--
    } 
    if d == 1 {
        sb.WriteString("1")
    }
    return reverseString(sb.String())
}

func helper(a byte, b byte, d byte) (byte, byte) {
	ans := a - '0' + b - '0' + d
	return ans % 2, ans / 2
}

func reverseString(s string) string {
    var sb strings.Builder
    for i := len(s) - 1; i >= 0; i-- {
        sb.WriteString(string(s[i]))
    }    
    return sb.String()
}
