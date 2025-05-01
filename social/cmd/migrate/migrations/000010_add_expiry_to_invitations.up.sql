alter table
  invitations
add
  column expiry timestamp(0) with time zone not null;