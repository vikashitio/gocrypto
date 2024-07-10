--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-07-10 09:38:17

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 219 (class 1259 OID 16562)
-- Name: client_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.client_details (
    client_id smallint NOT NULL,
    title character varying(5) DEFAULT NULL::character varying,
    gender character varying(1) DEFAULT NULL::character varying,
    birth_date character varying(30) DEFAULT NULL::character varying,
    country_code character varying(5) DEFAULT NULL::character varying,
    mobile character varying(20) DEFAULT NULL::character varying,
    address_line1 character varying(100) DEFAULT NULL::character varying,
    address_line2 character varying(100) DEFAULT NULL::character varying,
    city character varying(50) DEFAULT NULL::character varying,
    state character varying(50) DEFAULT NULL::character varying,
    country character varying(50) DEFAULT NULL::character varying,
    pincode character varying(10) DEFAULT NULL::character varying,
    add_date character varying(30) DEFAULT NULL::character varying,
    json_log_history character varying(22302) DEFAULT NULL::character varying,
    id smallint NOT NULL
);


ALTER TABLE public.client_details OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16640)
-- Name: client_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.client_details ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.client_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 4878 (class 0 OID 16562)
-- Dependencies: 219
-- Data for Name: client_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.client_details (client_id, title, gender, birth_date, country_code, mobile, address_line1, address_line2, city, state, country, pincode, add_date, json_log_history, id) FROM stdin;
149	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	21
150	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	22
151	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	23
72	Ms	F	2024-06-12	355	09555538740	Noida Sec 65, ggggg	ggggg	\N	\N	\N	\N	\N	\N	1
\.


--
-- TOC entry 4885 (class 0 OID 0)
-- Dependencies: 222
-- Name: client_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.client_details_id_seq', 23, true);


--
-- TOC entry 4732 (class 2606 OID 16648)
-- Name: client_details client_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_details
    ADD CONSTRAINT client_details_pkey PRIMARY KEY (id);


--
-- TOC entry 4734 (class 2606 OID 16637)
-- Name: client_details constraint_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client_details
    ADD CONSTRAINT constraint_name UNIQUE (client_id);


-- Completed on 2024-07-10 09:38:17

--
-- PostgreSQL database dump complete
--

