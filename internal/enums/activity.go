package enums

type ActivityField int

const (
	_ ActivityField = iota
	ActivityMenuIngredient
	ActivityMenu
	ActivityIngredient
	ActivityUser
	ActivityTransaction
	ActivityRole
	ActivityVariant
	ActivityVariantMenu
	ActivityCategory
	ActivityMenuCategory
	ActivityCompany
	ActivityUserGroup
	ActivityFranchise
	ActivityAttachmentFile
)

type ActivityAction string

const (
	ActivityActionCreate       ActivityAction = "create"
	ActivityActionUpdate       ActivityAction = "update"
	ActivityActionDelete       ActivityAction = "delete"
	ActivityActionLogin        ActivityAction = "login"
	ActivityActionLogout       ActivityAction = "logout"
	ActivityActionAuthenticate ActivityAction = "authenticate"
)
