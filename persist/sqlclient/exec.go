package sqlclient

var (
	createTable = "CREATE TABLE IF NOT EXISTS movieInfo(" +
		"ID VARCHAR(255)  CHARACTER SET utf8mb4 NOT NULL UNIQUE," +
		"Director VARCHAR(255) CHARACTER SET utf8mb4," +
		"Screenwriter VARCHAR(255) CHARACTER SET utf8mb4," +
		"Starring VARCHAR(255) CHARACTER SET utf8mb4," +
		"Type VARCHAR(255) CHARACTER SET utf8mb4," +
		"Region VARCHAR(255) CHARACTER SET utf8mb4," +
		"Language VARCHAR(255) CHARACTER SET utf8mb4," +
		"ReleaseDate VARCHAR(255) CHARACTER SET utf8mb4," +
		"TimeLong VARCHAR(30) CHARACTER SET utf8mb4," +
		"Sname VARCHAR(255) CHARACTER SET utf8mb4," +
		"Title VARCHAR(255) CHARACTER SET utf8mb4," +
		"URL VARCHAR(255) UNIQUE," +
		"PRIMARY KEY (ID)," +
		"INDEX idx_Title (Title)," +
		"INDEX idx_TimeLong (TimeLong)," +
		"INDEX idx_ReleaseDate (ReleaseDate)," +
		"INDEX idx_Region (Region)," +
		"INDEX idx_Language (Language)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"

	insertData = `INSERT INTO movieInfo (ID, Director, Screenwriter, Starring, Type, Region, Language, ReleaseDate, TimeLong, Sname, Title, URL) 
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
					ON DUPLICATE KEY UPDATE
        			Director = VALUES(Director),
        			Screenwriter = VALUES(Screenwriter),
        			Starring = VALUES(Starring),
        			Type = VALUES(Type),
        			Region = VALUES(Region),
        			Language = VALUES(Language),
        			ReleaseDate = VALUES(ReleaseDate),
        			TimeLong = VALUES(TimeLong),
        			Sname = VALUES(Sname),
        			Title = VALUES(Title),
        			URL = VALUES(URL)`

	insertData1 = `INSERT INTO movieInfo (ID, Director, Screenwriter, Starring, Type, Region, Language, ReleaseDate, TimeLong, Sname, Title, URL)
		VALUES ("%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s");`

	updateData = `UPDATE movieInfo
            SET
                Director = ?,
                Screenwriter = ?,
                Starring = ?,
                Type = ?,
                Region = ?,
                Language = ?,
                ReleaseDate = ?,
                TimeLong = ?,
                Sname = ?,
                Title = ?,
                URL = ?
            WHERE ID = ?`

	queryData = `SELECT * FROM student;`
)
