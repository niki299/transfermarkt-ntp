package repository

var defaultQuery = `
	SELECT off.offer_id, off.date, pl.name, pl.surname,
	pl.age, pl.foot, pl.club, pl.playervalue, pl.playerposition
	FROM offer off, player pl
	WHERE off.player_id = pl.player_id
	AND LOWER(pl.name) LIKE '%' || LOWER($1) || '%'
	AND LOWER(pl.surname) LIKE '%' || LOWER($2) || '%'
	AND LOWER(pl.club) LIKE '%' || LOWER($3) || '%'
	AND pl.playerPosition LIKE '%' || $4 || '%'
	AND pl.playervalue BETWEEN $5 AND $6
	`

var averageRate = `
	SELECT offer.offer_id, avg(ratevalue)
	FROM offer, rate
	WHERE offer.offer_id = rate.offer_id
	GROUP BY offer.offer_id
	HAVING offer.offer_id = $1
	`
var commentQuery = `
	SELECT content
	FROM offer, comment
	WHERE offer.offer_id = comment.offer_id AND offer.offer_id = $1
	`

//dodati za cenu i poziciju/
