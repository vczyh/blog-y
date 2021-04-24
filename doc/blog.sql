use blog;
set foreign_key_checks = 0;

-- post
drop table if exists post;
drop  table if exists post_tag;
create table post
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3)   not null,
    updated_at datetime(3)   not null,
    deleted_at datetime(3)   null,
    post_id    varchar(10)   not null,
    title      varchar(100)  not null,
    content    longtext      not null,
    cover_url  varchar(2083) not null,
    cover_desc varchar(100)  not null,
    subtitle   varchar(100)  not null
);

create unique index uk_post_id on post (post_id);

create index idx_deleted_at on post (deleted_at);

-- tag
drop table if exists tag;
create table tag
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3)  not null,
    updated_at datetime(3)  not null,
    deleted_at datetime(3)  null,
    tag_id     varchar(10)  not null,
    name       varchar(100) not null
);

create unique index uk_tag_id on tag (tag_id);

create unique index uk_name on tag (name);

create index idx_deleted_at on tag (deleted_at);

-- post_tag
drop table if exists post_tag;
create table post_tag
(
    post_id varchar(10) not null,
    tag_id  varchar(10) not null,
    primary key (post_id, tag_id),
    constraint fk_post_id
        foreign key (post_id) references post (post_id),
    constraint fk_tag_id
        foreign key (tag_id) references tag (tag_id)
);

-- config
drop table if exists config;
create table config
(
    id            bigint unsigned auto_increment primary key,
    created_at    datetime(3)  not null,
    updated_at    datetime(3)  not null,
    deleted_at    datetime(3)  null,
    name          varchar(100) not null,
    initial_value varchar(100) null,
    current_value varchar(100) null
);
create unique index uk_name on config (name);

create index idx_deleted_at
    on config (deleted_at);

set foreign_key_checks = 1;



