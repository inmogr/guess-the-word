--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2018-11-09 12:41:32

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE sample;
--
-- TOC entry 2835 (class 1262 OID 16393)
-- Name: sample; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE sample WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1252' LC_CTYPE = 'English_United States.1252';


ALTER DATABASE sample OWNER TO postgres;

\connect sample

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 7 (class 2615 OID 16394)
-- Name: sample; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA sample;


ALTER SCHEMA sample OWNER TO postgres;

--
-- TOC entry 2 (class 3079 OID 16404)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 2836 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 199 (class 1259 OID 16433)
-- Name: scores; Type: TABLE; Schema: sample; Owner: postgres
--

CREATE TABLE sample.scores (
    id uuid NOT NULL,
    score double precision NOT NULL,
    uid uuid NOT NULL
);


ALTER TABLE sample.scores OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 16423)
-- Name: users; Type: TABLE; Schema: sample; Owner: postgres
--

CREATE TABLE sample.users (
    uid uuid NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);


ALTER TABLE sample.users OWNER TO postgres;

--
-- TOC entry 2829 (class 0 OID 16433)
-- Dependencies: 199
-- Data for Name: scores; Type: TABLE DATA; Schema: sample; Owner: postgres
--

INSERT INTO sample.scores VALUES ('f4d706c0-ab64-4f6f-a17d-113dbad5b591', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('a74af41e-54b0-4fb1-9010-fc53e1e15152', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('70e82691-0c56-4ac6-883f-53ec93ad15bc', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('074991f8-545d-4738-b28d-738f8a8836bb', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('0c94dc12-d075-4c03-952d-7b1c8fb5f438', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('1efb13b4-1dd1-4943-b68c-acf2ac1788e8', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');
INSERT INTO sample.scores VALUES ('908cfb73-7e4f-4c12-b0ae-7033720cfd2f', 0.800000000000000044, 'e3199a60-58cd-4582-a749-718ad31424f5');


--
-- TOC entry 2828 (class 0 OID 16423)
-- Dependencies: 198
-- Data for Name: users; Type: TABLE DATA; Schema: sample; Owner: postgres
--

INSERT INTO sample.users VALUES ('e3199a60-58cd-4582-a749-718ad31424f5', 'morhaf', 'shamia');


--
-- TOC entry 2705 (class 2606 OID 16437)
-- Name: scores scores_pkey; Type: CONSTRAINT; Schema: sample; Owner: postgres
--

ALTER TABLE ONLY sample.scores
    ADD CONSTRAINT scores_pkey PRIMARY KEY (id);


--
-- TOC entry 2701 (class 2606 OID 16432)
-- Name: users username_is_unqiue; Type: CONSTRAINT; Schema: sample; Owner: postgres
--

ALTER TABLE ONLY sample.users
    ADD CONSTRAINT username_is_unqiue UNIQUE (username);


--
-- TOC entry 2703 (class 2606 OID 16430)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: sample; Owner: postgres
--

ALTER TABLE ONLY sample.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (uid);


--
-- TOC entry 2706 (class 2606 OID 16438)
-- Name: scores uid_is_user_uid; Type: FK CONSTRAINT; Schema: sample; Owner: postgres
--

ALTER TABLE ONLY sample.scores
    ADD CONSTRAINT uid_is_user_uid FOREIGN KEY (uid) REFERENCES sample.users(uid);


-- Completed on 2018-11-09 12:41:32

--
-- PostgreSQL database dump complete
--

