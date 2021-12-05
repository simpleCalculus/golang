func lengthOfLastWord(s string) int {
    ln := len(s) - 1
    for ln >= 0 && s[ln] == ' ' {
        ln--
    }
    ans := 0
    for ln >= 0 && s[ln] != ' ' {
        ans++
        ln--
    }
    return ans
}
