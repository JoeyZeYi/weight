package weight

import (
	"fmt"
	"testing"
)

type Gift struct {
	Id     uint32 //礼物ID
	Weight int    //礼物权重值
	Num    int    //中奖数量
}

//实现权重接口的两个方法
func (gift *Gift) GetId() uint32 {
	return gift.Id
}
func (gift *Gift) GetWeight() int {
	return gift.Weight
}
func (gift *Gift) GetNum() int {
	return gift.Num
}

func TestNewWeighted(t *testing.T) {
	//实例化三个礼物
	gift1 := new(Gift)
	gift1.Id = 1
	gift1.Weight = 5
	gift1.Num = 1
	gift2 := new(Gift)
	gift2.Id = 2
	gift2.Weight = 3
	gift2.Num = 1
	gift3 := new(Gift)
	gift3.Id = 3
	gift3.Weight = 2
	gift3.Num = 1
	//将三个礼物放到接口切片里
	servers := make([]Weighted, 0)
	servers = append(servers, gift1, gift2, gift3)
	//循环18次  会有10次ID为1的礼物  5次ID为2的礼物  3次ID为3的礼物
	//实例化礼物的权重池
	load := NewPool(servers)

	for _, v := range load.Training {
		fmt.Println("抽奖之前", v)
	}

	for i := 1; i <= 10; i++ {
		weighted := load.Draw()
		fmt.Println(weighted)

	}
	for _, v := range load.Training {
		fmt.Println("抽奖之后", v)
	}

}
