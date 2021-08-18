package types

type UiCompiledInstruction struct {
	ProgramIDIndex uint8   `json:"programIdIndex"`
	Accounts       []uint8 `json:"accounts"`
	Data           string  `json:"data"`
}

type UiInnerInstruction struct {
	Index        uint8                   `json:"index"`
	Instructions []UiCompiledInstruction `json:"instructions"`
}
