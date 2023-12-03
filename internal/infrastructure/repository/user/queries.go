package user

const (
	InsertUserQuery = `
		INSERT INTO users (
			is_super_admin,
			is_admin,
			email,
			password,
			fullname,
			phone_number,
			company_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`

	GetUserByEmailQuery = `
		SELECT
			id,
			is_super_admin,
			is_admin,
			email,
			password,
			refresh_token,
			fullname,
			phone_number,
			company_id,
			franchise_id
		FROM users WHERE email = $1 AND is_active = true
	`

	GetRoleScopeByUserIDQuery = `
		SELECT
			array_agg(ug.id),
			array_agg(r.id),
			array_agg(r.name),
			array_agg(DISTINCT s.id),
			array_agg(DISTINCT s.name)
		FROM user_group ug
		JOIN users u on ug.user_id = u.id
		JOIN role r on ug.role_id = r.id
		JOIN role_scope rs on r.id = rs.role_id
		JOIN scope s on rs.scope_id = s.id
		WHERE u.id = $1 AND u.is_active = true
		GROUP BY u.id
	`

	GetUserByIDQuery = `
		SELECT
			id,
			is_super_admin,
			is_admin,
			email,
			password,
			refresh_token,
			fullname,
			phone_number,
			company_id,
			franchise_id
		FROM users WHERE id = $1 AND is_active = true
	`

	UpdateRefreshTokenQuery = `
		UPDATE users SET refresh_token = $1 WHERE id = $2
	`

	GetDetailUserByIDQuery = `
		SELECT
			u.id,
			u.is_super_admin,
			u.is_admin,
			u.email,
			u.password,
			u.refresh_token,
			u.fullname,
			u.phone_number,
			c.id,
			c.brand_name,
			c.address_1,
			c.address_2,
			ct.id,
			ct.name,
			co.id,
			co.name,
			json_agg(
			    json_build_object(
					'id', f.id,
					'name', f.name
				)
            )
		FROM users u
		JOIN company c on u.company_id = c.id
		JOIN city ct on c.city_id = ct.id
		JOIN country co on ct.country_id = co.id
		JOIN franchise f on c.id = f.company_id
		WHERE u.id = $1 AND u.is_active = true
		GROUP BY u.id, c.id, ct.id, co.id;
	`

	GetUserByRefreshTokenQuery = `
		SELECT
			id,
			is_super_admin,
			is_admin,
			email,
			password,
			refresh_token,
			fullname,
			phone_number,
			company_id,
			franchise_id
		FROM users WHERE refresh_token = $1 AND is_active = true
	`
)
