create table if not exists comments (
  id bigserial primary key,
  post_id bigserial not null references posts(id),
  user_id bigserial not null references users(id),
  content text not null,
  created_at timestamp with time zone default now()
);

