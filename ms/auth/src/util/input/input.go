package input

import "flag"

var (
	dbDriver string
)

type InputParams struct {
	dbDriver string
}

func NewInput() *InputParams {
	flag.StringVar(&dbDriver, "mysql", "mysql", "mysql or postgres")

	flag.Parse()

	return &InputParams{
		dbDriver: dbDriver,
	}
}
