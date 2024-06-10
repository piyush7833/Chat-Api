CREATE TABLE IF NOT EXISTS  "User" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "name" VARCHAR(255),
    "image" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "phone" VARCHAR(255) UNIQUE,
    "password" VARCHAR(255),
    "socketId" VARCHAR(255),
    "lastSeen" TIMESTAMPTZ,
    "online" BOOLEAN NOT NULL DEFAULT FALSE,
    "phoneVerified" BOOLEAN NOT NULL DEFAULT FALSE,
    "emailVerified" BOOLEAN NOT NULL DEFAULT FALSE,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
