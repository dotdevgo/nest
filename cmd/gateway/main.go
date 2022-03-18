package main

import (
	"github.com/dotdevgo/nest/examples/app/dataform"
	nest "github.com/dotdevgo/nest/pkg/core"
	"github.com/dotdevgo/nest/pkg/core/kernel"
)

func main() {
	e := nest.New(
		kernel.Provider(),
		dataform.Provider(),
	)

	// Router
	e.InvokeFn(dataform.Router)

	e.Logger.Fatal(e.Start(":1323"))
}
