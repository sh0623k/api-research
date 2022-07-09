package errors

import "github.com/pkg/errors"

func CombineErrors(originalErr, secondErr error) error {
	if originalErr != nil && secondErr != nil {
		return errors.Wrap(originalErr, secondErr.Error())
	} else if secondErr != nil {
		return secondErr
	} else if originalErr != nil {
		return originalErr
	}
	return nil
}
