CREATE TABLE IF NOT EXISTS "notificationId" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255) NOT NULL,
    "deviceId" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);
