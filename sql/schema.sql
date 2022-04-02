create table events
(
    id           varchar(32) primary key,
    owner        varchar(32),
    title        text,
    notes        text,
    start_at     timestamp with time zone,
    end_at       timestamp with time zone,
    alert_before bigint
);

create index owner_idx on events (owner);
create index start_idx on events using btree (start_at);
create index end_idx on events using btree (end_at);