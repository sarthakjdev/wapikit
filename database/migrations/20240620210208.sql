-- Create "OrganisationIntegration" table
CREATE TABLE "public"."OrganisationIntegration" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  PRIMARY KEY ("UniqueId")
);

-- Create enum type "OrganisationMemberRole"
CREATE TYPE "public"."OrganisationMemberRole" AS ENUM ('owner', 'admin', 'member');

-- Create enum type "OauthProviderEnum"
CREATE TYPE "public"."OauthProviderEnum" AS ENUM ('google');

-- Create enum type "ContactStatus"
CREATE TYPE "public"."ContactStatus" AS ENUM ('active', 'inactive');

-- Create enum type "UserAccountStatusEnum"
CREATE TYPE "public"."UserAccountStatusEnum" AS ENUM ('active', 'deleted', 'suspended');

-- Create "User" table
CREATE TABLE "public"."User" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Name" text NOT NULL,
  "Email" text NOT NULL,
  "PhoneNumber" text NULL,
  "Username" text NOT NULL,
  "Password" text NOT NULL,
  "OauthProvider" "public"."OauthProviderEnum" NULL,
  "ProfilePictureUrl" text NULL,
  "Status" "public"."UserAccountStatusEnum" NOT NULL,
  PRIMARY KEY ("UniqueId")
);

-- Create index "UserEmailIndex" to table: "User"
CREATE UNIQUE INDEX "UserEmailIndex" ON "public"."User" ("Email");

-- Create index "UserUsernameIndex" to table: "User"
CREATE UNIQUE INDEX "UserUsernameIndex" ON "public"."User" ("Username");

-- Create enum type "AccessLogType"
CREATE TYPE "public"."AccessLogType" AS ENUM ('web_interface', 'api_access');

-- Create enum type "CampaignStatus"
CREATE TYPE "public"."CampaignStatus" AS ENUM (
  'draft',
  'running',
  'finished',
  'paused',
  'cancelled'
);

-- Create enum type "MessageStatus"
CREATE TYPE "public"."MessageStatus" AS ENUM (
  'sent',
  'delivered',
  'read',
  'failed',
  'undelivered'
);

-- Create enum type "MessageDirection"
CREATE TYPE "public"."MessageDirection" AS ENUM ('inbound', 'outbound');

-- Create enum type "ConversationInitiatedEnum"
CREATE TYPE "public"."ConversationInitiatedEnum" AS ENUM ('contact', 'campaign');

-- Create "Notification" table
CREATE TABLE "public"."Notification" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  PRIMARY KEY ("UniqueId")
);

-- Create "Integration" table
CREATE TABLE "public"."Integration" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  PRIMARY KEY ("UniqueId")
);

-- Create "Organisation" table
CREATE TABLE "public"."Organisation" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Name" text NOT NULL,
  "WebsiteUrl" text NULL,
  "LogoUrl" text NULL,
  "FaviconUrl" text NOT NULL,
  PRIMARY KEY ("UniqueId")
);

-- Create "OrganisationMember" table
CREATE TABLE "public"."OrganisationMember" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Role" "public"."OrganisationMemberRole" NOT NULL,
  "OrganisationId" uuid NOT NULL,
  "UserId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "OrganisationMemberToUserForeignKey" FOREIGN KEY ("UserId") REFERENCES "public"."User" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "OrganisationToOrganisationMemberForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "OrganisationMemberOrganisationIdIndex" to table: "OrganisationMember"
CREATE INDEX "OrganisationMemberOrganisationIdIndex" ON "public"."OrganisationMember" ("OrganisationId");

-- Create index "OrganisationMemberUserIdIndex" to table: "OrganisationMember"
CREATE INDEX "OrganisationMemberUserIdIndex" ON "public"."OrganisationMember" ("UserId");

-- Create "ApiKey" table
CREATE TABLE "public"."ApiKey" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "MemberId" uuid NOT NULL,
  "Key" text NOT NULL,
  "OrganisationId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "ApiKeyToOrganisationMemberForeignKey" FOREIGN KEY ("MemberId") REFERENCES "public"."OrganisationMember" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "ApiKeyToOrganizationForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "ApiKeyIndex" to table: "ApiKey"
CREATE UNIQUE INDEX "ApiKeyIndex" ON "public"."ApiKey" ("Key");

-- Create index "ApiKeyOrganisationIdIndex" to table: "ApiKey"
CREATE INDEX "ApiKeyOrganisationIdIndex" ON "public"."ApiKey" ("OrganisationId");

