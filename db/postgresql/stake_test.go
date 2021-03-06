package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveStakeAccount() {

	testCases := []struct {
		name     string
		data     dbtypes.StakeAccountRow
		expected dbtypes.StakeAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewStakeAccountRow(
				"address", 1, "staker", "withdrawer",
			),
			expected: dbtypes.NewStakeAccountRow(
				"address", 1, "staker", "withdrawer",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewStakeAccountRow(
				"address", 0, "pre_staker", "withdrawer",
			),
			expected: dbtypes.NewStakeAccountRow(
				"address", 1, "staker", "withdrawer",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewStakeAccountRow(
				"address", 1, "curr_staker", "withdrawer",
			),
			expected: dbtypes.NewStakeAccountRow(
				"address", 1, "curr_staker", "withdrawer",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewStakeAccountRow(
				"address", 2, "new_staker", "withdrawer",
			),
			expected: dbtypes.NewStakeAccountRow(
				"address", 2, "new_staker", "withdrawer",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveStakeAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.StakeAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM stake_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteStakeAccount() {
	err := suite.database.SaveStakeAccount(
		dbtypes.NewStakeAccountRow(
			"address",
			0,
			"staker",
			"withdrawer",
		),
	)
	suite.Require().NoError(err)
	err = suite.database.SaveStakeLockup(
		dbtypes.NewStakeLockupRow(
			"address",
			0,
			"custodian",
			0,
			0,
		),
	)
	suite.Require().NoError(err)
	err = suite.database.SaveStakeDelegation(
		dbtypes.NewStakeDelegationRow(
			"address",
			0,
			0,
			0,
			1,
			"validator",
			0,
		),
	)
	suite.Require().NoError(err)
	accountRows := []dbtypes.StakeAccountRow{}

	err = suite.database.Sqlx.Select(&accountRows, "SELECT * FROM stake_account")
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1)
	accountRows = nil

	lockupRows := []dbtypes.StakeLockupRow{}
	err = suite.database.Sqlx.Select(&lockupRows, "SELECT * FROM stake_lockup")
	suite.Require().NoError(err)
	suite.Require().Len(lockupRows, 1)
	lockupRows = nil

	delegationRows := []dbtypes.StakeDelegationRow{}
	err = suite.database.Sqlx.Select(&delegationRows, "SELECT * FROM stake_delegation")
	suite.Require().NoError(err)
	suite.Require().Len(delegationRows, 1)
	delegationRows = nil

	err = suite.database.DeleteStakeAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&accountRows, "SELECT * FROM stake_account")
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 0)

	err = suite.database.Sqlx.Select(&lockupRows, "SELECT * FROM stake_lockup")
	suite.Require().NoError(err)
	suite.Require().Len(lockupRows, 0)
	accountRows = nil

	err = suite.database.Sqlx.Select(&delegationRows, "SELECT * FROM stake_delegation")
	suite.Require().NoError(err)
	suite.Require().Len(delegationRows, 0)
}

func (suite *DbTestSuite) TestSaveStakeLockup() {
	testCases := []struct {
		name     string
		data     dbtypes.StakeLockupRow
		expected dbtypes.StakeLockupRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
			expected: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
			expected: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
			expected: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
			expected: dbtypes.NewStakeLockupRow(
				"address", 1, "custodian", 1, time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC).Unix(),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveStakeAccount(
				dbtypes.NewStakeAccountRow(
					tc.data.Address,
					0,
					"staker",
					"withdrawer",
				),
			)
			suite.Require().NoError(err)
			err = suite.database.SaveStakeLockup(tc.data)
			suite.Require().NoError(err)
			// Verify the data
			rows := []dbtypes.StakeLockupRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM stake_lockup")

			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().True(tc.expected.Equal(rows[0]))
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
			err := suite.database.SaveStakeAccount(
				dbtypes.NewStakeAccountRow(
					tc.data.Address,
					0,
					"staker",
					"withdrawer",
				),
			)
			suite.Require().NoError(err)
			err = suite.database.SaveStakeDelegation(
				dbtypes.NewStakeDelegationRow(
					tc.data.Address,
					tc.data.Slot,
					tc.data.ActivationEpoch,
					tc.data.DeactivationEpoch,
					tc.data.Stake,
					tc.data.Voter,
					tc.data.Rate,
				),
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
