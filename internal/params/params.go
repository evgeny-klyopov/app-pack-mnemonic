package params

import (
	"errors"
	"fmt"
	mnemonic "github.com/evgeny-klyopov/go-pack-mnemonic"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strconv"
	"strings"
)

type Params struct {
	Phrase    string
	Operation string
	Lang      string
	Base      string
}

func (p *Params) GetFlags() []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "phrase",
			Required:    true,
			HasBeenSet:  false,
			Aliases:     []string{"p"},
			Usage:       "mnemonic phrase",
			Destination: &p.Phrase,
		},
		&cli.StringFlag{
			Name:        "operation",
			Required:    false,
			Value:       PackOperation,
			Aliases:     []string{"o"},
			Usage:       fmt.Sprintf("operation (%s)", strings.Join(ValidValues["operation"], "/")),
			HasBeenSet:  true,
			Destination: &p.Operation,
		},
		&cli.StringFlag{
			Name:        "base",
			Required:    false,
			Value:       strconv.Itoa(mnemonic.Base10),
			Aliases:     []string{"b"},
			Usage:       fmt.Sprintf("base (%s)", strings.Join(ValidValues["base"], "/")),
			HasBeenSet:  true,
			Destination: &p.Base,
		},
		&cli.StringFlag{
			Name:        "lang",
			Required:    false,
			Value:       mnemonic.English,
			Aliases:     []string{"l"},
			Usage:       fmt.Sprintf("lang (%s)", strings.Join(ValidValues["lang"], "/")),
			HasBeenSet:  true,
			Destination: &p.Lang,
		},
	}
}

func (p *Params) Validate() error {
	var err error
	var errs []string

	for key, values := range ValidValues {
		r := reflect.ValueOf(p)
		f := reflect.Indirect(r).FieldByName(cases.Title(language.English).String(key))

		if slices.Index(values, f.String()) == -1 {
			errs = append(errs, fmt.Sprintf("not support %s - %s", key, f.String()))
		}
	}

	if len(errs) > 0 {
		err = errors.New(strings.Join(errs, "\n"))
	}
	return err
}

func NewParams() Params {
	return Params{}
}
