package bpfloader

type LoaderInstruction uint16

const (
	Write LoaderInstruction = iota
	Finalize
)
