create database yu_bolog;

use yu_bolog;

create table advert
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3) null comment '创建时间',
    updated_at datetime(3) null comment '更新时间',
    title      varchar(32) null comment '广告标题',
    href       longtext    null comment '跳转链接',
    images     longtext    null comment '图片',
    is_show    tinyint(1)  null comment '是否展示'
);

create table banner
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3)      null comment '创建时间',
    updated_at datetime(3)      null comment '更新时间',
    path       longtext         null comment '图片路径',
    hash       longtext         null comment '图片的hash值',
    name       varchar(38)      null comment '图片名称',
    image_type bigint default 1 null comment '图片的类型，本地还是七牛,1本地，2七牛，默认是1'
);

create table chat
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3)  null comment '创建时间',
    updated_at datetime(3)  null comment '更新时间',
    nick_name  varchar(15)  null comment '昵称',
    avatar     varchar(128) null comment '头像',
    content    varchar(256) null comment '内容',
    ip         varchar(32)  null comment 'ip',
    addr       varchar(64)  null comment '地址',
    is_group   tinyint(1)   null comment '是否是群组消息',
    msg_type   tinyint      null comment '消息类型'
);

create table feedback
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3)  null comment '创建时间',
    updated_at datetime(3)  null comment '更新时间',
    email      varchar(64)  null,
    content    varchar(256) null
);

create table menu
(
    id            bigint unsigned auto_increment comment 'id' primary key,
    created_at    datetime(3) null comment '创建时间',
    updated_at    datetime(3) null comment '更新时间',
    title         varchar(32) null comment '菜单标题',
    path          varchar(32) null comment '菜单路径',
    slogan        varchar(64) null comment 'slogan',
    abstract      longtext    null comment '简介，按照换行去切割为数组',
    abstract_time bigint      null comment '简介的切换时间',
    banner_time   bigint      null comment 'banner图的切换时间',
    sort          smallint    null comment '顺序'
);

create table menu_banner
(
    menu_id   bigint unsigned null comment '菜单的id',
    banner_id bigint unsigned null comment 'banner图的id',
    sort      smallint        null comment '序号',
    constraint fk_banner_menus_banner foreign key (banner_id) references banner (id),
    constraint fk_menu_banner_menu_model foreign key (menu_id) references menu (id)
);

create table tag
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3) null comment '创建时间',
    updated_at datetime(3) null comment '更新时间',
    title      varchar(16) null comment '标签的名称'
);

create table user
(
    id          bigint unsigned auto_increment comment 'id' primary key,
    created_at  datetime(3)       null comment '创建时间',
    updated_at  datetime(3)       null comment '更新时间',
    nick_name   varchar(36)       null comment '昵称',
    user_name   varchar(36)       null comment '用户名',
    password    varchar(128)      null comment '密码',
    avatar      varchar(256)      null comment '头像',
    email       varchar(128)      null comment '邮箱',
    tel         varchar(18)       null comment '手机号',
    addr        varchar(64)       null comment '地址',
    token       varchar(64)       null comment '其他平台的唯一id',
    ip          varchar(20)       null comment 'ip',
    role        tinyint default 1 null comment '权限，1管理员，2普通用户，3游客',
    sign_status bigint            null comment '注册来源，1qq，3邮箱',
    integral    bigint  default 0 null comment '我的积分',
    sign        varchar(128)      null comment '我的签名',
    link        varchar(128)      null comment '我的链接地址'
);

create table collect
(
    id         bigint unsigned auto_increment primary key,
    user_id    bigint unsigned null comment '用户id',
    article_id varchar(32)     null comment '文章的es id',
    created_at datetime(3)     null comment '收藏的时间',
    constraint fk_collect_user_model foreign key (user_id) references user (id)
);

create table comment
(
    id                bigint unsigned auto_increment comment 'id'
        primary key,
    created_at        datetime(3)       null comment '创建时间',
    updated_at        datetime(3)       null comment '更新时间',
    parent_comment_id bigint unsigned   null comment '父评论id',
    content           varchar(256)      null comment '评论内容',
    digg_count        tinyint default 0 null comment '点赞数',
    comment_count     tinyint default 0 null comment '子评论数',
    article_id        varchar(32)       null comment '文章id',
    user_id           bigint unsigned   null comment '关联的用户id',
    constraint fk_comment_sub_comments  foreign key (parent_comment_id) references comment (id),
    constraint fk_comment_user foreign key (user_id) references user (id)
);

create table login_data
(
    id         bigint unsigned auto_increment comment 'id' primary key,
    created_at datetime(3)     null comment '创建时间',
    updated_at datetime(3)     null comment '更新时间',
    user_id    bigint unsigned null comment '用户id',
    ip         varchar(20)     null comment 'ip',
    nick_name  varchar(42)     null comment '昵称',
    token      varchar(256)    null comment 'token',
    device     varchar(256)    null comment '登录失败',
    addr       varchar(64)     null comment '地址',
    login_type tinyint         null comment '登录方式，1QQ，3邮箱',
    constraint fk_login_data_user_model foreign key (user_id) references user (id)
);

create table message
(
    id                  bigint unsigned auto_increment comment 'id',
    created_at          datetime(3)          null comment '创建时间',
    updated_at          datetime(3)          null comment '更新时间',
    send_user_id        bigint unsigned      not null comment '发送人id',
    send_user_nick_name varchar(42)          null comment '发送人昵称',
    send_user_avatar    longtext             null comment '发送人头像',
    rev_user_id         bigint unsigned      not null comment '接收人id',
    rev_user_nick_name  varchar(42)          null comment '接收人昵称',
    rev_user_avatar     longtext             null comment '接收人头像',
    is_read             tinyint(1) default 0 null comment '接收人是否查看',
    content             longtext             null comment '消息内容',
    primary key (id, send_user_id, rev_user_id),
    constraint fk_message_rev_user_model foreign key (rev_user_id) references user (id),
    constraint fk_message_send_user_model foreign key (send_user_id) references user (id)
);

