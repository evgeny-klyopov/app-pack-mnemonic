package app

import (
	"fmt"
	"github.com/evgeny-klyopov/bashColor"
	mnemonic "github.com/evgeny-klyopov/go-pack-mnemonic"
	"github.com/evgeny-klyopov/pack-mnemonic-app/internal/params"
	"strconv"
	"strings"
)

type app struct {
	color  bashColor.Colorer
	params params.Params
	packer mnemonic.Packer
}

func (a *app) Run() error {
	var convert mnemonic.Converter
	var err error

	err = a.params.Validate()
	if err != nil {
		return err
	}

	switch a.params.Operation {
	case params.PackOperation:
		convert, err = a.packer.Pack()
	case params.UnPackOperation:
		base, _ := strconv.Atoi(a.params.Base)
		convert, err = a.packer.UnPack(base)
	}

	if err == nil {
		a.printResult(convert)
	}

	return err
}

func (a *app) printResult(convert mnemonic.Converter) {
	fmt.Println(a.color.Teal("Phrase:"), a.packer.GetPhrase())
	fmt.Println(a.color.Teal("Lang:"), a.packer.GetLang())
	fmt.Println("")
	fmt.Println(a.color.Red("Mnemonic original:"), strings.Join(a.packer.GetMnemonicOriginal(), " "))
	fmt.Println(a.color.Red("Mnemonic:"), strings.Join(convert.GetMnemonic(), " "))
	fmt.Println(a.color.Red("Mnemonic short:"), strings.Join(convert.GetMnemonicShort(), " "))
	fmt.Println(a.color.Red("Mnemonic number:"), strings.Join(convert.GetNumberMnemonic(), " "))
	fmt.Println("")
	for _, val := range params.ValidValues["base"] {
		base, _ := strconv.Atoi(val)
		fmt.Println(a.color.Yellow("Base"+val+":"), convert.Get(base))
	}
}

type Apper interface {
	Run() error
}

func NewApp(p params.Params, color bashColor.Colorer) Apper {
	return &app{
		params: p,
		color:  color,
		packer: mnemonic.New(p.Phrase, p.Lang),
	}
}
