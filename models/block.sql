CREATE TABLE IF NOT EXISTS "Block" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "blockerId" VARCHAR(255),
    "blockedId" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "unique_blocker_blocked_pair" UNIQUE ("blockerId", "blockedId"),
    PRIMARY KEY ("id")
);