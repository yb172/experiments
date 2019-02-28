package cfg

// Config defines config structure
type Config struct {
	Default struct {
		RPS int
	}

	Gen struct {
		Gateway struct {
			Service struct {
				Host string
				Port string
			}
		}
	}
}
