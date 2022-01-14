package epoch

func (m *Module) RegisterService(service EpochService) {
	m.services = append(m.services, service)
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
