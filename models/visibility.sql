CREATE TYPE visibility_type AS ENUM ('profile', 'lastseen','status','readreceipts');
CREATE TABLE IF NOT EXISTS "visibility" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255) NOT NULL,
    "type" visibility_type,
    "friendIds" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT unique_user_friend UNIQUE ("userId", "friendIds")
);