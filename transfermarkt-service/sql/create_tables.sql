CREATE TABLE IF NOT EXISTS player (
  player_id varchar(250) NOT NULL,
  name varchar(250) NOT NULL,
  surname varchar(250) NOT NULL,
  age INT NOT NULL,
  foot varchar(250) NOT NULL,
  club varchar(250) NOT NULL,
  playerValue INT NOT NULL,
 
  PRIMARY KEY (player_id)
);