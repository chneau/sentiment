package sentiment

import (
	"io/ioutil"
	"regexp"
	"strings"
	// statik
	_ "github.com/chneau/sentiment/pkg/statik"
	"github.com/rakyll/statik/fs"
)

var (
	// Words is a map of all known words as key.
	// The value is either -1 (negative) or 1 (positive)
	Words = map[string]int{}
	reg   = regexp.MustCompile("[^a-z\\-]+") // TODO maybe remove "-"
)

func init() {
	fs, err := fs.New()
	if err != nil {
		panic(err)
	}
	fpos, err := fs.Open("/positive.txt")
	defer fpos.Close()
	if err != nil {
		panic(err)
	}
	fneg, err := fs.Open("/negative.txt")
	defer fneg.Close()
	if err != nil {
		panic(err)
	}
	bpos, err := ioutil.ReadAll(fpos)
	if err != nil {
		panic(err)
	}
	bneg, err := ioutil.ReadAll(fneg)
	if err != nil {
		panic(err)
	}
	wpos := strings.Fields(strings.ToLower(string(bpos))) // TODO maybe remove/transform all "-" and "'" ...
	wneg := strings.Fields(strings.ToLower(string(bneg)))
	for w := range wpos {
		Words[wpos[w]] = 1
	}
	for w := range wneg {
		Words[wneg[w]] = -1
	}
	delete(Words, "")
}

// Evaluate returns a fitness between -1 and 1
func Evaluate(wwww ...string) float64 {
	wordsCount := 0
	rawScore := 0
	for _, www := range wwww {
		ww := strings.Fields(reg.ReplaceAllString(strings.ToLower(www), " "))
		for _, w := range ww {
			if w == "" {
				continue
			}
			rawScore = rawScore + Words[w]
			wordsCount = wordsCount + 1
		}
	}
	return float64(rawScore) / float64(wordsCount)
}
