package main

import (
	"fmt"
)

type hero interface {
	getName()
	say()
	getSkill()
} //选英雄接口
type Attack interface {
	attack() int
	getHit(x int) bool
} //攻击接口

type aly struct {
	name  string
	HP    int
	skill map[string]int
} //队友结构体d
type enemy struct {
	name  string
	HP    int
	skill map[string]int
} //敌人结构体

func (a *aly) attack() int {
	var m string
	for {
		_, err := fmt.Scanln(&m)
		if err != nil {
			return 0
		} else if a.skill[m] > 0 {
			a.HP += a.skill[m]
			fmt.Println("Your HP", a.HP)
			return 0
		} else {
			return a.skill[m]
		}
	}
} //队友选择出手
func (a *aly) getHit(x int) bool {
	a.HP += x
	fmt.Println("Your HP:", a.HP)
	if a.HP <= 0 {
		fmt.Println("You lose")
		return true
	}
	return false
} //受到伤害
func (e *enemy) attack() int {
	var m string
	for {
		_, err := fmt.Scanln(&m)
		if err != nil {
			fmt.Println("Please choose correct skill")
			continue
		} else if e.skill[m] > 0 {
			e.HP += e.skill[m]
			fmt.Println("Enemy's HP", e.HP)
			return 0
		} else {
			return e.skill[m]
		}
	}
} //敌人选择出手
func (e *enemy) getHit(x int) bool {
	e.HP += x
	fmt.Println("Enemy's HP", e.HP)
	if e.HP <= 0 {
		fmt.Println("You win")
		return true
	}
	return false
} //受到伤害

func (a *aly) getName() {
	fmt.Print("请选择你的英雄：")
	fmt.Scanln(&a.name)
	a.HP = 500
} //选队友英雄，获得基础生命值
func (a *aly) say() {
	switch a.name {
	case "yasuo":
		fmt.Println("往事如风，常伴吾身")
	case "ruiwen":
		fmt.Println("前车之鉴，后车之师")
	case "gailun":
		fmt.Println("德玛西亚，永世长存")
	case "jianji":
		fmt.Println("所有人都尝试过，但只有我成功")
	}
} //英雄语音
func (a *aly) getSkill() {
	a.skill = map[string]int{}
	a.skill["q"] = -100
	a.skill["w"] = 60
	a.skill["e"] = -20
	a.skill["r"] = -170
} //获得技能

func (e *enemy) getName() {
	fmt.Print("请选择你的英雄：")
	fmt.Scanln(&e.name)
	e.HP = 500
} //选敌人英雄，获得基础生命值
func (e *enemy) say() {
	switch e.name {
	case "yasuo":
		fmt.Println("往事如风，常伴吾身")
	case "ruiwen":
		fmt.Println("前车之鉴，后车之师")
	case "gailun":
		fmt.Println("德玛西亚，永世长存")
	case "jianji":
		fmt.Println("所有人都尝试过，但只有我成功")
	}
} //英雄语音
func (e *enemy) getSkill() {
	e.skill = map[string]int{}
	e.skill["q"] = -70
	e.skill["w"] = 40
	e.skill["e"] = -110
	e.skill["r"] = -150
} //获得技能

func alyChoose(aly hero) {
	aly.getName()
	aly.say()
	aly.getSkill()
} //实现队友选英雄
func enemyChoose(enemy hero) {
	enemy.getName()
	enemy.say()
	enemy.getSkill()
} //实现敌方选英雄

func alyAttack(aly Attack, x int) int {
	if x != 0 {
		if aly.getHit(x) {
			return 0
		}
	}
	return aly.attack()
} //实现队友攻击
func enemyAttack(enemy Attack, x int) int {
	if x != 0 {
		if enemy.getHit(x) {
			return 0
		}
	}
	return enemy.attack()
} //实现敌人攻击

func main() {
	m := aly{}
	n := enemy{}
	alyChoose(&m)
	enemyChoose(&n)
	for i := 0; ; i++ {
		x := 0
		x = alyAttack(&m, x)
		if m.HP <= 0 {
			break
		}
		x = enemyAttack(&n, x)
		if n.HP <= 0 {
			break
		}
	}
}
