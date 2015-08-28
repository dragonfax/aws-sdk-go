//Package elasticloadbalancing provides gucumber integration tests suppport.
package elasticloadbalancing

import (
	"github.com/dragonfax/aws-sdk-go/internal/features/shared"
	"github.com/dragonfax/aws-sdk-go/service/elb"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elasticloadbalancing", func() {
		World["client"] = elb.New(nil)
	})
}
