CREATE TABLE public.brsm_profiles (
    user_id uuid NOT NULL,
    full_name text NOT NULL,
    subrole text NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    "position" character varying(255),
    phone character varying(50)
);

CREATE TABLE public.chat_members (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    chat_id uuid NOT NULL,
    user_id uuid NOT NULL,
    role character varying(50) DEFAULT 'member'::character varying NOT NULL,
    joined_at timestamp with time zone DEFAULT now() NOT NULL,
    last_read_at timestamp with time zone DEFAULT now(),
    unread_count integer DEFAULT 0
);

CREATE TABLE public.chat_message_contents (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    message_id uuid NOT NULL,
    type character varying(50) NOT NULL,
    content text NOT NULL,
    metadata jsonb
);

CREATE TABLE public.chat_messages (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    chat_id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.chats (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type character varying(50) DEFAULT 'direct'::character varying NOT NULL,
    name text,
    avatar text,
    created_by uuid,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.customer_profiles (
    user_id uuid NOT NULL,
    full_name text NOT NULL,
    enterprise_id uuid,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    "position" character varying(255),
    phone character varying(50),
    subrole text
);

CREATE TABLE public.departments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.employment_participant_statuses (
    id integer NOT NULL,
    name text NOT NULL,
    description text
);

CREATE TABLE public.employment_participants (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    request_id uuid NOT NULL,
    user_id uuid NOT NULL,
    status_id integer DEFAULT 1 NOT NULL,
    applied_at timestamp with time zone DEFAULT now() NOT NULL,
    contracted_at timestamp with time zone
);

CREATE TABLE public.employment_requests (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    enterprise_id uuid NOT NULL,
    university_department_id uuid NOT NULL,
    title text NOT NULL,
    description text,
    requirements text,
    salary text,
    schedule text,
    max_participants integer DEFAULT 1 NOT NULL,
    status_id integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT employment_requests_max_participants_check CHECK ((max_participants > 0))
);

CREATE TABLE public.employment_statuses (
    id integer NOT NULL,
    name text NOT NULL,
    description text
);

CREATE TABLE public.enterprises (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    address text,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.squad_statuses (
    id integer NOT NULL,
    name text NOT NULL,
    description text
);

CREATE TABLE public.student_profiles (
    user_id uuid NOT NULL,
    full_name text NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    "position" character varying(255),
    phone character varying(50),
    specialty character varying(255),
    grade double precision,
    university_department_id uuid
);

CREATE TABLE public.student_squad_participants (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    squad_id uuid NOT NULL,
    user_id uuid NOT NULL,
    joined_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.student_squads (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    organizer_id uuid NOT NULL,
    title text NOT NULL,
    description text,
    profile text,
    max_participants integer DEFAULT 1 NOT NULL,
    status_id integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT student_squads_max_participants_check CHECK ((max_participants > 0))
);

CREATE TABLE public.universities (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE public.university_departments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    address text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    university_id uuid NOT NULL,
    department_id uuid NOT NULL
);

CREATE TABLE public.university_profiles (
    user_id uuid NOT NULL,
    full_name text NOT NULL,
    subrole text NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    "position" character varying(255),
    phone character varying(50),
    university_department_id uuid
);

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    login character varying(100) NOT NULL,
    password_hash text NOT NULL,
    role character varying(50) DEFAULT 'user'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    avatar text
);