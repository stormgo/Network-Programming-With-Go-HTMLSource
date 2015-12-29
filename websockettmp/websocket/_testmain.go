package main

import target "websocket"
import "testing"
import "regexp"

var tests = []testing.InternalTest{
	{"websocket.TestHixie76Challenge", target.TestHixie76Challenge},
	{"websocket.TestHixie76ClientHandshake", target.TestHixie76ClientHandshake},
	{"websocket.TestHixie76ServerHandshake", target.TestHixie76ServerHandshake},
	{"websocket.TestHixie76SkipLengthFrame", target.TestHixie76SkipLengthFrame},
	{"websocket.TestHixie76SkipNoUTF8Frame", target.TestHixie76SkipNoUTF8Frame},
	{"websocket.TestHixie76ClosingFrame", target.TestHixie76ClosingFrame},
	{"websocket.TestSecWebSocketAccept", target.TestSecWebSocketAccept},
	{"websocket.TestHybiClientHandshake", target.TestHybiClientHandshake},
	{"websocket.TestHybiClientHandshakeHybi08", target.TestHybiClientHandshakeHybi08},
	{"websocket.TestHybiServerHandshake", target.TestHybiServerHandshake},
	{"websocket.TestHybiServerHandshakeHybi08", target.TestHybiServerHandshakeHybi08},
	{"websocket.TestHybiServerHandshakeHybiBadVersion", target.TestHybiServerHandshakeHybiBadVersion},
	{"websocket.TestHybiShortTextFrame", target.TestHybiShortTextFrame},
	{"websocket.TestHybiShortMaskedTextFrame", target.TestHybiShortMaskedTextFrame},
	{"websocket.TestHybiShortBinaryFrame", target.TestHybiShortBinaryFrame},
	{"websocket.TestHybiControlFrame", target.TestHybiControlFrame},
	{"websocket.TestHybiLongFrame", target.TestHybiLongFrame},
	{"websocket.TestHybiClientRead", target.TestHybiClientRead},
	{"websocket.TestHybiShortRead", target.TestHybiShortRead},
	{"websocket.TestHybiServerRead", target.TestHybiServerRead},
	{"websocket.TestHybiServerReadWithoutMasking", target.TestHybiServerReadWithoutMasking},
	{"websocket.TestHybiClientReadWithMasking", target.TestHybiClientReadWithMasking},
	{"websocket.TestHybiServerFirefoxHandshake", target.TestHybiServerFirefoxHandshake},
	{"websocket.TestHybiErrSubProtocol", target.TestHybiErrSubProtocol},
	{"websocket.TestHybiNoSuitableSubProtocol", target.TestHybiNoSuitableSubProtocol},
	{"websocket.TestEcho", target.TestEcho},
	{"websocket.TestAddr", target.TestAddr},
	{"websocket.TestCount", target.TestCount},
	{"websocket.TestWithQuery", target.TestWithQuery},
	{"websocket.TestWithProtocol", target.TestWithProtocol},
	{"websocket.TestHTTP", target.TestHTTP},
	{"websocket.TestTrailingSpaces", target.TestTrailingSpaces},
	{"websocket.TestSmallBuffer", target.TestSmallBuffer},
}

var benchmarks = []testing.InternalBenchmark{}
var examples = []testing.InternalExample{}

var matchPat string
var matchRe *regexp.Regexp

func matchString(pat, str string) (result bool, err error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = regexp.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func main() {
	testing.Main(matchString, tests, benchmarks, examples)
}
