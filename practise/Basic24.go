package main

/**
	1. A type assertion provides access to an interface value's underlying concrete value
	2. t := i.(T)
	3. This statement asserts that the interface value i holds the concrete type T
		and assigns the underlying T value to the variable t.
	4. t,ok = i.(T), if i holds a T, then t will be underlying value and ok will be true
	5. If it doesn't, t will be the underlying value and ok will have false.
 */