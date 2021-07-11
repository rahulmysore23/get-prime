package seive

import "testing"

func TestCheckPrime(t *testing.T) {
	s := NewSeive(100)

	assertCorrectMessage := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("checking if 3 is prime", func(t *testing.T) {
		got := s.CheckPrime(3)
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("checking if 2 is prime", func(t *testing.T) {
		got := s.CheckPrime(2)
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("checking if 44 is prime", func(t *testing.T) {
		got := s.CheckPrime(44)
		want := false
		assertCorrectMessage(t, got, want)
	})

}

func TestGetPrimeTrue(t *testing.T) {
	s := NewSeive(100)

	assertCorrectMessage := func(t testing.TB, gotPrime, wantPrime int64, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		if gotPrime != wantPrime {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Get highest prime number lesser than 3", func(t *testing.T) {
		gotPrime, got := s.GetPrime(3)
		wantPrime, want := int64(2), true
		assertCorrectMessage(t, gotPrime, wantPrime, got, want)
	})

	t.Run("Get highest prime number lesser than 2", func(t *testing.T) {
		gotPrime, got := s.GetPrime(2)
		wantPrime, want := int64(0), true
		assertCorrectMessage(t, gotPrime, wantPrime, got, want)
	})

	t.Run("Get highest prime number lesser than 55", func(t *testing.T) {
		gotPrime, got := s.GetPrime(55)
		wantPrime, want := int64(53), false
		assertCorrectMessage(t, gotPrime, wantPrime, got, want)
	})
}
