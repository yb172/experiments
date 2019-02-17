package cfg

// Config defines config structure
type Config struct {
	Own struct {
		Port int
	}

	Word struct {
		Service struct {
			Host string
			Port int
		}
	}

	Number struct {
		Service struct {
			Host string
			Port int
		}
	}

	Internets struct {
		Service struct {
			Host string
			Port int
		}
	}
}
