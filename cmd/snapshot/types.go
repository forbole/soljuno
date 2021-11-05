package snapshot

type AccountInfo struct {
	Pubkey string
	Detail AccountDetail `yaml:"account"`
}

type AccountDetail struct {
	Balance string `yaml:"balance"`
	Owner   string `yaml:"owner"`
	Slot    uint64 `yaml:"slot"`
}
