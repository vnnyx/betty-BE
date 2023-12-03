package dbutils

import "database/sql"

func CommitOrRollback(tx *sql.Tx, err error) error {
	// If the transaction was already committed or rolled back, return the error
	if tx == nil {
		return sql.ErrTxDone
	}

	if err != nil {
		// If there is an error, rollback the transaction and return the error
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}
