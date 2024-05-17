package enum

type Env string

var (
	Dev  Env = "development"
	Test Env = "test"
	Stag Env = "staging"
	Prod Env = "production"
)
