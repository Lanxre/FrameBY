package frender

import "context"

type Render interface {
	Name() string
	Ext() string
	Render(
		ctx context.Context,
		outputPath string,
	) error
}
