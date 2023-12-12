package sqs

import (
	"fmt"
	"time"
)

func (ss *SessionStorage) createSessionsTable() error {
	q := `CREATE TABLE IF NOT EXISTS ` + ss.tableName + ` (
		token      TEXT,
        user_id    INTEGER,
		created_at INTEGER,
		valid_to   INTEGER,
		is_invalid INTEGER
	)`
	_, err := ss.db.Exec(q, nil)
	if err != nil {
		return fmt.Errorf("create sessions table: %v", err)
	}
	return nil
}

func (ss *SessionStorage) insertSession(session *Session) (token string, err error) {
	t := generateToken(32)
	q := `INSERT INTO ` + ss.tableName + ` (token, user_id, created_at, valid_to, is_invalid) VALUES (?, ?, ?, ?, ?)`
	createdAtInt := session.CreatedAt.Unix()
	validToInt := session.ValidTo.Unix()
	isInvalidInt := boolToInt(session.IsInvalid)

	_, erro := ss.db.Exec(q, t, session.UserID, createdAtInt, validToInt, isInvalidInt)
	if erro != nil {
		return "", fmt.Errorf("insert session: %v", erro)
	}
	return t, nil
}

func (ss *SessionStorage) selectSessionByToken(token string) (session *Session, err error) {
	q := `SELECT user_id, created_at, valid_to, is_invalid FROM ` + ss.tableName + ` WHERE token=? LIMIT 1`
	row := ss.db.QueryRow(q, token)
	if e := row.Err(); e != nil {
		return nil, fmt.Errorf("select session by token: %v", e)
	}
	var v Session
	var createdAt int64
	var validTo int64
	var isInvalid int
	if e := row.Scan(&v.UserID, &createdAt, &validTo, &isInvalid); e != nil {
		return nil, fmt.Errorf("select session - scanning: %v", e)
	}
	v.CreatedAt = time.Unix(createdAt, 0)
	v.ValidTo = time.Unix(validTo, 0)
	v.IsInvalid = intToBool(isInvalid)

	return &v, nil
}
