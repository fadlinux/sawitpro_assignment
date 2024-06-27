CREATE TABLE estate (
    id uuid,
    width integer,
    length integer,
    create_time timestamptz,
    update_time timestamptz
);

create table tree
(
    id         uuid,
    width      int,
    length     int,
    height     int,
    estate_id  uuid,
    create_time timestamptz,
    update_time timestamptz
);

create table stats
(
    tree_id    uuid,
    width      int,
    length     int,
    height     int,
    estate_id  uuid,
    create_time timestamptz,
    update_time timestamptz
);