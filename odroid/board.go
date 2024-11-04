//go:build linux

// Package odroid implements a odroid based board.
package odroid

import (
	"context"
	"errors"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "odroid"

// Model for viam supported hardkernel odroid board.
var Model = resource.NewModel("viam", "hardkernel", "odroid")

func init() {
	gpioMappings, err := genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		logging.Global().Debugw("error getting odroid GPIO board mapping", "error", err)
	}

	resource.RegisterComponent(
		board.API,
		Model,
		resource.Registration[board.Board, *genericlinux.Config]{
			Constructor: func(
				ctx context.Context,
				_ resource.Dependencies,
				conf resource.Config,
				logger logging.Logger,
			) (board.Board, error) {
				return genericlinux.NewBoard(ctx, conf, genericlinux.ConstPinDefs(gpioMappings), logger)
			},
		})
}
