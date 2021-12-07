package keybase_test

import (
	"testing"

	"github.com/forbole/soljuno/apis/keybase"
	"github.com/stretchr/testify/require"
)

func TestGetAvatarURL(t *testing.T) {
	client := keybase.NewClient()
	url, err := client.GetAvatarURL("forbole")
	require.NoError(t, err)
	require.Equal(t, "https://s3.amazonaws.com/keybase_processed_uploads/f5b0771af36b2e3d6a196a29751e1f05_360_360.jpeg", url)
}
