ALTER TABLE "reservations" ADD COLUMN "num_nights" int;
ALTER TABLE "reservations" ADD COLUMN "num_guests" int;
ALTER TABLE "reservations" ADD COLUMN "status" varchar NOT NULL DEFAULT 'unconfirmed';
ALTER TABLE "reservations" ADD COLUMN "total_price" int;
ALTER TABLE "reservations" ADD COLUMN "extras_price" int NOT NULL DEFAULT 0;
ALTER TABLE "reservations" ADD COLUMN "is_paid" bool NOT NULL DEFAULT false;
ALTER TABLE "reservations" ADD COLUMN "has_breakfast" bool NOT NULL DEFAULT false;