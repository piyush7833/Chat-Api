CREATE TABLE IF NOT EXISTS "reminders" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255) NOT NULL,
    "recieverId" VARCHAR(255) NOT NULL,
    "time" TIMESTAMPTZ,
    "tune" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);