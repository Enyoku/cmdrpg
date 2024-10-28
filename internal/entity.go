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

type Enemies struct {
	Id int
	Enemy
}

type Player struct {
	Name string
	HP   int
	MP   int
	DMG  int
	LVL  int
	EXP  int
}

// Enemy struct methods
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

// Player struct methods
func (p *Player) Attack() int {
	return p.DMG
}

func (p *Player) IsAlive() bool {
	return p.HP > 0
}

func (p *Player) IsManaEmpty() bool {
	return p.MP > 0
}

func (p *Player) TakingDamage(dmg int) int {
	return dmg - p.HP
}

func (p *Player) Spell() int {
	if p.IsManaEmpty() {
		p.MP -= 5
		return p.DMG + 10
	}
	return 0
}

func (p *Player) LvlUp() int {
	if p.EXP >= 25 {
		p.LVL += 1
	}
	return 0
}

func (p *Player) ExpGain(e *Enemy) int {
	if !e.IsAlive() {
		p.EXP += 5
	}
	return 0
}
