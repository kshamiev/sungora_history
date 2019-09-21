CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type status_orders as enum ('DRAFT', 'NEW', 'WORK', 'CANCEL', 'CLOSE');