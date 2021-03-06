//Package devicefarm provides gucumber integration tests suppport.
package devicefarm

import (
	"github.com/dragonfax/aws-sdk-go/aws"
	"github.com/dragonfax/aws-sdk-go/internal/features/shared"
	"github.com/dragonfax/aws-sdk-go/service/devicefarm"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@devicefarm", func() {
		// FIXME remove custom region
		World["client"] = devicefarm.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
