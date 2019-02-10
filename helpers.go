package h

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"
)

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	l := len(runes)
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(l)]
	}
	return string(b)
}

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

func LogOnError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, e.Error())
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

var whitespaceRegex = regexp.MustCompile(`[\s\t]+`)

func TrimWhitespace(s string) string {
	return whitespaceRegex.ReplaceAllString(s, " ")
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

//ex: SendMail("127.0.0.1:25", (&mail.Address{"from name", "from@example.com"}).String(), "Email Subject", "message body", []string{(&mail.Address{"to name", "to@example.com"}).String()})
func SendMail(addr, from, subject, body string, to []string) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func RegexReplaceAllStringFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func IsPrivateIp(ip string) (bool, error) {
	if strings.Contains(ip, ".") {
		return IsPrivateIpv4(ip)
	}
	return IsPrivateIPv6(ip), nil
}

func IsPrivateIpv4(ipv4 string) (bool, error) {
	ip := net.ParseIP(ipv4)
	if ip == nil {
		errors.New(fmt.Sprintf("Invalid IP %s", ipv4))
	} else if ip.IsLoopback() {
		return true, nil
	} else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		return private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip), nil
	}
	return false, nil
}

func IsPublicIpv6(ipv6 string) bool {
	if strings.Contains(ipv6, ".") {
		return false
	}
	return !IsPrivateIPv6(ipv6)
}

func IsPrivateIPv6(ipv6 string) bool {
	if ipv6 == "::" || ipv6 == "::1" {
		return true
	}
	s := strings.Split(ipv6, ":")
	if len(s) < 1 {
		return false
	}
	firstWord := s[0]
	if (strings.HasPrefix(firstWord, "fc") || strings.HasPrefix(firstWord, "fd")) && len(firstWord) >= 4 {
		return true
	} else if firstWord == "fe80" || firstWord == "100" {
		return true
	}

	return false
}

func NetworkInterfaces() (interfaces []net.Interface, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, networkInterface := range netInterfaces {
		if networkInterface.Flags&net.FlagUp == 0 || networkInterface.Flags&net.FlagLoopback != 0 {
			continue
		}
		interfaces = append(interfaces, networkInterface)
	}
	return
}

func MapSearch(i interface{}, keys ...string) (m map[string]interface{}, b bool) {
	m, b = i.(map[string]interface{})
	if !b {
		return
	}
	for _, k := range keys {
		m, b = m[k].(map[string]interface{})
		if !b {
			return
		}
	}
	return
}

type Flock struct {
	*os.File
}

func NewFlock(path string) (*Flock, error) {
	f, err := os.OpenFile(path, syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, 0666)
	if err != nil {
		return nil, err
	}

	err = syscall.FcntlFlock(
		f.Fd(),
		syscall.F_SETLK,
		&syscall.Flock_t{
			Type:   syscall.F_WRLCK,
			Whence: io.SeekStart,
			Start:  0,
			Len:    0,
		})
	if err != nil {
		return nil, err
	}

	return &Flock{f}, nil
}

func (f *Flock) UnLock() error {
	err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	if err != nil {
		return err
	}
	return f.Close()
}

type WaitGroupRunner struct {
	*sync.WaitGroup
}

func NewWgExec() *WaitGroupRunner {
	return &WaitGroupRunner{}
}

func (w *WaitGroupRunner) Run(f func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		f()
	}()
}

func GetUrl(u string) (c string, err error) {
	r, err := http.Get(u)
	if err != nil {
		return
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	c = string(b)

	return
}

func JsonUrl(u string, t interface{}) (err error) {
	r, err := http.Get(u)
	if err != nil {
		return
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(t)
}
