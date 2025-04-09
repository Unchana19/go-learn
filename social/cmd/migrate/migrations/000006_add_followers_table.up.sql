create table if not exists followers (
  user_id bigint not null,
  follower_id bigint not null,
  created_at timestamp with time zone default now(),
  primary key (user_id, follower_id),
  foreign key (user_id) references users(id) on delete cascade,
  foreign key (follower_id) references users(id) on delete cascade
);

