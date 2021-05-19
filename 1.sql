--
-- PostgreSQL database dump
--

-- Dumped from database version 11.11 (Debian 11.11-1.pgdg100+1)
-- Dumped by pg_dump version 13.2 (Debian 13.2-1.pgdg100+1)

-- Started on 2021-05-17 09:38:45 MSK

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

--
-- TOC entry 197 (class 1259 OID 16619)
-- Name: CP; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."CP" (
    cp_id integer NOT NULL,
    date date NOT NULL,
    status integer NOT NULL,
    tz_id integer NOT NULL,
    proj text,
    o_id integer NOT NULL,
    price integer NOT NULL,
    group_id integer NOT NULL,
    type_id integer NOT NULL,
    kind_id integer NOT NULL,
    task_id integer NOT NULL,
    pay_cond_id integer NOT NULL,
    end_date date NOT NULL,
    info text,
    history text
);


ALTER TABLE public."CP" OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 16617)
-- Name: CP_cp_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."CP_cp_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CP_cp_id_seq" OWNER TO postgres;

--
-- TOC entry 3115 (class 0 OID 0)
-- Dependencies: 196
-- Name: CP_cp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."CP_cp_id_seq" OWNED BY public."CP".cp_id;


--
-- TOC entry 198 (class 1259 OID 16634)
-- Name: CP_docs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."CP_docs" (
    file_name text NOT NULL,
    cp_id integer NOT NULL
);


ALTER TABLE public."CP_docs" OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16855)
-- Name: Costs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Costs" (
    cost_id integer NOT NULL,
    name_id integer NOT NULL,
    metr_id integer NOT NULL,
    count integer NOT NULL,
    tz_id integer NOT NULL,
    cp_id integer NOT NULL,
    ppu numeric,
    info text
);


ALTER TABLE public."Costs" OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 16853)
-- Name: Costs_cost_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Costs_cost_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Costs_cost_id_seq" OWNER TO postgres;

--
-- TOC entry 3116 (class 0 OID 0)
-- Dependencies: 227
-- Name: Costs_cost_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Costs_cost_id_seq" OWNED BY public."Costs".cost_id;


--
-- TOC entry 202 (class 1259 OID 16660)
-- Name: Countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Countries" (
    country_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Countries" OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16658)
-- Name: Countries_country_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Countries_country_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Countries_country_id_seq" OWNER TO postgres;

--
-- TOC entry 3117 (class 0 OID 0)
-- Dependencies: 201
-- Name: Countries_country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Countries_country_id_seq" OWNED BY public."Countries".country_id;


--
-- TOC entry 216 (class 1259 OID 16737)
-- Name: Metrics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Metrics" (
    metr_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Metrics" OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16735)
-- Name: Metrics_metr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Metrics_metr_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Metrics_metr_id_seq" OWNER TO postgres;

--
-- TOC entry 3118 (class 0 OID 0)
-- Dependencies: 215
-- Name: Metrics_metr_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Metrics_metr_id_seq" OWNED BY public."Metrics".metr_id;


--
-- TOC entry 218 (class 1259 OID 16755)
-- Name: Org_countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Org_countries" (
    o_id integer NOT NULL,
    country_id integer NOT NULL
);


ALTER TABLE public."Org_countries" OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16767)
-- Name: Orgs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Orgs" (
    o_id integer NOT NULL,
    name text NOT NULL,
    group_id integer NOT NULL,
    site text,
    phone text,
    email text,
    adress text,
    info text,
    login text NOT NULL,
    hashed_pwd text NOT NULL,
    status boolean NOT NULL,
    history text
);


ALTER TABLE public."Orgs" OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16765)
-- Name: Orgs_o_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Orgs_o_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Orgs_o_id_seq" OWNER TO postgres;

--
-- TOC entry 3119 (class 0 OID 0)
-- Dependencies: 220
-- Name: Orgs_o_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Orgs_o_id_seq" OWNED BY public."Orgs".o_id;


--
-- TOC entry 219 (class 1259 OID 16760)
-- Name: Orgs_specs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Orgs_specs" (
    o_id integer NOT NULL,
    spec_id integer NOT NULL
);


ALTER TABLE public."Orgs_specs" OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16671)
-- Name: Pack_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_groups" (
    group_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_groups" OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16669)
