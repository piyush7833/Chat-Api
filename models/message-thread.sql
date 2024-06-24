CREATE TABLE IF NOT EXISTS "messageThreads" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "messageId" VARCHAR(255),
    "createdBy" VARCHAR(255) NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);