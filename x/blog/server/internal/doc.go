// Package internal is a convenience to allow us to use exported types to signal importance.
// Thus, it prevents the parent server package from having to always use unexported types to hide implementation
// from callers.
//
// This also allows us to unit test exported types and methods only.
// We never test private types, methods, or functions. Private behavior should exhibit itself through a public,
// tested interface. Therefore, we can change private code quickly minimizing overhead of refactoring unit tests.
package internal
