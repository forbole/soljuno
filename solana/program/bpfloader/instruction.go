package bpfloader

type LoaderInstruction uint32

const (
	Write LoaderInstruction = iota
	Finalize
)
