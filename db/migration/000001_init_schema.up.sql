CREATE TABLE public.users (
	id SERIAL NOT NULL,
	first_name varchar DEFAULT ''::character varying NOT NULL,
	last_name varchar DEFAULT ''::character varying NOT NULL,
	email varchar NOT NULL,
	phone varchar DEFAULT ''::character varying NOT NULL,
	"password" varchar NOT NULL,
	access_level int DEFAULT 1 NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);

CREATE TABLE public.rooms (
	id SERIAL NOT NULL,
	room_name_sr varchar NOT NULL,
	room_name_en varchar NOT NULL,
	room_name_bg varchar NOT NULL,
	room_short_des_sr varchar NOT NULL,
	room_short_des_en varchar NOT NULL,
	room_short_des_bg varchar NOT NULL,
	room_description_sr varchar NOT NULL,
	room_description_en varchar NOT NULL,
	room_description_bg varchar NOT NULL,
	room_pictures_folder varchar NOT NULL,
	room_guest_number int NOT NULL,
	room_price_en int NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	CONSTRAINT rooms_pkey PRIMARY KEY (id)
);

CREATE TABLE public.restrictions (
	id SERIAL NOT NULL,
	restriction_name_sr varchar NOT NULL,
	restriction_name_en varchar NOT NULL,
	restriction_name_bg varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	CONSTRAINT restrictions_pkey PRIMARY KEY (id)
);

CREATE TABLE public.reservations (
	id SERIAL NOT NULL,
	room_id int NOT NULL,
	first_name varchar DEFAULT ''::character varying NOT NULL,
	last_name varchar DEFAULT ''::character varying NOT NULL,
	email varchar NOT NULL,
	phone varchar DEFAULT ''::character varying NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	processed int DEFAULT 0 NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	CONSTRAINT reservations_pkey PRIMARY KEY (id)
);
CREATE INDEX reservations_email_idx ON public.reservations USING btree (email);
CREATE INDEX reservations_last_name_idx ON public.reservations USING btree (last_name);


-- public.reservations foreign keys

ALTER TABLE public.reservations ADD CONSTRAINT reservations_rooms_id_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE;

CREATE TABLE public.room_restrictions (
	id SERIAL NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	room_id int NOT NULL,
	reservation_id int,
	restriction_id int NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
	CONSTRAINT room_restrictions_pkey PRIMARY KEY (id)
);
CREATE INDEX room_restrictions_reservation_id_idx ON public.room_restrictions USING btree (reservation_id);
CREATE INDEX room_restrictions_room_id_idx ON public.room_restrictions USING btree (room_id);
CREATE INDEX room_restrictions_start_date_end_date_idx ON public.room_restrictions USING btree (start_date, end_date);


-- public.room_restrictions foreign keys

ALTER TABLE public.room_restrictions ADD CONSTRAINT room_restrictions_reservations_id_fk FOREIGN KEY (reservation_id) REFERENCES public.reservations(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.room_restrictions ADD CONSTRAINT room_restrictions_restrictions_id_fk FOREIGN KEY (restriction_id) REFERENCES public.restrictions(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.room_restrictions ADD CONSTRAINT room_restrictions_rooms_id_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE;
