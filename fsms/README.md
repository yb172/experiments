# Finite state machines

Purpose of this experiment is to learn how [FSM](https://github.com/looplab/fsm) works. We're planning to use it in one of our projects to group actions available for objects: user can create tasklists and send them for execution, and there are certain restrictions - e.g. user could add task only before list is submitted for execution, no task updates are allowed after list is "timed out", etc. It is possible to implement this logic with traditinal "service" that provides methods but then any unwanted transition should be coded explicitly (if list is started - do not allow addition of new tasks). FSM theoretically blacklists any transitions except for allowed ones.

## Most useless machine

[Most useless machine](https://www.youtube.com/watch?v=Z86V_ICUCD4) implemented in Go!

```bash
go build main.go && ./main useless
```
