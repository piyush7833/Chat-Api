CREATE TYPE call_type AS ENUM ('video', 'audio');
CREATE TABLE IF NOT EXISTS "calls" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "callerId" VARCHAR(255) NOT NULL,
    "receiverId" VARCHAR(255) NOT NULL,
    "type" call_type NOT NULL,
    "duration" INT,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);