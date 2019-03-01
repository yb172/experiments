package cfg

// Config defines config structure
type Config struct {
	Default struct {
		RPS int
	}

	Gateway struct {
		Service struct {
			Address string
		}
	}
}
