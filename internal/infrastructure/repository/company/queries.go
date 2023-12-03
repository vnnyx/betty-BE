package company

const (
	// InsertCompany is a query to insert company data into database.
	InsertCompanyQuery = `
		WITH inserted_company AS (
			INSERT INTO company (
				brand_name,
				franchise_name,
				address_1,
				address_2,
				city_id,
				country_id,
				postal_code
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *
		)
		SELECT
			ic.id,
			ic.brand_name,
			ic.franchise_name,
			ic.address_1,
			ic.address_2,
			c.*,
			co.*,
			ic.postal_code
		FROM inserted_company ic
		JOIN city c ON c.id = ic.city_id
		JOIN country co ON co.id = ic.country_id
	`

	InsertFranchiseQuery = `
		INSERT INTO franchise (
			name,
			company_id
		)
		VALUES ($1, $2) RETURNING id, name, company_id;
	`
)
