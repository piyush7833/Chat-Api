CREATE TABLE IF NOT EXISTS "tags" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "messageId" VARCHAR(255),
    "reminderId" VARCHAR(255),
    "title" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);