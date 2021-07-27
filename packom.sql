--
-- PostgreSQL database dump
--

-- Dumped from database version 11.12
-- Dumped by pg_dump version 11.12

-- Started on 2021-07-27 09:19:02

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
-- TOC entry 1 (class 3079 OID 16384)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 3065 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 16395)
-- Name: CP; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."CP" (
    cp_id integer NOT NULL,
    date date NOT NULL,
    tz_id integer NOT NULL,
    proj text,
    o_id integer NOT NULL,
    pay_cond_id integer NOT NULL,
    end_date date NOT NULL,
    info text,
    history text,
    cp_st integer NOT NULL
);


ALTER TABLE public."CP" OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 16401)
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
-- TOC entry 3066 (class 0 OID 0)
-- Dependencies: 198
-- Name: CP_cp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."CP_cp_id_seq" OWNED BY public."CP".cp_id;


--
-- TOC entry 199 (class 1259 OID 16403)
-- Name: CP_docs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."CP_docs" (
    file_name text NOT NULL,
    cp_id integer NOT NULL,
    active boolean DEFAULT true
);


ALTER TABLE public."CP_docs" OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 16753)
-- Name: Calendar; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Calendar" (
    cal_id integer NOT NULL,
    name_id integer NOT NULL,
    period integer NOT NULL,
    term integer NOT NULL,
    tz_id integer,
    cp_id integer,
    active boolean,
    cp_period integer DEFAULT 0
);


ALTER TABLE public."Calendar" OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 16751)
-- Name: Calendar_cal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Calendar_cal_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Calendar_cal_id_seq" OWNER TO postgres;

--
-- TOC entry 3067 (class 0 OID 0)
-- Dependencies: 229
-- Name: Calendar_cal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Calendar_cal_id_seq" OWNED BY public."Calendar".cal_id;


--
-- TOC entry 200 (class 1259 OID 16409)
-- Name: Costs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Costs" (
    cost_id integer NOT NULL,
    metr_id integer,
    count integer,
    tz_id integer,
    cp_id integer,
    ppu numeric,
    info text,
    task_id integer NOT NULL,
    sum numeric,
    active boolean
);


ALTER TABLE public."Costs" OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16415)
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
-- TOC entry 3068 (class 0 OID 0)
-- Dependencies: 201
-- Name: Costs_cost_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Costs_cost_id_seq" OWNED BY public."Costs".cost_id;


--
-- TOC entry 202 (class 1259 OID 16417)
-- Name: Countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Countries" (
    country_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Countries" OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16423)
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
-- TOC entry 3069 (class 0 OID 0)
-- Dependencies: 203
-- Name: Countries_country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Countries_country_id_seq" OWNED BY public."Countries".country_id;


--
-- TOC entry 204 (class 1259 OID 16425)
-- Name: Metrics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Metrics" (
    metr_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Metrics" OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16431)
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
-- TOC entry 3070 (class 0 OID 0)
-- Dependencies: 205
-- Name: Metrics_metr_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Metrics_metr_id_seq" OWNED BY public."Metrics".metr_id;


--
-- TOC entry 206 (class 1259 OID 16433)
-- Name: Org_countries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Org_countries" (
    o_id integer NOT NULL,
    country_id integer NOT NULL
);


ALTER TABLE public."Org_countries" OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 16436)
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
-- TOC entry 232 (class 1259 OID 24596)
-- Name: Orgs_docs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Orgs_docs" (
    doc_id integer NOT NULL,
    file_path text NOT NULL,
    file_name text NOT NULL,
    o_id integer NOT NULL
);


ALTER TABLE public."Orgs_docs" OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 24594)
-- Name: Orgs_docs_doc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Orgs_docs_doc_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Orgs_docs_doc_id_seq" OWNER TO postgres;

--
-- TOC entry 3071 (class 0 OID 0)
-- Dependencies: 231
-- Name: Orgs_docs_doc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Orgs_docs_doc_id_seq" OWNED BY public."Orgs_docs".doc_id;


--
-- TOC entry 208 (class 1259 OID 16442)
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
-- TOC entry 3072 (class 0 OID 0)
-- Dependencies: 208
-- Name: Orgs_o_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Orgs_o_id_seq" OWNED BY public."Orgs".o_id;


--
-- TOC entry 233 (class 1259 OID 24610)
-- Name: Orgs_orgs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Orgs_orgs" (
    o_id integer,
    f_o_id integer,
    active boolean DEFAULT true
);


ALTER TABLE public."Orgs_orgs" OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16444)
-- Name: Orgs_specs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Orgs_specs" (
    o_id integer NOT NULL,
    spec_id integer NOT NULL,
    active boolean DEFAULT true NOT NULL
);


ALTER TABLE public."Orgs_specs" OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16447)
-- Name: Pack_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_groups" (
    group_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_groups" OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16453)
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
-- TOC entry 3073 (class 0 OID 0)
-- Dependencies: 211
-- Name: Pack_groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_groups_group_id_seq" OWNED BY public."Pack_groups".group_id;


--
-- TOC entry 212 (class 1259 OID 16455)
-- Name: Pack_kinds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_kinds" (
    kind_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_kinds" OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16461)
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
-- TOC entry 3074 (class 0 OID 0)
-- Dependencies: 213
-- Name: Pack_kinds_kind_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_kinds_kind_id_seq" OWNED BY public."Pack_kinds".kind_id;


--
-- TOC entry 214 (class 1259 OID 16463)
-- Name: Pack_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pack_types" (
    type_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pack_types" OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16469)
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
-- TOC entry 3075 (class 0 OID 0)
-- Dependencies: 215
-- Name: Pack_types_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pack_types_type_id_seq" OWNED BY public."Pack_types".type_id;


--
-- TOC entry 227 (class 1259 OID 16701)
-- Name: Pay_conds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Pay_conds" (
    pay_cond_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Pay_conds" OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16699)
-- Name: Pay_conds_pay_cond_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Pay_conds_pay_cond_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Pay_conds_pay_cond_id_seq" OWNER TO postgres;

--
-- TOC entry 3076 (class 0 OID 0)
-- Dependencies: 226
-- Name: Pay_conds_pay_cond_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Pay_conds_pay_cond_id_seq" OWNED BY public."Pay_conds".pay_cond_id;


--
-- TOC entry 216 (class 1259 OID 16479)
-- Name: Specs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Specs" (
    spec_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Specs" OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16485)
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
-- TOC entry 3077 (class 0 OID 0)
-- Dependencies: 217
-- Name: Specs_spec_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Specs_spec_id_seq" OWNED BY public."Specs".spec_id;


--
-- TOC entry 235 (class 1259 OID 32814)
-- Name: Task_kinds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Task_kinds" (
    tk_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Task_kinds" OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 32812)
-- Name: Task_kinds_tk_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Task_kinds_tk_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Task_kinds_tk_id_seq" OWNER TO postgres;

