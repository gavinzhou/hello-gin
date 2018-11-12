--
-- PostgreSQL database dump
--

-- Dumped from database version 10.5
-- Dumped by pg_dump version 10.5

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
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- Name: blog_article_seq; Type: SEQUENCE; Schema: public; Owner: gavin
--

CREATE SEQUENCE public.blog_article_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_article_seq OWNER TO gavin;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: blog_article; Type: TABLE; Schema: public; Owner: gavin
--

CREATE TABLE public.blog_article (
    id integer DEFAULT nextval('public.blog_article_seq'::regclass) NOT NULL,
    tag_id integer DEFAULT 0,
    title character varying(100) DEFAULT ''::character varying,
    content text,
    created_on integer,
    created_by character varying(100) DEFAULT ''::character varying,
    modified_on integer DEFAULT 0,
    modified_by character varying(255) DEFAULT ''::character varying,
    state smallint DEFAULT '1'::smallint,
    description character varying(255),
    CONSTRAINT blog_article_id_check CHECK ((id > 0)),
    CONSTRAINT blog_article_modified_on_check CHECK ((modified_on > 0)),
    CONSTRAINT blog_article_state_check CHECK ((state > 0)),
    CONSTRAINT blog_article_tag_id_check CHECK ((tag_id > 0))
);


ALTER TABLE public.blog_article OWNER TO gavin;

--
-- Name: blog_auth_seq; Type: SEQUENCE; Schema: public; Owner: gavin
--

CREATE SEQUENCE public.blog_auth_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_auth_seq OWNER TO gavin;

--
-- Name: blog_auth; Type: TABLE; Schema: public; Owner: gavin
--

CREATE TABLE public.blog_auth (
    id integer DEFAULT nextval('public.blog_auth_seq'::regclass) NOT NULL,
    username character varying(50) DEFAULT ''::character varying,
    password character varying(50) DEFAULT ''::character varying,
    CONSTRAINT blog_auth_id_check CHECK ((id > 0))
);


ALTER TABLE public.blog_auth OWNER TO gavin;

--
-- Name: blog_tag_seq; Type: SEQUENCE; Schema: public; Owner: gavin
--

CREATE SEQUENCE public.blog_tag_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_tag_seq OWNER TO gavin;

--
-- Name: blog_tag; Type: TABLE; Schema: public; Owner: gavin
--

CREATE TABLE public.blog_tag (
    id integer DEFAULT nextval('public.blog_tag_seq'::regclass) NOT NULL,
    name character varying(100) DEFAULT ''::character varying,
    created_on integer DEFAULT 0,
    created_by character varying(100) DEFAULT ''::character varying,
    modified_on integer DEFAULT 0,
    modified_by character varying(100) DEFAULT ''::character varying,
    state smallint DEFAULT '1'::smallint,
    CONSTRAINT blog_tag_created_on_check CHECK ((created_on > 0)),
    CONSTRAINT blog_tag_id_check CHECK ((id > 0)),
    CONSTRAINT blog_tag_modified_on_check CHECK ((modified_on > 0)),
    CONSTRAINT blog_tag_state_check CHECK ((state > 0))
);


ALTER TABLE public.blog_tag OWNER TO gavin;

--
-- Name: blog_article blog_article_pkey; Type: CONSTRAINT; Schema: public; Owner: gavin
--

ALTER TABLE ONLY public.blog_article
    ADD CONSTRAINT blog_article_pkey PRIMARY KEY (id);


--
-- Name: blog_auth blog_auth_pkey; Type: CONSTRAINT; Schema: public; Owner: gavin
--

ALTER TABLE ONLY public.blog_auth
    ADD CONSTRAINT blog_auth_pkey PRIMARY KEY (id);


--
-- Name: blog_tag blog_tag_pkey; Type: CONSTRAINT; Schema: public; Owner: gavin
--

ALTER TABLE ONLY public.blog_tag
    ADD CONSTRAINT blog_tag_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

