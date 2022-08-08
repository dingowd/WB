create table if not exists delivery
(
    id      serial primary key,
    name    varchar(50)  not null,
    phone   varchar(12)  not null,
    zip     varchar(10)  not null,
    city    varchar(30)  not null,
    address varchar(255) not null,
    region  varchar(30)  not null,
    email   varchar(255) not null
);

create table if not exists payment
(
    transaction   varchar(255) primary key not null,
    request_id    varchar(255)             not null,
    currency      varchar(3)               not null,
    provider      varchar(50)              not null,
    amount        int                      not null,
    payment_dt    int                      not null,
    bank          varchar(50)              not null,
    delivery_cost int                      not null,
    goods_total   int                      not null,
    custom_fee    int                      not null
);

create table if not exists orders
(
    order_uid          varchar(255) primary key,
    track_number       varchar(50)              not null,
    entry              varchar(25)              not null,
    locale             varchar(2)               not null,
    internal_signature varchar(255)             not null,
    customer_id        varchar(50)              not null,
    delivery_service   varchar(50)              not null,
    shardkey           varchar(50)              not null,
    sm_id              int2                     not null,
    date_created       timestamp with time zone not null,
    oof_shard          varchar(2)               not null,
    delivery_id        int references delivery (id),
    transaction        varchar(255) references payment (transaction)
);

create table if not exists items
(
    chrt_id      int primary key,
    track_number varchar(30) not null,
    price        int         not null,
    rid          varchar(50) not null,
    name         varchar(50) not null,
    sale         int2        not null,
    size         varchar(5)  not null,
    total_price  int         not null,
    nm_id        int         not null,
    brand        varchar(50) not null,
    status       int2        not null,
    order_uid    varchar(255) references orders (order_uid)
        on delete cascade
);