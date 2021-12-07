package keybase

type Client interface {
	// GetAvatarURL returns the avatar URL from the given identity.
	// If no identity is found, it returns an empty string instead.
	GetAvatarURL(identity string) (string, error)
}
