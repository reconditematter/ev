package ev

import (
	"os"
)

// Value -- if an environment variable with the given
// name exists, then Value returns ("<value>",true);
// otherwise, it returns ("",false).
func Value(name string) (value string, exists bool) {
	value, exists = os.LookupEnv(name)
	return value, exists
}

// ValueOrDefault -- if an environment variable with the
// given name exists, then ValueOrDefault returns its value;
// otherwise, it returns defval.
func ValueOrDefault(name, defval string) (value string) {
	var exists bool
	value, exists = os.LookupEnv(name)
	if exists {
		return value
	}
	return defval
}

// Exists -- if an environment variable with the given name
// exists, then Exists returns true; otherwise, it returns false.
func Exists(name string) (exists bool) {
	_, exists = os.LookupEnv(name)
	return exists
}

// Set -- first Set clears an environment variable with the
// given name, and then it defines a new environment variable
// with the given name and value. Ir returns an error, if any.
func Set(name, value string) error {
	return os.Setenv(name, value)
}

// Clear -- if an environment variable with the given name exists,
// then Clear deletes it; otherwise, it does nothing.
func Clear(name string) {
	_ = os.Unsetenv(name)
}

// ClearAll -- deletes all existing environment variables.
func ClearAll() {
	os.Clearenv()
}