-- Name: Pack_groups_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Pack_groups_group_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Pack_groups_group_id_seq" OWNER TO postgres;

--
-- TOC entry 3120 (class 0 OID 0)
-- Dependencies: 203
-- Name: Pack_groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_groups_group_id_seq" OWNED BY public."Pack_groups".group_id;


--
-- TOC entry 208 (class 1259 OID 16693)
-- Name: Pack_kinds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_kinds" (
    kind_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_kinds" OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 16691)
-- Name: Pack_kinds_kind_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Pack_kinds_kind_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Pack_kinds_kind_id_seq" OWNER TO postgres;

--
-- TOC entry 3121 (class 0 OID 0)
-- Dependencies: 207
-- Name: Pack_kinds_kind_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_kinds_kind_id_seq" OWNED BY public."Pack_kinds".kind_id;


--
-- TOC entry 206 (class 1259 OID 16682)
-- Name: Pack_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_types" (
    type_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_types" OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16680)
-- Name: Pack_types_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Pack_types_type_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Pack_types_type_id_seq" OWNER TO postgres;

--
-- TOC entry 3122 (class 0 OID 0)
-- Dependencies: 205
-- Name: Pack_types_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_types_type_id_seq" OWNED BY public."Pack_types".type_id;


--
-- TOC entry 212 (class 1259 OID 16715)
-- Name: Pay_conds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pay_conds" (
    pay_conds_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pay_conds" OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16713)
-- Name: Pay_conds_pay_conds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Pay_conds_pay_conds_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Pay_conds_pay_conds_id_seq" OWNER TO postgres;

--
-- TOC entry 3123 (class 0 OID 0)
-- Dependencies: 211
-- Name: Pay_conds_pay_conds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pay_conds_pay_conds_id_seq" OWNED BY public."Pay_conds".pay_conds_id;


--
-- TOC entry 200 (class 1259 OID 16649)
-- Name: Specs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Specs" (
    spec_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Specs" OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 16647)
-- Name: Specs_spec_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Specs_spec_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Specs_spec_id_seq" OWNER TO postgres;

--
-- TOC entry 3124 (class 0 OID 0)
-- Dependencies: 199
-- Name: Specs_spec_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Specs_spec_id_seq" OWNED BY public."Specs".spec_id;


--
-- TOC entry 226 (class 1259 OID 16843)
-- Name: Sup_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Sup_list" (
    o_id integer NOT NULL,
    sup_id integer NOT NULL
);


ALTER TABLE public."Sup_list" OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16726)
-- Name: Task_names; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Task_names" (
    name_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Task_names" OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16724)
-- Name: Task_names_name_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Task_names_name_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Task_names_name_id_seq" OWNER TO postgres;

--
-- TOC entry 3125 (class 0 OID 0)
-- Dependencies: 213
-- Name: Task_names_name_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_names_name_id_seq" OWNED BY public."Task_names".name_id;


--
-- TOC entry 210 (class 1259 OID 16704)
-- Name: Tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tasks" (
    task_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Tasks" OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16702)
-- Name: Tasks_task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Tasks_task_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Tasks_task_id_seq" OWNER TO postgres;

--
-- TOC entry 3126 (class 0 OID 0)
-- Dependencies: 209
-- Name: Tasks_task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tasks_task_id_seq" OWNED BY public."Tasks".task_id;


--
-- TOC entry 217 (class 1259 OID 16747)
-- Name: Tech_docs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tech_docs" (
    tz_id integer NOT NULL,
    file_name text NOT NULL
);


ALTER TABLE public."Tech_docs" OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16783)
-- Name: Techs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Techs" (
    tz_id integer NOT NULL,
    date date NOT NULL,
    o_id integer NOT NULL,
    end_date date NOT NULL,
    proj text,
    group_id integer NOT NULL,
    kind_id integer NOT NULL,
    type_id integer NOT NULL,
    task_id integer NOT NULL,
    tz_st integer NOT NULL,
    tender_st integer,
    count integer NOT NULL,
    cp_st integer,
    pay_cond_id integer NOT NULL,
    private boolean NOT NULL,
    info text,
    history text
);


ALTER TABLE public."Techs" OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16781)
-- Name: Techs_tz_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Techs_tz_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Techs_tz_id_seq" OWNER TO postgres;

