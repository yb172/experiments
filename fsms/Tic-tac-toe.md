# Tic tac toe

To create this game we would need:

* Server
* Two players (CLI)

Game works this way: player X runs CLI and game is created. We can't start playing until another person joins so game is in state "waiting for second player".
Player O runs CLI and game moves to state "Player X's move" when player X should move. Player O's CLI allows to issue commands but server should not allow B to make move until move is done by player X. But player O could exit the game.

```console

    0   1   2
  |-----------|
0 |   |   |   |
  |---|---|---|
1 |   |   |   |
  |---|---|---|
2 |   |   |   |
  |-----------|
```

Let's formalize a few scenarios

## Scenarios

### A wins

1. Player X connects to the server - game is created
1. Player O connects to the server - game is started
1. X makes move 1,1
1. O makes move 2,1
1. X makes move 1,2
1. O makes move 2,2
1. X makes move 1,0
1. Game ends, X wins

```console

    0   1   2
  |-----------|
0 |   |   |   |
  |---|---|---|
1 | x | x | x |
  |---|---|---|
2 |   | o | o |
  |-----------|
```

### B wins

1. Player X connects to the server - game is created
1. Player O connects to the server - game is started
1. X makes move 1,1
1. O makes move 1,2
1. X makes move 0,0
1. O makes move 2,2
1. X makes move 0,2
1. O makes move 2,0
1. Game ends, O wins

```console

    0   1   2
  |-----------|
0 | x |   | o |
  |---|---|---|
1 |   | x | o |
  |---|---|---|
2 | x |   | o |
  |-----------|
```

### Not allowed actions

1. X connects to server - game is created
1. X sends command `move 1,1` - server replies that move is not allowed, game hasn't started yet
1. O connects to server
1. O sends command `move 1,2` - server replies with error
1. X sends command `move 1,2` - ok
1. X sends command `move 1,0` - error, O's move
1. O sends command `move 1,0` - ok
1. X sends command `exit` - game exits for X and pauses for O
1. O sends command `move 0,0` - error, game is paused
1. X connects to server - game continues
1. X sends command `move 0,0` - error, O's move
1. O sends command `move 0,0` - ok

## Server

Server would keep track of game progress and provide updates to users. And it would be nice to send game updates as soon as possible so we're not going to use http and rest, but would use grpc instead.

Since we're not actually building a game as a product we would only support two users.
