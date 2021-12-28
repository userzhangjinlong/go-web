package Config

import "sync"

type Configer interface {
	GetInstance() *config
	SetConfig(map[string]map[string]string) bool
	GetConfig() map[string]map[string]string
}

type config struct {
	conf map[string]map[string]string
}

var (
	once sync.Once
	instance *config
)

func (this *config) GetInstance() *config  {
	once.Do(func() {
		instance = &config{
			map[string]map[string]string{},
		}
	})

	return instance
}

func (this *config) SetConfig(configMap map[string]map[string]string) (result bool) {
	this.conf = configMap
	return true
}

func (this *config) GetConfig () map[string]map[string]string {
	return this.conf
}

