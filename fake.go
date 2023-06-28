package ghsource

import "github.com/go-faker/faker/v4"

// contrive a dependency
func getFake() string {
	return faker.FirstName()
}
