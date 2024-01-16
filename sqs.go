package sqs

import (
	"database/sql"
	"errors"
	"time"
)

type SessionStorage struct {
	db        *sql.DB
	cache     map[string]*Session
	cacheCap  int
	tableName string
}

func NewSessionStorage(db *sql.DB, tableName string) (*SessionStorage, error) {
	ss := &SessionStorage{
		db:        db,
		cache:     map[string]*Session{},
		cacheCap:  1024, // TODO:
		tableName: tableName,
	}
	err := ss.createSessionsTable()
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (ss *SessionStorage) cacheCleanup() {
	for k, v := range ss.cache {
		if v == nil {
			delete(ss.cache, k)
		}
	}
}

func (ss *SessionStorage) FlushCache() {
	ss.cache = make(map[string]*Session)
}

func (ss *SessionStorage) AddSession(userID int, validTo time.Time) (token string, err error) {
	session := NewSession(userID, validTo)
	token, e := ss.insertSession(session)
	if e != nil {
		return "", e
	}

	ss.cache[token] = session

	return token, nil
}

func (ss *SessionStorage) GetSession(token string) (*Session, error) {
	cached, ok := ss.cache[token]
	if !ok {
		session, err := ss.selectSessionByToken(token)
		if err != nil {
			// HACK:
			if len(ss.cache) > ss.cacheCap {
				ss.cacheCleanup()
			}
			ss.cache[token] = nil
			return nil, err
		}
		ss.cache[token] = session.Clone()
		return session, nil
	}

	if cached == nil {
		return nil, errors.New("Session not found")
	}

	return cached.Clone(), nil
}

func (ss *SessionStorage) RemoveSession(token string) error {
	err := ss.removeSession(token)
	if err != nil {
		return err
	}
	delete(ss.cache, token)
	return nil
}
