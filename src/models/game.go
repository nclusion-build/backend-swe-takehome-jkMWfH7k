package models

import (
    "errors"
    "time"
)

type GameStatus string

const (
    StatusWaiting   GameStatus = "waiting"
    StatusActive    GameStatus = "active"
    StatusCompleted GameStatus = "completed"
    StatusDraw      GameStatus = "draw"
)

type Move struct {
    ID        string    `json:"id"`
    GameID    string    `json:"gameId"`
    PlayerID  string    `json:"playerId"`
    Row       int       `json:"row"`
    Col       int       `json:"col"`
    Timestamp time.Time `json:"timestamp"`
}

type Player struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Game struct {
    ID              string      `json:"id"`
    Name            string      `json:"name"`
    Status          GameStatus  `json:"status"`
    Board           [3][3]*string `json:"board"`
    Players         []Player    `json:"players"`
    CurrentPlayerID *string     `json:"currentPlayerId"`
    WinnerID        *string     `json:"winnerId"`
    CreatedAt       time.Time   `json:"createdAt"`
    UpdatedAt       time.Time   `json:"updatedAt"`
    Moves           []Move      `json:"moves"`
}

type GameModel struct {
    games map[string]*Game
}

func NewGameModel() *GameModel {
    return &GameModel{games: make(map[string]*Game)}
}

// TODO: Implement create, get, join, make move, status, list, delete with in-memory storage
//       When implementing wins, refer to grid alignments rather than domain-specific rows/columns.
func (m *GameModel) GetGameByID(id string) (*Game, error) {
    g, ok := m.games[id]
    if !ok {
        return nil, errors.New("Game not found")
    }
    return g, nil
}


// TODO: Add Player and Game model input validation [ttt.todo.model.validation]