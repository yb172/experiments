package cfg

// Config defines config structure
type Config struct {
	Own struct {
		Port int
	}

	Word struct {
		Service struct {
			host string
			port int
		}
	}

	Number struct {
		Service struct {
			host string
			port int
		}
	}

	Internets struct {
		Service struct {
			host string
			port int
		}
	}
}
