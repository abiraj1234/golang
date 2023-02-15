package main 

import "fmt"






func findMiddleIndex(nums []int) int {
  
    for ind,_:=range nums{
        res := getsum(nums,ind)
        if res[0]==res[1]{
            return ind
        }
    }
    return -1
}
func getsum(lst []int,key int) []int{
    sum1 := 0
    sum2 := 0
    for i:=0;i<key;i++{
        sum1+=lst[i]
    }
    for i:=key+1;i<len(lst);i++{
        sum2+=lst[i]
    }
    return []int {sum1,sum2}
}

func main (){ 
	fmt.Print(findMiddleIndex(ind))
	fmt.Scan(&findMiddleIndex)
}