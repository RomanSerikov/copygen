// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/cyclic/domain"
	"github.com/switchupcb/copygen/examples/cyclic/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	ModelsToDomain(models.Account, models.User) *domain.Account
}