package termplt

import "regexp"

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorGray = "\033[37m"
const ColorWhite = "\033[97m"
const ColorCrossed = "\033[9m"

// Copied from https://github.com/acarl005/stripansi/blob/master/stripansi.go
const colorEscapeCodes = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var colorEscapeCodesRE = regexp.MustCompile(colorEscapeCodes)

// StripColor removes color escape codes from string
func StripColor(str string) string {
	return colorEscapeCodesRE.ReplaceAllString(str, "")
}