-- Create index "ApiKeyOrganisationMemberIdIndex" to table: "ApiKey"
CREATE UNIQUE INDEX "ApiKeyOrganisationMemberIdIndex" ON "public"."ApiKey" ("MemberId");

-- Create "Campaign" table
CREATE TABLE "public"."Campaign" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Name" text NOT NULL,
  "Status" "public"."CampaignStatus" NOT NULL DEFAULT 'draft',
  "CreatedByOrganisationMemberId" uuid NOT NULL,
  "OrganisationId" uuid NOT NULL,
  "MessageTemplateId" text NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "CampaignToOrganisationMemberForeignKey" FOREIGN KEY ("CreatedByOrganisationMemberId") REFERENCES "public"."OrganisationMember" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "OrganisationToCampaignForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "CampaignCreatedByOrganisationMemberIdIndex" to table: "Campaign"
CREATE INDEX "CampaignCreatedByOrganisationMemberIdIndex" ON "public"."Campaign" ("CreatedByOrganisationMemberId");

-- Create index "CampaignMessageTemplateIndex" to table: "Campaign"
CREATE INDEX "CampaignMessageTemplateIndex" ON "public"."Campaign" ("MessageTemplateId");

-- Create "ContactList" table
CREATE TABLE "public"."ContactList" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "OrganisationId" uuid NOT NULL,
  "Name" text NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "OrganisationToContactListForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "ContactListOrganisationIdIndex" to table: "ContactList"
CREATE INDEX "ContactListOrganisationIdIndex" ON "public"."ContactList" ("OrganisationId");

-- Create "CampaignList" table
CREATE TABLE "public"."CampaignList" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ContactListId" uuid NOT NULL,
  "CampaignId" uuid NOT NULL,
  PRIMARY KEY ("ContactListId", "CampaignId"),
  CONSTRAINT "CampaignListToCampaignForeignKey" FOREIGN KEY ("CampaignId") REFERENCES "public"."Campaign" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "CampaignListToContactListForeignKey" FOREIGN KEY ("ContactListId") REFERENCES "public"."ContactList" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "Tag" table
CREATE TABLE "public"."Tag" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Label" text NOT NULL,
  "slug" text NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "UniqueSlug" UNIQUE ("slug")
);

-- Create index "slugIndex" to table: "Tag"
CREATE INDEX "slugIndex" ON "public"."Tag" ("slug");

-- Create "CampaignTag" table
CREATE TABLE "public"."CampaignTag" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "CampaignId" uuid NOT NULL,
  "TagId" uuid NOT NULL,
  PRIMARY KEY ("CampaignId", "TagId"),
  CONSTRAINT "CampaignTagToCampaignForeignKey" FOREIGN KEY ("CampaignId") REFERENCES "public"."Campaign" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "CampaignTagToTagForeignKey" FOREIGN KEY ("TagId") REFERENCES "public"."Tag" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "Contact" table
CREATE TABLE "public"."Contact" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "OrganisationId" uuid NOT NULL,
  "Status" "public"."ContactStatus" NOT NULL,
  "Name" text NOT NULL,
  "PhoneNumber" text NOT NULL,
  "Attributes" jsonb NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "OrganisationToContactForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "ContactOrganisationIdIndex" to table: "Contact"
CREATE INDEX "ContactOrganisationIdIndex" ON "public"."Contact" ("OrganisationId");

-- Create index "ContactPhoneNumberIndex" to table: "Contact"
CREATE UNIQUE INDEX "ContactPhoneNumberIndex" ON "public"."Contact" ("PhoneNumber");

-- Create "ContactListContact" table
CREATE TABLE "public"."ContactListContact" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ContactListId" uuid NOT NULL,
  "ContactId" uuid NOT NULL,
  PRIMARY KEY ("ContactListId", "ContactId"),
  CONSTRAINT "ContactListContactToContactForeignKey" FOREIGN KEY ("ContactId") REFERENCES "public"."Contact" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "ContactListContactToContactListForeignKey" FOREIGN KEY ("ContactListId") REFERENCES "public"."ContactList" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "ContactListTag" table
CREATE TABLE "public"."ContactListTag" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ContactListId" uuid NOT NULL,
  "TagId" uuid NOT NULL,
  PRIMARY KEY ("ContactListId", "TagId"),
  CONSTRAINT "ContactListTagToContactListForeignKey" FOREIGN KEY ("ContactListId") REFERENCES "public"."ContactList" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "ContactListTagToTagForeignKey" FOREIGN KEY ("TagId") REFERENCES "public"."Tag" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "WhatsappBusinessAccount" table
