package linkrelation

// LinkRelation represents an IANA Link Relation type as defined by [RFC-8288⤴].
//
// [RFC-8288⤴]: https://datatracker.ietf.org/doc/html/rfc8288
//
//go:generate go run ./internal/tools/generator
type LinkRelation string
