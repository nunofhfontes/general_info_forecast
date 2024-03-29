-- type Portfolio struct {
-- 	Id       int64
-- 	UserId   int64
-- 	Type     string // needed?
-- 	Security string // this should be a multi valur column, json, jsonb, array // -> JSONB
-- 	// slight slower to input, but faster to process from that point, since input will not be that frequent...
-- 	Created time.Time `xorm:"created"`
-- 	Updated time.Time `xorm:"updated"`
-- }

DROP TABLE "portfolio";

CREATE TABLE "portfolio" (
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   type TEXT,
   security jsonb
);