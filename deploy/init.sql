DROP DATABASE IF EXISTS timetable;    

CREATE DATABASE timetable;

\c timetable;        

CREATE TABLE Users (
    ID bigint GENERATED ALWAYS AS IDENTITY,
    Name character varying(50) NOT NULL,
    PRIMARY KEY(ID)
);

CREATE TABLE Timeslots (
    ID bigint GENERATED ALWAYS AS IDENTITY,
    StartAt bigint NOT NULL,
    EndAt bigint NOT NULL,
    UserId bigint,
    PRIMARY KEY(ID, StartAt, Endat),
    CONSTRAINT fk_id
      FOREIGN KEY(UserId) 
	  REFERENCES Users(ID)
	  ON DELETE CASCADE
);

create or replace function add_timeslot(user_id int, start_at int, end_at int)
returns int
language plpgsql
as
$$
declare
   overlapped_count integer;
   new_id integer;
begin
   select * 
   into overlapped_count
   from public.timeslots
   where EndAt > start_at and StartAt < end_at and UserID = user_id;
   
  if found then
     new_id := -1;
  else
     insert into public.timeslots (UserID, StartAt, EndAt) VALUES (user_id, start_at, end_at) RETURNING ID into new_id; 
  end if;

   return new_id;
end;
$$;

