CREATE TABLE  IF NOT EXISTS "Call"    (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "callerId" VARCHAR(255),
    "receiverId" VARCHAR(255),
    "duration" INT4,
    "type" VARCHAR(255),
    "status" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "unique_call_at_time_by_same_users" UNIQUE ("callerId", "receiverId", "createdAt"),
    PRIMARY KEY ("id")
);