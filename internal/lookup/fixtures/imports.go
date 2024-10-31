package fixtures

import (
	"fmt"
	mlog "log"
	"log/slog"

	. "github.com/davecgh/go-spew/spew"
	_ "github.com/davecgh/go-spew/spew"

	"github.com/davecgh/go-spew/spew"
)

func SomeTestFunction() {
	fmt.Println("msg")
	mlog.Println("msg")
	_ = slog.Logger{}

	_ = Config

	_ = anythingelse.Exported{}

	_ = spew.Config
}
