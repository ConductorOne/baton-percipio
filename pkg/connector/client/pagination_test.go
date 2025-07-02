package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	ctx := context.Background()
	
	t.Run("GetUserNextToken", func(t *testing.T) {
		testCases := []struct {
			message  string
			offset   int
			limit    int
			total    int
			expected string
		}{
			{
				message:  "next page",
				offset:   80,
				limit:    10,
				total:    95,
				expected: "{\"offset\":90}",
			},
			{
				message:  "no more results",
				offset:   0,
				limit:    100,
				total:    100,
				expected: "",
			},
			{
				message:  "standard pagination",
				offset:   0,
				limit:    100,
				total:    200,
				expected: "{\"offset\":100}",
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.message, func(t *testing.T) {
				actual := GetUserNextToken(
					ctx,
					testCase.offset,
					testCase.limit,
					testCase.total,
				)
				require.Equal(t, testCase.expected, actual)
			})
		}
	})

	t.Run("GetContentNextToken", func(t *testing.T) {
		testCases := []struct {
			message         string
			currentOffset   int
			limit           int
			finalOffset     int
			pagingRequestId string
			expected        string
		}{
			{
				message:         "next page with pagingRequestId",
				currentOffset:   0,
				limit:           100,
				finalOffset:     200,
				pagingRequestId: "example",
				expected:        "{\"offset\":100,\"pagingRequestId\":\"example\",\"finalOffset\":200}",
			},
			{
				message:         "no more results",
				currentOffset:   100,
				limit:           100,
				finalOffset:     200,
				pagingRequestId: "example",
				expected:        "{\"offset\":200,\"pagingRequestId\":\"example\",\"finalOffset\":200}",
			},
			{
				message:         "pagination complete - beyond final offset",
				currentOffset:   200,
				limit:           100,
				finalOffset:     200,
				pagingRequestId: "example",
				expected:        "",
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.message, func(t *testing.T) {
				actual := GetContentNextToken(
					ctx,
					testCase.currentOffset,
					testCase.limit,
					testCase.finalOffset,
					testCase.pagingRequestId,
				)
				require.Equal(t, testCase.expected, actual)
			})
		}
	})
}
