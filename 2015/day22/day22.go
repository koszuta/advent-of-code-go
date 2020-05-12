package main

import (
	"fmt"
	"time"
)

type player struct {
	hp, mana, armor int
	activeSpells    map[string]int
}

type enemy struct {
	hp, damage int
}

type spell struct {
	cost, duration int
	// playerHPInstantEffect,
	// enemyHPInstantEffect,
	// playerManaEffect,
	// playerArmorInstantEffect int
	instantEffect func(player, enemy)
	turnEffect    func(player, enemy)
	afterEffect   func(player, enemy)
}

var allSpells map[string]spell
var noop func(player, enemy) = func(player, enemy) {}

func init() {
	allSpells = make(map[string]spell, 5)
	allSpells["Magic Missile"] = spell{
		53,
		-1,
		func(player player, enemy enemy) {
			enemy.hp -= 4
		},
		noop,
		noop,
	}
	allSpells["Drain"] = spell{
		73,
		-1,
		func(player player, enemy enemy) {
			player.hp += 2
			enemy.hp -= 2
		},
		noop,
		noop,
	}
	allSpells["Shield"] = spell{
		113,
		6,
		func(player player, enemy enemy) {
			player.armor += 7
		},
		noop,
		func(player player, enemy enemy) {
			player.armor -= 7
		},
	}
	allSpells["Poison"] = spell{
		173,
		6,
		noop,
		func(player player, enemy enemy) {
			enemy.hp -= 3
		},
		noop,
	}
	allSpells["Recharge"] = spell{
		229,
		5,
		noop,
		func(player player, enemy enemy) {
			player.mana += 101
		},
		noop,
	}
}

func main() {
	activeSpells := make(map[string]int, len(allSpells))
	for spellName := range allSpells {
		activeSpells[spellName] = -1
	}

	wizard := player{50, 500, 0, activeSpells}
	boss := enemy{58, 9}

	start := time.Now()
	minMana := 999999999
	{
		var fight func(player, enemy, bool)
		fight = func(p player, e enemy, playersTurn bool) {
			if p.mana >= 0 {
				if boss.hp <= 0 && p.mana < minMana {
					minMana = p.mana
					fmt.Printf("Winna winna chicken dinna %d\n", minMana)
					return
				}
				for spellName, spell := range allSpells {
					if p.activeSpells[spellName] < 0 {
						for j := range p.activeSpells {
							if p.activeSpells[j] > 0 {
								// If a spell is active, do its effect
								allSpells[j].turnEffect(p, e)
								p.activeSpells[j]--
							} else if p.activeSpells[j] == 0 {
								// If a spell has become inactive, do its after effect
								allSpells[j].afterEffect(p, e)
								p.activeSpells[j]--
							}
						}
						if playersTurn {
							if p.activeSpells[spellName] < 0 {
								p.mana -= spell.cost
								p.activeSpells[spellName] = spell.duration
								// If a spell has become active, do its instant effect
								spell.instantEffect(p, e)
							}
						} else {
							d := e.damage - p.armor
							if d < 1 {
								d = 1
							}
							p.hp -= d
						}
						p := player{p.hp, p.mana, p.armor, p.activeSpells}
						e := enemy{e.hp, e.damage}
						fight(p, e, !playersTurn)
					}
				}
			}
		}
		fight(wizard, boss, true)
	}
	fmt.Printf("Part 1 took %v\n", time.Since(start))
	fmt.Printf("minMana=%d\n", minMana)
}
