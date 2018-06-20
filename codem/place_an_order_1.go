package main

import (
	"fmt"
	"sort"
)

type Good struct {
	Price     float64
	IsSpecial bool
}

type Goods []Good

func (g Goods) Less(i, j int) bool {
	return g[i].Price < g[i].Price
}

func (g Goods) Len() int {
	return len(g)
}

func (g Goods) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

type Offer struct {
	Exceed float64
	Cut    float64
}

type Offers []Offer

func (o Offers) Less(i, j int) bool {
	if o[i].Cut == o[j].Cut {
		return o[i].Exceed < o[j].Exceed
	}
	return o[i].Cut > o[j].Cut
}

func (o Offers) Len() int {
	return len(o)
}

func (o Offers) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func findClosestAndNotLessNumber(arr Goods, target float64) float64 {
	if target == 0.0 {
		return 0.0
	}
	N := len(arr)
	// 注意，如果target是真正的float，就不好使了！！！
	// 而且good的price也都得是 int，而不能是 float！！！
	// （因为是价格，所以最多到分，真是float的话，就乘以100，再算）
	W := int(target)
	V := make([][]float64, N+1, N+1)
	for i := 0; i < N+1; i++ {
		V[i] = make([]float64, W+1, W+1)
	}
	for item := 1; item <= N; item++ {
		for weight := 1; weight <= W; weight++ {
			if arr[item-1].Price <= float64(weight) {
				V[item][weight] = max(arr[item-1].Price+V[item-1][weight-int(arr[item-1].Price)], V[item-1][weight])
			} else {
				V[item][weight] = V[item-1][weight]
			}
		}
	}
	fmt.Println(V)
	return V[N][W]
}

func main() {
	n := 0
	m := 0
	discount := 0.80

	x, _ := fmt.Scan(&n, &m)
	if x == 0 {
		return
	}
	goods := make(Goods, 0, n)
	specialGoods := make(Goods, 0)
	offers := make(Offers, 0, m)

	originalTotalPrice := 0.0
	originalTotalPriceForNormal := 0.0
	originalTotalPriceForSpecial := 0.0
	for i := 0; i < n; i++ {
		price := 0.0
		isSpecialOffer := 0
		x, _ := fmt.Scan(&price, &isSpecialOffer)
		if x == 0 {
			return
		}
		goods = append(goods, Good{Price: price, IsSpecial: isSpecialOffer == 1})
		if isSpecialOffer != 1 {
			originalTotalPriceForNormal += price
		} else {
			originalTotalPriceForSpecial += price
			specialGoods = append(specialGoods, Good{Price: price, IsSpecial: true})
		}
		originalTotalPrice += price
	}

	for i := 0; i < m; i++ {
		exceed := 0.0
		cut := 0.0
		x, _ := fmt.Scan(&exceed, &cut)
		if x == 0 {
			return
		}
		offers = append(offers, Offer{Exceed: exceed, Cut: cut})
	}

	// sort special goods by price asc
	sort.Sort(specialGoods)
	// sort offers by cut desc, exceed asc
	sort.Sort(offers)

	// remove some bad offers, which have the same cut but higher exceed
	numberToRemove := 0
	i := 1
	for i < m-numberToRemove {
		if offers[i].Cut == offers[i-1].Cut {
			offers = append(offers[:i], offers[i+1:]...)
			numberToRemove++
		} else {
			i++
		}
	}

	// remove those offers that will never be used
	i = 0
	for i < m-numberToRemove {
		if offers[i].Exceed > originalTotalPrice {
			offers = append(offers[:i], offers[i+1:]...)
			numberToRemove++
		} else {
			i++
		}
	}

	ret := originalTotalPrice
	// 这里假设一定会使用某张券
	for _, offer := range offers {
		/*
			// 如果不需要可以特殊优惠的商品帮忙，剩下的只能参加满减活动的商品，就可以参加这一张券了，说明，后面的券都用不上了，
			// 因为后面的券的优惠力度都比这张要小，所以就可以 break 了
			if originalTotalPriceForNormal >= offer.Exceed {
				ret = min(ret, originalTotalPriceForSpecial*discount+originalTotalPriceForNormal-offer.Cut)
				break
			}

			// 下面是最复杂的情况，即需要能够特殊优惠的商品去凑单
			// 那么就需要找出 special goods 里，价格之和刚好能凑齐这张优惠券，或者只多一点点。
			// 因为满足条件的情况下，优惠的力度是固定的，那么就要让剩下参加自己本身特殊优惠的商品金额尽量大！
			// 现在问题转化为：从 special goods的list里，找到若干个（关键是这个数量不定）商品，使得它们的和最小，且它们的和大于等于 offer.Exceed - originalTotalPriceForNormal
			// 因为 n <= 10，所以可以用遍历的方式去做，即从 special goods里选1个，选2个，选3个，。。选10个，使得那个和最小，同时大于等于一个target
			//
			// 上面的遍历的思路不好，虽然可以做。现在把问题转换为背包问题：
			// 现在要找若干个商品，使得它们的和最大，且它们的和小于等于 originalTotalPrice - offer.Exceed
			// 背包问题：小于等于10个的商品，价值是 price，重量也是price，要求总重量小于等于 originalTotalPrice - offer.Exceed 的最大price和
			// 注意：因为这里的价格、优惠都是正整数，所以可以采用一个二维数组的方式去求解（http://www.importnew.com/13072.html），否则不行
			// 找到这个和 (X)之后，则需要支付：originalTotalPrice - X*(1-discount) - offer.Cut

			// 上面全错！不能一部分满减、一部分特价优惠。"满减优惠方式只有在所有物品都不选择特价优惠时才能使用，且最多只可以选择最多一款"。注意这里是所有！！！
			x := findClosestAndNotLessNumber(specialGoods, originalTotalPrice-offer.Exceed)
			ret = min(ret, originalTotalPrice-x*(1-discount)-offer.Cut)
		*/

		if originalTotalPrice >= offer.Exceed {
			ret = min(ret, originalTotalPrice-offer.Cut)
			break
		}
	}
	// 不使用优惠券，或者说所有的券都不合适，都被删掉了，也就是上面的 for 循环空了
	ret = min(ret, originalTotalPriceForSpecial*discount+originalTotalPriceForNormal)
	fmt.Printf("%.2f\n", ret)
}
