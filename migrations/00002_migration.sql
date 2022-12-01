-- +goose Up

create table image(
	id text primary key,
	owner_username text not null references users(username) on delete cascade,
	byte_size bigint not null,
	content_type text not null
);

-- +goose Down

