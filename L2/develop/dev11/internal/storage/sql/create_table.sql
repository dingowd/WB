create table events
(
    id         serial primary key,
    owner      bigint,
    title      text,
    descr      text,
    start_date date not null,
    start_time time not null ,
    end_date   date not null,
    end_time   time not null
);
create index owner_idx on events (owner);
create index start_idx on events using btree (start_date, start_time);