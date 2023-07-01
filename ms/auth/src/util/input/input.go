package input

import "flag"

var (
	dbDriver string
)

type InputParams struct {
	dbDriver string
}

func New() *InputParams {
	flag.StringVar(&dbDriver, "mysql", "mysql", "mysql or postgres")

	flag.Parse()

	return &InputParams{
		dbDriver: dbDriver,
	}
}
