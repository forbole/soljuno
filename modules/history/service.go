package history

func (m *Module) RegisterService(services ...HistroyService) {
	for _, s := range services {
		m.services = append(m.services, s)
	}
}

func (m *Module) RunServices() error {
	for _, service := range m.services {
		err := service.ExecHistory()
		if err != nil {
			return err
		}
	}
	return nil
}
