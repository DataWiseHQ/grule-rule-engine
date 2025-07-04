package benchmark

import (
	"bufio"
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	words    []string
	dupCheck map[string]bool
)

func init() {
	f, err := os.Open("words.txt")
	if err != nil {
		panic(err.Error())
	}
	defer func(f *os.File) {
		closeErr := f.Close()
		if closeErr != nil {
			panic(closeErr.Error())
		}
	}(f)
	buf := bufio.NewReader(f)
	words = make([]string, 0)
	for {
		str, readStringErr := buf.ReadString('\n')
		if readStringErr != nil {
			break
		}
		words = append(words, strings.TrimSpace(str))
	}
	dupCheck = make(map[string]bool)
}

// GetWord get a random english word
func GetWord(t bool) string {
	w := words[rand.Intn(len(words))]
	if t {
		return cases.Title(language.English).String(w)
	}
	return w
}

// MakeRule make a single dummy rule
func MakeRule(seq int) string {
	var rname string
	for {
		rname = GetWord(true) + GetWord(true) + GetWord(true)
		if _, ok := dupCheck[rname]; !ok {
			break
		}
	}
	buff := &bytes.Buffer{}
	buff.WriteString("rule ")
	buff.WriteString(rname)
	buff.WriteString(" \"")
	buff.WriteString(strconv.Itoa(seq))
	buff.WriteString(" ")
	buff.WriteString(GetWord(true))
	buff.WriteString(" ")
	buff.WriteString(GetWord(true))
	buff.WriteString("\"")
	buff.WriteString(" salience ")
	buff.WriteString(strconv.Itoa(rand.Intn(100) + 10))
	buff.WriteString(" {\n\t")
	buff.WriteString("when\n\t\t")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString(" == ")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString("() && \n\t\t")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString(" <= ")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString("()\n\tthen\n\t\t")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString("();\n\t\t")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString(" = ")
	buff.WriteString(GetWord(false))
	buff.WriteString(".")
	buff.WriteString(GetWord(true))
	buff.WriteString("();")
	buff.WriteString("\n}\n\n")

	return buff.String()
}

// GenRandomRule simply generate count number of simple parse-able rule into a file
func GenRandomRule(fileName string, count int) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		closeErr := f.Close()
		if closeErr != nil {
			panic(closeErr.Error())
		}
	}(f)
	for i := 1; i <= count; i++ {
		_, writeStringErr := f.WriteString(MakeRule(i))
		if writeStringErr != nil {
			return writeStringErr
		}
	}
	return nil
}
