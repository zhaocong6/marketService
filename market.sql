-- auto-generated definition
create table market
(
    id         bigint unsigned         not null
        primary key,
    organize   varchar(255) default '' not null comment '交易所',
    symbol     varchar(255) default '' not null comment '币对',
    type       tinyint      default 1  not null comment '交易类型:1币币/现货, 2期货/交割, 3永续, 4期权',
    expire     timestamp               null comment '到期时间',
    status     tinyint      default 1  not null comment '状态:1运行中, 2暂停',
    created_at timestamp               not null,
    updated_at timestamp               not null,
    deleted_at timestamp               null,
    constraint market_organize_type_symbol_uindex
        unique (organize, type, symbol)
)
    collate = utf8mb4_unicode_ci;

