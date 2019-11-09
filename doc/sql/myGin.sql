-- auto-generated definition
create table user
(
    id        int unsigned auto_increment primary key,
    name      varchar(30)      default ''            not null comment '姓名',
    age       tinyint unsigned default 0                 not null comment '年龄',
    created_at datetime         default CURRENT_TIMESTAMP not null,
    updated_at datetime         default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
)
    comment '用户表';

create index k_created_at
    on user (created_at);

create index k_name
    on user (name);

create index k_updated_at
    on user (updated_at);

