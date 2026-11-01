package auth

type CredentialsProvider uint8

const (
	CredentialsProviderEpic CredentialsProvider = iota
	CredentialsProviderEpicLauncher
	CredentialsProviderMSStore
)

type CredentialsEpic struct {
	RefreshToken string `json:"refreshToken"`
}
