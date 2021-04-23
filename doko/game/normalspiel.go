package game

func IsNormalTrump(k Card) bool {
	return k.Suit == Karo || k.Rank == Bube || k.Rank == Dame || k == Dulle()
}

func NormalGameSuit(c Card) GameSuit {
	if IsNormalTrump(c) {
		return Trumpf
	}
	return c.Suit.AsFehl()
}

type NormalspielMode struct {
	parties    [NumPlayers]Party
	silentSolo bool
}

func NewNormalspiel(dealtCards Cards) NormalspielMode {
	var parties [NumPlayers]Party
	silentSolo := false
	for _, p := range Players() {
		alte := dealtCards[p].NumAlte()
		if alte >= 1 {
			parties[p] = ReParty
			if alte == 2 {
				silentSolo = true
			}
		} else {
			parties[p] = ContraParty
		}
	}
	return NormalspielMode{parties, silentSolo}
}

func (n NormalspielMode) IsSilentSolo() bool {
	return n.silentSolo
}

func (NormalspielMode) OnCompletedTrick(Trick, int) {
}

func (NormalspielMode) Klaerungsstich() int {
	return -1
}

func (n NormalspielMode) GameSuit(k Card) GameSuit {
	return NormalGameSuit(k)
}

func (n NormalspielMode) PartyOf(p Player) Party {
	return n.parties[p]
}

func (n NormalspielMode) Type() AnnouncedGameType {
	return NormalGameType
}
