package history

type HistroyService interface {
	Name() string
	ExecHistory() error
	Cron() string
}
