-- +goose Up

create index idx_store__title on store using gin (title gin_trgm_ops);
create index idx_store__affordability on store(affordability);
create index idx_store__cuisine on store(cuisine);
create index idx_store__avg_rating on store(avg_rating);

create index idx_store__rating on store_review(rating);

create index idx_product__price_cents on product(price_cents);
create index idx_product__calories on product(calories);

-- +goose Down
