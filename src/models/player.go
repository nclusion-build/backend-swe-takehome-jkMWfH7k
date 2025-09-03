package models

import (
    "errors"
    "regexp"
    "time"
)

type PlayerStats struct {
    GamesPlayed        int     `json:"gamesPlayed"`
    GamesWon           int     `json:"gamesWon"`
    GamesLost          int     `json:"gamesLost"`
    GamesDrawn         int     `json:"gamesDrawn"`
    TotalMoves         int     `json:"totalMoves"`
    AverageMovesPerWin float64 `json:"averageMovesPerWin"`
    WinRate            float64 `json:"winRate"`
    Efficiency         float64 `json:"efficiency"`
}

type PlayerModel struct {
    players map[string]*Player
}

func NewPlayerModel() *PlayerModel {
    return &PlayerModel{players: make(map[string]*Player)}
}

type Player struct {
    ID        string      `json:"id"`
    Name      string      `json:"name"`
    Email     string      `json:"email"`
    Stats     PlayerStats `json:"stats"`
    CreatedAt time.Time   `json:"createdAt"`
    UpdatedAt time.Time   `json:"updatedAt"`
}

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

// TODO: Implement create, get, update, delete, search with validations
func (m *PlayerModel) GetPlayerByID(id string) (*Player, error) {
    p, ok := m.players[id]
    if !ok {
        return nil, errors.New("Player not found")
    }
    return p, nil
}


// TODO: Add Player and Game model input validation [ttt.todo.model.validation]
// TODO: Add player email validation [ttt.todo.validation.player-email]