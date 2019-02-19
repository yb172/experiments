package cfg

// Config defines config structure
type Config struct {
	Default struct {
		RPS int
	}

	Service struct {
		Address string
	}
}
