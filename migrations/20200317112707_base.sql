-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TYPE status_orders AS enum (
	'DRAFT',
	'NEW',
	'WORK',
	'CANCEL',
	'CLOSE'
);

CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	login VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	is_online bool NOT NULL DEFAULT FALSE,
	sample_js jsonb NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
	deleted_at TIMESTAMP NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE public.roles (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	code VARCHAR NOT NULL,
	description VARCHAR NOT NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id)
);

CREATE TABLE public.users_roles (
	user_id uuid NOT NULL,
	role_id uuid NOT NULL,
	CONSTRAINT users_roles_pk PRIMARY KEY (
		user_id,
		role_id
	)
);

ALTER TABLE public.users_roles ADD CONSTRAINT users_roles_fk FOREIGN KEY (user_id) REFERENCES users(id) ON
UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE public.users_roles ADD CONSTRAINT users_roles_fk_1 FOREIGN KEY (role_id) REFERENCES roles(id) ON
UPDATE CASCADE ON DELETE CASCADE;

CREATE TABLE public.orders (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	user_id uuid NOT NULL,
	"number" int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	status status_orders NOT NULL DEFAULT 'NEW'::status_orders,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
	deleted_at TIMESTAMP NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);

ALTER TABLE public.orders ADD CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES users(id) ON
UPDATE CASCADE ON DELETE RESTRICT;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE public.orders;

DROP TABLE public.users_roles;

DROP TABLE public.users;

DROP TABLE public.roles;

DROP TYPE status_orders;
