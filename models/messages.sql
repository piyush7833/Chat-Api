CREATE TYPE message_type AS ENUM ('message', 'video', 'audio', 'image');
CREATE TABLE IF NOT EXISTS "messages" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255) NOT NULL,
    "receiverId" VARCHAR(255),
    "threadId" VARCHAR(255),
    "content" TEXT,
    "type" message_type,
    "isPinned" BOOLEAN NOT NULL DEFAULT FALSE,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
