CREATE TABLE IF NOT EXISTS "Media" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255),
    "receiverId" VARCHAR(255),
    "imageUrl" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
