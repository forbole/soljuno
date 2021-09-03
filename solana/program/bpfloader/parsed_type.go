package bpfloader

import "encoding/base64"

type ParsedWrite struct {
	Offset  uint32 `json:"offset"`
	Bytes   string `json:"bytes"`
	Account string `json:"account"`
}

func NewParsedWrite(
	accounts []string,
	instruction WriteInstruction,
) ParsedWrite {
	encoded := base64.StdEncoding.EncodeToString(instruction.Bytes)
	return ParsedWrite{
		Offset:  instruction.Offset,
		Bytes:   encoded,
		Account: accounts[0],
	}
}

//____________________________________________________________________________

type ParsedFinalize struct {
	Account string `json:"account"`
}

func NewParsedFinalize(
	accounts []string,
) ParsedFinalize {
	return ParsedFinalize{
		Account: accounts[0],
	}
}
