/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Environment struct {
	Key
	Attributes    EnvironmentAttributes     `json:"attributes"`
	Relationships *EnvironmentRelationships `json:"relationships,omitempty"`
}
type EnvironmentResponse struct {
	Data     Environment `json:"data"`
	Included Included    `json:"included"`
}

type EnvironmentListResponse struct {
	Data     []Environment `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustEnvironment - returns Environment from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEnvironment(key Key) *Environment {
	var environment Environment
	if c.tryFindEntry(key, &environment) {
		return &environment
	}
	return nil
}