CREATE TABLE "public"."WhatsappBusinessAccount" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "AccountId" text NOT NULL,
  "OrganisationId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "WhatsappBusinessAccountToOrganisationForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "WhatsappBusinessAccountAccountIdIndex" to table: "WhatsappBusinessAccount"
CREATE UNIQUE INDEX "WhatsappBusinessAccountAccountIdIndex" ON "public"."WhatsappBusinessAccount" ("AccountId");

-- Create index "WhatsappBusinessAccountOrganisationIdIndex" to table: "WhatsappBusinessAccount"
CREATE INDEX "WhatsappBusinessAccountOrganisationIdIndex" ON "public"."WhatsappBusinessAccount" ("OrganisationId");

-- Create "WhatsappBusinessAccountPhoneNumber" table
CREATE TABLE "public"."WhatsappBusinessAccountPhoneNumber" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "WhatsappBusinessAccountId" uuid NOT NULL,
  "MetaTitle" text NOT NULL,
  "MetaDescription" text NOT NULL,
  "PhoneNumber" text NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "PhoneNumberToWhatsappBusinessAccountForeignKey" FOREIGN KEY ("WhatsappBusinessAccountId") REFERENCES "public"."WhatsappBusinessAccount" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "PhoneNumberPhoneNumberIndex" to table: "WhatsappBusinessAccountPhoneNumber"
CREATE UNIQUE INDEX "PhoneNumberPhoneNumberIndex" ON "public"."WhatsappBusinessAccountPhoneNumber" ("PhoneNumber");

-- Create index "PhoneNumberWhatsappBusinessAccountIdIndex" to table: "WhatsappBusinessAccountPhoneNumber"
CREATE INDEX "PhoneNumberWhatsappBusinessAccountIdIndex" ON "public"."WhatsappBusinessAccountPhoneNumber" ("WhatsappBusinessAccountId");

-- Create "Conversation" table
CREATE TABLE "public"."Conversation" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ContactId" uuid NOT NULL,
  "WhatsappBusinessAccountPhoneNumberId" uuid NOT NULL,
  "InitiatedBy" "public"."ConversationInitiatedEnum" NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "ConversationToContactForeignKey" FOREIGN KEY ("ContactId") REFERENCES "public"."Contact" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "ConversationToWhatsappBusinessAccountPhoneNumberForeignKey" FOREIGN KEY ("WhatsappBusinessAccountPhoneNumberId") REFERENCES "public"."WhatsappBusinessAccountPhoneNumber" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "ConversationContactIdIndex" to table: "Conversation"
CREATE INDEX "ConversationContactIdIndex" ON "public"."Conversation" ("ContactId");

-- Create index "ConversationWhatsappBusinessAccountPhoneNumberIdIndex" to table: "Conversation"
CREATE INDEX "ConversationWhatsappBusinessAccountPhoneNumberIdIndex" ON "public"."Conversation" ("WhatsappBusinessAccountPhoneNumberId");

-- Create "ConversationTag" table
CREATE TABLE "public"."ConversationTag" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ConversationId" uuid NOT NULL,
  "TagId" uuid NOT NULL,
  PRIMARY KEY ("ConversationId", "TagId"),
  CONSTRAINT "ConversationTagToConversationForeignKey" FOREIGN KEY ("ConversationId") REFERENCES "public"."Conversation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "ConversationTagToTagForeignKey" FOREIGN KEY ("TagId") REFERENCES "public"."Tag" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "Message" table
CREATE TABLE "public"."Message" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "ConversationId" uuid NULL,
  "CampaignId" uuid NULL,
  "ContactId" uuid NOT NULL,
  "WhatsappBusinessAccountPhoneNumberId" uuid NOT NULL,
  "Direction" "public"."MessageDirection" NOT NULL,
  "Content" text NOT NULL,
  "Status" "public"."MessageStatus" NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "MessageToCampaignForeignKey" FOREIGN KEY ("CampaignId") REFERENCES "public"."Campaign" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "MessageToContactForeignKey" FOREIGN KEY ("ContactId") REFERENCES "public"."Contact" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "MessageToConversationForeignKey" FOREIGN KEY ("ConversationId") REFERENCES "public"."Conversation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "MessageToWhatsappBusinessAccountPhoneNumberForeignKey" FOREIGN KEY ("WhatsappBusinessAccountPhoneNumberId") REFERENCES "public"."WhatsappBusinessAccountPhoneNumber" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "MessageCampaignIdIndex" to table: "Message"
