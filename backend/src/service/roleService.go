package service

import (
	"errors"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateRole Creates a role
func CreateRole(data model.Role) error {
	result := storage.DB.Create(&data)
	return result.Error
}

// GetRole Function to return a specific role. Not for the API controller
func GetRole(id uint8) *model.Role {
	role := model.Role{}
	storage.DB.Find(&role, id)

	permissions := []*model.Permission{}

	storage.DB.Model(&role).Association("Permissions").Find(&permissions)

	role.Permissions = permissions

	return &role
}

// GetRoles Gets all roles
func GetRoles() *[]model.Role {
	roles := []model.Role{}
	storage.DB.Find(&roles)
	return &roles
}

// UpdateRole Updates a role
func UpdateRole(updatedRole model.Role) error {
	output := storage.DB.Save(&updatedRole)
	return output.Error
}

// DeleteRole Deletes a role
func DeleteRole(id uint8) error {
	role := GetRole(id)

	if role.ID == 0 {
		return errors.New("Not Found")
	} else {
		result := storage.DB.Delete(&role)
		return result.Error
	}
}

// PopulateRoles Function to populate the Role model
func PopulateRoles() {
	adminRole := model.Role{
		Description: "Admin",
	}

	trainerRole := model.Role{
		Description: "Trainer",
	}

	userRole := model.Role{
		Description: "User",
	}

	systemRole := model.Role{
		Description: "System",
	}

	roles := []model.Role{
		adminRole,
		trainerRole,
		userRole,
		systemRole,
	}

	storage.DB.Create(&roles)
}

// AssociatePermissions Function to associate permissions to roles
func AssociatePermissions() {
	adminPermissions := []*model.Permission{
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
		GetPermission(46),
		GetPermission(47),
		GetPermission(48),
		GetPermission(49),
		GetPermission(50),
	}

	trainerPermissions := []*model.Permission{
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
		GetPermission(46),
		GetPermission(47),
		GetPermission(48),
		GetPermission(49),
		GetPermission(50),
	}

	userPermissions := []*model.Permission{
		GetPermission(10),
		GetPermission(11),
	}

	systemPermissions := []*model.Permission{}

	storage.DB.Model(GetRole(1)).Association("Permissions").Append(adminPermissions)
	storage.DB.Model(GetRole(2)).Association("Permissions").Append(trainerPermissions)
	storage.DB.Model(GetRole(3)).Association("Permissions").Append(userPermissions)
	storage.DB.Model(GetRole(4)).Association("Permissions").Append(systemPermissions)
}
