CREATE TYPE message_type AS ENUM ('message', 'video', 'audio', 'image');
CREATE TABLE IF NOT EXISTS "Message" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255),
    "receiverId" VARCHAR(255),
    "content" text,
    "type" message_type,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
