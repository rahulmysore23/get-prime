package prime

type Prime interface {
	CheckPrime(number int64) bool
	GetPrime(number int64) (int64, bool)
}
