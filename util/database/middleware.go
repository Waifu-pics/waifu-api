package database

import (
	"database/sql"
	"errors"
	"strings"
)

// Database : database type
type Database struct {
	db *sql.DB
}

func (m Database) CreateAdmin(username, hashpass string) error {
	_, err := m.db.Exec("INSERT INTO admins (Username, Password) VALUES (?, ?)", username, hashpass)
	if err != nil {
		return err
	}
	return nil
}

func (m Database) GetAdminHash(username string) (string, error) {
	var hash string

	row := m.db.QueryRow("SELECT Password FROM admins WHERE Username = ?", username)

	err := row.Scan(&hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

var (
	ErrorMD5Exists      = errors.New("file hash already exists")
	ErrorFileNameExists = errors.New("filename already exists")
)

func (m Database) CreateFileInDB(file, md5, endp string, verified bool) error {
	var count int
	m.db.QueryRow("SELECT COUNT(*) FROM uploads WHERE md5 = ?", md5).Scan(&count)
	if count != 0 {
		return ErrorMD5Exists
	}

	_, err := m.db.Exec("INSERT INTO uploads (file. md5, type, verified) VALUES (? ? ? ?)", file, md5, endp, verified)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ErrorFileNameExists
		}
		return err
	}
	return nil
}

// GetFilesAdmin : get files for specified endpoint admin
func (m Database) GetFilesAdmin(endp string, verified bool) ([]string, error) {
	rows, err := m.db.Query("SELECT file FROM uploads WHERE verified = ?", verified)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var files []string

	for rows.Next() {
		var file string

		err := rows.Scan(&file)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (m Database) GetFiles(endpoint string, notIN []string, limit int) ([]string, error) {
	var err error
	var rows *sql.Rows

	if len(notIN) == 0 || notIN == nil {
		rows, err = m.db.Query("SELECT file FROM uploads WHERE type = ? AND verified = 1 ORDER BY RAND() LIMIT ?", limit)
	} else {
		args := make([]interface{}, len(notIN))
		for i, fn := range notIN {
			args[i] = fn
		}
		sqlstr := "SELECT file FROM uploads WHERE file NOT IN (?" + strings.Repeat(",?", len(args)-1) + `)` + " AND type = ? AND verified = 1 ORDER BY RAND() LIMIT ?"
		args = append(args, endpoint, limit)
		rows, err = m.db.Query(sqlstr, args...)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var files []string

	for rows.Next() {
		var file string

		rows.Scan(&file)

		files = append(files, file)
	}

	return files, nil
}

func (m Database) VerifyFile(filename string) error {
	r, err := m.db.Exec("UPDATE uploads SET verified = 1 WHERE file = ?", filename)
	rf, _ := r.RowsAffected()
	if err != nil || rf == 0 {
		return err
	}

	return nil
}

func (m Database) DeleteFile(filename string) error {
	r, err := m.db.Exec("DELETE FROM uploads WHERE file = ?", filename)
	rf, _ := r.RowsAffected()
	if err != nil || rf == 0 {
		return err
	}

	return nil
}
