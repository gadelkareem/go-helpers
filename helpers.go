package h

import (
	"crypto/md5"
	"encoding/hex"
	"html"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

func RandomNumber(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func LiftRLimits() (rLimit syscall.Rlimit, err error) {
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return
}

type NoRetryError struct {
	err error
}

func NewNoRetryError(err error) NoRetryError {
	return NoRetryError{err}
}
func (e NoRetryError) Error() string { return e.err.Error() }

func Retry(retryFunc func() error, maxRetries int) (err error) {
	retry := 0
	for retry < maxRetries {
		err = retryFunc()
		if err == nil {
			return
		}
		if nre, isNoRetry := err.(NoRetryError); isNoRetry {
			return nre.err
		}
		retry++
	}
	return
}

func Utf8Encode(s string) string {
	if utf8.ValidString(s) {
		return s
	}
	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		v = append(v, r)
	}
	return string(v)
}

func CleanString(s string) string {
	s = Utf8Encode(strings.TrimSpace(html.UnescapeString(s)))
	s = strings.Replace(s, "\\", "", -1)
	return s
}

var slugReplacer = strings.NewReplacer(" ", "-", "\t", "-", "/", "-", "\\", "-")

func Slug(s string) string {
	return slugReplacer.Replace(s)
}

func SubString(s string, l int) string {
	if len(s) <= l {
		return s
	}
	return s[0:l]
}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

var nonAlphaNumRegex = regexp.MustCompile(`[^0-9A-Za-z]+`)

func Tags(s string, minChar, maxNum int) (tags []string) {
	words := nonAlphaNumRegex.Split(s, -1)
	if len(words) == 0 {
		return
	}
	for _, w := range words {
		if len(w) > minChar && !IsStopWord(w) {
			tags = append(tags, w)
		}
		if len(tags) > maxNum {
			break
		}
	}
	return

}

var stopWords = []string{"a", "about", "above", "across", "after", "afterwards", "again", "against", "all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "amoungst", "amount", "an", "and", "another", "any", "anyhow", "anyone", "anything", "anyway", "anywhere", "are", "around", "as", "at", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "behind", "being", "below", "beside", "besides", "between", "beyond", "bill", "both", "bottom", "but", "by", "call", "can", "cannot", "cant", "co", "con", "could", "couldnt", "cry", "de", "describe", "detail", "do", "done", "down", "due", "during", "each", "eg", "eight", "either", "eleven", "else", "elsewhere", "empty", "enough", "etc", "even", "ever", "every", "everyone", "everything", "everywhere", "except", "few", "fifteen", "fify", "fill", "find", "fire", "first", "five", "for", "former", "formerly", "forty", "found", "four", "from", "front", "full", "further", "get", "give", "go", "had", "has", "hasnt", "have", "he", "hence", "her", "here", "hereafter", "hereby", "herein", "hereupon", "hers", "herself", "him", "himself", "his", "how", "however", "hundred", "ie", "if", "in", "inc", "indeed", "interest", "into", "is", "it", "its", "itself", "keep", "last", "latter", "latterly", "least", "less", "ltd", "made", "many", "may", "me", "meanwhile", "might", "mill", "mine", "more", "moreover", "most", "mostly", "move", "much", "must", "my", "myself", "name", "namely", "neither", "never", "nevertheless", "next", "nine", "no", "nobody", "none", "noone", "nor", "not", "nothing", "now", "nowhere", "of", "off", "often", "on", "once", "one", "only", "onto", "or", "other", "others", "otherwise", "our", "ours", "ourselves", "out", "over", "own", "part", "per", "perhaps", "please", "put", "rather", "re", "same", "see", "seem", "seemed", "seeming", "seems", "serious", "several", "she", "should", "show", "side", "since", "sincere", "six", "sixty", "so", "some", "somehow", "someone", "something", "sometime", "sometimes", "somewhere", "still", "such", "system", "take", "ten", "than", "that", "the", "their", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "therefore", "therein", "thereupon", "these", "they", "thickv", "thin", "third", "this", "those", "though", "three", "through", "throughout", "thru", "thus", "to", "together", "too", "top", "toward", "towards", "twelve", "twenty", "two", "un", "under", "until", "up", "upon", "us", "very", "via", "was", "we", "well", "were", "what", "whatever", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "whereupon", "wherever", "whether", "which", "while", "whither", "who", "whoever", "whole", "whom", "whose", "why", "will", "with", "within", "without", "would", "yet", "you", "your", "yours", "yourself", "yourselves", "able", "dear", "did", "does", "got", "i", "just", "let", "like", "likely", "said", "say", "says", "tis", "twas", "wants"}

func IsStopWord(s string) bool {
	s = strings.ToLower(s)
	for _, w := range stopWords {
		if s == w {
			return true
		}
	}
	return false
}

func RunEvery(ttl time.Duration, f func()) {
	ticker := time.NewTicker(ttl)
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}

func ReCaptcha(secret, response string) (bool, error) {
	resp, err := http.PostForm(
		"https://www.google.com/recaptcha/api/siteverify",
		url.Values{
			"secret":   {secret},
			"response": {response},
		},
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return strings.Contains(string(b), "success\": true"), nil
}
