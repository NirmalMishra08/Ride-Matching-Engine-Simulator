
CREATE TYPE user_role as ENUM('DRIVER', 'USER');

CREATE TYPE provider as ENUM (
    'GOOGLE',
    'APPLE',
    'PASSWORD'
);

CREATE Table if not exists users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   name text not null,
   email text UNIQUE NOT NULL,
   role user_role DEFAULT 'USER',
   password_hash TEXT ,
    provider provider NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);