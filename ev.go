package ev

import (
	"os"
	"strings"
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
// with the given name and value. It returns an error, if any.
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

// Iterate -- calls the given function for each existing environment
// variable, passing the name and value of that environment variable.
func Iterate(p func(name, value string)) {
	envs := os.Environ()
	for _, env := range envs {
		k := strings.Index(env, "=")
		if k > 0 && k < len(env)-1 {
			p(env[:k], env[k+1:])
		}
	}
}
