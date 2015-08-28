//Package swf provides gucumber integration tests suppport.
package swf

import (
	"github.com/dragonfax/aws-sdk-go/internal/features/shared"
	"github.com/dragonfax/aws-sdk-go/service/swf"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@swf", func() {
		World["client"] = swf.New(nil)
	})
}
