package ghsource

import "github.com/go-faker/faker/v4"

// contrive a dependency
func GetFake() string {
	return faker.FirstName()
}
