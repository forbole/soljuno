package upgradable_loader

type UpgradeableLoaderInstruction uint16

const (
	InitializeBuffer UpgradeableLoaderInstruction = iota
	Write
	DeployWithMaxDataLen
	Upgrade
	SetAuthority
	Close
)
