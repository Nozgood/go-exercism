package parsinglogs

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {

	textSplit := strings.Split(text, " ")
	stringToTest := textSplit[0]

	trcRegex := regexp.MustCompile(`[TRC]`)
	dbgRegex := regexp.MustCompile(`[DBG]`)
	infRegex := regexp.MustCompile(`[INF]`)
	wrnRegex := regexp.MustCompile(`[WRN]`)
	errRegex := regexp.MustCompile(`[ERR]`)
	ftlRegex := regexp.MustCompile(`[FTL]`)

	sliceReg := []*regexp.Regexp{trcRegex, dbgRegex, infRegex, wrnRegex, errRegex, ftlRegex}

	for _, value := range sliceReg {
		test := value.MatchString(stringToTest)
		fmt.Print(test)
	}
	return false
}