--
-- TOC entry 3127 (class 0 OID 0)
-- Dependencies: 222
-- Name: Techs_tz_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Techs_tz_id_seq" OWNED BY public."Techs".tz_id;


--
-- TOC entry 225 (class 1259 OID 16824)
-- Name: Tenders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tenders" (
    tender_id integer NOT NULL,
    date date NOT NULL,
    selected_cp integer NOT NULL,
    tz_id integer NOT NULL,
    history text
);


ALTER TABLE public."Tenders" OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16822)
-- Name: Tenders_tender_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Tenders_tender_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Tenders_tender_id_seq" OWNER TO postgres;

--
-- TOC entry 3128 (class 0 OID 0)
-- Dependencies: 224
-- Name: Tenders_tender_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tenders_tender_id_seq" OWNED BY public."Tenders".tender_id;


--
-- TOC entry 2889 (class 2604 OID 16622)
-- Name: CP cp_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP" ALTER COLUMN cp_id SET DEFAULT nextval('public."CP_cp_id_seq"'::regclass);


--
-- TOC entry 2902 (class 2604 OID 16858)
-- Name: Costs cost_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs" ALTER COLUMN cost_id SET DEFAULT nextval('public."Costs_cost_id_seq"'::regclass);


--
-- TOC entry 2891 (class 2604 OID 16663)
-- Name: Countries country_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Countries" ALTER COLUMN country_id SET DEFAULT nextval('public."Countries_country_id_seq"'::regclass);


--
-- TOC entry 2898 (class 2604 OID 16740)
-- Name: Metrics metr_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Metrics" ALTER COLUMN metr_id SET DEFAULT nextval('public."Metrics_metr_id_seq"'::regclass);


--
-- TOC entry 2899 (class 2604 OID 16770)
-- Name: Orgs o_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs" ALTER COLUMN o_id SET DEFAULT nextval('public."Orgs_o_id_seq"'::regclass);


--
-- TOC entry 2892 (class 2604 OID 16674)
-- Name: Pack_groups group_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_groups" ALTER COLUMN group_id SET DEFAULT nextval('public."Pack_groups_group_id_seq"'::regclass);


--
-- TOC entry 2894 (class 2604 OID 16696)
-- Name: Pack_kinds kind_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_kinds" ALTER COLUMN kind_id SET DEFAULT nextval('public."Pack_kinds_kind_id_seq"'::regclass);


--
-- TOC entry 2893 (class 2604 OID 16685)
-- Name: Pack_types type_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_types" ALTER COLUMN type_id SET DEFAULT nextval('public."Pack_types_type_id_seq"'::regclass);


--
-- TOC entry 2896 (class 2604 OID 16718)
-- Name: Pay_conds pay_conds_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pay_conds" ALTER COLUMN pay_conds_id SET DEFAULT nextval('public."Pay_conds_pay_conds_id_seq"'::regclass);


--
-- TOC entry 2890 (class 2604 OID 16652)
-- Name: Specs spec_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Specs" ALTER COLUMN spec_id SET DEFAULT nextval('public."Specs_spec_id_seq"'::regclass);


--
-- TOC entry 2897 (class 2604 OID 16729)
-- Name: Task_names name_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_names" ALTER COLUMN name_id SET DEFAULT nextval('public."Task_names_name_id_seq"'::regclass);


--
-- TOC entry 2895 (class 2604 OID 16707)
-- Name: Tasks task_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tasks" ALTER COLUMN task_id SET DEFAULT nextval('public."Tasks_task_id_seq"'::regclass);


--
-- TOC entry 2900 (class 2604 OID 16786)
-- Name: Techs tz_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs" ALTER COLUMN tz_id SET DEFAULT nextval('public."Techs_tz_id_seq"'::regclass);


--
-- TOC entry 2901 (class 2604 OID 16827)
-- Name: Tenders tender_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders" ALTER COLUMN tender_id SET DEFAULT nextval('public."Tenders_tender_id_seq"'::regclass);


--
-- TOC entry 3078 (class 0 OID 16619)
-- Dependencies: 197
-- Data for Name: CP; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."CP" (cp_id, date, status, tz_id, proj, o_id, price, group_id, type_id, kind_id, task_id, pay_cond_id, end_date, info, history) FROM stdin;
\.


