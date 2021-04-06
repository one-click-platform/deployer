/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AccountObject struct {
	Key
	Attributes AccountObjectAttributes `json:"attributes"`
}
type AccountObjectResponse struct {
	Data     AccountObject `json:"data"`
	Included Included      `json:"included"`
}

type AccountObjectListResponse struct {
	Data     []AccountObject `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustAccountObject - returns AccountObject from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccountObject(key Key) *AccountObject {
	var accountObject AccountObject
	if c.tryFindEntry(key, &accountObject) {
		return &accountObject
	}
	return nil
}