--
-- TOC entry 3078 (class 0 OID 0)
-- Dependencies: 234
-- Name: Task_kinds_tk_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_kinds_tk_id_seq" OWNED BY public."Task_kinds".tk_id;


--
-- TOC entry 218 (class 1259 OID 16490)
-- Name: Task_names; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Task_names" (
    name_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Task_names" OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16496)
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
-- TOC entry 3079 (class 0 OID 0)
-- Dependencies: 219
-- Name: Task_names_name_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_names_name_id_seq" OWNED BY public."Task_names".name_id;


--
-- TOC entry 225 (class 1259 OID 16685)
-- Name: Tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tasks" (
    task_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public."Tasks" OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16683)
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
-- TOC entry 3080 (class 0 OID 0)
-- Dependencies: 224
-- Name: Tasks_task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tasks_task_id_seq" OWNED BY public."Tasks".task_id;


--
-- TOC entry 220 (class 1259 OID 16506)
-- Name: Tech_docs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tech_docs" (
    tz_id integer NOT NULL,
    file_name text NOT NULL,
    active boolean DEFAULT true NOT NULL
);


ALTER TABLE public."Tech_docs" OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16512)
-- Name: Techs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Techs" (
    date date NOT NULL,
    o_id integer NOT NULL,
    end_date date NOT NULL,
    proj text,
    group_id integer NOT NULL,
    kind_id integer NOT NULL,
    type_id integer NOT NULL,
    tender_st integer,
    cp_st integer,
    pay_cond_id integer NOT NULL,
    private boolean NOT NULL,
    info text,
    history text,
    tz_id integer NOT NULL,
    task_name text,
    active boolean DEFAULT true NOT NULL,
    tz_st integer DEFAULT 0 NOT NULL,
    selected_cp integer DEFAULT 0 NOT NULL
);


ALTER TABLE public."Techs" OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16715)
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
-- TOC entry 3081 (class 0 OID 0)
-- Dependencies: 228
-- Name: Techs_tz_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Techs_tz_id_seq" OWNED BY public."Techs".tz_id;


--
-- TOC entry 222 (class 1259 OID 16520)
-- Name: Tenders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tenders" (
    tender_id integer NOT NULL,
    date date NOT NULL,
    selected_cp integer,
    tz_id integer NOT NULL,
    history text
);


ALTER TABLE public."Tenders" OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16526)
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
-- TOC entry 3082 (class 0 OID 0)
-- Dependencies: 223
-- Name: Tenders_tender_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tenders_tender_id_seq" OWNED BY public."Tenders".tender_id;


--
-- TOC entry 2820 (class 2604 OID 16528)
-- Name: CP cp_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP" ALTER COLUMN cp_id SET DEFAULT nextval('public."CP_cp_id_seq"'::regclass);


--
-- TOC entry 2840 (class 2604 OID 16756)
-- Name: Calendar cal_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Calendar" ALTER COLUMN cal_id SET DEFAULT nextval('public."Calendar_cal_id_seq"'::regclass);


--
-- TOC entry 2822 (class 2604 OID 16529)
-- Name: Costs cost_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs" ALTER COLUMN cost_id SET DEFAULT nextval('public."Costs_cost_id_seq"'::regclass);


--
-- TOC entry 2823 (class 2604 OID 16530)
-- Name: Countries country_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Countries" ALTER COLUMN country_id SET DEFAULT nextval('public."Countries_country_id_seq"'::regclass);


--
-- TOC entry 2824 (class 2604 OID 16531)
-- Name: Metrics metr_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Metrics" ALTER COLUMN metr_id SET DEFAULT nextval('public."Metrics_metr_id_seq"'::regclass);


--
-- TOC entry 2825 (class 2604 OID 16532)
-- Name: Orgs o_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs" ALTER COLUMN o_id SET DEFAULT nextval('public."Orgs_o_id_seq"'::regclass);


--
-- TOC entry 2842 (class 2604 OID 24599)
-- Name: Orgs_docs doc_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs_docs" ALTER COLUMN doc_id SET DEFAULT nextval('public."Orgs_docs_doc_id_seq"'::regclass);


--
-- TOC entry 2827 (class 2604 OID 16533)
-- Name: Pack_groups group_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_groups" ALTER COLUMN group_id SET DEFAULT nextval('public."Pack_groups_group_id_seq"'::regclass);


--
-- TOC entry 2828 (class 2604 OID 16534)
-- Name: Pack_kinds kind_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_kinds" ALTER COLUMN kind_id SET DEFAULT nextval('public."Pack_kinds_kind_id_seq"'::regclass);


--
-- TOC entry 2829 (class 2604 OID 16535)
-- Name: Pack_types type_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_types" ALTER COLUMN type_id SET DEFAULT nextval('public."Pack_types_type_id_seq"'::regclass);


--
-- TOC entry 2839 (class 2604 OID 16704)
-- Name: Pay_conds pay_cond_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pay_conds" ALTER COLUMN pay_cond_id SET DEFAULT nextval('public."Pay_conds_pay_cond_id_seq"'::regclass);


--
-- TOC entry 2830 (class 2604 OID 16537)
-- Name: Specs spec_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Specs" ALTER COLUMN spec_id SET DEFAULT nextval('public."Specs_spec_id_seq"'::regclass);


--
-- TOC entry 2844 (class 2604 OID 32817)
-- Name: Task_kinds tk_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_kinds" ALTER COLUMN tk_id SET DEFAULT nextval('public."Task_kinds_tk_id_seq"'::regclass);


--
-- TOC entry 2831 (class 2604 OID 16538)
-- Name: Task_names name_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_names" ALTER COLUMN name_id SET DEFAULT nextval('public."Task_names_name_id_seq"'::regclass);


--
-- TOC entry 2838 (class 2604 OID 16688)
-- Name: Tasks task_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tasks" ALTER COLUMN task_id SET DEFAULT nextval('public."Tasks_task_id_seq"'::regclass);


--
-- TOC entry 2833 (class 2604 OID 16717)
-- Name: Techs tz_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs" ALTER COLUMN tz_id SET DEFAULT nextval('public."Techs_tz_id_seq"'::regclass);


--
-- TOC entry 2837 (class 2604 OID 16541)
-- Name: Tenders tender_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders" ALTER COLUMN tender_id SET DEFAULT nextval('public."Tenders_tender_id_seq"'::regclass);


--
-- TOC entry 3021 (class 0 OID 16395)
-- Dependencies: 197
-- Data for Name: CP; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."CP" (cp_id, date, tz_id, proj, o_id, pay_cond_id, end_date, info, history, cp_st) FROM stdin;
23	2021-06-14	45	8	22	1	2021-12-05	1	 \n Обновлен график: Разработка концепта длительностью 1 кн. Дата: 2021-06-15 \n Обновлен график: Изготовление серии длительностью 2 кн. Дата: 2021-06-15 \n Обновлена стоимость: Изготовление серии в количестве undefined undefined. Дата: 2021-06-15 \n Обновлена стоимость: Доставка в количестве undefined undefined. Дата: 2021-06-15 \n Обновлена стоимость: Единичный в количестве undefined undefined. Дата: 2021-06-15 \n Изменена конечная дата: 2021-12-05 Дата: 2021-06-15 \n Изменена общая информация: 1 Дата: 2021-06-15	1
24	2021-07-27	47	ГК 6059	22	2	2021-08-23	4	 \n Изменены условия оплаты: 100% предоплата Дата: 2021-07-27 \n Изменена общая информация: 4 Дата: 2021-07-27	1
27	2021-07-27	49	2	22	1	2021-09-19			1
28	2021-07-27	49	2	4	3	2021-08-19	1		1
\.


