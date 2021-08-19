package upgradable_loader

type UpgradeableLoaderInstruction uint32

const (
	InitializeBuffer UpgradeableLoaderInstruction = iota
	Write
	DeployWithMaxDataLen
	Upgrade
	SetAuthority
	Close
)