--
-- TOC entry 3079 (class 0 OID 16634)
-- Dependencies: 198
-- Data for Name: CP_docs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."CP_docs" (file_name, cp_id) FROM stdin;
\.


--
-- TOC entry 3109 (class 0 OID 16855)
-- Dependencies: 228
-- Data for Name: Costs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Costs" (cost_id, name_id, metr_id, count, tz_id, cp_id, ppu, info) FROM stdin;
\.


--
-- TOC entry 3083 (class 0 OID 16660)
-- Dependencies: 202
-- Data for Name: Countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Countries" (country_id, name) FROM stdin;
\.


--
-- TOC entry 3097 (class 0 OID 16737)
-- Dependencies: 216
-- Data for Name: Metrics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Metrics" (metr_id, name) FROM stdin;
\.


--
-- TOC entry 3099 (class 0 OID 16755)
-- Dependencies: 218
-- Data for Name: Org_countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Org_countries" (o_id, country_id) FROM stdin;
\.


--
-- TOC entry 3102 (class 0 OID 16767)
-- Dependencies: 221
-- Data for Name: Orgs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs" (o_id, name, group_id, site, phone, email, adress, info, login, hashed_pwd, status, history) FROM stdin;
\.


--
-- TOC entry 3100 (class 0 OID 16760)
-- Dependencies: 219
-- Data for Name: Orgs_specs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs_specs" (o_id, spec_id) FROM stdin;
\.


--
-- TOC entry 3085 (class 0 OID 16671)
-- Dependencies: 204
-- Data for Name: Pack_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_groups" (group_id, name) FROM stdin;
\.


--
-- TOC entry 3089 (class 0 OID 16693)
-- Dependencies: 208
-- Data for Name: Pack_kinds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_kinds" (kind_id, name) FROM stdin;
\.


--
-- TOC entry 3087 (class 0 OID 16682)
-- Dependencies: 206
-- Data for Name: Pack_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_types" (type_id, name) FROM stdin;
\.


--
-- TOC entry 3093 (class 0 OID 16715)
-- Dependencies: 212
-- Data for Name: Pay_conds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pay_conds" (pay_conds_id, name) FROM stdin;
\.


--
-- TOC entry 3081 (class 0 OID 16649)
-- Dependencies: 200
-- Data for Name: Specs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Specs" (spec_id, name) FROM stdin;
\.


--
-- TOC entry 3107 (class 0 OID 16843)
-- Dependencies: 226
-- Data for Name: Sup_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Sup_list" (o_id, sup_id) FROM stdin;
\.


--
-- TOC entry 3095 (class 0 OID 16726)
-- Dependencies: 214
-- Data for Name: Task_names; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Task_names" (name_id, name) FROM stdin;
\.


--
-- TOC entry 3091 (class 0 OID 16704)
-- Dependencies: 210
-- Data for Name: Tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tasks" (task_id, name) FROM stdin;
\.


--
-- TOC entry 3098 (class 0 OID 16747)
-- Dependencies: 217
-- Data for Name: Tech_docs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tech_docs" (tz_id, file_name) FROM stdin;
\.


--
-- TOC entry 3104 (class 0 OID 16783)
-- Dependencies: 223
-- Data for Name: Techs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Techs" (tz_id, date, o_id, end_date, proj, group_id, kind_id, type_id, task_id, tz_st, tender_st, count, cp_st, pay_cond_id, private, info, history) FROM stdin;
\.


--
-- TOC entry 3106 (class 0 OID 16824)
-- Dependencies: 225
-- Data for Name: Tenders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tenders" (tender_id, date, selected_cp, tz_id, history) FROM stdin;
\.


--
-- TOC entry 3129 (class 0 OID 0)
-- Dependencies: 196
-- Name: CP_cp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."CP_cp_id_seq"', 1, false);


--
-- TOC entry 3130 (class 0 OID 0)
-- Dependencies: 227
-- Name: Costs_cost_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Costs_cost_id_seq"', 1, false);


--
-- TOC entry 3131 (class 0 OID 0)
-- Dependencies: 201
-- Name: Countries_country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Countries_country_id_seq"', 1, false);


--
-- TOC entry 3132 (class 0 OID 0)
-- Dependencies: 215
-- Name: Metrics_metr_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Metrics_metr_id_seq"', 1, false);