--
-- TOC entry 3023 (class 0 OID 16403)
-- Dependencies: 199
-- Data for Name: CP_docs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."CP_docs" (file_name, cp_id, active) FROM stdin;
1.docx	23	t
pgadmin.log	23	t
\.


--
-- TOC entry 3054 (class 0 OID 16753)
-- Dependencies: 230
-- Data for Name: Calendar; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Calendar" (cal_id, name_id, period, term, tz_id, cp_id, active, cp_period) FROM stdin;
1	1	1	7	24	\N	t	0
2	2	2	9	24	\N	t	0
3	3	1	10	24	\N	t	0
4	1	1	7	25	\N	t	0
5	2	2	9	25	\N	t	0
6	3	1	10	25	\N	t	0
16	1	1	7	26	\N	t	0
17	2	2	9	26	\N	t	0
18	3	1	10	26	\N	t	0
25	1	1	49	27	\N	t	0
26	4	3	52	27	\N	t	0
27	1	1	49	28	\N	t	0
28	5	2	51	28	\N	t	0
66	1	1	49	36	\N	t	0
67	2	3	52	36	\N	t	0
68	3	1	53	36	\N	t	0
69	1	3	51	37	\N	t	0
70	3	1	52	37	\N	t	0
88	1	1	42	44	\N	t	0
89	2	2	44	44	\N	t	0
90	3	1	45	44	\N	t	0
29	1	1	49	29	\N	f	0
30	3	2	51	29	\N	f	0
31	2	1	0	29	\N	t	0
32	3	1	0	29	\N	t	0
33	1	2	0	29	\N	t	0
34	3	5	0	30	\N	t	0
91	1	1	49	45	\N	f	0
92	2	2	51	45	\N	f	0
96	2	2	0	45	\N	f	0
97	2	2	0	45	\N	f	0
58	2	1	49	35	\N	t	0
61	3	2	0	35	\N	t	0
119	1	2	36	47	\N	t	0
120	4	3	39	47	\N	t	0
99	1	2	0	45	\N	t	0
121	5	1	40	47	\N	t	0
95	3	1	0	45	\N	f	0
100	3	1	0	45	\N	f	0
107	2	3	0	45	\N	t	0
204	1	1	34	48	\N	t	0
125	2	2	35	2	\N	t	0
126	4	3	38	2	\N	t	0
127	5	1	39	2	\N	t	0
205	2	2	36	48	\N	f	0
206	5	1	37	48	\N	f	0
207	1	2	43	46	\N	t	0
128	1	2	35	2	\N	t	0
208	2	5	48	46	\N	t	0
209	5	1	49	46	\N	t	0
222	1	2	39	49	\N	t	0
223	2	5	44	49	\N	t	0
224	3	2	46	49	\N	t	0
225	4	1	47	49	\N	t	0
226	5	2	49	49	\N	t	0
7	1	0	6	\N	1	t	0
8	2	2	8	\N	1	t	2
9	3	1	9	\N	1	t	1
10	1	0	6	\N	2	t	0
11	2	2	8	\N	2	t	2
12	3	1	9	\N	2	t	1
13	1	1	6	\N	3	t	1
14	2	1	8	\N	3	t	1
15	3	1	9	\N	3	t	1
19	1	1	6	\N	4	t	1
20	2	1	8	\N	4	t	1
21	3	1	9	\N	4	t	1
22	1	1	6	\N	5	t	1
23	2	1	8	\N	5	t	1
24	3	1	9	\N	5	t	1
71	1	2	0	\N	21	t	2
72	2	2	0	\N	21	t	2
73	3	2	0	\N	21	t	2
35	0	1	7	\N	13	t	1
36	0	2	9	\N	13	t	2
37	0	1	10	\N	13	t	1
38	0	1	7	\N	14	t	1
39	0	2	9	\N	14	t	2
40	0	1	10	\N	14	t	1
41	3	0	0	\N	15	t	0
42	3	0	0	\N	16	t	0
43	1	5	0	\N	17	t	5
44	2	1	0	\N	17	t	1
45	3	1	0	\N	17	t	1
46	1	5	0	\N	18	f	5
47	2	1	0	\N	18	f	1
48	3	1	0	\N	18	f	1
49	1	5	0	\N	18	f	5
50	2	1	0	\N	18	f	1
51	3	1	0	\N	18	f	1
52	1	6	0	\N	18	f	6
53	2	0	0	\N	18	f	0
54	3	1	0	\N	18	f	1
55	1	6	0	\N	18	t	6
56	2	0	0	\N	18	t	0
57	3	1	0	\N	18	t	1
59	2	2	0	\N	19	f	2
64	2	2	0	\N	19	t	2
65	3	3	0	\N	19	t	3
93	1	0	0	\N	23	f	0
94	2	1	0	\N	23	f	1
98	3	1	0	\N	23	f	1
60	2	2	0	\N	20	f	2
62	2	0	0	\N	20	f	0
63	3	3	0	\N	20	f	3
78	2	1	0	\N	20	f	1
79	3	0	0	\N	20	f	0
80	2	1	0	\N	20	f	1
81	3	2	0	\N	20	f	2
82	2	0	0	\N	20	f	0
83	3	0	0	\N	20	f	0
84	2	4	0	\N	20	t	4
85	3	4	0	\N	20	t	4
86	1	1	0	\N	22	t	1
87	3	1	0	\N	22	t	1
101	3	0	0	\N	23	f	0
102	1	1	0	\N	23	f	1
103	3	3	0	\N	23	f	3
104	3	1	0	\N	23	f	1
105	1	2	0	\N	23	f	2
106	3	1	0	\N	23	f	1
108	1	1	0	\N	23	f	1
109	2	2	0	\N	23	f	2
110	1	2	0	\N	23	f	2
111	2	3	0	\N	23	f	3
112	1	1	0	\N	23	f	1
113	2	2	0	\N	23	f	2
114	1	1	0	\N	23	t	1
115	2	2	0	\N	23	t	2
210	1	1	0	\N	24	f	1
211	4	2	0	\N	24	f	2
212	5	1	0	\N	24	f	1
213	1	1	0	\N	24	f	1
214	4	2	0	\N	24	f	2
215	5	1	0	\N	24	f	1
216	1	2	0	\N	24	f	2
217	4	3	0	\N	24	f	3
218	5	1	0	\N	24	f	1
219	1	1	0	\N	24	t	1
220	4	2	0	\N	24	t	2
221	5	1	0	\N	24	t	1
242	1	2	0	0	27	t	4
243	2	5	0	0	27	t	10
244	3	2	0	0	27	t	4
245	4	1	0	0	27	t	2
246	5	2	0	0	27	t	4
247	1	2	0	49	28	t	1
248	2	5	0	49	28	t	1
249	3	2	0	49	28	t	1
250	4	1	0	49	28	t	1
251	5	2	0	49	28	t	1
\.


