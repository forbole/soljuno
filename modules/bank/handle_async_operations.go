package bank

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		m.consumeTasks()
	}
}

func (m *Module) consumeTasks() {
	task := <-m.tasks
	task()
}
