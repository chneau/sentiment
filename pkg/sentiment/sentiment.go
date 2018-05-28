package sentiment

import (
	"bufio"
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
	reg   = regexp.MustCompile("[^a-z0-9]+")
	files = map[string]int{
		"/badwords.txt": -1,
		"/negative.txt": -1,
		"/positive.txt": 1,
	}
	inMemory, _ = fs.New()
)

func sanitize(word string, replacer string) string {
	return reg.ReplaceAllString(strings.ToLower(word), replacer)
}

func init() {
	for k, v := range files {
		f, _ := inMemory.Open(k)
		scanner := bufio.NewScanner(f) // `\n` and `\r\n` proof
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			w := sanitize(scanner.Text(), "")
			if len(w) > 1 {
				Words[w] = v
			}
		}
	}
}

// Evaluate returns a fitness between -1 and 1
func Evaluate(wwww ...string) float64 {
	wordsCount := 0
	totalScore := 0
	for _, www := range wwww {
		for _, w := range strings.Fields(sanitize(www, " ")) {
			if w == "" {
				continue
			}
			if score, known := Words[w]; known == true {
				totalScore = totalScore + score
				wordsCount = wordsCount + 1
			}
		}
	}
	if wordsCount == 0 {
		return 0
	}
	return float64(totalScore) / float64(wordsCount)
}
