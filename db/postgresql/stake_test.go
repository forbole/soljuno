package postgresql_test

import "time"

func (suite *DbTestSuite) TestSaveStake() {
	type StakeRow struct {
		Address    string `db:"address"`
		Slot       uint64 `db:"slot"`
		Staker     string `db:"staker"`
		Withdrawer string `db:"withdrawer"`
		State      string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     StakeRow
		expected StakeRow
	}{
		{
			name: "initialize the data",
			data: StakeRow{
				"address", 1, "staker", "withdrawer", "initialized",
			},
			expected: StakeRow{
				"address", 1, "staker", "withdrawer", "initialized",
			},
		},
		{
			name: "update with lower slot",
			data: StakeRow{
				"address", 0, "pre_staker", "withdrawer", "initialized",
			},
			expected: StakeRow{
				"address", 1, "staker", "withdrawer", "initialized",
			},
		},
		{
			name: "update with same slot",
			data: StakeRow{
				"address", 1, "curr_staker", "withdrawer", "initialized",
			},
			expected: StakeRow{
				"address", 1, "curr_staker", "withdrawer", "initialized",
			},
		},
		{
			name: "update with higher slot",
			data: StakeRow{
				"address", 2, "new_staker", "withdrawer", "initialized",
			},
			expected: StakeRow{
				"address", 2, "new_staker", "withdrawer", "initialized",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveStake(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Staker,
				tc.data.Withdrawer,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []StakeRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM stake")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveStakeLockup() {
	type LockupRow struct {
		Address       string    `db:"address"`
		Slot          uint64    `db:"slot"`
		Custodian     string    `db:"custodian"`
		Epoch         uint64    `db:"epoch"`
		UnixTimestamp time.Time `db:"unix_timestamp"`
	}

	testCases := []struct {
		name     string
		data     LockupRow
		expected LockupRow
	}{
		{
			name: "initialize the data",
			data: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
			expected: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
		},
		{
			name: "update with lower slot",
			data: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
			expected: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
		},
		{
			name: "update with same slot",
			data: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
			expected: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
		},
		{
			name: "update with higher slot",
			data: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
			expected: LockupRow{
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveStakeLockup(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Custodian,
				tc.data.Epoch,
				tc.data.UnixTimestamp.Unix(),
			)
			suite.Require().NoError(err)
			// Verify the data
			rows := []LockupRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM stake_lockup")

			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.T().Log(tc.data.UnixTimestamp)
			suite.T().Log(rows[0].UnixTimestamp)
			suite.Require().Equal(tc.expected.Address, rows[0].Address)
			suite.Require().Equal(tc.expected.Slot, rows[0].Slot)
			suite.Require().Equal(tc.expected.Custodian, rows[0].Custodian)
			suite.Require().Equal(tc.expected.Epoch, rows[0].Epoch)
			suite.Require().Equal(tc.expected.UnixTimestamp, rows[0].UnixTimestamp)
		})
	}
}

func (suite *DbTestSuite) TestStakeDelegation() {
	type DelegationRow struct {
		Address           string  `db:"address"`
		Slot              uint64  `db:"slot"`
		ActivationEpoch   uint64  `db:"activation_epoch"`
		DeactivationEpoch uint64  `db:"deactivation_epoch"`
		Stake             uint64  `db:"stake"`
		Voter             string  `db:"voter"`
		Rate              float64 `db:"warmup_cooldown_rate"`
	}

	testCases := []struct {
		name     string
		data     DelegationRow
		expected DelegationRow
	}{
		{
			name: "initialize the data",
			data: DelegationRow{
				"address", 1, 0, 0, 100, "voter", 0.25,
			},
			expected: DelegationRow{
				"address", 1, 0, 0, 100, "voter", 0.25,
			},
		},
		{
			name: "update with lower slot",
			data: DelegationRow{
				"address", 0, 0, 0, 100, "pre_voter", 0.25,
			},
			expected: DelegationRow{
				"address", 1, 0, 0, 100, "voter", 0.25,
			},
		},
		{
			name: "update with same slot",
			data: DelegationRow{
				"address", 1, 0, 0, 100, "curr_voter", 0.25,
			},
			expected: DelegationRow{
				"address", 1, 0, 0, 100, "curr_voter", 0.25,
			},
		},
		{
			name: "update with higher slot",
			data: DelegationRow{
				"address", 2, 0, 0, 100, "new_voter", 0.25,
			},
			expected: DelegationRow{
				"address", 2, 0, 0, 100, "new_voter", 0.25,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveStakeDelegation(
				tc.data.Address,
				tc.data.Slot,
				tc.data.ActivationEpoch,
				tc.data.DeactivationEpoch,
				tc.data.Stake,
				tc.data.Voter,
				tc.data.Rate,
			)
			suite.Require().NoError(err)
			// Verify the data
			rows := []DelegationRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM stake_delegation")

			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
