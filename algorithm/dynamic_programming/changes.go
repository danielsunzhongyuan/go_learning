package main

import "fmt"

/*
把100块钱换成 1，2，5的零钱，有多少种换法？
*/

func changesDP(smallChanges []int, money int) int {
	tmp := make([]int, money+1, money+1)
	fmt.Println("init tmp:", tmp)
	tmp[0] = 1
	for _, i := range smallChanges {
		for j := i; j <= money; j++ {
			tmp[j] += tmp[j-i]
		}
		fmt.Println(tmp)
	}
	return tmp[money]
}

func main() {
	fmt.Println(changesDP([]int{1, 2, 5, 10, 20, 50}, 100))
}

/*
def DP_Min(exlist,target):
    fan=[0]+[target]*target
    record=['']*(target+1)
    #存储最后一位
    lic=[]
    for i in exlist:
        for j in range(i,target+1):
            if fan[j-i]!=target:#如果等于target，说明之前没有数的结合可以达到这个数
                fan[j]=min(fan[j-i]+1,fan[j])
                #记录
                if fan[j-i]+1<=fan[j]:
                    record[j]=record[j-i]+'%d*'%i
                else:
                    record[j]=record[j]
        if record[-1]!='':
            lic.append(record[-1])
    if len(lic)==0:
        return '无解'
    else:
        return lic,fan[target]

#处理字符串
def Handle(exstr,num):
    for i in set(exstr):
        hu=i.split('*')
        hu=list(filter(lambda x:x,hu))#去除掉split留下的空格
        if len(hu)==num:
            shuoming=''
            for j in set(hu):
                shuoming+='%s个%s '%(hu.count(j),j)
            print(shuoming)
    return 'Done'

ggg=DP_Min([1,3,5],11)
print(Handle(ggg[0],ggg[1]))
#结果：最少为3。组合方法：2个5 1个1
*/
