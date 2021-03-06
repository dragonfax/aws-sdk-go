//Package support provides gucumber integration tests suppport.
package support

import (
	"github.com/dragonfax/aws-sdk-go/internal/features/shared"
	"github.com/dragonfax/aws-sdk-go/service/support"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@support", func() {
		World["client"] = support.New(nil)
	})
}
