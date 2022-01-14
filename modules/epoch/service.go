package epoch

func (m *Module) RegisterService(services ...EpochService) {
	for _, s := range services {
		m.services = append(m.services, s)
	}
}

func (m *Module) RunServices() error {
	for _, service := range m.services {
		err := service.ExecEpoch(m.epoch)
		if err != nil {
			return err
		}
	}
	return nil
}