--
-- TOC entry 3024 (class 0 OID 16409)
-- Dependencies: 200
-- Data for Name: Costs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Costs" (cost_id, metr_id, count, tz_id, cp_id, ppu, info, task_id, sum, active) FROM stdin;
28	\N	\N	\N	\N	5000	1	3	\N	t
29	\N	\N	\N	\N	50	1	1	\N	t
30	\N	\N	\N	\N	1	1	3	\N	t
31	\N	\N	\N	\N	1	1	1	\N	t
43	\N	\N	\N	\N	170		2	\N	f
50	\N	\N	\N	\N	100		2	\N	t
51	\N	\N	\N	\N	80		1	\N	t
52	\N	\N	\N	\N	5000		3	\N	t
44	\N	\N	\N	\N	900		2	\N	f
47	\N	\N	\N	\N	5000		2	\N	f
48	\N	\N	\N	\N	250		1	\N	f
49	\N	\N	\N	\N	600		3	\N	f
64	\N	\N	\N	\N	500		2	\N	f
65	\N	\N	\N	\N	200		1	\N	f
15	1	120	\N	\N	\N		4	\N	t
16	3	1	\N	\N	\N		2	\N	t
17	1	1	\N	\N	\N		4	\N	t
18	3	15	\N	\N	\N		5	\N	t
32	\N	\N	\N	\N	500	1	1	\N	t
33	\N	\N	\N	\N	3000	1	3	\N	t
34	\N	\N	\N	\N	500	1	1	\N	f
35	\N	\N	\N	\N	3000	1	3	\N	f
36	\N	\N	\N	\N	500		1	\N	f
37	\N	\N	\N	\N	3000		3	\N	f
38	\N	\N	\N	\N	500		1	\N	f
39	\N	\N	\N	\N	3000		3	\N	f
66	\N	\N	\N	\N	1		3	\N	f
67	\N	\N	\N	\N	500		2	\N	f
68	\N	\N	\N	\N	2		1	\N	f
40	\N	\N	\N	\N	500		1	\N	t
41	\N	\N	\N	\N	3000		3	\N	t
2	1	150	\N	\N	\N		1	\N	t
3	2	1	\N	\N	\N		3	\N	t
9	1	10	\N	\N	\N		1	\N	t
10	2	3	\N	\N	\N		3	\N	t
19	1	5	\N	\N	\N		1	\N	f
20	1	10	\N	\N	\N		1	\N	f
21	1	15	\N	\N	\N		1	\N	t
22	2	1	\N	\N	\N		3	\N	t
23	2	1	\N	\N	\N		3	\N	t
24	1	150	\N	\N	\N		1	\N	t
69	\N	\N	\N	\N	2		3	\N	f
70	\N	\N	\N	\N	5		2	\N	f
71	\N	\N	\N	\N	5		1	\N	f
72	\N	\N	\N	\N	5		3	\N	f
73	\N	\N	\N	\N	4		2	\N	t
74	\N	\N	\N	\N	4		1	\N	t
75	\N	\N	\N	\N	4		3	\N	t
58	\N	\N	\N	\N	250		1	\N	t
5	1	\N	\N	\N	150.00		1	\N	t
6	2	\N	\N	\N	3000.00		3	\N	t
7	1	\N	\N	\N	100.00		1	\N	t
8	2	\N	\N	\N	5000.00		3	\N	t
11	1	\N	\N	\N	100.00		1	\N	t
12	2	\N	\N	\N	5000.00		3	\N	t
13	1	\N	\N	\N	100.00		1	\N	t
14	2	\N	\N	\N	5000.00		3	\N	t
26	\N	\N	\N	\N	80		1	\N	t
27	\N	\N	\N	\N	8000		3	\N	t
59	\N	\N	\N	\N	500		3	\N	t
76	\N	\N	\N	\N	500		1	\N	t
77	\N	\N	\N	\N	5		2	\N	t
78	\N	\N	\N	\N	1		3	\N	t
55	1	150	\N	\N	\N		1	\N	t
56	1	1	\N	\N	\N		2	\N	t
57	2	2	\N	\N	\N		3	\N	t
53	1	30	\N	\N	\N		1	\N	t
54	2	2	\N	\N	\N		3	\N	t
42	1	1	\N	\N	\N		2	\N	t
45	1	150	\N	\N	\N		1	\N	t
46	2	2	\N	\N	\N		3	\N	t
79	1	10	\N	\N	\N		1	\N	t
80	2	1	\N	\N	\N		3	\N	t
100	2	1	45	\N	\N		3	\N	f
97	2	1	45	\N	\N		3	\N	f
82	2	1	45	\N	\N		3	\N	f
87	2	1	45	\N	\N		3	\N	f
81	1	80	45	\N	\N		1	\N	f
91	4	50	45	\N	\N		1	\N	f
92	1	50	45	\N	\N		1	\N	f
94	1	80	45	\N	\N		1	\N	f
98	1	50	45	\N	\N		1	\N	t
96	1	1	45	\N	\N		2	\N	f
88	1	1	45	\N	\N		2	\N	f
85	1	1	45	\N	\N		2	\N	f
86	1	1	45	\N	\N		2	\N	f
90	1	1	45	\N	\N		2	\N	f
89	2	1	45	\N	\N		3	\N	f
95	2	1	45	\N	\N		3	\N	f
99	2	1	45	\N	\N		3	\N	f
264	1	1	49	\N	\N		1	\N	t
266	1	1	49	\N	\N		3	\N	t
102	2	1	45	\N	\N		3	\N	t
103	1	1	45	\N	\N		2	\N	t
268	1	500	49	\N	\N		6	\N	t
93	1	1	45	\N	\N		2	\N	f
83	\N	\N	\N	23	50		1	\N	f
84	\N	\N	\N	23	5000		3	\N	f
101	\N	\N	\N	23	500		1	\N	f
104	\N	\N	\N	23	50		1	\N	f
105	\N	\N	\N	23	3000		3	\N	f
106	\N	\N	\N	23	500		2	\N	f
107	\N	\N	\N	23	0		1	\N	f
108	\N	\N	\N	23	0		3	\N	f
109	\N	\N	\N	23	0		2	\N	f
110	\N	\N	\N	23	500		1	\N	f
111	\N	\N	\N	23	500		3	\N	f
112	\N	\N	\N	23	500		2	\N	f
113	\N	\N	\N	23	500		1	\N	f
114	\N	\N	\N	23	500		3	\N	f
115	\N	\N	\N	23	500		2	\N	f
244	3	150	48	\N	\N		8	\N	t
246	1	250	46	\N	\N		1	\N	t
248	\N	\N	\N	24	5000		1	\N	f
249	\N	\N	\N	24	2300		3	\N	f
250	\N	\N	\N	24	200		5	\N	f
251	\N	\N	\N	24	300		8	\N	f
256	\N	\N	\N	24	5000		1	\N	f
116	\N	\N	\N	23	50		1	\N	f
117	\N	\N	\N	23	500		3	\N	f
118	\N	\N	\N	23	10		2	\N	f
119	\N	\N	\N	23	50		1	\N	t
120	\N	\N	\N	23	1500		3	\N	t
121	\N	\N	\N	23	15		2	\N	t
265	1	20	49	\N	\N		2	\N	t
267	1	1	49	\N	\N		4	\N	t
124	1	1	47	\N	\N		1	\N	t
125	1	5	47	\N	\N		3	\N	t
126	1	15	47	\N	\N		5	\N	t
127	2	1	47	\N	\N		8	\N	t
269	2	2	49	\N	\N		8	\N	t
282	\N	\N	\N	27	2		4	\N	t
283	\N	\N	\N	27	2		3	\N	t
284	\N	\N	\N	27	2		1	\N	t
285	\N	\N	\N	27	4		8	\N	t
286	\N	\N	\N	27	40		2	\N	t
287	\N	\N	\N	27	1000		6	\N	t
288	\N	\N	\N	28	5000		1	\N	t
289	\N	\N	\N	28	5000		3	\N	t
290	\N	\N	\N	28	20		6	\N	t
291	\N	\N	\N	28	200		2	\N	t
292	\N	\N	\N	28	3000		4	\N	t
293	\N	\N	\N	28	5000		8	\N	t
243	1	1250	48	\N	\N		2	\N	f
245	1	20	48	\N	\N		7	\N	t
247	2	1	46	\N	\N		3	\N	t
257	\N	\N	\N	24	0		3	\N	f
258	\N	\N	\N	24	0		5	\N	f
259	\N	\N	\N	24	0		8	\N	f
252	\N	\N	\N	24	400	1	1	\N	f
253	\N	\N	\N	24	6000	1	3	\N	f
254	\N	\N	\N	24	2400	1	5	\N	f
255	\N	\N	\N	24	300	1	8	\N	f
260	\N	\N	\N	24	5000		1	\N	t
261	\N	\N	\N	24	2300		3	\N	t
262	\N	\N	\N	24	500		5	\N	t
263	\N	\N	\N	24	200		8	\N	t
\.


