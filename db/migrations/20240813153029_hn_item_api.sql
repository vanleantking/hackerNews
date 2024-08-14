-- +goose Up
-- +goose StatementBegin
CREATE TABLE hn_item (
    id serial,
    hn_item_id BIGINT,
    deleted boolean,
    item_type VARCHAR(20),
    item_by VARCHAR(200),
    created_time BIGINT,
    item_content text,
    parent BIGINT,
    kids json,
    item_url text,
    item_score int,
    category_id BIGINT,
    item_title text,
    descendants BIGINT,
    label VARCHAR(20),
    item_status int,
    updated_at BIGINT,
    created_at BIGINT,
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX idx_hn_item_ids ON hn_item (hn_item_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hn_item;
DROP INDEX idx_hn_item_ids;
-- +goose StatementEnd
