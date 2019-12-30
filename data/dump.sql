--
-- PostgreSQL database dump
--

-- Dumped from database version 10.11 (Ubuntu 10.11-1.pgdg18.04+1)
-- Dumped by pg_dump version 11.6 (Ubuntu 11.6-1.pgdg18.04+1)

-- Started on 2019-12-30 16:12:38 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 133304)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 2979 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 605 (class 1247 OID 133316)
-- Name: status_orders; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.status_orders AS ENUM (
    'DRAFT',
    'NEW',
    'WORK',
    'CANCEL',
    'CLOSE'
);


ALTER TYPE public.status_orders OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 198 (class 1259 OID 133297)
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goose_db_version OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 133295)
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goose_db_version_id_seq OWNER TO postgres;

--
-- TOC entry 2980 (class 0 OID 0)
-- Dependencies: 197
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- TOC entry 203 (class 1259 OID 133365)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    number integer NOT NULL,
    status public.status_orders DEFAULT 'NEW'::public.status_orders NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 133363)
-- Name: orders_number_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.orders ALTER COLUMN number ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_number_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 200 (class 1259 OID 133339)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    code character varying NOT NULL,
    description character varying NOT NULL
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 133327)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    login character varying NOT NULL,
    email character varying NOT NULL,
    is_online boolean DEFAULT false NOT NULL,
    sample_js jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 133348)
-- Name: users_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_roles (
    user_id uuid NOT NULL,
    role_id uuid NOT NULL
);


ALTER TABLE public.users_roles OWNER TO postgres;

--
-- TOC entry 2822 (class 2604 OID 133300)
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- TOC entry 2968 (class 0 OID 133297)
-- Dependencies: 198
-- Data for Name: goose_db_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
1	0	t	2019-09-23 10:03:45.864129
2	20190921191020	t	2019-09-23 10:03:49.166408
\.


--
-- TOC entry 2973 (class 0 OID 133365)
-- Dependencies: 203
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, user_id, number, status, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 2970 (class 0 OID 133339)
-- Dependencies: 200
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, code, description) FROM stdin;
\.


--
-- TOC entry 2969 (class 0 OID 133327)
-- Dependencies: 199
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, login, email, is_online, sample_js, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 2971 (class 0 OID 133348)
-- Dependencies: 201
-- Data for Name: users_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users_roles (user_id, role_id) FROM stdin;
\.


--
-- TOC entry 2981 (class 0 OID 0)
-- Dependencies: 197
-- Name: goose_db_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goose_db_version_id_seq', 2, true);


--
-- TOC entry 2982 (class 0 OID 0)
-- Dependencies: 202
-- Name: orders_number_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_number_seq', 1, false);


--
-- TOC entry 2834 (class 2606 OID 133303)
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- TOC entry 2842 (class 2606 OID 133373)
-- Name: orders orders_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pk PRIMARY KEY (id);


--
-- TOC entry 2838 (class 2606 OID 133347)
-- Name: roles roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pk PRIMARY KEY (id);


--
-- TOC entry 2836 (class 2606 OID 133338)
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 2840 (class 2606 OID 133352)
-- Name: users_roles users_roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_roles
    ADD CONSTRAINT users_roles_pk PRIMARY KEY (user_id, role_id);


--
-- TOC entry 2845 (class 2606 OID 133374)
-- Name: orders orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2843 (class 2606 OID 133353)
-- Name: users_roles users_roles_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_roles
    ADD CONSTRAINT users_roles_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2844 (class 2606 OID 133358)
-- Name: users_roles users_roles_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_roles
    ADD CONSTRAINT users_roles_fk_1 FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2019-12-30 16:12:38 MSK

--
-- PostgreSQL database dump complete
--

