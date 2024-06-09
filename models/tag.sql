CREATE TABLE IF NOT EXISTS "Tag" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "messageId" VARCHAR(255),
    "mediaId" VARCHAR(255),
    "reminderId" VARCHAR(255),
    "tagContent" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
