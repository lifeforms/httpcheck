package urlcheck

type Testable interface {
	Test() error
}

type Manifest []Server
