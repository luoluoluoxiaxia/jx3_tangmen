package main

import (
	"math"
	"math/rand"
)

var maxWuShuang = 33.00

func main() {

}

func DamageCalculation() (DamageEffect int, Damage float64) {
	//命中
	hitProbability := 0.00
	//无双
	wuShuangProbability := 0.00
	//攻击 面板总攻击
	totalAttack := 0.00
	//基础攻击
	baseAttack := 0.00
	//武器伤害最低值
	minWeaponDamage := 0.00
	//武器伤害最高值
	maxWeaponDamage := 0.00
	//破防百分比
	increaseDamageCoefficient := 0.00
	//会心率
	criticalStrikeProbability := 0.00
	//会心效果
	criticalStrikeAddition := 0.00

	//秘籍和奇穴总加成
	otherBonusCoefficient := 0.00
	//技能加成
	skillAddition := 0.00
	//技能伤害系数
	SkillCoefficient := 0.00
	//技能武器伤害系数
	skillWeaponCoefficient := 0.00
	//技能最小数值伤害
	minSkillDamage := 0.00
	//技能最大数值伤害
	maxSkillDamage := 0.00

	//判断是否命中
	if IfHit(hitProbability) {
		return 0, 0
	}

	basheHurt := ((totalAttack+baseAttack*skillAddition+baseAttack*skillAddition+baseAttack*skillAddition)*SkillCoefficient +
		(minWeaponDamage+maxWeaponDamage)/2*skillWeaponCoefficient + (minSkillDamage+maxSkillDamage)/2) *
		(1 + increaseDamageCoefficient + increaseDamageCoefficient*skillAddition + increaseDamageCoefficient*skillAddition + increaseDamageCoefficient*skillAddition) *
		(1 + otherBonusCoefficient)

	//*  (是否会心，若会心，则计算会效)
	if IfCriticalStrike(criticalStrikeProbability) {
		basheHurt = basheHurt * criticalStrikeAddition
		DamageEffect = 2
	}
	//是否识破
	if ifSeeThrough(wuShuangProbability) {
		DamageEffect = 3
		basheHurt = basheHurt * 0.25
	}

	return DamageEffect, basheHurt
}

//是否会心
func IfCriticalStrike(criticalStrikeProbability float64) bool {
	probability := rand.Intn(10000)
	criticalStrikeProbability = criticalStrikeProbability * 100
	i := int(math.Floor(criticalStrikeProbability))
	if probability <= i {
		return true
	}
	return false
}

//是否命中
func IfHit(hitProbability float64) bool {
	probability := rand.Intn(100)
	hitProbability = hitProbability - 10
	i := int(math.Floor(hitProbability))
	if probability <= i {
		return true
	}
	return false
}

//是否识破
func ifSeeThrough(wushuangProbability float64) bool {
	if wushuangProbability >= maxWuShuang {
		return false
	}
	probability := rand.Intn(100)
	wushuangProbability = (1 - (wushuangProbability / maxWuShuang)) * 100
	i := int(math.Floor(wushuangProbability))
	if probability <= i {
		return true
	}
	return false
}
