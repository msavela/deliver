package deliver

type Context struct {
	Keys map[string]interface{}
}

// Initialize new context.
func NewContext() *Context {
	return &Context{}
}

// Sets new key/value pair.
// Initializes a new hash table in case not already specified.
func (c *Context) Set(key string, item interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = item
}

// Returns value for the given key.
func (c *Context) Get(key string) interface{} {
	if c.Keys != nil {
		value, ok := c.Keys[key]
		if ok {
			return value
		}
	}
	return nil
}

// Returns value for the given key.
// Returns error in case the key does not exist.
func (c *Context) GetOk(key string) (interface{}, bool) {
	if c.Keys != nil {
		value, ok := c.Keys[key]
		if ok {
			return value, ok
		}
	}
	return nil, false
}

// Does context have the given key specified.
func (c *Context) Has(key string) bool {
	if c.Keys != nil {
		_, ok := c.Keys[key]
		return ok
	}
	return false
}