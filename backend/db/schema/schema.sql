
CREATE TYPE user_role as ENUM('DRIVER', 'USER');

CREATE TYPE provider as ENUM (
    'GOOGLE',
    'APPLE',
    'PASSWORD'
);

CREATE type driver_status as ENUM('ONLINE','OFFLINE','BUSY');

CREATE type ride_request_status as ENUM('SEARCHING', 'MATCHED','CANCELLED','COMPLETED');

CREATE TYPE trip_status as ENUM ('COMPLETED', 'ONGOING', 'REJECTED');

CREATE type vehicle_type as enum('SEDAN', 'SUV');

CREATE table users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name text not NULL,
    role user_role NOT NULL,
    email text not NULL UNIQUE,
    phone VARCHAR(15),
    password_hash TEXT,
    provider provider,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE Table rider_profile (
   user_id UUID PRIMARY KEY REFERENCES users(id),
   rating DOUBLE precision
);

CREATE Table if not exists driver_profile (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    license_number INT,
    status driver_status,
    rating DOUBLE precision,
    updated_at timestamptz
);


CREATE Table if not exists vehicles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    driver_id UUID REFERENCES driver_profile(user_id),
    model TEXT,
    plate_number TEXT,
    type vehicle_type
);

CREATE Table driver_location (
    driver_id UUID REFERENCES driver_profile(user_id),
    latitude DOUBLE precision,
    longitude DOUBLE precision
);



CREATE Table ride_requests(
    id UUID PRIMARY key DEFAULT gen_random_uuid(),
    rider_id UUID REFERENCES rider_profile(user_id),
    pickup_lat DOUBLE precision,
    pickup_long DOUBLE precision,
    destination_lat DOUBLE precision,
    destination_long DOUBLE precision,
    status ride_request_status,
    created_at TIMESTAMPtZ
);

CREATE Table trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ride_request_id UUID REFERENCES ride_requests(id),
    driver_id UUID REFERENCES  driver_profile(id),
    status trip_status,
    started_at TIMESTAMPtZ DEFAULT now(),
    completed_at TIMESTAMPtZ
);

