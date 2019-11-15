CREATE TABLE IF NOT EXISTS drinks(
   id UUID PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   volume FLOAT,
   price FLOAT
);

CREATE TABLE IF NOT EXISTS reservations(
   id UUID PRIMARY KEY,
   phone VARCHAR (50) NOT NULL,
   name VARCHAR (50) NOT NULL,
   start_at TIMESTAMP,
   end_at TIMESTAMP,
   deposit FLOAT,
   table_id UUID --FOREIGN KEY
);

CREATE TABLE IF NOT EXISTS reserved_drinks(
    reservation_id UUID,
    drink_id UUID
);

CREATE TABLE IF NOT EXISTS tables(
   id UUID PRIMARY KEY,
   sits INT
);