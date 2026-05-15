IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = @p1)
	BEGIN
		DECLARE @sql NVARCHAR(MAX) = N'CREATE DATABASE ' + QUOTENAME(@p1);
		EXEC(@sql);
			PRINT 'Database created successfully.'
		END
	ELSE
		BEGIN
	    	PRINT 'Database already exists.';
		END
