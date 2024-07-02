package main

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/teran/go-random"
)

func TestCharacterSetByName(t *testing.T) {
	type testCase struct {
		input     string
		expOutput []rune
		expError  error
	}

	tcs := []testCase{
		{
			input:     "numeric",
			expOutput: random.Numeric,
		},
		{
			input:     "special",
			expOutput: random.Special,
		},
		{
			input:     "numeric+special",
			expOutput: append(random.Numeric, random.Special...),
		},
		{
			input:    "",
			expError: errors.New("unexpected set name: ``"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			r := require.New(t)

			res, err := characterSetByName(tc.input)
			if tc.expError != nil {
				r.Error(err)
				r.Equal(tc.expError.Error(), err.Error())
			} else {
				r.NoError(err)
				r.Equal(tc.expOutput, res)
			}
		})
	}
}
