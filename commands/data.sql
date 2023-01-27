--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

-- Started on 2022-01-20 18:27:29

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
-- TOC entry 215 (class 1259 OID 24627)
-- Name: Clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Clients" (
    id bigint NOT NULL,
    first_name character varying(128) NOT NULL,
    last_name character varying(128) NOT NULL,
    identified boolean NOT NULL,
    wallet_id character varying(128) NOT NULL,
    passport_number character varying
);


ALTER TABLE public."Clients" OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 24632)
-- Name: Clients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public."Clients" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."Clients_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 209 (class 1259 OID 24593)
-- Name: Partners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Partners" (
    id bigint NOT NULL,
    name character varying(256) NOT NULL,
    address character varying(256),
    iin character varying(12)
);


ALTER TABLE public."Partners" OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 24598)
-- Name: Partners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public."Partners" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."Partners_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 218 (class 1259 OID 24643)
-- Name: Payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Payments" (
    id bigint NOT NULL,
    src character varying NOT NULL,
    dest character varying,
    sum numeric NOT NULL,
    created_at date,
    updated_at date,
    processed_at date,
    status bigint NOT NULL,
    description character varying,
    ext_id character varying NOT NULL
);


ALTER TABLE public."Payments" OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 24642)
-- Name: Payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Payments_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Payments_id_seq" OWNER TO postgres;

--
-- TOC entry 3350 (class 0 OID 0)
-- Dependencies: 217
-- Name: Payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Payments_id_seq" OWNED BY public."Payments".id;


--
-- TOC entry 211 (class 1259 OID 24604)
-- Name: Users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Users" (
    id integer NOT NULL,
    first_name character varying(128) NOT NULL,
    last_name character varying(128) NOT NULL,
    login character varying(128) NOT NULL,
    password character varying(256) NOT NULL,
    salt character varying(256) NOT NULL,
    partner_id bigint NOT NULL,
    wallet_id bigint NOT NULL
);


ALTER TABLE public."Users" OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 24609)
-- Name: Users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public."Users" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."Users_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 213 (class 1259 OID 24610)
-- Name: Wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Wallets" (
    id bigint NOT NULL,
    login character varying NOT NULL,
    sum numeric DEFAULT 0.00 NOT NULL,
    created_at date,
    updated_at date
);


ALTER TABLE public."Wallets" OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 24616)
-- Name: Wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public."Wallets" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."Wallets_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3185 (class 2604 OID 24646)
-- Name: Payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Payments" ALTER COLUMN id SET DEFAULT nextval('public."Payments_id_seq"'::regclass);


--
-- TOC entry 3341 (class 0 OID 24627)
-- Dependencies: 215
-- Data for Name: Clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Clients" (id, first_name, last_name, identified, wallet_id, passport_number) FROM stdin;
1	Ёров	Шерали	t	1	somepassid
2	Ёрзода	Бехзад	f	2	behid
\.


--
-- TOC entry 3335 (class 0 OID 24593)
-- Dependencies: 209
-- Data for Name: Partners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Partners" (id, name, address, iin) FROM stdin;
1	Alif Bank	Peshi spartak	12345667
2	Eskhata	Peshi Circ\n	122233554
\.


--
-- TOC entry 3344 (class 0 OID 24643)
-- Dependencies: 218
-- Data for Name: Payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Payments" (id, src, dest, sum, created_at, updated_at, processed_at, status, description, ext_id) FROM stdin;
75	alif_partner_77	917196244	11	2022-01-18	\N	\N	200	Success	IS4cMWnTSI
76	alif_partner_77	917196244	11	2022-01-18	2022-01-18	2022-01-18	200	Success	TzYlYWJslM
77	alif_partner_77	917196244	11	2022-01-18	2022-01-18	2022-01-18	200	Success	wcz1PLsCBEFdt7aNWgdl
78	alif_partner_77	917196244	11	2022-01-18	2022-01-18	2022-01-18	200	Success	vaF9qh9amWPJoemw92lP
79	alif_partner_77	917196244	11	2022-01-18	2022-01-18	2022-01-18	200	Success	WeOnwzZtwK7NK9S8lBqV
74	alif_partner_77	917196244	11.23	2022-01-18	\N	\N	200	Success	k5XZWOagik
80	alif_partner_77	917220068	100	2022-01-18	2022-01-18	2022-01-18	200	Success	Umn7J8KVYCyyFDtLvwMZ
81	alif_partner_77	917220068	100	2022-01-18	2022-01-18	2022-01-18	200	Success	aGxJ0Nz1ijRzYfLnkktq
82	alif_partner_77	917220068	100.82	2022-01-18	2022-01-18	2022-01-18	200	Success	GUW0ufXc0YCquo21OWMg
83	alif_partner_77	917220068	100.82	2022-01-18	2022-01-18	2022-01-18	200	Success	C0Mecr4XRqJ2MeOwfHIA
\.


--
-- TOC entry 3337 (class 0 OID 24604)
-- Dependencies: 211
-- Data for Name: Users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Users" (id, first_name, last_name, login, password, salt, partner_id, wallet_id) FROM stdin;
1	Чахонгир	Чамолов	jjamolov	9CGIGaKKrMRg1IfszVJrANk8i-0g7Q4y2HgE0U9BRF4=	Gl3!#j	1	3
2	Фаррух	Суллейманов	farrukh1111	288519f953133ade284a2f0b03dc15e1c9273e3ebdb1af4e327f705975c59cfb	b3otee	2	4
\.


--
-- TOC entry 3339 (class 0 OID 24610)
-- Dependencies: 213
-- Data for Name: Wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Wallets" (id, login, sum, created_at, updated_at) FROM stdin;
4	ekshata1111	116	2022-01-15	\N
2	917196244	67.38	2022-01-15	\N
3	alif_partner_77	19530.98	2022-01-15	\N
1	917220068	401.64	2022-01-15	\N
\.


--
-- TOC entry 3351 (class 0 OID 0)
-- Dependencies: 216
-- Name: Clients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Clients_id_seq"', 2, true);


--
-- TOC entry 3352 (class 0 OID 0)
-- Dependencies: 210
-- Name: Partners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Partners_id_seq"', 2, true);


--
-- TOC entry 3353 (class 0 OID 0)
-- Dependencies: 217
-- Name: Payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Payments_id_seq"', 83, true);


--
-- TOC entry 3354 (class 0 OID 0)
-- Dependencies: 212
-- Name: Users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Users_id_seq"', 2, true);


--
-- TOC entry 3355 (class 0 OID 0)
-- Dependencies: 214
-- Name: Wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Wallets_id_seq"', 4, true);


--
-- TOC entry 3193 (class 2606 OID 24634)
-- Name: Clients Clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Clients"
    ADD CONSTRAINT "Clients_pkey" PRIMARY KEY (id);


--
-- TOC entry 3187 (class 2606 OID 24620)
-- Name: Partners Partners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Partners"
    ADD CONSTRAINT "Partners_pkey" PRIMARY KEY (id);


--
-- TOC entry 3195 (class 2606 OID 24650)
-- Name: Payments Payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Payments"
    ADD CONSTRAINT "Payments_pkey" PRIMARY KEY (id);


--
-- TOC entry 3189 (class 2606 OID 24624)
-- Name: Users Users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Users"
    ADD CONSTRAINT "Users_pkey" PRIMARY KEY (id);


--
-- TOC entry 3191 (class 2606 OID 24626)
-- Name: Wallets Wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Wallets"
    ADD CONSTRAINT "Wallets_pkey" PRIMARY KEY (id);


-- Completed on 2022-01-20 18:27:29

--
-- PostgreSQL database dump complete
--