--
-- TOC entry 3026 (class 0 OID 16417)
-- Dependencies: 202
-- Data for Name: Countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Countries" (country_id, name) FROM stdin;
286	Japan
0	?
287	Belgium
4	Украина
2	Беларусь
1	Россия
5	Азербайджан
17	Abkhazia
18	Afghanistan
19	Albania
20	Algeria
21	Argentina
22	Armenia
23	Australia
24	Austria
25	Azerbaijan
26	Bahamas
27	Bahrain
28	Bangladesh
29	Belarus
30	Belgium
31	Bermuda Islands
32	Bolivia
33	Brazil
34	Bulgaria
35	Burundi
36	Cambodia
37	Cameroon
38	Canada
39	Chile
40	China
41	Colombia
42	Congo
43	Costa Rica
44	Cuba
45	Cyprus
46	Czech Republic
47	Denmark
48	Dominican Republic
49	Ecuador
50	Egypt
51	El Salvador
52	Estonia
53	Ethiopia
54	Finland
55	France
56	Georgia
57	Germany
58	Ghana
59	Gibraltar
60	Great Britain / United Kingdom
61	Greece
62	Guatemala
63	Guinea
64	Haiti
65	Hawaii
66	Honduras
67	Hong Kong
68	Hungary
69	Iceland
70	India
71	Indonesia
72	Iran
73	Iraq
74	Ireland
75	Israel
76	Italy
77	Ivory Coast
78	Jamaica
80	Kazakhstan
81	Kenya
82	Kuwait
83	Kyrgyzstan
84	Latvia
85	Lebanon
86	Liberia
87	Libya
88	Lithuania
89	Luxemburg
90	Madagascar
91	Malawi
92	Malaysia
93	Malta
94	Mexico
95	Moldova
96	Monaco
97	Mongolia
98	Morocco
99	Nepal
100	Netherlands / Holland
101	New Zeland
102	Nicaragua
103	Nigeria
104	North Korea
105	Norway
106	Oman
107	Pakistan
108	Panama
109	Papua New Guinea
110	Paraguay
111	Peru
112	Philippines
113	Poland
114	Portugal
115	Puerto Rico
116	Romania
117	Russia
118	Rwanda
119	Saudi Arabia
120	Senegal
121	Sierra Leone
122	Singapore
123	Slovakia
124	Slovenia
125	South Africa
126	South Korea
127	South Ossetia
128	Spain
129	Sri Lanka
130	Sudan
131	Sweden
132	Switzerland
133	Syria
134	Taiwan
135	Tajikistan
136	Thailand
137	Togo
138	Tunisia
139	Turkey
140	Turkmenistan
141	Uganda
142	Ukraine
143	United Arab Emirates
144	United States of America / USA
145	Uruguay
146	Uzbekistan
147	Venezuela
148	Yemen
149	Zaire
150	Zambia
285	Zimbabwe
\.


--
-- TOC entry 3028 (class 0 OID 16425)
-- Dependencies: 204
-- Data for Name: Metrics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Metrics" (metr_id, name) FROM stdin;
1	шт.
2	рейс
3	м.
4	м. кв.
5	партия
6	рулон
7	паллет
\.


--
-- TOC entry 3030 (class 0 OID 16433)
-- Dependencies: 206
-- Data for Name: Org_countries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Org_countries" (o_id, country_id) FROM stdin;
4	1
8	2
10	1
11	1
13	1
28	4
29	5
32	8
33	9
34	10
35	11
36	12
37	13
38	14
39	15
40	16
41	286
42	287
44	29
45	28
46	29
47	29
48	29
49	83
50	95
51	84
52	20
53	27
54	21
55	18
56	2
57	5
20	1
21	1
22	1
23	1
24	1
25	1
5	1
31	1
\.


