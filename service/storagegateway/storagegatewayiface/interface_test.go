// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package storagegatewayiface_test

import (
	"testing"

	"github.com/dragonfax/aws-sdk-go/service/storagegateway"
	"github.com/dragonfax/aws-sdk-go/service/storagegateway/storagegatewayiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*storagegatewayiface.StorageGatewayAPI)(nil), storagegateway.New(nil))
}
