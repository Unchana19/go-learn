create table if not exists invitations (
  token bytea primary key,
  user_id bigint not null references users(id) on delete cascade
);