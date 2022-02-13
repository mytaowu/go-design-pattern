package _2_facade

var m *Manager

type Manager struct {
	config *ConfigModel
}

// GetConfigDao 获取配置
func (m *Manager) GetConfigDao() *ConfigModel {
	return m.config
}

// GetManagerInstance 获取Manager实例
func GetManagerInstance() *Manager {
	if m == nil {
		m = &Manager{
			config: &ConfigModel{
				NeedGenPresentation: true,
				NeedGenBusiness:     true,
				NeedGenDAO:          true,
			},
		}
	}

	return m
}