--
-- TOC entry 3031 (class 0 OID 16436)
-- Dependencies: 207
-- Data for Name: Orgs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs" (o_id, name, group_id, site, phone, email, adress, info, login, hashed_pwd, status, history) FROM stdin;
52	56	2						56	657769727566687170336939343738356173646654ceb91256e8190e474aa752a6e0650a2df5ba37	f	
53	57	2						57	65776972756668717033693934373835617364669109c85a45b703f87f1413a405549a2cea9ab556	f	
54	58	2						58	6577697275666871703369393437383561736466667be543b02294b7624119adc3a725473df39885	f	
22	2	2	1	+7 919 037 04 81	daria	ул. Николо	q	2	6577697275666871703369393437383561736466da4b9237bacccdf19c0760cab7aec4a8359010b0	f	   \n Изменена общая информация: q Дата: 2021-07-13       \n Изменена общая информация: asdfasdf Дата: 2021-07-12   \n Изменена общая информация: " Дата: 2021-07-12                               \n Изменена специализация: Металлоконтейнеры Дата: 2021-06-14  \n Изменена специализация: Металлокерамика Дата: 2021-06-14  \n Изменен адрес: ул. Николо Дата: 2021-06-14  \n Изменен адрес: ул. Николо-Козинская Дата: 2021-06-14 \n Изменен телефон: +7 919 037 04 81 Дата: 2021-06-14 \n Изменен адрес электронной почты: daria Дата: 2021-06-14 \n Изменен сайт: 1 Дата: 2021-06-14 1
8	Организация 4\n	1	poga_i_kopyta.ru	8 800 555 35 35	mail@mail.ru	ул. Циолковского, 1a	?	daria40tim@gmail.com	6577697275666871703369393437383561736466a0f1490a20d0211c997b44bc357e1972deab8ae3	f	&
4	'Организация 1'	3	mysite.com	533286	daria140tim@gmail.com	ул. Николо		daria	657769727566687170336939343738356173646687920921c770b54d62fd553645a7fc579e1a9a7f	f	6
21	1	1	сайт	+7 919 037 04 81	daria40tim@gmail.com	ул. Николо-Козинская		1	6577697275666871703369393437383561736466356a192b7913b04c54574d18c28d46e6395428ab	f	 \n Изменен сайт: сайт Дата: 2021-06-14  \n Изменен сайт: 8 Дата: 2021-06-14
55	59	3						59	65776972756668717033693934373835617364665a5b0f9b7d3f8fc84c3cef8fd8efaaa6c70d75ab	f	
31	5	3						5	6577697275666871703369393437383561736466ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4	f	
10	Организация 10	1	main.com	+79190385798	mail@gmail.com	Адрес	?	?	?	f	
11	Организация 11	1	main.com	+79190385798	mail@gmail.com	Адрес	?	?	?	f	
13	Организация 13	2	main.com	+79190385798	mail@gmail.com	Адрес	?	?	?	f	
46	8	3						8	6577697275666871703369393437383561736466fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f	f	
23	3	1						3	657769727566687170336939343738356173646677de68daecd823babbb58edb1c8e14d7106e83bb	f	
47	9	3						10	65776972756668717033693934373835617364660ade7c2cf97f75d009975f4d720d1fa6c19f4897	f	
42	21	3						21	6577697275666871703369393437383561736466472b07b9fcf2c2451e8781e944bf5f77cd8457c8	f	
48	11	2						11	657769727566687170336939343738356173646617ba0791499db908433b80f37c5fbc89b870084b	f	
45	32	1						32	6577697275666871703369393437383561736466cb4e5208b4cd87268b208e49452ed6e89a68e0b8	f	
49	12	3						12	65776972756668717033693934373835617364667b52009b64fd0a2a49e6d8a939753077792b0554	f	
50	13	2						13	6577697275666871703369393437383561736466bd307a3ec329e10a2cff8fb87480823da114f8f4	f	
51	14	1						14	6577697275666871703369393437383561736466fa35e192121eabf3dabf9f5ea6abdbcbc107ac3b	f	
56	60	1						60	6577697275666871703369393437383561736466e6c3dd630428fd54834172b8fd2735fed9416da4	f	
57	62	3						62	6577697275666871703369393437383561736466511a418e72591eb7e33f703f04c3fa16df6c90bd	f	
\.


--
-- TOC entry 3056 (class 0 OID 24596)
-- Dependencies: 232
-- Data for Name: Orgs_docs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs_docs" (doc_id, file_path, file_name, o_id) FROM stdin;
3		1.txt	31
4		1.txt	31
5		Dz.docx	31
1	D:\\packom	Ustav.pdf	4
2	D:\\packom	Example.doc	4
6		1.docx	4
7		1.docx	21
8		pgadmin.log	21
\.


--
-- TOC entry 3057 (class 0 OID 24610)
-- Dependencies: 233
-- Data for Name: Orgs_orgs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs_orgs" (o_id, f_o_id, active) FROM stdin;
4	10	t
4	11	t
21	4	t
4	10	t
4	10	t
4	22	t
4	21	f
4	13	f
4	13	f
4	13	f
4	13	f
21	22	t
\.


--
-- TOC entry 3033 (class 0 OID 16444)
-- Dependencies: 209
-- Data for Name: Orgs_specs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Orgs_specs" (o_id, spec_id, active) FROM stdin;
46	0	t
47	0	t
10	1	t
48	0	t
13	1	t
49	0	t
50	0	t
51	0	t
52	0	t
53	0	t
54	0	t
55	0	t
56	0	t
8	3	t
11	3	t
57	0	t
19	0	t
20	0	t
23	0	t
24	0	t
25	0	t
4	0	t
30	0	t
31	1	t
21	2	t
32	0	t
33	0	t
34	0	t
35	0	t
36	0	t
37	0	t
38	0	t
39	0	t
40	0	t
41	0	t
42	0	t
44	0	t
45	0	t
22	15	f
22	19	f
22	17	f
22	18	f
22	20	f
22	3	f
22	22	f
22	21	f
22	23	f
22	1	t
22	2	t
\.


--
-- TOC entry 3034 (class 0 OID 16447)
-- Dependencies: 210
-- Data for Name: Pack_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_groups" (group_id, name) FROM stdin;
1	Гофрокороб
2	Металлоконтейнер
3	Пленка
4	Паллет деревянный
5	Паллет пластиковый
6	Евроконтейнер
7	Расходные материалы
8	Ящик деревянный
9	Крышка пластиковая
\.


--
-- TOC entry 3036 (class 0 OID 16455)
-- Dependencies: 212
-- Data for Name: Pack_kinds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_kinds" (kind_id, name) FROM stdin;
1	Стандартная
2	Специальная
\.


--
-- TOC entry 3038 (class 0 OID 16463)
-- Dependencies: 214
-- Data for Name: Pack_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pack_types" (type_id, name) FROM stdin;
1	Одноразовая
2	Многооборотная
\.


--
-- TOC entry 3051 (class 0 OID 16701)
-- Dependencies: 227
-- Data for Name: Pay_conds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Pay_conds" (pay_cond_id, name) FROM stdin;
1	100% постоплата
2	100% предоплата
3	30% предоплата
4	50% предоплата
5	постоплата каждой партии
\.


--
-- TOC entry 3040 (class 0 OID 16479)
-- Dependencies: 216
-- Data for Name: Specs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Specs" (spec_id, name) FROM stdin;
1	Гофроупаковка
2	Металлокерамика
3	Металлоконтейнеры
0	?
7	Гофрокороб
14	2
15	Кристина
16	1
17	q
18	w
19	e
20	r
21	t
22	y
23	u
\.


