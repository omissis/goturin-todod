create table todos
(
    uuid          uuid not null constraint todos_pk primary key,
    title         varchar(255) not null,
    description   text not null
);

alter table todos owner to todod;
