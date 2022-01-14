package epoch

type EpochService interface {
	Name() string
	ExecEpoch(epoch uint64) error
}
