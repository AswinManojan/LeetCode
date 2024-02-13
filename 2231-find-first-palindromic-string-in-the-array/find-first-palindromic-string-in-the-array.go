func firstPalindrome(words []string) string {
    for i:=0;i<len(words);i++{
        flag:=false
        for j:=0;j<=(len(words[i])/2);j++{
            if words[i][j]!=words[i][len(words[i])-j-1]{
                flag=true
                break
            }
        }
        if flag==false{
            return words[i]
        }
    }
    return ""
}