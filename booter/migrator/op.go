package migrator

import (
	"fmt"

	"github.com/invisible-train-40/eosio-boot/config"
	bootops "github.com/invisible-train-40/eosio-boot/ops"
	"github.com/zhongshuwen/zswchain-go/ecc"
)

func init() {
	bootops.Register("migration.inject", &OpMigration{})
}

type OpMigration struct {
	DataDir string `json:"data_dir"`
}

func (op *OpMigration) RequireValidation() bool {
	return false
}

func (op *OpMigration) Actions(opPubkey ecc.PublicKey, c *config.OpConfig, in chan interface{}) error {
	impt := newImporter(opPubkey, op.DataDir, in, c.Logger)

	err := impt.init()
	if err != nil {
		return fmt.Errorf("faile to initialize migrator: %w", err)
	}

	err = impt.inject()
	if err != nil {
		return fmt.Errorf("unable to inject data on chain: %w", err)
	}

	return nil
}