--
-- TOC entry 3133 (class 0 OID 0)
-- Dependencies: 220
-- Name: Orgs_o_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Orgs_o_id_seq"', 1, false);


--
-- TOC entry 3134 (class 0 OID 0)
-- Dependencies: 203
-- Name: Pack_groups_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_groups_group_id_seq"', 1, false);


--
-- TOC entry 3135 (class 0 OID 0)
-- Dependencies: 207
-- Name: Pack_kinds_kind_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_kinds_kind_id_seq"', 1, false);


--
-- TOC entry 3136 (class 0 OID 0)
-- Dependencies: 205
-- Name: Pack_types_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_types_type_id_seq"', 1, false);


--
-- TOC entry 3137 (class 0 OID 0)
-- Dependencies: 211
-- Name: Pay_conds_pay_conds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pay_conds_pay_conds_id_seq"', 1, false);


--
-- TOC entry 3138 (class 0 OID 0)
-- Dependencies: 199
-- Name: Specs_spec_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Specs_spec_id_seq"', 1, false);


--
-- TOC entry 3139 (class 0 OID 0)
-- Dependencies: 213
-- Name: Task_names_name_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Task_names_name_id_seq"', 1, false);


--
-- TOC entry 3140 (class 0 OID 0)
-- Dependencies: 209
-- Name: Tasks_task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Tasks_task_id_seq"', 1, false);


--
-- TOC entry 3141 (class 0 OID 0)
-- Dependencies: 222
-- Name: Techs_tz_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Techs_tz_id_seq"', 1, false);


--
-- TOC entry 3142 (class 0 OID 0)
-- Dependencies: 224
-- Name: Tenders_tender_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Tenders_tender_id_seq"', 1, false);


--
-- TOC entry 2904 (class 2606 OID 16627)
-- Name: CP CP_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP"
    ADD CONSTRAINT "CP_pkey" PRIMARY KEY (cp_id);


--
-- TOC entry 2940 (class 2606 OID 16863)
-- Name: Costs Costs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT "Costs_pkey" PRIMARY KEY (cost_id);


--
-- TOC entry 2910 (class 2606 OID 16668)
-- Name: Countries Countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Countries"
    ADD CONSTRAINT "Countries_pkey" PRIMARY KEY (country_id);


--
-- TOC entry 2924 (class 2606 OID 16745)
-- Name: Metrics Metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Metrics"
    ADD CONSTRAINT "Metrics_pkey" PRIMARY KEY (metr_id);


--
-- TOC entry 2928 (class 2606 OID 16759)
-- Name: Org_countries Org_countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Org_countries"
    ADD CONSTRAINT "Org_countries_pkey" PRIMARY KEY (o_id, country_id);


--
-- TOC entry 2932 (class 2606 OID 16775)
-- Name: Orgs Orgs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs"
    ADD CONSTRAINT "Orgs_pkey" PRIMARY KEY (o_id);


--
-- TOC entry 2930 (class 2606 OID 16764)
-- Name: Orgs_specs Orgs_specs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs_specs"
    ADD CONSTRAINT "Orgs_specs_pkey" PRIMARY KEY (o_id, spec_id);


--
-- TOC entry 2912 (class 2606 OID 16679)
-- Name: Pack_groups Pack_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_groups"
    ADD CONSTRAINT "Pack_groups_pkey" PRIMARY KEY (group_id);


--
-- TOC entry 2916 (class 2606 OID 16701)
-- Name: Pack_kinds Pack_kinds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_kinds"
    ADD CONSTRAINT "Pack_kinds_pkey" PRIMARY KEY (kind_id);


--
-- TOC entry 2914 (class 2606 OID 16690)
-- Name: Pack_types Pack_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_types"
    ADD CONSTRAINT "Pack_types_pkey" PRIMARY KEY (type_id);


--
-- TOC entry 2920 (class 2606 OID 16723)
-- Name: Pay_conds Pay_conds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pay_conds"
    ADD CONSTRAINT "Pay_conds_pkey" PRIMARY KEY (pay_conds_id);


--
-- TOC entry 2908 (class 2606 OID 16657)
-- Name: Specs Specs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Specs"
    ADD CONSTRAINT "Specs_pkey" PRIMARY KEY (spec_id);


--
-- TOC entry 2938 (class 2606 OID 16847)
-- Name: Sup_list Sup_list_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Sup_list"
    ADD CONSTRAINT "Sup_list_pkey" PRIMARY KEY (o_id, sup_id);


