-- create database ankets;
--
-- CREATE TABLE public.bun_migration_locks (
-- 	id serial4 NOT NULL,
-- 	table_name varchar NULL,
-- 	CONSTRAINT bun_migration_locks_pkey PRIMARY KEY (id),
-- 	CONSTRAINT bun_migration_locks_table_name_key UNIQUE (table_name)
-- );
--
-- CREATE TABLE public.bun_migrations (
-- 	id serial4 NOT NULL,
-- 	"name" varchar NULL,
-- 	group_id int4 NULL,
-- 	migrated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
-- 	CONSTRAINT bun_migrations_pkey PRIMARY KEY (id)
-- );
--
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
--

INSERT INTO public.value_types (id, "name", value_relation) VALUES
(
    '9f61f302-6821-40b9-94bc-78dedf955a11'::uuid,
    'string',
    'simple'::public."value_type_value_relation_enum"
),
(
    '6fd80c4c-ece3-479c-94b6-005ccebcfe73'::uuid,
    'set',
    'manyToMany'::public."value_type_value_relation_enum"
),
(
    'fc00cc5a-f7a5-4974-ad57-9432656d5e0e'::uuid,
    'radio',
    'oneToMany'::public."value_type_value_relation_enum"
),
(
    '47affcc5-5d32-4b1f-bf07-33382ed06cda'::uuid,
    'number',
    'simple'::public."value_type_value_relation_enum"
),
(
    'efdd456c-091b-49d9-ac32-d0d345f88e64'::uuid,
    'date',
    'simple'::public."value_type_value_relation_enum"
),
(
    '9fa59e5f-b5f4-4dd0-821f-b9aa6eb25a10'::uuid,
    'file',
    'simple'::public."value_type_value_relation_enum"
),
(
    '6fe180c8-c40e-4d7a-8b0a-9d9e22ae9c61'::uuid,
    'text',
    'simple'::public."value_type_value_relation_enum"
),
(
    '841bb273-e78c-409c-b442-598e3de6a2b2'::uuid,
    'files',
    'manyToMany'::public."value_type_value_relation_enum"
);
