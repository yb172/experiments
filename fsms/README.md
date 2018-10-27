# Finite state machines

Purpose of this experiment is to learn how [FSM](https://github.com/looplab/fsm) works. We're planning to use it in one of our projects to group actions available for objects: user can create tasklists and send them for execution, and there are certain restrictions - e.g. user could add task only before list is submitted for execution, no task updates are allowed after list is "timed out", etc. It is possible to implement this logic with traditinal "service" that provides methods but then any unwanted transition should be coded explicitly (if list is started - do not allow addition of new tasks). FSM theoretically blacklists any transitions except for allowed ones.

## Most useless machine

[Most useless machine](https://www.youtube.com/watch?v=Z86V_ICUCD4) implemented in Go!

```bash
go build main.go && ./main useless
```

## Step 1

First commit is simple copy of example from documentation. We have FSM with two states: `off` and `on` and two events: `turn-on` and `turn-off`.

Interesting - we don't define list of state explicitly - set of states is a collection of destinations for events and initial state.

All right. We could fire `turn-on` when in `off` state and vice versa.
But that is not yet fully functionall "most useless machine" - it doesn't turn itself off...

## Step 2

So now we want our FSM to fire `turn-off` event immediately after `turn-on` event is fired.

There is `Callbacks` property which calls a function every time some state change is happening, so technically we could check if destination state is `on` - then fire `turn-off` event. But it won't be great to have if/switch here - there must be an API for this use case...
All right - here is [the description of what could be callback](https://godoc.org/github.com/looplab/fsm#NewFSM). Looks like `turn-on` is what we need:

```go
    fsm.Callbacks{
        "enter_state": func(e *fsm.Event) { m.enterState(e) },
        "turn-on":     func(_ *fsm.Event) { m.onTurnOn() },
    },
...
func (m *MostUselessMachine) onTurnOn() {
    m.FSM.Event("turn-off")
}
```

For some reason it doesn't work - hangs after producing output `Moving from off to on`. To see what's going on let's add all non-specific callbacks with function `log`:

```go
    fsm.Callbacks{
        "before_event": m.log,
        "leave_state":  m.log,
        "enter_state":  m.log,
        "after_event":  m.log,
        "turn-on":      func(_ *fsm.Event) { m.onTurnOn() },
    },
...
func (m *MostUselessMachine) log(e *fsm.Event) {
    fmt.Printf("Event: %v\n", e)
}
```

Hmm, that wasn't very helpful:

```console
$ go build main.go && ./main useless

Turning on...
Event: &{0xc4200c4360 turn-on off on <nil> [] false false}
Event: &{0xc4200c4360 turn-on off on <nil> [] false false}
Event: &{0xc4200c4360 turn-on off on <nil> [] false false}
```

Maybe I should use `Async` [docs](https://godoc.org/github.com/looplab/fsm#Event.Async)? No, it seems to do different thing. Maybe some error is happening? No, no errors. It looks like deadlock is happening here: [fsm.go#270](https://github.com/looplab/fsm/blob/master/fsm.go#L270) - lock is not released until state transition is complete and it would be complete when we transition to a new state. But to transition to a new state we need lock...

It is actually not necessarily a bad thing. Auto-firing events could potentially lead to an endless loop. There might be a way to get around that but it's probably simpler to just not allow that. But! In this case it would be nice not to hang in deadlock but to give a meaningful error message. Opened [issue](https://github.com/looplab/fsm/issues/36) and PR to fix it in project repository, lets see what would author say (if anything).

## Step 3

So apparently we could not create most useless FSM using FSM API, but we could do a hack to still implement the machine: use goroutine which waits.

And it works! Though it looks like a very bad practice to use this approach anywhere except for this example - FSM for which state transitions happen only outside of FSM would be more simple and easier to test.

## Step 4

Code works but program doesn't feel like most useless machine - only few lines of text are printed. So let's make it interacitve - add a way to submit `on` and `off` commands to interact with the machine.

Try running:

```bash
go build main.go && ./main useless
```
