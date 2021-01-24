package main

type Player struct {
}

func (p *Player) Attack(target *BeAttackable) {
	target.BeAttacked()
}

type Monstor struct {
}

func (m *Monster) Attack(target *BeAttackable) {
	target.BeAttacked()
}

type Chest struct {
}

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	BeAttacked()
}

func Attack(attacker *Attack, defender *BeAttackable) {
	attacker.Attack(b)
}
