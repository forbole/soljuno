package votestatus

func (m *Module) ExecEpoch(epoch uint64) error {
	return UpdateValidatorSkipRates(epoch-1, m.db, m.client)
}
