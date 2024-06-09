CREATE TABLE IF NOT EXISTS "Friend" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255),
    "friendId" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "unique_friend_pair" UNIQUE ("userId", "friendId"),
    PRIMARY KEY ("id")
);
