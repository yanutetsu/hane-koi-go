package main

import (
	"fmt"
	"math/rand"
	"time"
)

var giveFoodCount int
var giveTrainCount int

type Koiking struct {
	Name  string
	Level int
	CP    int
}

func (k *Koiking) Status() {
	fmt.Printf("今の'%s'のステータス %+v\n", k.Name, k)
}

func (k *Koiking) Bite(feed int) {
	fmt.Printf("CPが %d 上がった\n", feed)
	k.CP += feed
}

func (k *Koiking) Train(trainingEffect int) {
	fmt.Printf("とっくんで、%sのはねる力(CP)が強くなった！\n", k.Name)
	fmt.Printf("CPが %d あがった！\n", trainingEffect)
	k.CP += trainingEffect
}

func (k *Koiking) Retire() {
}

type You struct {
	Level    int
	Coin     int
	Daiamond int
	Koiking  Koiking
}

func (y *You) Status() {
	fmt.Printf("今のあなたのすてーたす %+v\n", y)
}

func (y *You) Fish(fishingRod string) Koiking {
	fmt.Printf("%sでコイキングをつります\n", fishingRod)
	fmt.Println("コイキングがつれました")
	fmt.Println("個性はまぁまぁです")
	return Koiking{
		Name:  "コイ太郎",
		Level: 1,
		CP:    1,
	}
}

func (y *You) GiveFood() int {
	giveFoodCount += 1
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(25)
}

func (y *You) GiveTrain() int {
	giveTrainCount += 1
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(75)
}

func (y *You) Retire(k Koiking) {
	k.Retire()
	fmt.Println("最終スコア")
	fmt.Printf("  CP %d\n", k.CP)
	fmt.Printf("  たべものを食べた回数 %d\n", giveFoodCount)
	fmt.Printf("  とっくんした回数 %d\n", giveTrainCount)
	fmt.Printf("  リーグで勝った回数 1\n")
}

type League struct {
	RivalKoiking Koiking
}

func (l *League) Battle(yourKoiking Koiking) bool {
	fmt.Println("バトル準備完了！スタート！")
	fmt.Println("はねろ！コイキング！")
	if yourKoiking.CP < l.RivalKoiking.CP {
		return false
	} else {
		return true
	}
}

func judge(result bool) {
	if result {
		fmt.Println("ブリーダーに勝利した！")
	} else {
		fmt.Println("負けてしまった...")
	}
}

func main() {
	fmt.Println("はねろ！コイキング")

	you := You{
		Level:    1,
		Coin:     0,
		Daiamond: 0,
	}
	you.Status()

	koiking := you.Fish("ボロつりざお")
	koiking.Status()

	koiking.Bite(you.GiveFood())
	koiking.Status()

	for koiking.CP < 100 {
		koiking.Bite(you.GiveFood())
	}
	koiking.Status()

	koiking.Train(you.GiveTrain())
	koiking.Status()

	for koiking.CP < 300 {
		koiking.Train(you.GiveTrain())
	}
	koiking.Status()

	league := League{
		RivalKoiking: Koiking{
			Name:  "コイッチョ",
			Level: 20,
			CP:    550,
		},
	}

	result := league.Battle(koiking)
	judge(result)

	for koiking.CP < 600 {
		koiking.Bite(you.GiveFood())
		koiking.Train(you.GiveTrain())
	}

	fmt.Println("最大レベルまで育ったよ！")
	fmt.Println("最後のリーグへ挑戦しよう！")

	result = league.Battle(koiking)
	judge(result)

	fmt.Println("最後の試合、おつかれさま！")
	fmt.Printf("%sは引退するよ!", koiking.Name)

	you.Retire(koiking)
}
