SELECT id, api_id, name, short_name, country_code, country_name, age, plays, turned_pro
FROM players
WHERE id = @p1
