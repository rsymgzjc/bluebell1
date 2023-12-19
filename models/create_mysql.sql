drop table if exists user;

create table user(
    id bigint(20) not null auto_increment,
    user_id bigint(20) not null ,
    username varchar(64) collate utf8mb4_general_ci not null,
    password varchar(64) collate utf8mb4_general_ci not null,
    email varchar(64) null collate utf8mb4_general_ci,
    gender tinyint(4) not null default '0',
    create_time timestamp null default current_timestamp,
    update_time timestamp null default current_timestamp on update current_timestamp,
    PRIMARY KEY  (id),
    UNIQUE KEY idx_username (username) USING BTREE ,
    UNIQUE KEY idx_user_id (user_id) USING BTREE
)engine =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

drop table if exists  community;
create table  community(
    id int(11) not null auto_increment,
    community_id int(10) unsigned not null ,
    community_name varchar(128) collate utf8mb4_general_ci not null ,
    introduction varchar(256) collate utf8mb4_general_ci not null ,
    creat_time timestamp not null default current_timestamp,
    update_time timestamp NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    primary key (id),
    unique key idx_community_id (community_id),
    unique key idx_community_name (community_name)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;

insert into community value ('1','1','Go','Golang','2016-11-01 08:10:10','2016-11-01 08:10:10');
insert into community value ('2','2','leetcode','刷刷刷刷刷题','2020-01-01 08:00:00','2020-01-01 08:00:00');
insert into community value ('3','3','CS:GO','Rush B....','2018-08-07 08:30:00','2018-08-07 08:30:00');
insert into community value ('4','4','LOL','欢迎来到英雄联盟','2016-01-01 08:00:00','2016-01-01 08:00:00');

drop table if exists post;
create table post (
                      id bigint(20) not null auto_increment,
                      post_id bigint(20)  not null comment '帖子id',
                      title varchar(128) collate utf8mb4_general_ci not null comment '标题',
                      content varchar(8192) collate utf8mb4_general_ci not null comment '内容',
                      author_id bigint(20) not null comment '作者的用户id',
                      community_id bigint(20) not null comment '所属社区',
                      status tinyint(4) not null default '1' comment '帖子状态',
                      creat_time timestamp null default current_timestamp comment '创建时间',
                      update_time timestamp null DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
                      primary key (id),
                      unique key idx_post_id (post_id),
                      key idx_community_id (community_id),
                      key idx_author_id (author_id)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COLLATE=utf8mb4_general_ci;