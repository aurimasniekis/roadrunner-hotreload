package hotreload

import (
	"github.com/spiral/roadrunner/service"
	"time"
)

type Config struct {
	Enable bool
	Paths  []string
	Files  string
	Tick   *time.Duration
}

// Hydrate must populate Config values using given Config source. Must return error if Config is not valid.
func (c *Config) Hydrate(cfg service.Config) error {
	if err := cfg.Unmarshal(c); err != nil {
		return err
	}

	if c.Paths == nil || len(c.Paths) < 1 {
		c.Paths = []string{
			".",
		}
	}
	if len(c.Files) == 0 {
		c.Files = "*.php"
	}
	if c.Tick == nil {
		refresh := time.Duration(500)
		c.Tick = &refresh
	}

	return nil
}
