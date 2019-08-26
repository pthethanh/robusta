package closeutil

import "github.com/pthethanh/robusta/internal/pkg/log"

type (
	CloseFunc = func() error
	// Closer is a helper for closeing up underlying resources
	// base on the order of the close functions
	Closer struct {
		funcs []CloseFunc
	}
)

// NewCloser return new closer
func NewCloser() *Closer {
	return &Closer{
		funcs: make([]CloseFunc, 0),
	}
}

// Close close invokes all close functions by order
func (c Closer) Close() error {
	log.Infof("closing %d resources", len(c.funcs))
	success := 0
	failed := 0
	for _, close := range c.funcs {
		if err := close(); err != nil {
			failed++
			log.Errorf("failed to close resource: %v", err)
		}
		success++
	}
	log.Infof("finished closing %d resources, %d succeed, %d failed", len(c.funcs), success, failed)
	return nil
}

func (c *Closer) Append(nc *Closer) {
	c.funcs = append(c.funcs, nc.funcs...)
}

func (c *Closer) Add(f CloseFunc) {
	c.funcs = append(c.funcs, f)
}

func (c *Closer) AddFunc(f func()) {
	c.funcs = append(c.funcs, func() error {
		f()
		return nil
	})
}
