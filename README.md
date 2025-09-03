# Backend SWE Take-Home Assignment - Go

## Overview

This is a **1-2 hour take-home assignment**. You will build a small, network-accessible backend web service that manages a turn-based, grid-driven game from pre-defined rules. Your assignment is tailored: a randomized (but reproducible) set of TODOs, features, and bugs has been embedded inline.

You should focus on:
- Clear, maintainable API handlers and service logic
- Robust input validation and error handling
- Simple, reliable tests (unit and integration)
- Helpful logs/metrics stubs where applicable

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
go mod tidy
```

### Running the Application

```bash
go run src/main.go
```

The application will start on port 3000.

### Running Tests

```bash
go test ./...
```

### Running the Simulation

> Optional: You may create a simple simulation script or test that spins up your server, plays multiple sessions concurrently, and prints a small leaderboard summary.

## Project Structure

```
src/
├── models/
├── routes/
├── services/
└── main.go
```

## What You Need to Implement

### Selected Tasks

#### TODOs
- Add Player and Game model input validation
- Implement and validate core game logic (join/move/win/draw)
- Add player email validation
- Add validation for game creation and move inputs
- Extend leaderboard endpoints (pagination, filters)
- Add Go unit tests for Game and Player models

#### Feature Requests
- Add basic rate limiting middleware

#### Bugs To Fix

### Core Requirements (high-level)

1. Turn-based rules on a finite grid with obvious invalid-move conditions
2. Multiple sessions can run concurrently; two players start a session
3. End a session on win or draw; expose session status
4. Leaderboard endpoint returning top users by wins or "efficiency" (lower moves per win is better)
5. A small simulation or test path that exercises the API

Additionally, look for inline TODOs in language-appropriate files. Examples:
- Go: `src/routes/*.go`, `src/models/*.go`, `src/services/*.go`, `src/main.go`

> Focus on correctness, quality, and clarity. If you finish early, feel free to polish or extend.

## Notes

- Inline TODOs are your primary guide. GitHub Issues are intentionally disabled.
- Keep commits small and frequent with clear messages.
- You may add libraries if they help you implement tasks cleanly.

## Quick API Examples

Assuming your server is running on http://localhost:3000

Create a game
```bash
curl -s -X POST http://localhost:3000/games -H 'Content-Type: application/json' -d '{"name":"Sample"}' | jq .
```

Join the game
```bash
GAME_ID=<paste-from-create>
curl -s -X POST http://localhost:3000/games/$GAME_ID/join -H 'Content-Type: application/json' -d '{"playerId":"player-1"}' | jq .
curl -s -X POST http://localhost:3000/games/$GAME_ID/join -H 'Content-Type: application/json' -d '{"playerId":"player-2"}' | jq .
```

Make a move and get status
```bash
curl -s -X POST http://localhost:3000/games/$GAME_ID/moves -H 'Content-Type: application/json' -d '{"playerId":"player-1","row":0,"col":0}' | jq .
curl -s http://localhost:3000/games/$GAME_ID/status | jq .
```

Leaderboard (optional)
```bash
curl -s http://localhost:3000/leaderboard | jq .
```

## Submission

1. Ensure tests pass
2. Run the simulation script
3. Update this README with any setup notes
4. Submit your repository URL

Good luck! 🚀
