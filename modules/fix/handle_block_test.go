package fix_test

import (
	"github.com/forbole/soljuno/modules/fix"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestEnqueueMissingSlots() {
	testCases := []struct {
		name     string
		height   uint64
		start    uint64
		end      uint64
		expected []uint64
	}{
		{
			name:     "no missing height does not enqueue anything",
			height:   0,
			expected: []uint64{},
		},
		{
			name:     "no range returning does not enqueue anything",
			height:   1,
			end:      0,
			expected: []uint64{},
		},
		{
			name:     "enqueue missing height properly",
			height:   1,
			start:    0,
			end:      1,
			expected: []uint64{1},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			db := NewMockDb(tc.height, tc.start, tc.end)
			queue := types.NewQueue(1)
			fix.EnqueueMissingSlots(db, queue, &MockClient{}, tc.start, tc.end)
			suite.Require().Len(queue, len(tc.expected))
			if len(queue) == 0 {
				return
			}
			slots := []uint64{<-queue}
			suite.Require().Equal(tc.expected, slots)
		})
	}
}
