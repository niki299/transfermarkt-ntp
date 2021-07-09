package repository

var defaultQuery = `
	SELECT off.offer_id, off.date, pl.name, pl.surname,
	pl.age, pl.foot, pl.club
	FROM offer off, player pl
	WHERE off.player_id = pl.player_id
	AND LOWER(pl.name) LIKE '%' || $1 || '%'
	AND LOWER(pl.surname) LIKE '%' || $2 || '%'
	AND LOWER(pl.club) LIKE '%' || $3 || '%'
	`