--
-- TOC entry 2922 (class 2606 OID 16734)
-- Name: Task_names Task_names_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_names"
    ADD CONSTRAINT "Task_names_pkey" PRIMARY KEY (name_id);


--
-- TOC entry 2918 (class 2606 OID 16712)
-- Name: Tasks Tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tasks"
    ADD CONSTRAINT "Tasks_pkey" PRIMARY KEY (task_id);


--
-- TOC entry 2926 (class 2606 OID 16754)
-- Name: Tech_docs Tech_docs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tech_docs"
    ADD CONSTRAINT "Tech_docs_pkey" PRIMARY KEY (tz_id, file_name);


--
-- TOC entry 2934 (class 2606 OID 16791)
-- Name: Techs Techs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT "Techs_pkey" PRIMARY KEY (tz_id);


--
-- TOC entry 2936 (class 2606 OID 16832)
-- Name: Tenders Tenders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders"
    ADD CONSTRAINT "Tenders_pkey" PRIMARY KEY (tender_id);


--
-- TOC entry 2906 (class 2606 OID 16641)
-- Name: CP_docs cp_docs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP_docs"
    ADD CONSTRAINT cp_docs_pk PRIMARY KEY (file_name, cp_id);


--
-- TOC entry 2941 (class 2606 OID 16642)
-- Name: CP_docs cp_docs_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP_docs"
    ADD CONSTRAINT cp_docs_fk FOREIGN KEY (cp_id) REFERENCES public."CP"(cp_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2949 (class 2606 OID 16833)
-- Name: Tenders cp_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders"
    ADD CONSTRAINT cp_fk FOREIGN KEY (selected_cp) REFERENCES public."CP"(cp_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2955 (class 2606 OID 16879)
-- Name: Costs cp_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT cp_fk FOREIGN KEY (cp_id) REFERENCES public."CP"(cp_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2951 (class 2606 OID 16848)
-- Name: Sup_list fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Sup_list"
    ADD CONSTRAINT fk FOREIGN KEY (o_id) REFERENCES public."Orgs"(o_id);


--
-- TOC entry 2942 (class 2606 OID 16776)
-- Name: Orgs gr_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs"
    ADD CONSTRAINT gr_fk FOREIGN KEY (group_id) REFERENCES public."Pack_groups"(group_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2943 (class 2606 OID 16792)
-- Name: Techs gr_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT gr_fk FOREIGN KEY (group_id) REFERENCES public."Pack_groups"(group_id);


--
-- TOC entry 2945 (class 2606 OID 16802)
-- Name: Techs kind_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT kind_fk FOREIGN KEY (kind_id) REFERENCES public."Pack_kinds"(kind_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2953 (class 2606 OID 16869)
-- Name: Costs metr_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT metr_fk FOREIGN KEY (metr_id) REFERENCES public."Metrics"(metr_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2952 (class 2606 OID 16864)
-- Name: Costs name_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT name_fk FOREIGN KEY (name_id) REFERENCES public."Task_names"(name_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2944 (class 2606 OID 16797)
-- Name: Techs o_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT o_fk FOREIGN KEY (o_id) REFERENCES public."Orgs"(o_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2948 (class 2606 OID 16817)
-- Name: Techs pay_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT pay_fk FOREIGN KEY (pay_cond_id) REFERENCES public."Pay_conds"(pay_conds_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2947 (class 2606 OID 16812)
-- Name: Techs task_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT task_fk FOREIGN KEY (task_id) REFERENCES public."Tasks"(task_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2946 (class 2606 OID 16807)
-- Name: Techs type_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT type_fk FOREIGN KEY (type_id) REFERENCES public."Pack_types"(type_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2950 (class 2606 OID 16838)
-- Name: Tenders tz_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders"
    ADD CONSTRAINT tz_fk FOREIGN KEY (tz_id) REFERENCES public."Techs"(tz_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2954 (class 2606 OID 16874)
-- Name: Costs tz_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT tz_fk FOREIGN KEY (tz_id) REFERENCES public."Techs"(tz_id) ON UPDATE CASCADE ON DELETE SET NULL;


-- Completed on 2021-05-17 09:38:45 MSK

--
-- PostgreSQL database dump complete
--