--
-- TOC entry 3059 (class 0 OID 32814)
-- Dependencies: 235
-- Data for Name: Task_kinds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Task_kinds" (tk_id, name) FROM stdin;
1	Разработка и прототип
2	Изготовление серии
3	Изготовление доп. к серии
4	Ремонт
5	Модификация
6	Утилизация
\.


--
-- TOC entry 3042 (class 0 OID 16490)
-- Dependencies: 218
-- Data for Name: Task_names; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Task_names" (name_id, name) FROM stdin;
1	Разработка концепта
2	Изготовление серии
0	
3	Изготовление прототипа
4	Доработка прототипа
5	Доставка
\.


--
-- TOC entry 3049 (class 0 OID 16685)
-- Dependencies: 225
-- Data for Name: Tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tasks" (task_id, name) FROM stdin;
1	Разработка концепта
2	Изготовление оснастки
3	Изготовление прототипа
4	Доработка прототипа
5	Изготовление предсерии
6	Изготовление серии
7	Модернизация серии
8	Доставка
\.


--
-- TOC entry 3044 (class 0 OID 16506)
-- Dependencies: 220
-- Data for Name: Tech_docs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tech_docs" (tz_id, file_name, active) FROM stdin;
35	?	t
43	1	t
43	2	t
36	1.doc	t
36	2.doc	t
45	1.docx	t
45	pgadmin.log	t
48	Концептуальная схема.png	t
\.


--
-- TOC entry 3045 (class 0 OID 16512)
-- Dependencies: 221
-- Data for Name: Techs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Techs" (date, o_id, end_date, proj, group_id, kind_id, type_id, tender_st, cp_st, pay_cond_id, private, info, history, tz_id, task_name, active, tz_st, selected_cp) FROM stdin;
2021-06-14	4	2021-12-05	12	2	1	2	0	0	3	f		 \n Добавлен график: Изготовление серии длительностью 3 кн. Дата: 2021-06-15 \n Добавлен график: Разработка концепта длительностью 2 кн. Дата: 2021-06-15 \n Добавлен график: Доставка длительностью 1 кн. Дата: 2021-06-15 \n Добавлена стоимость: Доставка в количестве 1 рейс. Дата: 2021-06-15 \n Добавлена стоимость: Единичный в количестве 1 шт.. Дата: 2021-06-15 \n Добавлена стоимость: Доставка в количестве 1 рейс. Дата: 2021-06-15 \n Добавлена стоимость: Доставка в количестве 1 рейс. Дата: 2021-06-15 \n Добавлена стоимость: Изготовление серии в количестве 50 шт.. Дата: 2021-06-15 \n 	45	Изготовление серии	t	0	0
2021-07-23	23	2021-08-23	ГК 6059	3	2	1	0	0	2	f	1		47	Разработка и прототип	t	0	0
2021-07-25	21	2021-08-19		2	2	2	0	0	3	f	1	 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-26 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-26 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-26 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-26 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-26 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25 \n Изменена конечная дата: 2021-08-19 Дата: 2021-07-25	48	Модификация	t	0	0
2021-06-15	21	2021-07-25	13	3	1	1	0	0	3	t		 \n Изменена конечная дата: 2021-07-25 Дата: 2021-07-26 \n Изменена конечная дата: 2021-10-15 Дата: 2021-07-25 \n Изменена конечная дата: 2021-10-15 Дата: 2021-07-25	46	Изготовление серии	t	0	0
2021-07-25	21	2021-09-19	2	8	1	1	0	0	1	f	1		49	Разработка и прототип	t	0	0
\.


--
-- TOC entry 3046 (class 0 OID 16520)
-- Dependencies: 222
-- Data for Name: Tenders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tenders" (tender_id, date, selected_cp, tz_id, history) FROM stdin;
14	2021-12-03	0	45	
15	2021-10-15	0	46	
16	2021-08-23	0	47	
17	2021-08-19	0	48	
18	2021-09-19	0	49	
\.


--
-- TOC entry 3083 (class 0 OID 0)
-- Dependencies: 198
-- Name: CP_cp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."CP_cp_id_seq"', 28, true);


--
-- TOC entry 3084 (class 0 OID 0)
-- Dependencies: 229
-- Name: Calendar_cal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Calendar_cal_id_seq"', 251, true);


--
-- TOC entry 3085 (class 0 OID 0)
-- Dependencies: 201
-- Name: Costs_cost_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Costs_cost_id_seq"', 293, true);


--
-- TOC entry 3086 (class 0 OID 0)
-- Dependencies: 203
-- Name: Countries_country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Countries_country_id_seq"', 287, true);


--
-- TOC entry 3087 (class 0 OID 0)
-- Dependencies: 205
-- Name: Metrics_metr_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Metrics_metr_id_seq"', 7, true);


--
-- TOC entry 3088 (class 0 OID 0)
-- Dependencies: 231
-- Name: Orgs_docs_doc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Orgs_docs_doc_id_seq"', 8, true);


--
-- TOC entry 3089 (class 0 OID 0)
-- Dependencies: 208
-- Name: Orgs_o_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Orgs_o_id_seq"', 57, true);


--
-- TOC entry 3090 (class 0 OID 0)
-- Dependencies: 211
-- Name: Pack_groups_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_groups_group_id_seq"', 9, true);


--
-- TOC entry 3091 (class 0 OID 0)
-- Dependencies: 213
-- Name: Pack_kinds_kind_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_kinds_kind_id_seq"', 3, true);


--
-- TOC entry 3092 (class 0 OID 0)
-- Dependencies: 215
-- Name: Pack_types_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pack_types_type_id_seq"', 2, true);


--
-- TOC entry 3093 (class 0 OID 0)
-- Dependencies: 226
-- Name: Pay_conds_pay_cond_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Pay_conds_pay_cond_id_seq"', 5, true);


--
-- TOC entry 3094 (class 0 OID 0)
-- Dependencies: 217
-- Name: Specs_spec_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Specs_spec_id_seq"', 23, true);


--
-- TOC entry 3095 (class 0 OID 0)
-- Dependencies: 234
-- Name: Task_kinds_tk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Task_kinds_tk_id_seq"', 6, true);


--
-- TOC entry 3096 (class 0 OID 0)
-- Dependencies: 219
-- Name: Task_names_name_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Task_names_name_id_seq"', 5, true);


--
-- TOC entry 3097 (class 0 OID 0)
-- Dependencies: 224
-- Name: Tasks_task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Tasks_task_id_seq"', 8, true);


--
-- TOC entry 3098 (class 0 OID 0)
-- Dependencies: 228
-- Name: Techs_tz_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Techs_tz_id_seq"', 49, true);


--
-- TOC entry 3099 (class 0 OID 0)
-- Dependencies: 223
-- Name: Tenders_tender_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Tenders_tender_id_seq"', 18, true);


--
-- TOC entry 2846 (class 2606 OID 16543)
-- Name: CP CP_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP"
    ADD CONSTRAINT "CP_pkey" PRIMARY KEY (cp_id);


