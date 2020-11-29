package model

import (
	"log"

	"github.com/fullstacktf/fitness-backend/storage"
	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	Description string        `gorm:"column:description;type:varchar(20);unique"`
	Permissions []*Permission `gorm:"many2many:roles_permissions;"`
}

// TableName Function to change the name of a table.
func (r *Role) TableName() string {
	return "roles"
}

// PopulateRoles Function to populate the Role model
func PopulateRoles() {
	adminRole := Role{
		Description: "Admin",
	}

	trainerRole := Role{
		Description: "Trainer",
	}

	userRole := Role{
		Description: "User",
	}

	systemRole := Role{
		Description: "System",
	}

	roles := []Role{
		adminRole,
		trainerRole,
		userRole,
		systemRole,
	}

	storage.DB.Create(&roles)
}

// CreateRole Creates a role
func (r *Role) CreateRole() string {
	return "Create a Role"
}

// GetRole Function to return a specific role. Not for the API controller
func GetRole(id uint8) *Role {
	role := Role{}
	result := storage.DB.Find(&role, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &role
}

// GetRole Gets a role
func (r *Role) GetRole() string {
	return "getRole"
}

// GetRoles Gets all roles
func (r *Role) GetRoles() string {
	return "Get all Roles"
}

// UpdateRole Updates a role
func (r *Role) UpdateRole() string {
	return "Update a Role"
}

// DeleteRole Deletes a role
func (r *Role) DeleteRole() string {
	return "Delete a Role"
}

// AssociatePermissions Function to associate permissions to roles
func AssociatePermissions() {
	adminPermissions := []*Permission{
		GetPermission(1),
		GetPermission(2),
		GetPermission(3),
		GetPermission(4),
		GetPermission(5),
		GetPermission(6),
		GetPermission(7),
		GetPermission(8),
		GetPermission(9),
		GetPermission(10),
		GetPermission(11),
		GetPermission(12),
		GetPermission(13),
		GetPermission(14),
		GetPermission(15),
		GetPermission(16),
		GetPermission(17),
		GetPermission(18),
		GetPermission(19),
		GetPermission(20),
		GetPermission(21),
		GetPermission(22),
		GetPermission(23),
		GetPermission(24),
		GetPermission(25),
		GetPermission(26),
		GetPermission(27),
		GetPermission(28),
		GetPermission(29),
		GetPermission(30),
		GetPermission(31),
		GetPermission(32),
		GetPermission(33),
		GetPermission(34),
		GetPermission(35),
		GetPermission(36),
		GetPermission(37),
		GetPermission(38),
		GetPermission(39),
		GetPermission(40),
		GetPermission(41),
		GetPermission(42),
		GetPermission(43),
		GetPermission(44),
		GetPermission(45),
	}

	trainerPermissions := []*Permission{
		GetPermission(9),
		GetPermission(10),
		GetPermission(11),
		GetPermission(12),
		GetPermission(13),
		GetPermission(14),
		GetPermission(15),
		GetPermission(16),
		GetPermission(17),
		GetPermission(18),
		GetPermission(19),
		GetPermission(20),
		GetPermission(21),
		GetPermission(22),
		GetPermission(23),
		GetPermission(24),
		GetPermission(25),
		GetPermission(26),
		GetPermission(27),
		GetPermission(28),
		GetPermission(29),
		GetPermission(30),
		GetPermission(31),
		GetPermission(32),
		GetPermission(33),
		GetPermission(34),
		GetPermission(35),
		GetPermission(36),
		GetPermission(37),
		GetPermission(38),
		GetPermission(39),
		GetPermission(40),
		GetPermission(41),
		GetPermission(42),
		GetPermission(43),
		GetPermission(44),
		GetPermission(45),
	}

	userPermissions := []*Permission{
		GetPermission(10),
		GetPermission(11),
	}

	systemPermissions := []*Permission{}

	storage.DB.Model(GetRole(1)).Association("Permissions").Append(adminPermissions)
	storage.DB.Model(GetRole(2)).Association("Permissions").Append(trainerPermissions)
	storage.DB.Model(GetRole(3)).Association("Permissions").Append(userPermissions)
	storage.DB.Model(GetRole(4)).Association("Permissions").Append(systemPermissions)
}
