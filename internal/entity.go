package internal

type Entity interface {
	Attack() int
	TakingDamage() int
	Spell() int
	IsAlive() bool
	IsManaEmpty() bool
}

type Enemy struct {
	Name string
	HP   int
	DMG  int
	MP   int
}

type Enemys[E Enemy] struct{}

func (e *Enemy) Attack() int {
	return e.DMG
}

func (e *Enemy) IsAlive() bool {
	return e.HP > 0
}

func (e *Enemy) IsManaEmpty() bool {
	return e.MP > 0
}

func (e *Enemy) TakingDamage(dmg int) int {
	return dmg - e.HP
}

func (e *Enemy) Spell() int {
	if e.IsManaEmpty() {
		e.MP -= 5
		return e.DMG + 10
	}
	return 0
}
