package searchV2

import (
	"strings"

	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
	"github.com/blugelabs/bluge/analysis/tokenizer"
)

type punctuationCharFilter struct{}

var punctuationReplacer *strings.Replacer

func init() {
	var punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	args := make([]string, 0, len(punctuation)*2)
	for _, r := range punctuation {
		args = append(args, string(r))
		args = append(args, " ")
	}
	punctuationReplacer = strings.NewReplacer(args...)
}

const ngramEdgeFilterMaxLength = 7

func (t *punctuationCharFilter) Filter(input []byte) []byte {
	return []byte(punctuationReplacer.Replace(string(input)))
}

var ngramIndexAnalyzer = &analysis.Analyzer{
	CharFilters: []analysis.CharFilter{&punctuationCharFilter{}},
	Tokenizer:   tokenizer.NewWhitespaceTokenizer(),
	TokenFilters: []analysis.TokenFilter{
		token.NewCamelCaseFilter(),
		token.NewLowerCaseFilter(),
		token.NewEdgeNgramFilter(token.FRONT, 1, ngramEdgeFilterMaxLength),
	},
}

var ngramQueryAnalyzer = &analysis.Analyzer{
	CharFilters: []analysis.CharFilter{&punctuationCharFilter{}},
	Tokenizer:   tokenizer.NewWhitespaceTokenizer(),
	TokenFilters: []analysis.TokenFilter{
		token.NewCamelCaseFilter(),
		token.NewLowerCaseFilter(),
	},
}
