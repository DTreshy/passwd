package passwd_test

import (
	"testing"

	"github.com/DTreshy/passwd"
	"github.com/stretchr/testify/require"
)

func TestHashAndCheck(t *testing.T) {
	hash, err := passwd.Hash("myverystrongpassword123!")
	require.NoError(t, err)

	resultTrue := passwd.Check("myverystrongpassword123!", hash)
	require.True(t, resultTrue)

	resultFalse := passwd.Check("wrongpassword", hash)
	require.False(t, resultFalse)

	_, err = passwd.HashWithCost("wrong cost test", 50)
	require.Error(t, err)
}
