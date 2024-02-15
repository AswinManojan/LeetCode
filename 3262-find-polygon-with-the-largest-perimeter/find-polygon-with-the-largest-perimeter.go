func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    if len(nums)<3{
        return -1
    }else if len(nums)==3{
        if nums[0]+nums[1]<=nums[2]{
            return -1
        }else{
            sum:=nums[0]+nums[1]+nums[2]
            return int64(sum)
        }
    }else{
        sum:=nums[0]+nums[1]+nums[2]
        i:=3
        j:=0
        flag:=false
        flag2:=false
        // flag3:=false
        sum2:=sum
        for ;i<len(nums);i++{
            if sum>nums[i] ||sum2>nums[i]{
                if sum2>nums[i]{
                    sum=sum2
                }
                sum+=nums[i]
                sum2=sum
                flag=true
            }else{
                sum2+=nums[i]
                if !flag2{
                    j=i
                    // flag3=true
                }
                flag2=true
            }
        }
        fmt.Println(j)
        if j==3 && flag2{
            if flag {
                return int64(sum)
            }else{
                if nums[0]+nums[1]<=nums[2]{
                    return -1
                }else{
                    return int64(sum)
                }
            }
        }
        return int64(sum)
    }
    return -1
}