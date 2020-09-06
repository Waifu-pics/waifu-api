package database

import (
	"database/sql"
	"errors"
	"fmt"
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

func (m Database) CreateFileInDB(file, md5, endp string, nsfw bool) error {
	var count int
	m.db.QueryRow("SELECT COUNT(*) FROM uploads WHERE md5 = ?", md5).Scan(&count)
	if count != 0 {
		return ErrorMD5Exists
	}

	_, err := m.db.Exec("INSERT INTO uploads (file, md5, type, nsfw) VALUES (?, ?, ?, ?)", file, md5, endp, nsfw)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ErrorFileNameExists
		}
		return err
	}
	return nil
}

// GetFilesAdmin : get files for specified endpoint admin
func (m Database) GetFilesAdmin(endp, query string, verified, nsfw bool) ([]string, error) {
	var args []interface{}
	args = append(args,
		verified,
		endp,
		nsfw,
	)

	var qstring = "SELECT file FROM uploads WHERE verified = ? AND type = ? AND nsfw = ?"

	if query != "" {
		qstring += " AND file LIKE CONCAT('%', ?, '%')"
		args = append(args, query)
	}

	rows, err := m.db.Query(qstring, args...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var files []string

	for rows.Next() {
		var file string

		err := rows.Scan(&file)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (m Database) GetFiles(endpoint string, nsfw bool, notIN []string, limit int) ([]string, error) {
	if notIN != nil && len(notIN) > 0 {
		tx, err := m.db.Begin()
		if err != nil {
			return nil, err
		}

		defer tx.Rollback()

		_, err = tx.Exec("CREATE TEMPORARY TABLE IF NOT EXISTS ignorefile (file TEXT)")
		if err != nil {
			return nil, err
		}

		// Chunked to prevent string over 2048 characters
		var splice []string
		for i, v := range notIN {
			splice = append(splice, v)
			// If index is divisible by 100
			if i%100 == 0 || i == len(notIN)-1 {
				args := make([]interface{}, len(splice))
				for i, fn := range splice {
					args[i] = fn
				}

				_, err := tx.Exec("INSERT INTO ignorefile (file) VALUES (?)"+strings.Repeat(",(?)", len(splice)-1), args...)
				if err != nil {
					return nil, err
				}

				splice = []string{}
			}
		}

		rows, err := tx.Query("SELECT file FROM uploads WHERE type = ? AND verified = 1 AND nsfw = ? AND file NOT IN (SELECT file FROM ignorefile) ORDER BY RAND() LIMIT ?",
			endpoint, nsfw, limit)
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

	rows, err := m.db.Query("SELECT file FROM uploads WHERE type = ? AND verified = 1 AND nsfw = ? ORDER BY RAND() LIMIT ?", endpoint, nsfw, limit)
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
