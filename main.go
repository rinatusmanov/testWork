package main

import (
	"fmt"
	"sort"
)

/*
2.
Реализовать функцию, которая принимает текущие координаты шахматного коня,
расположенного на пустой доске и возвращает список возможных ходов в алгебраической нотации
*/

func convert(a, b byte) string {
	return string(a+64) + string(b+48)
}

// Func2 - задача 2
func Func2(a string) (res []string) {
	temp := []byte(a)
	res = []string{}
	if (temp[0] < 65 || temp[0] > 72) && (temp[0] < 97 || temp[0] > 104) {
		return
	} else if temp[0] > 72 {
		temp[0] -= 32
	}
	temp[0] -= 64
	if temp[1] < 49 || temp[1] > 56 {
		return
	}
	temp[1] -= 48
	x := temp[0]
	y := temp[1]
	if x > 1 {
		if y > 2 {
			res = append(res, convert(x-1, y-2))
		}
		if y < 7 {
			res = append(res, convert(x-1, y+2))
		}
	}
	if x > 2 {
		if y > 1 {
			res = append(res, convert(x-2, y-1))
		}
		if y < 8 {
			res = append(res, convert(x-2, y+1))
		}
	}
	if x < 8 {
		if y > 2 {
			res = append(res, convert(x+1, y-2))
		}
		if y < 7 {
			res = append(res, convert(x+1, y+2))
		}
	}
	if x < 7 {
		if y > 1 {
			res = append(res, convert(x+2, y-1))
		}
		if y < 8 {
			res = append(res, convert(x+2, y+1))
		}
	}
	return
}

/*
3.
Реализовать объект "игральные карты" и на нём реализовать сортировку любым методом, написать тесты

(дополнение: Колода 36 карт + 2 джокера, джокеры всегда старше любой из мастей, старшинство
	масти всегда сильнее старшинства карты, козырной может быть только одна масть (а может и не быть)
	и её можно передать на входе, порядок старшинства значений карт тоже можно передать на входе.
	Передать на входе означает что это параметры метода сортировки.)
*/
type Card struct {
	Joker  bool
	Amount int
	Mast   int
}
type Koloda struct {
	Cards []Card
	Trump int
}

func (k Koloda) Len() int      { return len(k.Cards) }
func (k Koloda) Swap(i, j int) { k.Cards[i], k.Cards[j] = k.Cards[j], k.Cards[i] }
func (k Koloda) Less(i, j int) bool {
	switch true {
	default:
		return k.Cards[i].Amount < k.Cards[j].Amount
	case k.Cards[i].Joker:
		return true
	case k.Cards[j].Joker:
		return false
	case k.Cards[i].Mast == k.Trump:
		if (k.Cards[j].Mast == k.Trump) && (k.Cards[j].Amount > k.Cards[i].Amount) {
			return false
		}
		return true
	case k.Cards[j].Mast == k.Trump:
		return false
	}
}

func main() {
	kk := Koloda{[]Card{{false, 10, 2}, {false, 3, 3}, {false, 5, 0}, {false, 1, 2}}, 2}
	fmt.Println(kk)
	sort.Sort(kk)
	fmt.Println(kk)
	fmt.Println(Func2("D4"))
	fmt.Println(Func2("d4"))
	fmt.Println(Func2("a4"))
	fmt.Println(Func2("a1"))
}
