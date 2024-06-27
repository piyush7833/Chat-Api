CREATE TABLE IF NOT EXISTS "reminders" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255) NOT NULL,
    "receiverId" VARCHAR(255) NOT NULL,
    "time" TIMESTAMPTZ  NOT NULL,
    "tune" VARCHAR(255) DEFAULT 'default' NOT NULL,
    "message" TEXT DEFAULT 'reminder' NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);