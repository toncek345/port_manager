CREATE TABLE ports(
  id bigserial NOT NULL PRIMARY KEY,
  id_str text NOT NULL,
  name text NOT NULL,
  city text NOT NULL,
  country text NOT NULL,
  coord_long double precision,
  coord_lat double precision,
  province text NOT NULL,
  timezone text NOT NULL,
  code text NOT NULL,
  regions text NOT NULL,
  unlocs text NOT NULL,
  alias text NOT NULL
);

CREATE UNIQUE INDEX ports_id_str on ports(id_str);
