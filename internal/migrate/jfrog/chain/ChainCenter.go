package chain

type ChainCenter struct {
}

func (cc ChainCenter) AssembleChainHandle() Handle {
	jfHandle := &JfrogHandle{}
	npmHandle := &NpmHandle{}
	jfHandle.handle = npmHandle

	return jfHandle
}
