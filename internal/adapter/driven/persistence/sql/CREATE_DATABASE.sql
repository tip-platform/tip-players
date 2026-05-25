IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = @p1)
BEGIN
  DECLARE @sql NVARCHAR(MAX);
  SET @sql = N'CREATE DATABASE ' + QUOTENAME(@p1);
  EXEC sp_executesql @sql;
  PRINT 'Database created successfully.';
END
ELSE
BEGIN
  PRINT 'Database already exists.';
END
