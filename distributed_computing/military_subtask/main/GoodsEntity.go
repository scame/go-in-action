package main


type GoodsHolder struct {
	goods []int
}

func (goodsHolder *GoodsHolder) initGoods(goodsNumber int) {
	goodsHolder.goods = make([]int, goodsNumber)

	for i := 0; i < goodsNumber; i++ {
		goodsHolder.goods[i] = i
	}
}
