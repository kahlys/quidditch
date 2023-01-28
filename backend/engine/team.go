package engine

type Team struct {
	Name  string
	Squad Squad
}

type Squad struct {
	Seeker  Player
	Chaser1 Player
	Chaser2 Player
	Chaser3 Player
	Beater1 Player
	Beater2 Player
	Keeper  Player
}
