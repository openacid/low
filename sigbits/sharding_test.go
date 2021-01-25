package sigbits

import (
	"strings"
	"testing"

	"github.com/openacid/testkeys"
	"github.com/stretchr/testify/require"
)

func TestShardByPrefix(t *testing.T) {

	ta := require.New(t)

	split := func(ks string) []string {
		ks = strings.Trim(ks, "\n ")
		return strings.Split(ks, "\n")
	}

	makeRst := func(keys []string, prefLens, starts []int32) []string {
		rst := make([]string, 0, len(keys))
		for j := 0; j < len(starts)-1; j++ {
			s := starts[j]
			e := starts[j+1]
			prefLen := prefLens[j]

			rst = append(rst, "="+keys[s])

			for k := s + 1; k < e; k++ {
				rst = append(rst, "-"+strings.Repeat(" ", int(prefLen))+keys[k][prefLen:])
			}

		}

		return rst
	}

	keysAbe_Abo := `
Abe
Abel
Abelia
Abelian
Abelicea
Abelite
Abelmoschus
Abelonian
Abencerrages
Aberdeen
Aberdonian
Aberia
Abhorson
Abie
Abies
Abietineae
Abiezer
Abigail
Abipon
Abitibi
Abkhas
Abkhasian
Ablepharus
Abnaki
Abner
Abo
`

	keysAaron_Abuta := `
Aaron
Aaronic
Aaronical
Aaronite
Aaronitic
Aaru
Ab
Ababdeh
Ababua
Abadite
Abama
Abanic
Abantes
Abarambo
Abaris
Abasgi
Abassin
Abatua
Abba
Abbadide
Abbasside
Abbie
Abby
Abderian
Abderite
Abdiel
Abdominales
Abe
Abel
Abelia
Abelian
Abelicea
Abelite
Abelmoschus
Abelonian
Abencerrages
Aberdeen
Aberdonian
Aberia
Abhorson
Abie
Abies
Abietineae
Abiezer
Abigail
Abipon
Abitibi
Abkhas
Abkhasian
Ablepharus
Abnaki
Abner
Abo
Abobra
Abongo
Abraham
Abrahamic
Abrahamidae
Abrahamite
Abrahamitic
Abram
Abramis
Abranchiata
Abrocoma
Abroma
Abronia
Abrus
Absalom
Absaroka
Absi
Absyrtus
Abu
Abundantia
Abuta
`

	cases := []struct {
		keys    []string
		maxSize int32
	}{
		{
			keys:    split(keysAbe_Abo),
			maxSize: 5,
		},
		{
			keys:    split(keysAbe_Abo),
			maxSize: 7,
		},
		{
			keys:    split(keysAaron_Abuta),
			maxSize: 7,
		},
		{
			keys:    testkeys.Load("200kweb2"),
			maxSize: 16,
		},
		{
			keys:    testkeys.Load("200kweb2"),
			maxSize: 64,
		},
	}

	for _, c := range cases {
		keys := c.keys
		prefLens, starts := ShardByPrefix(keys, c.maxSize)

		got := makeRst(keys, prefLens, starts)

		// fmt.Println(strings.Join(got, "\n"))

		ta.Equal(len(keys), len(got))

		// check if prefixes are all in ascending order.
		for j := 1; j < len(prefLens); j++ {
			prev := keys[starts[j-1]][:prefLens[j-1]]
			shd := keys[starts[j]][:prefLens[j]]

			ta.GreaterOrEqual(shd, prev, "compare %d th pref with %d th", j-1, j)
		}

		// check keys can be rebuilt.
		for j := 0; j < len(prefLens); j++ {

			ta.LessOrEqual(starts[j+1]-starts[j], c.maxSize)

			pref := keys[starts[j]][:prefLens[j]]

			for k := starts[j]; k < starts[j+1]; k++ {
				ta.Equal(pref, keys[k][:len(pref)], "%d th key", k)
			}
		}

		// fmt.Println(len(keys) / len(prefLens))
	}
}
