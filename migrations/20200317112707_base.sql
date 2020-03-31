-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE public.users
(
    id         uuid      NOT NULL DEFAULT uuid_generate_v4(),
    login      varchar   NOT NULL,
    email      varchar   NOT NULL,
    price      numeric   NOT NULL DEFAULT 0,
    summa_one  float4    NOT NULL DEFAULT 0,
    summa_two  float8    NOT NULL DEFAULT 0,
    cnt2       int2      NOT NULL DEFAULT 0,
    cnt4       int4      NOT NULL DEFAULT 0,
    cnt8       int8      NOT NULL DEFAULT 0,
    is_online  bool      NOT NULL DEFAULT false,
    alias      _text     NULL,
    data_byte  bytea     NULL,
    metrika    jsonb     NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE TRIGGER users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

CREATE TABLE public.roles
(
    id          uuid    NOT NULL DEFAULT uuid_generate_v4(),
    code        VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    CONSTRAINT roles_pk PRIMARY KEY (id)
);

CREATE TABLE public.users_roles
(
    user_id uuid NOT NULL,
    role_id uuid NOT NULL,
    CONSTRAINT users_roles_pk PRIMARY KEY (
                                           user_id,
                                           role_id
        )
);
ALTER TABLE public.users_roles
    ADD CONSTRAINT users_roles_fk FOREIGN KEY (user_id) REFERENCES users (id) ON
        UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE public.users_roles
    ADD CONSTRAINT users_roles_fk_1 FOREIGN KEY (role_id) REFERENCES roles (id) ON
        UPDATE CASCADE ON DELETE CASCADE;

CREATE TABLE public.orders
(
    id         uuid      NOT NULL DEFAULT uuid_generate_v4(),
    user_id    uuid      NULL,
    "number"   int4      NOT NULL GENERATED ALWAYS AS IDENTITY,
    status     text      NOT NULL DEFAULT 'DRAFT',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT orders_pk PRIMARY KEY (id)
);
CREATE TRIGGER orders_updated_at
    BEFORE UPDATE
    ON orders
    FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);
ALTER TABLE public.orders
    ADD CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES users (id) ON
        UPDATE CASCADE ON DELETE RESTRICT;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE public.orders;

DROP TABLE public.users_roles;

DROP TABLE public.users;

DROP TABLE public.roles;