CREATE INDEX "MessageCampaignIdIndex" ON "public"."Message" ("CampaignId");

-- Create index "MessageContactIdIndex" to table: "Message"
CREATE INDEX "MessageContactIdIndex" ON "public"."Message" ("ContactId");

-- Create index "MessageWhatsappBusinessAccountPhoneNumberIdIndex" to table: "Message"
CREATE INDEX "MessageWhatsappBusinessAccountPhoneNumberIdIndex" ON "public"."Message" ("WhatsappBusinessAccountPhoneNumberId");

-- Create "MessageReply" table
CREATE TABLE "public"."MessageReply" (
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "MessageId" uuid NOT NULL,
  "ReplyMessageId" uuid NOT NULL,
  PRIMARY KEY ("MessageId", "ReplyMessageId"),
  CONSTRAINT "MessageReplyToMessageForeignKey" FOREIGN KEY ("MessageId") REFERENCES "public"."Message" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "MessageReplyToReplyMessageForeignKey" FOREIGN KEY ("ReplyMessageId") REFERENCES "public"."Message" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "OrganisationRole" table
CREATE TABLE "public"."OrganisationRole" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "Name" text NOT NULL,
  "OrganisationId" uuid NOT NULL,
  "Permissions" text [] NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "OrganisationToOrganisationRoleForeignKey" FOREIGN KEY ("OrganisationId") REFERENCES "public"."Organisation" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "OrganisationRoleOrganisationIdIndex" to table: "OrganisationRole"
CREATE INDEX "OrganisationRoleOrganisationIdIndex" ON "public"."OrganisationRole" ("OrganisationId");

-- Create "RoleAssignment" table
CREATE TABLE "public"."RoleAssignment" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "OrganisationRoleId" uuid NOT NULL,
  "OrganisationMemberId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "OrganisationMemberToRoleAssignmentForeignKey" FOREIGN KEY ("OrganisationMemberId") REFERENCES "public"."OrganisationMember" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "OrganisationRoleToRoleAssignmentForeignKey" FOREIGN KEY ("OrganisationRoleId") REFERENCES "public"."OrganisationRole" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "RoleAssignmentOrganisationMemberIdIndex" to table: "RoleAssignment"
CREATE INDEX "RoleAssignmentOrganisationMemberIdIndex" ON "public"."RoleAssignment" ("OrganisationMemberId");

-- Create index "RoleAssignmentOrganisationRoleIdIndex" to table: "RoleAssignment"
CREATE INDEX "RoleAssignmentOrganisationRoleIdIndex" ON "public"."RoleAssignment" ("OrganisationRoleId");

-- Create "TrackLink" table
CREATE TABLE "public"."TrackLink" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "CampaignId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "TrackLinkToCampaignForeignKey" FOREIGN KEY ("CampaignId") REFERENCES "public"."Campaign" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "TrackLinkCampaignIdIndex" to table: "TrackLink"
CREATE INDEX "TrackLinkCampaignIdIndex" ON "public"."TrackLink" ("CampaignId");

-- Create "TrackLinkClick" table
CREATE TABLE "public"."TrackLinkClick" (
  "UniqueId" uuid NOT NULL,
  "CreatedAt" timestamp NOT NULL,
  "UpdatedAt" timestamp NOT NULL,
  "TrackLinkId" uuid NOT NULL,
  "ContactId" uuid NOT NULL,
  PRIMARY KEY ("UniqueId"),
  CONSTRAINT "TrackLinkClickToContactForeignKey" FOREIGN KEY ("ContactId") REFERENCES "public"."Contact" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "TrackLinkClickToTrackLinkForeignKey" FOREIGN KEY ("TrackLinkId") REFERENCES "public"."TrackLink" ("UniqueId") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create index "TrackLinkClickContactIdIndex" to table: "TrackLinkClick"
CREATE INDEX "TrackLinkClickContactIdIndex" ON "public"."TrackLinkClick" ("ContactId");

-- Create index "TrackLinkClickTrackLinkIdIndex" to table: "TrackLinkClick"
CREATE INDEX "TrackLinkClickTrackLinkIdIndex" ON "public"."TrackLinkClick" ("TrackLinkId");