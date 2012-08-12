package goravendb

type DocumentStore struct {
	url string
}

type DocumentSession struct {
	store *DocumentStore
}

func (ds *DocumentStore) Init(url string) (bool, error) {
	ds.url = url
	return true, nil
}

func (ds *DocumentStore) Connect() (*DocumentSession, error) {
	session := new(DocumentSession)
	session.store = ds
	return session, nil
}
