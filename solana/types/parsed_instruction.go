package types

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
