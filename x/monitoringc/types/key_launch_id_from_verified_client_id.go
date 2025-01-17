package types

const (
	// LaunchIDFromVerifiedClientIDKeyPrefix is the prefix to retrieve all LaunchIDFromVerifiedClientID
	LaunchIDFromVerifiedClientIDKeyPrefix = "LaunchIDFromVerifiedClientID/value/"
)

// LaunchIDFromVerifiedClientIDKey returns the store key to retrieve a LaunchIDFromVerifiedClientID from the index fields
func LaunchIDFromVerifiedClientIDKey(clientID string) []byte {
	return []byte(clientID + "/")
}
