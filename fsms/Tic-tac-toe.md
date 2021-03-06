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

## Step 1. Game FSM

All right, we have initial grpc configured but right now it doesn't have anything to talk to. So we need our game implemented. And the first step in it is how game starts. When first user connects - game is created. When second user connects existing game is used where first user is already waiting.

So game should provide method "connect" which would either create a new game or connect to existing game. Or to make it clearer we could have "connect" return label of the player if there is a game, or error if there is no game. In the second case code will call "create" method and new game would be created.

> Regarding returning an error for `Connect` method: it is tricky thing what is the right approach. On one hand it is a legitimate error from our requirements point of view - somebody is trying to connect when there is no game - that's an error. On the other hand we know how to deal with this error - create a new game. But let's say we would like to do some operation in `Connect` that could return different type of error - e.g. talking over network or working with files. Then function that calls `Connect` should distinguish between "there is no game" error and other types of errors. One way is to use "error values" but it's not great. There was actually a talk about this... There is an article by Rob Pike "[errors are values](https://blog.golang.org/errors-are-values)" but that's different.
>
> Here it is - "[GopherCon 2016: Dave Cheney - Dont Just Check Errors Handle Them Gracefully](https://youtu.be/lsBF58Q-DnY)". And way to handle different errors by having `if err == smth` he calls "[Sentinel errors](https://youtu.be/lsBF58Q-DnY?t=250)" and recommends to not use such approach. All right, let's don't use it.

What should `Connect` and `Create` return? These methods would be called by grpc handlers so they should return something that could be later used to call game methods, e.g. "move" - so what we return should know which player it represents. So let's call it `Player`. It would contain reference to a game and a label and provide methods to play the game.

```Go
// Player provides methods to play a game
type Player struct {
  game  *game
  Label rune
}
```

`game` is another struct that contains game logic, so it should contain FSM:

```Go
type game struct {
  fsm   *fsm.FSM
}
```

However `game` is more like one match of the game, not the game itself (it contains state). So it could make more sense to rename it to `GameMatch`. Also right now to simplify our life we're going to support only single instance of the game, but still this should be enforced somewhere. So in addition to `game.go` with `GameMatch` there will be a separate `playground.go` file which would contain `Start` and `Join` methods, and these two files would be located in separate `thegame` package (to isolate any package-level visible variables).

Moving closer to game logic: instance of the game is FSM. FSM consists of states and state changes are happening through events. Tic tac toe game could be represented as following FSM states:

* Created
* Waiting for second player
* X move
* X move check
* O move
* O move check
* X won
* O won
* Draw

And the second part - events that cause state change:

* X joins (Created -> Waiting for second player)
* O joins (Waiting for second player -> X move)
* X moves (X move -> X move check)
* Pass move to O (X move check -> O move)
* O moves (O move -> O move check)
* Pass move to X (O move check -> X move)
* X wins (X move check -> X won)
* O wins (O move check -> O won)
* Draw (X move check / O move check -> Draw)

If represented as a graph:

```text
                                     |-----------------------------------------|
                                     |                                         |
Created -> Waiting for second player -> X move -> X check -> O move -> O check -
                                                          |-> X wins           |-> O wins
                                                          |-> draw <-----------|
```

Interesting note here is that our FSM actually doesn't check the game board - it is only responsible for transitions between states. So when we receive "Move 1,1" from O we would first try to trigger FSM event and if that succeeds - we would check if that event caused O to win (then issue "O wins" event) or draw (then issue "draw" event). Else - let X make move (issue "pass move to X" event).

## Step 2. Game logic

So it turns out FSM could help us to "guard" state transitions but it doesn't have logic because we have to issue events outside of FSM (as was discovered in "[most useless machine](./Useless.md)"). What logic do we need? Let's consider what are the events that we have:

* X joins - no logic here
* O joins - no logic here either
* X moves - move should happen inside the board on an empty slot. Also after this event happens we have to decide what comes next:
  * If X have 3 consequtive Xs - then it's "X wins"
  * If there are no more free slots - then it's "Draw"
  * Otherwise "Pass move to O"
* O moves - same as for "X moves" but for O

And that is it. So we don't have a lot of logic.

Next question: which interface should `GameMatch` provide to the user of this interface? It should be same for both X and O but should behave differently for X and O, and provide methods that are related to events:

* `func (g *Game) Join() (*Player, error)` - returns `Player` instance. First caller would get instance for X, second - for O
* `func (p *Player) Move(row, col int) error` - make a move

One last thing that we wound need is game board where we would store moves and check if win or draw condition happened. `Board` would provide methods:

* `func Create() *Board` - creates new board
* `func (b *Board) Move(row, col int, label Label) error` - make a move
* `func (b *Board) CheckWin() bool` - check if move resulted in win
* `func (b *Board) CheckDraw() bool` - check if there are no more free slots. This actually could be enhanced to check if there are no winning moves possible.

`Board` actually has no need to be accessible outside of `thegame` package so it could be package private: `board`.

A few things came up while implementation was in progress:

* It seems better to have `Player` in a separate package so that it could not access `GameMatch` private fields and would only use public methods
* Check if move is valid (e.g. cell has not been used) should be performed in callback for `leave_<state>` key. Reason - if check is not passed - no transition should happen. However it is not very convenient to pass arguments to event, type assertion is needed
* Having set of events for X and for O resulted in some not-nicely-looking funcs in `event-helpers.go`. Alternative would be to have a single event: `join`, `move`, `win`, `pass`, but then we would have to guard move sequence ourselves. Which actually might be better, since we wouldn't have to have duplicate events, states, and these functions. Instead there will be one more field to track - last player, and few more checks for appropriate events. I'll try to refactor the code later.
