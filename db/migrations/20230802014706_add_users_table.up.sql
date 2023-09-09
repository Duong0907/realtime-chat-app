create table "user" (
    "id"        bigserial   primary key,
    "username"  varchar     not null,
    "email"     varchar     not null,
    "password"  varchar     not null 
)