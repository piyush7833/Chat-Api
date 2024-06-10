ALTER TABLE "Block" ADD CONSTRAINT "FK_Block_blockerId_User_id" FOREIGN KEY ("blockerId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Block" ADD CONSTRAINT "FK_Block_blockedId_User_id" FOREIGN KEY ("blockedId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Call" ADD CONSTRAINT "FK_Call_callerId_User_id" FOREIGN KEY ("callerId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Call" ADD CONSTRAINT "FK_Call_receiverId_User_id" FOREIGN KEY ("receiverId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "FriendRequest" ADD CONSTRAINT "FK_FriendRequest_senderId_User_id" FOREIGN KEY ("senderId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "FriendRequest" ADD CONSTRAINT "FK_FriendRequest_receiverId_User_id" FOREIGN KEY ("receiverId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Friend" ADD CONSTRAINT "FK_Friend_userId_User_id" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Friend" ADD CONSTRAINT "FK_Friend_friendId_User_id" FOREIGN KEY ("friendId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Media" ADD CONSTRAINT "FK_Media_senderId_User_id" FOREIGN KEY ("senderId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Media" ADD CONSTRAINT "FK_Media_receiverId_User_id" FOREIGN KEY ("receiverId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Message" ADD CONSTRAINT "FK_Message_senderId_User_id" FOREIGN KEY ("senderId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Message" ADD CONSTRAINT "FK_Message_receiverId_User_id" FOREIGN KEY ("receiverId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Notification" ADD CONSTRAINT "FK_Notification_userId_User_id" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Reminder" ADD CONSTRAINT "FK_Reminder_senderId_User_id" FOREIGN KEY ("senderId") REFERENCES "User" ("id") ON DELETE CASCADE;
ALTER TABLE "Reminder" ADD CONSTRAINT "FK_Reminder_receiverId_User_id" FOREIGN KEY ("receiverId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Status" ADD CONSTRAINT "FK_Status_userId_User_id" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE;

ALTER TABLE "Tag" ADD CONSTRAINT "FK_Tag_messageId_Message_id" FOREIGN KEY ("messageId") REFERENCES "Message" ("id") ON DELETE SET NULL;
ALTER TABLE "Tag" ADD CONSTRAINT "FK_Tag_mediaId_Media_id" FOREIGN KEY ("mediaId") REFERENCES "Media" ("id") ON DELETE SET NULL;
ALTER TABLE "Tag" ADD CONSTRAINT "FK_Tag_reminderId_Reminder_id" FOREIGN KEY ("reminderId") REFERENCES "Reminder" ("id") ON DELETE SET NULL;