--
-- TOC entry 2880 (class 2606 OID 16758)
-- Name: Calendar Calendar_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Calendar"
    ADD CONSTRAINT "Calendar_pkey" PRIMARY KEY (cal_id);


--
-- TOC entry 2850 (class 2606 OID 16545)
-- Name: Costs Costs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT "Costs_pkey" PRIMARY KEY (cost_id);


--
-- TOC entry 2852 (class 2606 OID 16547)
-- Name: Countries Countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Countries"
    ADD CONSTRAINT "Countries_pkey" PRIMARY KEY (country_id);


--
-- TOC entry 2854 (class 2606 OID 16549)
-- Name: Metrics Metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Metrics"
    ADD CONSTRAINT "Metrics_pkey" PRIMARY KEY (metr_id);


--
-- TOC entry 2856 (class 2606 OID 16551)
-- Name: Org_countries Org_countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Org_countries"
    ADD CONSTRAINT "Org_countries_pkey" PRIMARY KEY (o_id, country_id);


--
-- TOC entry 2882 (class 2606 OID 24604)
-- Name: Orgs_docs Orgs_docs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs_docs"
    ADD CONSTRAINT "Orgs_docs_pkey" PRIMARY KEY (doc_id);


--
-- TOC entry 2858 (class 2606 OID 16553)
-- Name: Orgs Orgs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs"
    ADD CONSTRAINT "Orgs_pkey" PRIMARY KEY (o_id);


--
-- TOC entry 2860 (class 2606 OID 16555)
-- Name: Orgs_specs Orgs_specs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs_specs"
    ADD CONSTRAINT "Orgs_specs_pkey" PRIMARY KEY (o_id, spec_id);


--
-- TOC entry 2862 (class 2606 OID 16557)
-- Name: Pack_groups Pack_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_groups"
    ADD CONSTRAINT "Pack_groups_pkey" PRIMARY KEY (group_id);


--
-- TOC entry 2864 (class 2606 OID 16559)
-- Name: Pack_kinds Pack_kinds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_kinds"
    ADD CONSTRAINT "Pack_kinds_pkey" PRIMARY KEY (kind_id);


--
-- TOC entry 2866 (class 2606 OID 16561)
-- Name: Pack_types Pack_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pack_types"
    ADD CONSTRAINT "Pack_types_pkey" PRIMARY KEY (type_id);


--
-- TOC entry 2878 (class 2606 OID 16709)
-- Name: Pay_conds Pay_conds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Pay_conds"
    ADD CONSTRAINT "Pay_conds_pkey" PRIMARY KEY (pay_cond_id);


--
-- TOC entry 2868 (class 2606 OID 16565)
-- Name: Specs Specs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Specs"
    ADD CONSTRAINT "Specs_pkey" PRIMARY KEY (spec_id);


--
-- TOC entry 2884 (class 2606 OID 32822)
-- Name: Task_kinds Task_kinds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_kinds"
    ADD CONSTRAINT "Task_kinds_pkey" PRIMARY KEY (tk_id);


--
-- TOC entry 2870 (class 2606 OID 16569)
-- Name: Task_names Task_names_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_names"
    ADD CONSTRAINT "Task_names_pkey" PRIMARY KEY (name_id);


--
-- TOC entry 2876 (class 2606 OID 16693)
-- Name: Tasks Tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tasks"
    ADD CONSTRAINT "Tasks_pkey" PRIMARY KEY (task_id);


--
-- TOC entry 2874 (class 2606 OID 16577)
-- Name: Tenders Tenders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders"
    ADD CONSTRAINT "Tenders_pkey" PRIMARY KEY (tender_id);


--
-- TOC entry 2848 (class 2606 OID 16579)
-- Name: CP_docs cp_docs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP_docs"
    ADD CONSTRAINT cp_docs_pk PRIMARY KEY (file_name, cp_id);


--
-- TOC entry 2872 (class 2606 OID 16725)
-- Name: Techs pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT pk PRIMARY KEY (tz_id);


--
-- TOC entry 2887 (class 2606 OID 16580)
-- Name: CP_docs cp_docs_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP_docs"
    ADD CONSTRAINT cp_docs_fk FOREIGN KEY (cp_id) REFERENCES public."CP"(cp_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2888 (class 2606 OID 16590)
-- Name: Costs cp_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT cp_fk FOREIGN KEY (cp_id) REFERENCES public."CP"(cp_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2892 (class 2606 OID 16605)
-- Name: Techs gr_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT gr_fk FOREIGN KEY (group_id) REFERENCES public."Pack_groups"(group_id);


--
-- TOC entry 2893 (class 2606 OID 16610)
-- Name: Techs kind_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT kind_fk FOREIGN KEY (kind_id) REFERENCES public."Pack_kinds"(kind_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2889 (class 2606 OID 16615)
-- Name: Costs metr_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT metr_fk FOREIGN KEY (metr_id) REFERENCES public."Metrics"(metr_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2898 (class 2606 OID 16759)
-- Name: Calendar name_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Calendar"
    ADD CONSTRAINT name_fk FOREIGN KEY (name_id) REFERENCES public."Task_names"(name_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2894 (class 2606 OID 16625)
-- Name: Techs o_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT o_fk FOREIGN KEY (o_id) REFERENCES public."Orgs"(o_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2899 (class 2606 OID 24605)
-- Name: Orgs_docs o_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Orgs_docs"
    ADD CONSTRAINT o_fk FOREIGN KEY (o_id) REFERENCES public."Orgs"(o_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2896 (class 2606 OID 16710)
-- Name: Techs pay_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT pay_fk FOREIGN KEY (pay_cond_id) REFERENCES public."Pay_conds"(pay_cond_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2886 (class 2606 OID 16741)
-- Name: CP pay_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP"
    ADD CONSTRAINT pay_fk FOREIGN KEY (pay_cond_id) REFERENCES public."Pay_conds"(pay_cond_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2891 (class 2606 OID 16746)
-- Name: Costs task_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT task_fk FOREIGN KEY (task_id) REFERENCES public."Tasks"(task_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2895 (class 2606 OID 16640)
-- Name: Techs type_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Techs"
    ADD CONSTRAINT type_fk FOREIGN KEY (type_id) REFERENCES public."Pack_types"(type_id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 2897 (class 2606 OID 16726)
-- Name: Tenders tz_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tenders"
    ADD CONSTRAINT tz_fk FOREIGN KEY (tz_id) REFERENCES public."Techs"(tz_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2890 (class 2606 OID 16731)
-- Name: Costs tz_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Costs"
    ADD CONSTRAINT tz_fk FOREIGN KEY (tz_id) REFERENCES public."Techs"(tz_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


--
-- TOC entry 2885 (class 2606 OID 16736)
-- Name: CP tz_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."CP"
    ADD CONSTRAINT tz_fk FOREIGN KEY (tz_id) REFERENCES public."Techs"(tz_id) ON UPDATE CASCADE ON DELETE SET NULL NOT VALID;


-- Completed on 2021-07-27 09:19:02

--
-- PostgreSQL database dump complete
--

