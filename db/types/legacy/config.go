package types

type ValidatorConfigRow struct {
	Address string `db:"address"`
	Slot    uint64 `db:"slot"`
	Owner   string `db:"owner"`

	Name            string `db:"name"`
	KeybaseUsername string `db:"keybase_username"`
	Website         string `db:"website"`
	Details         string `db:"details"`
}

type ParsedValidatorConfig struct {
	Name            string `json:"name"`
	KeybaseUsername string `json:"keybaseUsername"`
	Website         string `json:"website"`
	Details         string `json:"details"`
}

func NewParsedValidatorConfig(
	name string,
	keybaseUsername string,
	website string,
	details string,
) ParsedValidatorConfig {
	return ParsedValidatorConfig{
		Name:            name,
		KeybaseUsername: keybaseUsername,
		Website:         website,
		Details:         details,
	}
}
