-- +goose Up
create table users(
	username text primary key,
	full_name text not null,
	kind text not null,
	password_hash text not null,
	salt text not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	session_id text,
	session_expires_at timestamp
);

create table store(
	id bigserial primary key,
	title text not null,
	affordability text not null,
	cuisine text not null,
	owner_username text not null references users(username) on delete cascade,
	image_id text not null,
	avg_rating bigint not null,
	number_of_reviews bigint not null,
	created_at timestamp not null,
	updated_at timestamp not null
);

create table store_review(
	id bigserial primary key,
	author_username text not null references users(username) on delete cascade,
	store_id bigint not null references store(id) on delete cascade,
	rating bigint CONSTRAINT valid_rating CHECK (rating >= 0 and rating <= 5)
);

create table product(
	id bigserial primary key,
	name text not null,
	store_id bigint references store(id) on delete cascade,
	ingredients text[] not null,
	calories bigint not null,
	image_id text not null,
	created_at timestamp not null,
	updated_at timestamp not null
);

-- +goose Down
