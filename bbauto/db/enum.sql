BEGIN;

SET client_encoding = 'LATIN1';

CREATE TABLE domain_tld (
    id integer NOT NULL,
    domain text NOT NULL,
    date_add date NOT NULL,
    date_scan date
);

CREATE TABLE domain_scope (
    id integer NOT NULL,
    domain_tld integer NOT NULL,
    domain text NOT NULL,
    scope boolean NOT NULL

);

CREATE TABLE domain_dump (
    id integer NOT NULL,
    domain_tld integer NOT NULL,
    domain text NOT NULL,
    date_add date NOT NULL
);

INSERT into "domain_tld" (id, domain, date_add)
VALUES 
(1, 'example.com', current_timestamp),
(2, 'evil.com', current_timestamp);

INSERT into "domain_scope" (id, domain_tld, domain, scope)
VALUES
(1, 1, 'mail.example.com', false),
(2, 1, '*.example.com', true);


ALTER TABLE ONLY domain_tld
    ADD CONSTRAINT domain_tld_pkey PRIMARY KEY (id);

ALTER TABLE ONLY domain_scope
    ADD CONSTRAINT domain_scope_pkey PRIMARY KEY (id);

ALTER TABLE ONLY domain_scope
    ADD CONSTRAINT domain_tld_ref FOREIGN KEY (domain_tld) REFERENCES domain_tld(id);

ALTER TABLE ONLY domain_dump
    ADD CONSTRAINT domain_dump_pkey PRIMARY KEY (id);

ALTER TABLE ONLY domain_dump
    ADD CONSTRAINT domain_tld_ref FOREIGN KEY (domain_tld) REFERENCES domain_tld(id);
    
COMMIT;

ANALYZE domain_tld;
ANALYZE domain_scope;
ANALYZE domain_dump;

