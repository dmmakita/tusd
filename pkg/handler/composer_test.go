package handler_test

import (
	"github.com/dmmakita/tusd/pkg/filestore"
	"github.com/dmmakita/tusd/pkg/handler"
	"github.com/dmmakita/tusd/pkg/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := handler.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	config := handler.Config{
		StoreComposer: composer,
	}

	_, _ = handler.NewHandler(config)
}
