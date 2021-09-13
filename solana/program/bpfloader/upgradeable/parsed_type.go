package upgradable_loader

import "encoding/base64"

type ParsedInitializeBuffer struct {
	Account string `json:"account"`

	Authority string `json:"authority,omitempty"`
}

func NewParsedInitializeBuffer(
	accounts []string,
) ParsedInitializeBuffer {
	parsed := ParsedInitializeBuffer{
		Account: accounts[0],
	}
	if len(accounts) > 1 {
		parsed.Authority = accounts[1]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedWrite struct {
	Offset    uint32 `json:"offset"`
	Bytes     string `json:"bytes"`
	Account   string `json:"account"`
	Authority string `json:"authority"`
}

func NewParsedWrite(
	accounts []string,
	instruction WriteInstruction,
) ParsedWrite {
	encoded := base64.StdEncoding.EncodeToString(instruction.Bytes)
	return ParsedWrite{
		Offset:    instruction.Offset,
		Bytes:     encoded,
		Account:   accounts[0],
		Authority: accounts[1],
	}
}

//____________________________________________________________________________

type ParsedDeployWithMaxDataLen struct {
	MaxDataLen         uint32 `json:"maxDataLen"`
	PayerAccount       string `json:"payerAccount"`
	ProgramDataAccount string `json:"programDataAccount"`
	ProgramAccount     string `json:"programAccount"`
	BufferAccount      string `json:"bufferAccount"`
	RentSysvar         string `json:"rentSysvar"`
	ClockSysvar        string `json:"clockSysvar"`
	SystemProgram      string `json:"systemProgram"`
	Authority          string `json:"authority"`
}

func NewParsedDeployWithMaxDataLen(
	accounts []string,
	instruction DeployWithMaxDataLenInstruction,
) ParsedDeployWithMaxDataLen {
	return ParsedDeployWithMaxDataLen{
		MaxDataLen:         instruction.MaxDataLen,
		PayerAccount:       accounts[0],
		ProgramDataAccount: accounts[1],
		ProgramAccount:     accounts[2],
		BufferAccount:      accounts[3],
		RentSysvar:         accounts[4],
		ClockSysvar:        accounts[5],
		SystemProgram:      accounts[6],
		Authority:          accounts[7],
	}
}

//____________________________________________________________________________

type ParsedUpgrade struct {
	ProgramDataAccount string `json:"programDataAccount"`
	ProgramAccount     string `json:"programAccount"`
	BufferAccount      string `json:"bufferAccount"`
	SpillAccount       string `json:"spillAccount"`
	RentSysvar         string `json:"rentSysvar"`
	ClockSysvar        string `json:"clockSysvar"`
	Authority          string `json:"authority"`
}

func NewParsedUpgrade(
	accounts []string,
) ParsedUpgrade {
	return ParsedUpgrade{
		ProgramDataAccount: accounts[0],
		ProgramAccount:     accounts[1],
		BufferAccount:      accounts[2],
		SpillAccount:       accounts[3],
		RentSysvar:         accounts[4],
		ClockSysvar:        accounts[5],
		Authority:          accounts[6],
	}
}

//____________________________________________________________________________

type ParsedSetAuthority struct {
	Account   string `json:"account"`
	Authority string `json:"authority"`

	NewAuthority string `json:"newAuthority,omitempty"`
}

func NewParsedSetAuthority(
	accounts []string,
) ParsedSetAuthority {
	parsed := ParsedSetAuthority{
		Account:   accounts[0],
		Authority: accounts[1],
	}
	if len(accounts) > 2 {
		parsed.NewAuthority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedClose struct {
	Account   string `json:"account"`
	Recipient string `json:"recipient"`
	Authority string `json:"authority"`
}

func NewParsedClose(
	accounts []string,
) ParsedClose {
	return ParsedClose{
		Account:   accounts[0],
		Recipient: accounts[1],
		Authority: accounts[2],
	}
}
