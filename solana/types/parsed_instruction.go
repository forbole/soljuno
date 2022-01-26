package types

import "encoding/json"

func NewParsedInstruction(typ string, value interface{}) ParsedInstruction {
	return ParsedInstruction{
		Type:  typ,
		Value: value,
	}
}

type ParsedInstruction struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (p ParsedInstruction) GetValueJSON() interface{} {
	bz, err := json.Marshal(p.Value)
	if err != nil || bz == nil {
		return "{}"
	}
	return bz
}
