package migrator

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_retrieveAccounts(t *testing.T) {
	testDir := testMigrationDataDirPath("battlefield-snapshot")
	if !folderExists(testDir) {
		t.Skipf("Folder %q does not exist, skipping account retrieve test", testDir)
		return
	}

	i := &importer{
		common: common{dataDir: testDir},
		logger: zap.NewNop(),
	}

	accounts, err := i.retrieveAccounts(func(account *Account) error {
		return nil
	})
	require.NoError(t, err)

	expectedAccounts := map[string]bool{
		"battlefeeld4": false,
		"battlefield":  false,
		"battlefield1": true,
		"battlefield2": false,
		"battlefield3": true,
		"battlefield4": false,
		"battlefield5": false,
		"zswhq":        true,
		"zswhq.bpay":   false,
		"zswhq.msig":   true,
		"zswhq.ram":    false,
		"zswhq.token":  true,
		"zswhq.null":   false,
		"zswhq.prods":  false,
		"eosio2":       false,
		"eosio3":       false,
		"zswhq.names":  false,
		"zswhq.ramfee": false,
		"zswhq.saving": false,
		"zswhq.stake":  false,
		"zswhq.vpay":   false,
		"notified1":    false,
		"notified2":    true,
		"notified3":    false,
		"notified4":    false,
		"notified5":    false,
		"zzzzzzzzzzzz": false,
	}

	for _, account := range accounts {
		if _, found := expectedAccounts[account.name]; !found {
			assert.Fail(t, "Unable to find account in expected account list", "Account %q is not in expected account list", account.name)
		}

		assert.Equal(t, expectedAccounts[account.name], account.hasCode)
	}
}

func Test_walkScopes(t *testing.T) {
	testDir := testMigrationDataDirPath("battlefield-snapshot")
	if !folderExists(testDir) {
		t.Skipf("Folder %q does not exist, skipping scope walking test", testDir)
		return
	}
	scopes := []string{}
	accountPath, err := newAccountPath(testDir, "zswhq.token")
	require.NoError(t, err)

	walkScopes(fmt.Sprintf("%s/tables/accounts", accountPath), func(scope string) error {
		scopes = append(scopes, scope)
		return nil
	})

	assert.ElementsMatch(t, []string{
		"battlefeeld4",
		"battlefield1",
		"battlefield3",
		"zswhq",
		"zswhq.ram",
		"zswhq.ramfee",
		"zswhq.stake",
		"notified1",
		"notified2",
		"notified3",
		"notified4",
	}, scopes)

}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return info.IsDir()
}
