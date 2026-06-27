
CREATE TYPE user_role as ENUM('DRIVER', 'USER');

CREATE TYPE provider as ENUM (
    'GOOGLE',
    'APPLE',
    'PASSWORD'
);

CREATE type driver_status as ENUM('ONLINE','OFFLINE','BUSY');

CREATE type ride_request_status as ENUM('SEARCHING', 'MATCHED','CANCELLED','COMPLETED');

CREATE TYPE trip_status as ENUM ('COMPLETED', 'ONGOING', 'REJECTED');

CREATE table users (
    id UUID PRIMARY KEY gen_random__uuid(),
    name text not NULL,
    email text not NULL,
    phone INT,
    provider provider,
    created_at TIMESTAMPZ
);

CREATE Table if not exists riders (
   user_id UUID REFERENCES users(id),
   name TEXT,
   phone BIGINT

);



CREATE Table if not exists drivers (
    user_id UUID REFERENCES users(id),
    vehicle_id text,
    license_number INT,
    status driver_status,
    latitude TEXT ,
    longitude TEXT,
    updated_at timestampz
);

CREATE Table ride_requests(
    id UUID PRIMARY key DEFAULT gen_random_uuid(),
    rider_id UUID REFERENCES riders(id),
    pickup_lat TEXT,
    pickup_long TEXT,
    destination_lat TEXT,
    destination_lat TEXT,
    status ride_request_status,
    created_at TIMESTAMPZ
);

CREATE Table trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ride_request_id UUID REFERENCES ride_request(id),
    driver_id UUID REFERENCES  drivers(id),
    status trips_status,
    started_at TIMESTAMPZ DEFAULT now(),
    completed_at TIMESTAMPZ
);

