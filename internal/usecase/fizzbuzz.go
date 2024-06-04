package usecase

import (
	"errors"

	"github.com/dysdimas/internal/domain"
)

// FizzBuzzUsecase defines the use case interface for FizzBuzz operations.
type FizzBuzzUsecase interface {
	RangeFizzBuzz(from, to int) ([]string, error)
}

type fizzBuzzUsecase struct{}

// NewFizzBuzzUsecase creates a new FizzBuzz use case.
func NewFizzBuzzUsecase() FizzBuzzUsecase {
	return &fizzBuzzUsecase{}
}

// RangeFizzBuzz returns the FizzBuzz results for a range of numbers.
func (u *fizzBuzzUsecase) RangeFizzBuzz(from, to int) ([]string, error) {
	if from > to || to-from > 100 {
		return nil, errors.New("invalid range parameters")
	}

	results := make([]string, to-from+1)
	for i := from; i <= to; i++ {
		results[i-from] = domain.SingleFizzBuzz(i)
	}
	return results, nil
}
