//Package emr provides gucumber integration tests suppport.
package emr

import (
	"github.com/dragonfax/aws-sdk-go/internal/features/shared"
	"github.com/dragonfax/aws-sdk-go/service/emr"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@emr", func() {
		World["client"] = emr.New(nil)
	})
}
