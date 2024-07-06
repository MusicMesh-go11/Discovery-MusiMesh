CREATE TYPE genre AS ENUM ('rock', 'pop', 'jazz', 'classical', 'electronic', 'hip_hop', 'other');

CREATE TABLE composition_metadata (
                                      composition_id UUID NOT NULL,
                                      genre genre,
                                      tags varchar,
                                      listen_count INTEGER DEFAULT 0,
                                      like_count INTEGER DEFAULT 0
);


CREATE TABLE user_interactions (
                                   id uuid default gen_random_uuid() PRIMARY KEY,
                                   user_id uuid NOT NULL,
                                   composition_id uuid NOT NULL ,
                                   interaction_type VARCHAR(20),
                                   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                   deleted_at BIGINT default 0
);
