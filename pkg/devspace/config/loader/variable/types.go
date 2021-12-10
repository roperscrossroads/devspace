package variable

import (
	"github.com/loft-sh/devspace/pkg/devspace/config/versions/latest"
)

// Variable defines an interface to load a variable
type Variable interface {
	Load(definition *latest.Variable) (interface{}, error)
}

// Resolver defines an interface to resolve defined variables
type Resolver interface {
	// ConvertFlags converts several variables from input flags in the form of varname=value
	ConvertFlags(flags []string) (map[string]interface{}, error)

	// DefinedVars returns the defined variables
	DefinedVars() []*latest.Variable

	// UpdateVars sets the defined variables to use in the resolver
	UpdateVars(vars []*latest.Variable)

	// FindVariables returns all variable names that were found in the given map
	FindVariables(haystack interface{}) (map[string]bool, error)

	// FillVariables finds the used variables first and then fills in those in the haystack
	FillVariables(haystack interface{}) (interface{}, error)

	// FillVariablesExclude finds the used variables first and then fills in those that do not match the excluded paths in the haystack
	FillVariablesExclude(haystack interface{}, excluded []string) (interface{}, error)

	// ResolvedVariables returns the internal memory cache of the resolver with all resolved variables
	ResolvedVariables() map[string]interface{}
}
