package types

func NewParsedInstruction(typ string, data interface{}) ParsedInstruction {
	return ParsedInstruction{
		Type: typ,
		Data: data,
	}
}

type ParsedInstruction struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
