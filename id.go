package ddd

type Identifiable interface {
	Identify(Identifiable) bool
}

type Comparable interface {
	Less(Comparable) bool
}
