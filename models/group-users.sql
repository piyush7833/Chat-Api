CREATE TYPE group_user_role AS ENUM ('admin', 'owner', 'member','banned');
CREATE TABLE IF NOT EXISTS "groupUsers" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255) NOT NULL,
    "groupId" VARCHAR(255) NOT NULL,
    "role" group_user_role NOT NULL DEFAULT 'member',
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT unique_user_group UNIQUE ("userId", "groupId")
);
