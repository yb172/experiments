package thegame

import "fmt"

func joinFor(label Label) string {
	switch {
	case label == X:
		return xJoinsEvent
	case label == O:
		return oJoinsEvent
	default:
		panic(fmt.Sprintf("label %s is not supported", string(label)))
	}
}

func moveFor(label Label) string {
	switch {
	case label == X:
		return xMovesEvent
	case label == O:
		return oMovesEvent
	default:
		panic(fmt.Sprintf("label %s is not supported", string(label)))
	}
}

func winFor(label Label) string {
	switch {
	case label == X:
		return xWinsEvent
	case label == O:
		return oWinsEvent
	default:
		panic(fmt.Sprintf("label %s is not supported", string(label)))
	}
}

func passFor(label Label) string {
	switch {
	case label == X:
		return passMoveToOEvent
	case label == O:
		return passMoveToXEvent
	default:
		panic(fmt.Sprintf("label %s is not supported", string(label)))
	}
}
