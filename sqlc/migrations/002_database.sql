-- +goose Up
-- +goose StatementBegin


create table if not exists character
(
    id          uuid      default gen_random_uuid() not null
        constraint character_pk
            primary key,
    created_at  timestamp default now()             not null,
    owner_id  uuid                                     not null,
    deleted bool default false                      not null,
    public bool default false                       not null,
    name        text                                not null,
    biography text,
    nationality text                                not null,
    age         integer                             not null,
    body        text                                not null,
    breast      text                                not null,
    butt        text                                not null,
    eyes_color  text                                not null,
    hair_style  text                                not null,
    hair_color  text                                not null,
    meta_params text                                not null,
    main_image_id uuid
);

create table if not exists character_image
(
    id           uuid default gen_random_uuid() not null
        constraint character_image_pk
            primary key,
    created_at  timestamp default now()             not null,
    owner_id uuid not null,
    character_id uuid                           not null
        constraint character_image_character_fk
            references character on delete cascade on update cascade,
    file_hash   text                           not null,
    is_main bool default false not null
);

create table if not exists glossary
(
    id         uuid      default gen_random_uuid() not null
        constraint glossary_pk
            primary key,
    type  text not null,
    value text not null
);

create table if not exists glossary_localization
(
    id         uuid      default gen_random_uuid() not null
        constraint glossary_localization_pk
            primary key,
    glossary_id uuid not null
        constraint glossary_localization_glossary_fk
            references glossary on delete cascade on update cascade,
    local text not null,
    name text not null
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists character;
drop table if exists character_image;

drop table if exists glossary;
drop table if exists glossary_localization;
-- +goose StatementEnd
