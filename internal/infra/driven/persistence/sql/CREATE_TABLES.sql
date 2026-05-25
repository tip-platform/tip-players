IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='players' AND xtype='U')
	BEGIN
		CREATE TABLE players (
      id INT IDENTITY(1,1) PRIMARY KEY,
      api_id INT NOT NULL,
      name NVARCHAR(255) NOT NULL,
      short_name NVARCHAR(100),
      country_code NVARCHAR(10) NOT NULL,
      country_name NVARCHAR(100),
      age TINYINT,
      plays NVARCHAR(50),
      turned_pro DATETIME2
    );
  PRINT 'Table players created successfully in tip_players.';
END
