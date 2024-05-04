package wsx

import (
	"fmt"

	ma "github.com/multiformats/go-multiaddr"
)

// You **MUST** register your multicodecs with
// https://github.com/multiformats/multicodec before adding them here.
const (
	P_WS_WITH_PATH  = 4770 // non-standard
	P_WSS_WITH_PATH = 4780 // non-standard
)

var (
	protoWSX = ma.Protocol{
		Name:       "x-parity-ws",
		Code:       P_WS_WITH_PATH,
		VCode:      ma.CodeToVarint(P_WS_WITH_PATH),
		Size:       ma.LengthPrefixedVarSize,
		Path:       true,
		Transcoder: TranscoderWsPath,
	}
	protoWSSX = ma.Protocol{
		Name:       "x-parity-wss",
		Code:       P_WSS_WITH_PATH,
		VCode:      ma.CodeToVarint(P_WSS_WITH_PATH),
		Size:       ma.LengthPrefixedVarSize,
		Path:       true,
		Transcoder: TranscoderWsPath,
	}
)

func init() {
	for _, p := range []ma.Protocol{
		protoWSX,
		protoWSSX,
	} {
		if err := ma.AddProtocol(p); err != nil {
			panic(err)
		}
	}
}

var TranscoderWsPath = ma.NewTranscoderFromFunctions(wsxStB, wsxBtS, wsPathVal)

func wsPathVal(b []byte) error {
	if len(b) < 1 {
		return fmt.Errorf("empty websocket path")
	}
	// TODO: validate leading percent encoded slash
	return nil
}

func wsxStB(s string) ([]byte, error) {
	return []byte(s), nil
}

func wsxBtS(b []byte) (string, error) {
	return string(b), nil
}
