package interfaces

type IHasher interface {
	Hahser(text string) (string, error)
	Verify(originalText, hashedText string) bool
}
