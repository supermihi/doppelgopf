package common

import (
	"fmt"
	"github.com/supermihi/karlchencloud/api"
	"github.com/supermihi/karlchencloud/cloud"
	"github.com/supermihi/karlchencloud/doko/game"
	"github.com/supermihi/karlchencloud/doko/match"
)

func ToApiPlayer(p game.Player, mapNoPlayerToPlayer1 bool) api.Player {
	switch p {
	case game.Player1:
		return api.Player_PLAYER_1
	case game.Player2:
		return api.Player_PLAYER_2
	case game.Player3:
		return api.Player_PLAYER_3
	case game.Player4:
		return api.Player_PLAYER_4
	default:
		if mapNoPlayerToPlayer1 {
			return api.Player_PLAYER_1
		}
		panic("unexpected NoPlayer in ToApiPlayer")
	}
}

func ToApiBid(b match.Bid) api.BidType {
	switch b {
	case match.Re:
		return api.BidType_RE_BID
	case match.ReKeine90:
		return api.BidType_RE_NO_NINETY
	case match.ReKeine60:
		return api.BidType_CONTRA_NO_SIXTY
	case match.ReKeine30:
		return api.BidType_RE_NO_THIRTY
	case match.ReSchwarz:
		return api.BidType_RE_SCHWARZ
	case match.Contra:
		return api.BidType_CONTRA_BID
	case match.ContraKeine90:
		return api.BidType_CONTRA_NO_NINETY
	case match.ContraKeine60:
		return api.BidType_CONTRA_NO_SIXTY
	case match.ContraKeine30:
		return api.BidType_CONTRA_NO_THIRTY
	case match.ContraSchwarz:
		return api.BidType_CONTRA_SCHWARZ
	default:
		panic(fmt.Sprintf("unexpected bid %v in ToApiBid", b))
	}
}

func ToAuctionPhase(p match.AuctionPhase) api.AuctionPhase {
	switch p {
	case match.VorbehaltAbfrage:
		return api.AuctionPhase_DECLARATION
	case match.VorbehaltSpezifikation:
		return api.AuctionPhase_SPECIFICATION
	}
	panic("unsupported auction phase")
}

func toSoloType(t game.AnnouncedGameType) api.SoloType {
	switch t {
	case game.KaroSoloType:
		return api.SoloType_DIAMONDS_SOLO
	case game.HerzSoloType:
		return api.SoloType_HEARTS_SOLO
	case game.PikSoloType:
		return api.SoloType_SPADES_SOLO
	case game.KreuzSoloType:
		return api.SoloType_CLUBS_SOLO
	}
	panic(fmt.Sprintf("not a solo type: %v", t))
}
func ToApiMode(mode game.Mode) *api.Mode {
	var gameType api.GameType
	var soloInfo *api.SoloInfo
	switch mode.(type) {
	case game.NormalspielMode:
		gameType = api.GameType_NORMAL_GAME
	case game.Hochzeit:
		gameType = api.GameType_MARRIAGE
	default:
		gameType = api.GameType_VOLUNTARY_SOLO
		soloInfo = &api.SoloInfo{Soloist: ToApiPlayer(game.Soloist(mode), false), SoloType: toSoloType(mode.Type())}
	}
	return &api.Mode{Type: gameType, SoloInfo: soloInfo}
}

func ToApiSuit(s game.Suit) api.Suit {
	switch s {
	case game.Karo:
		return api.Suit_DIAMONDS
	case game.Herz:
		return api.Suit_HEARTS
	case game.Pik:
		return api.Suit_SPADES
	case game.Kreuz:
		return api.Suit_CLUBS
	}
	panic(fmt.Sprintf("unexpected suit %v", s))
}

func ToApiRank(r game.Rank) api.Rank {
	switch r {
	case game.Neun:
		return api.Rank_NINE
	case game.Bube:
		return api.Rank_JACK
	case game.Dame:
		return api.Rank_QUEEN
	case game.Koenig:
		return api.Rank_KING
	case game.Zehn:
		return api.Rank_TEN
	case game.Ass:
		return api.Rank_ACE
	}
	panic(fmt.Sprintf("unexpected rank %v", r))
}

func ToApiCard(c game.Card) *api.Card {
	return &api.Card{Suit: ToApiSuit(c.Suit), Rank: ToApiRank(c.Rank)}
}

func ToApiTrick(t *game.IncompleteTrick, m game.Mode) *api.Trick {
	result := &api.Trick{Forehand: ToApiPlayer(t.Forehand, false)}
	if c, ok := t.CardOf(game.Player1); ok {
		result.CardPlayer_1 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player2); ok {
		result.CardPlayer_2 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player3); ok {
		result.CardPlayer_3 = ToApiCard(c)
	}
	if c, ok := t.CardOf(game.Player4); ok {
		result.CardPlayer_4 = ToApiCard(c)
	}
	if t.IsComplete() {
		result.Winner = ToApiPlayer(game.WinnerOfTrick(t.CardsByPlayer(), t.Forehand, m), false)
	}
	return result
}

func toAuctionState(a *match.Auction) *api.AuctionState {
	declarations := make([]*api.Declaration, 0)
	for _, p := range game.Players() {
		decl := a.DeclarationOf(p)
		if decl != match.NotDeclared {
			apiDecl := &api.Declaration{Player: ToApiPlayer(p, false), Vorbehalt: decl == match.Vorbehalt}
			declarations = append(declarations, apiDecl)
		}
	}
	return &api.AuctionState{Phase: ToAuctionPhase(a.Phase()), Declarations: declarations}
}

func toApiBids(bids *match.Bids) []*api.Bid {
	var ans []*api.Bid
	for _, player := range game.Players() {
		bidsOf := bids.BidsOf(player)
		apiBidsOf := make([]*api.Bid, len(bidsOf))
		for i, b := range bidsOf {
			apiBidsOf[i] = &api.Bid{Player: ToApiPlayer(player, false), Bid: ToApiBid(b)}
		}
		ans = append(ans, apiBidsOf...)
	}
	return ans
}

func ToGameState(m *match.Match) *api.GameState {
	return &api.GameState{Mode: ToApiMode(m.Mode()),
		Bids:            toApiBids(m.Bids),
		CompletedTricks: int32(m.Game.NumCompletedTricks()),
		CurrentTrick:    ToApiTrick(m.Game.CurrentTrick, m.Mode())}
}
func ToMatchState(tm *cloud.TableMatch) *api.MatchState {
	m := tm.Match
	turn := ToApiPlayer(m.WhoseTurn(), true)
	players := &api.Players{
		Player_1: string(tm.Players[game.Player1]),
		Player_2: string(tm.Players[game.Player2]),
		Player_3: string(tm.Players[game.Player3]),
		Player_4: string(tm.Players[game.Player4]),
	}
	switch m.Phase() {
	case match.InAuction:
		auctionState := toAuctionState(m.Auction)
		return &api.MatchState{Phase: api.MatchPhase_AUCTION,
			Turn:    turn,
			Players: players,
			Details: &api.MatchState_AuctionState{AuctionState: auctionState}}
	case match.InGame:
		gameState := ToGameState(m)
		return &api.MatchState{Phase: api.MatchPhase_GAME,
			Turn:    turn,
			Players: players,
			Details: &api.MatchState_GameState{GameState: gameState}}
	case match.MatchFinished:
		return &api.MatchState{Phase: api.MatchPhase_FINISHED, Players: players}
	}
	panic(fmt.Sprintf("ToMatchState called with invalid match phase %v", m.Phase()))
}
