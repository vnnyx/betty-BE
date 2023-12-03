package activity

const (
	// InsertActivityQuery is a query to insert activity data into database.
	InsertActivityQuery = `
		INSERT INTO activity (
			activity_field_id,
			action,
			user_id
		)
		VALUES ($1, $2, $3) RETURNING id;
	`

	InsertActivityDetailQuery = `
		INSERT INTO activity_detail (
			activity_id,
			changed_id,
			old_value,
			new_value
		)
		VALUES ($1, $2, $3, $4)
	`
)
