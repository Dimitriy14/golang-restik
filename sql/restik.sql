CREATE TABLE IF NOT EXISTS foods(
   id UUID PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   type VARCHAR (50),
   description TEXT,
   image TEXT,
   price FLOAT
);

CREATE TABLE IF NOT EXISTS reservations(
   id UUID PRIMARY KEY,
   phone VARCHAR (50) NOT NULL,
   name VARCHAR (50) NOT NULL,
   start_at TIMESTAMP,
   end_at TIMESTAMP,
   table_id UUID --FOREIGN KEY
);

CREATE TABLE IF NOT EXISTS reserved_food(
    reservation_id UUID,
    food_id UUID,
    amount int
);

CREATE TABLE IF NOT EXISTS tables(
   id UUID PRIMARY KEY,
   sits INT
);