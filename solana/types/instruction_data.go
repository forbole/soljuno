package types

import "encoding/json"

type ParsedInstruction interface {
	Type() string
	JSON() []byte
}

func NewParsedInstruction(typ string, data interface{}) ParsedInstruction {
	return parsedInstruction{
		typ:  typ,
		data: data,
	}
}

type parsedInstruction struct {
	typ  string
	data interface{}
}

func (i parsedInstruction) Type() string {
	return i.typ
}

func (i parsedInstruction) JSON() []byte {
	if i.data == nil {
		return []byte{}
	}
	bz, _ := json.Marshal(i.data)
	return bz
}
