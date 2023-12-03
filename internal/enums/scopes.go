package enums

type Scope string

const (
	IngredientReadAccess       Scope = "ingredient:read"
	IngredientCreateAccess     Scope = "ingredient:create"
	IngredientUpdateAccess     Scope = "ingredient:update"
	IngredientDeleteAccess     Scope = "ingredient:delete"
	MenuReadAccess             Scope = "menu:read"
	MenuCreateAccess           Scope = "menu:create"
	MenuUpdateAccess           Scope = "menu:update"
	MenuDeleteAccess           Scope = "menu:delete"
	MenuIngredientReadAccess   Scope = "menu_ingredient:read"
	MenuIngredientCreateAccess Scope = "menu_ingredient:create"
	MenuIngredientUpdateAccess Scope = "menu_ingredient:update"
	MenuIngredientDeleteAccess Scope = "menu_ingredient:delete"
	FranchiseReadAccess        Scope = "franchise:read"
	FranchiseCreateAccess      Scope = "franchise:create"
	FranchiseUpdateAccess      Scope = "franchise:update"
	FranchiseDeleteAccess      Scope = "franchise:delete"
	EmployeeReadAccess         Scope = "employee:read"
	EmployeeCreateAccess       Scope = "employee:create"
	EmployeeUpdateAccess       Scope = "employee:update"
	EmployeeDeleteAccess       Scope = "employee:delete"
)
