// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package efsiface_test

import (
	"testing"

	"github.com/dragonfax/aws-sdk-go/service/efs"
	"github.com/dragonfax/aws-sdk-go/service/efs/efsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*efsiface.EFSAPI)(nil), efs.New(nil))
}
