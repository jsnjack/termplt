package main

import "regexp"

const ResetColor = "\033[0m"
const RedColor = "\033[31m"
const GreenColor = "\033[32m"
const YellowColor = "\033[33m"
const BlueColor = "\033[34m"
const PurpleColor = "\033[35m"
const CyanColor = "\033[36m"
const GrayColor = "\033[37m"
const WhiteColor = "\033[97m"
const CrossedColor = "\033[9m"

// Copied from https://github.com/acarl005/stripansi/blob/master/stripansi.go
const colorEscapeCodes = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var colorEscapeCodesRE = regexp.MustCompile(colorEscapeCodes)

// stripColor removes color escape codes from string
func stripColor(str string) string {
	return colorEscapeCodesRE.ReplaceAllString(str, "")
}
