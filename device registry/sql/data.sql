drop table nodes;
drop table users;
drop table regions;

create table if not exists regions
(
    id          varchar(50) not null primary key,
    name        varchar(30),
    description text        not null
);

alter table regions
    owner to postgres;

INSERT INTO regions (id, name, description)
VALUES ('AA004', 'Bagamoyo', 'Bagamoyo Area, iGrid Northern Zone Control Center');
INSERT INTO regions (id, name, description)
VALUES ('AA002', 'Kibamba', 'Kibamba Area, iGrid Eastern Zone Control Center');
INSERT INTO regions (id, name, description)
VALUES ('AA001', 'CoICT', 'CoICT Campus, Sayansi Kijitonyama Control Center');
INSERT INTO regions (id, name, description)
VALUES ('AA003', 'Tegeta', 'Wazo Hill Area, iGrid Western  Zone Control Center');

create table if not exists users
(
    id       varchar(100) not null primary key,
    name     varchar(100),
    email    varchar(100) not null,
    password varchar(100),
    ugroup   integer,
    region   varchar(50),
    created  date,
    foreign key (region) references regions(id)
);

alter table users
    owner to postgres;

INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('ours9489ho08', 'Carma Cumo', 'ccumo0@springer.com', 'u0WKt2JRaB', 1, 'AA004', '2020-02-07');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('glut8904no20', 'Horatio Eyckelberg', 'heyckelberg1@marriott.com', 'ijBAfoa', 1, 'AA003', '2020-09-04');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('ball8000us98', 'Syd Briant', 'sbriant2@patch.com', 'qHEcPqM8atbK', 1, 'AA001', '2020-03-18');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('feet3749am67', 'Sumner Eustace', 'seustace3@edublogs.org', 'RkAlCTBl4sG', 1, 'AA002', '2020-03-12');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('tail7534am70', 'Rolph Hissie', 'rhissie6@bloglines.com', 'w4zYGXZLa', 3, 'AA004', '2020-04-13');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('itch6823am90', 'Jacqueline Jonin', 'jjonin7@flavors.me', 'pyZhCMZXAJMR', 3, 'AA003', '2020-05-10');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('pant4217so33', 'Reeva Metson', 'rmetson9@php.net', 'UWXvAss', 3, 'AA003', '2020-02-24');
INSERT INTO users (id, name, email, password, ugroup, region, created)
VALUES ('638f1cf1-e7cf-4f1a-8064-bbc1053cbf49', 'Pius Alfred', 'me.pius1102@gmail.com',
        '$2a$10$dxLvrXCrVD9kpVZccrAXNenTeXI7N.KVrk1yWaPzVmNvwfPt4d4/6', 1, 'AA001', '2020-12-29');


create table if not exists nodes
(
    id      VARCHAR(600) NOT NULL PRIMARY KEY,
    addr    VARCHAR(60)  NOT NULL UNIQUE,
    name    VARCHAR(50)  NOT NULL,
    type    INT          NOT NULL,
    region  VARCHAR(5)   NOT NULL,
    lat     VARCHAR(50)  NOT NULL,
    long    VARCHAR(50)  NOT NULL,
    created VARCHAR(60)  NOT NULL,
    master  VARCHAR(60),
    FOREIGN KEY (region) REFERENCES regions (id)
);

alter table nodes
    owner to postgres;

insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('99f1773b-fb21-4ef8-9165-863e94301201', '89-19-60-34-8B-C3', 'igrid monitor', 1, 'AA004', 2.7239834,
        101.9476452, '2019-09-24T16:45:27Z', '1c1f6128-5afa-433f-bbc4-21a934b370a0');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('36fe015e-758a-4599-9818-26fe1b1b0a13', '3A-3E-71-A3-96-F4', 'igrid monitor', 3, 'AA002', 45.6103043, 79.467517,
        '2020-01-17T11:24:20Z', 'd9db57c0-7970-4513-9556-c6f88bc12d72');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('293d0a97-c0a2-400d-bd2b-9668bc1d3201', 'D3-6A-C0-84-09-75', 'oil sensor', 3, 'AA002', 48.6277459, 2.4381665,
        '2020-06-30T15:08:20Z', 'c7a520f0-1507-4081-b461-22a1fd540df7');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('7acc2d43-c5fb-4de6-8553-e99a8ddd1a87', 'F8-6D-62-42-8E-23', 'oil sensor', 3, 'AA003', 29.129743, 105.877112,
        '2019-01-02T15:31:25Z', '815dd439-a289-4f78-8f5c-11867c82efea');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('833981d1-1040-4c2d-ad9f-f44e26c8d17c', '10-13-2B-C1-BD-54', 'temp sensor', 1, 'AA001', 48.015883, 37.80285,
        '2019-09-18T02:01:25Z', 'aaf37f1a-55e4-4c9e-89e0-c4102bbd12a1');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('3d33d534-568f-4260-be9f-604d78f30d08', 'FB-7C-02-35-41-56', 'temp sensor', 2, 'AA001', 44.284636, 129.459707,
        '2019-10-19T12:05:40Z', '537c0859-0c66-484e-b728-fe3667ab46c4');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('9c0b50a8-50b0-4987-8cd0-56b91efa52d9', 'D0-85-C8-7C-4D-49', 'igrid monitor', 2, 'AA003', 37.5968793,
        -1.0346774, '2020-09-28T13:20:39Z', 'af58ae5c-740e-4daa-9d3b-0da33f9961eb');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('3b392372-6c7d-4eff-9d1f-4a5d21578b94', '28-CE-87-EE-09-FB', 'temp sensor', 1, 'AA002', -34.5656691,
        -58.7176959, '2020-11-28T14:03:01Z', 'ab022b43-eeba-4d58-b145-f28df6469b4d');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('7e1416ab-bf5d-447f-b300-5fd38e6324c6', 'E2-62-F8-58-40-A9', 'mqtt server', 1, 'AA003', 63.3063621, 18.7067796,
        '2019-04-01T11:52:00Z', 'a42c6281-d76a-4fca-9f9e-80be3e8edd87');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('9e130390-da5e-4fd5-b394-b0dcc7a55b10', 'F7-BC-A0-EF-14-29', 'oil sensor', 2, 'AA002', 52.4293273, 19.4619176,
        '2020-06-24T17:07:55Z', 'c7a7cd3e-d0d7-43aa-9115-b3c508fc31ba');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('f3f204c7-962b-440f-bd7b-5ed7d83eb874', '6F-5E-42-8B-36-B9', 'igrid monitor', 3, 'AA001', 43.2916776,
        -0.3696612, '2020-04-09T07:11:47Z', '195d620b-8556-4eb3-8085-6e9eaa09537a');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('9ef2e08b-53b1-4e5e-a306-6b8e1b6d3664', 'B0-B5-97-6A-21-A8', 'igrid monitor', 2, 'AA001', 24.6946241,
        70.1814258, '2019-08-07T19:15:19Z', 'a037af23-32b0-4c92-8612-4b9f9e50fd4f');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('1bbf45f1-908b-43ff-aedc-9b5e1f402298', '3B-B8-26-98-35-2D', 'temp sensor', 1, 'AA002', 45.8272842, 20.4615173,
        '2019-05-03T20:37:27Z', 'a46b01a5-96de-49b1-930e-7c5ab51bba78');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('ac8c7456-3117-495a-834d-3fb794c63192', '32-8A-E1-7E-26-84', 'mqtt server', 1, 'AA001', 7.3601663, 9.0377612,
        '2020-08-30T11:56:54Z', '2b0a2316-f116-4df9-b29e-d3bc5d8e135e');
insert into nodes (id, addr, name, type, region, lat, long, created, master)
values ('9ad69e46-4447-487c-809d-baba853a1fe5', 'BD-2E-AB-74-15-10', 'igrid monitor', 3, 'AA004', 41.2033027,
        22.5760759, '2020-07-20T22:06:25Z', '73309229-1edf-4e2f-ab5e-f7465c963014');
