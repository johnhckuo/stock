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

