package common

import "github.com/teris-io/shortid"

func GenerateId() (string, error) {
	return shortid.Generate()
}

func IdProcess(id *string) (*string, *AppError) {
	if id != nil && len(*id) != 0 {
		if len(*id) > 9 {
			return nil, ErrIdIsTooLong()
		}
		return id, nil
	} else {
		idGenerated, err := GenerateId()
		if err != nil {
			return nil, ErrInternal(err)
		}

		return &idGenerated, nil
	}
}
