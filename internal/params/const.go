package params

import (
	"fmt"
	mnemonic "github.com/evgeny-klyopov/go-pack-mnemonic"
)

const PackOperation = "pack"
const UnPackOperation = "unpack"

var ValidValues = map[string][]string{
	"base": []string{
		fmt.Sprintf("%d", mnemonic.Base10),
		fmt.Sprintf("%d", mnemonic.Base36),
		fmt.Sprintf("%d", mnemonic.Base62),
	},
	"lang": []string{
		mnemonic.English,
		mnemonic.Czech,
		mnemonic.French,
		mnemonic.Italian,
		mnemonic.Japanese,
		mnemonic.Korean,
		mnemonic.Spanish,
		mnemonic.ChineseTraditional,
		mnemonic.ChineseSimplified,
	},
	"operation": []string{
		PackOperation,
		UnPackOperation,
	},
}
