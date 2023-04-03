create table if not exists customers
(
    id           bigserial
        CONSTRAINT customer_primarykey PRIMARY KEY,
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp NULL     DEFAULT CURRENT_TIMESTAMP,
    customer_id  int       NOT NULL unique,
    name varchar not null,
    mobile_Number varchar   not null unique,
    is_active    boolean
);
create index customer_uindex
    on customers(customer_id);

create table if not exists accounts
(
    id          bigserial
        CONSTRAINT account_primarykey PRIMARY KEY,
    created_at   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp NULL     DEFAULT CURRENT_TIMESTAMP,
    account_id  int       NOT NULL unique,
    balance     BIGINT    NOT NULL,
    is_active   boolean,
    customer_id int       NOT NULL,
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id)
            REFERENCES customers (customer_id)
            ON DELETE RESTRICT
);
create unique index  accounts_uindex
    on accounts(account_id);

create table if not exists cards
(
    id         bigserial
        CONSTRAINT card_primarykey PRIMARY KEY,
    created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL     DEFAULT CURRENT_TIMESTAMP,
    card_id    BIGINT       NOT NULL unique,
    account_id int       NOT NULL,
    is_active  boolean,
    CONSTRAINT fk_account
        FOREIGN KEY (account_id)
            REFERENCES accounts (account_id)
            ON DELETE RESTRICT
);
create unique index  card_uindex
    on cards(card_id);

create table if not exists transaction_rules
(
    id         bigserial
        CONSTRAINT transactionRules_primarykey PRIMARY KEY,
    transaction varchar   not null ,
    created_at  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL     DEFAULT CURRENT_TIMESTAMP,
    fee        bigint  default 0,
    max_limit   bigint default 0,
    Min_limit   bigint default 0,
    template_sms  text default null,
    is_active  boolean
);
create unique index  transaction_rules_uindex
    on transaction_rules(transaction);

create table if not exists transactions
(
    id         bigserial
        CONSTRAINT transaction_primarykey PRIMARY KEY,
    created_at  timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NULL     DEFAULT current_timestamp,
    card_id BIGINT       NOT NULL,
    to_card_id BIGINT       NOT NULL,
    amount int not null ,
    is_active  boolean,
    CONSTRAINT fk_main_card
            FOREIGN KEY (card_id)
            REFERENCES cards (card_id)
            ON DELETE no action,
     CONSTRAINT fk_other_card
            FOREIGN KEY (to_card_id)
            REFERENCES cards (card_id)
            ON DELETE no action
);

create unique index  transactions_time_card_uindex
    on transactions(created_at desc ,card_id asc);


insert into customers(customer_id, name,mobile_number, is_active)
values (111, 'Pariya', 09364046601,true),(222, 'Sara',093640466011,true),(333, 'Steve',09915840507,true);

insert into accounts(account_id, balance, is_active, customer_id)
values(444,10090900,true,111) , (555,500000,true,111), (666,1600000,true,333),
      (777,1600000,true,333);

insert into cards(card_id, account_id, is_active)
values (6219861059454032,444,true),(6037703932049631,444,true),
       (6104337743428383,555,true),(6037997332745158,666,true);

insert into transaction_rules
    (transaction, fee, max_limit, Min_limit, is_active,template_sms)
values ('transfer',500,50000000,10000, true,
        'Dear NAME,'||chr(10)|| 'FUNC Amount: AMOUNT'||chr(10)||
        'Balance: BALANCE'||chr(10)||'Fee:FEE'||chr(10)||'Date:DATE')