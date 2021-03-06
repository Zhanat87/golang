-- go_restful created while created container, because it was in environments in docker-compose.yml
GRANT ALL PRIVILEGES ON DATABASE go_restful TO postgres;

create table news_shard_2(
    id bigint not null,
    category_id int not null constraint category_id_check check (category_id = 2),
    author character varying not null,
    rate int not null,
    title character varying,
    text text
);

create index news_shard_2_rate_idx on news_shard_2(rate);