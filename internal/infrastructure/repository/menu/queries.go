package menu

const (
	InsertIngredientQuery = `
		INSERT INTO ingredient (
			name,
			minimum_stock,
			recent_stock,
			price,
			unit_id,
			franchise_id
		)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
	`

	UpdatePopularityCountQuery = `
		UPDATE menu SET popularity_count = popularity_count + 1 WHERE id = $1;
	`

	InsertMenuQuery = `
		INSERT INTO menu (
			matchcode,
			name,
			price,
			description,
			franchise_id
		)
		VALUES ($1, $2, $3, $4, $5) RETURNING id, matchcode, name, description, price, stock, franchise_id;
	`

	InsertMenuIngredientQuery = `
		WITH inserted_menu_ingredient AS (
			INSERT INTO menu_ingredient (
				ingredient_id,
				menu_id,
				quantity
			)
			VALUES ($1, $2, $3) RETURNING *
		)
		SELECT
			imi.id,
			imi.ingredient_id,
			i.name,
			imi.menu_id,
			imi.quantity,
			u.id AS unit_id,
			u.is_international,
			u.conversion_rate,
			u.conversion_id,
			i.recent_stock
		FROM inserted_menu_ingredient imi
		JOIN ingredient i ON i.id = imi.ingredient_id
		JOIN unit u ON u.id = i.unit_id
	`

	InsertUnitQuery = `
		INSERT INTO unit (
			name,
			is_international,
			conversion_rate,
			conversion_id,
			franchise_id
		)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`

	GetMenuIngredientByMenuIDQuery = `
		SELECT
			mi.id,
			mi.ingredient_id,
			i.name,
			mi.menu_id,
			mi.quantity,
			u.id AS unit_id,
			u.is_international,
			u.conversion_rate,
			u.conversion_id,
			i.recent_stock
		FROM menu_ingredient mi
		JOIN ingredient i ON i.id = mi.ingredient_id
		JOIN unit u ON u.id = i.unit_id
		WHERE mi.menu_id = $1;
	`

	InsertCategoryQuery = `
		INSERT INTO category (
			name,
			color,
			franchise_id
		)
		VALUES ($1, $2, $3) RETURNING id, name, color;
	`

	InsertMenuCategoryQuery = `
		INSERT INTO menu_category (
			menu_id,
			category_id
		)
		VALUES ($1, $2) RETURNING id;
	`
)
