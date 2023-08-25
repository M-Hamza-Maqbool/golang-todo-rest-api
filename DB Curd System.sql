USE todo;
ALTER TABLE Pallet ALTER COLUMN CreatedByUserGuid NVARCHAR(MAX) NULL;


-- ALTER TABLE [User] ALTER COLUMN IsDeleted INT NULL;
-- ALTER TABLE [User] ALTER COLUMN CreatedOn DATETIME2(7) NULL;
-- ALTER TABLE [User] ALTER COLUMN CreatedByUserName NVARCHAR(MAX) NULL;
-- ALTER TABLE [User] ALTER COLUMN CreatedByUserGuid NVARCHAR(MAX) NULL;
-- ALTER TABLE [User] ALTER COLUMN ModifiedOn DATETIME2(7) NULL;
-- ALTER TABLE [User] ALTER COLUMN ModifiedByUserName NVARCHAR(MAX) NULL;
-- ALTER TABLE [User] ALTER COLUMN ModifiedByUserGuid NVARCHAR(MAX) NULL;



-- Create User table
CREATE TABLE [User] (
	LoginName NVARCHAR(255) NOT NULL PRIMARY KEY,
    DisplayName NVARCHAR(MAX) NOT NULL,
    Password NVARCHAR(MAX) NOT NULL,
    UserType INT NOT NULL,
    IsLocked INT NOT NULL,
    Status INT NOT NULL,
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserName NVARCHAR(MAX) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    ModifiedOn DATETIME2(7) NOT NULL,
    ModifiedByUserName NVARCHAR(MAX) NOT NULL,
    ModifiedByUserGuid UNIQUEIDENTIFIER NOT NULL
);


-- Create PalletTracking table
CREATE TABLE PalletTracking (
    PalletTrackingId INT NOT NULL PRIMARY KEY,
    PalletId INT NOT NULL,
    PalletLocation VARCHAR(50) NOT NULL,
    PalletType VARCHAR(50) NOT NULL,
    Remarks NVARCHAR(MAX),
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName VARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName VARCHAR(MAX)
);


-- Create MovePalletTransaction table
CREATE TABLE MovePalletTransaction (
    MovePalletTransactionId INT NOT NULL PRIMARY KEY,
    PalletTrackingId INT NOT NULL,
    MovePalletLocation NVARCHAR(50) NOT NULL,
    AssignedBy NVARCHAR(MAX) NOT NULL,
    AssignedOn DATETIME2(7) NOT NULL,
    Remarks NVARCHAR(MAX),
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName VARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName VARCHAR(MAX)
);

-- Create ClosePalletTransaction table
CREATE TABLE ClosePalletTransaction (
    ClosePalletTransactionId INT NOT NULL PRIMARY KEY,
    PalletTrackingId INT NOT NULL,
    Remarks NVARCHAR(MAX),
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName VARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName VARCHAR(MAX)
);


CREATE TABLE OpenPalletTransaction (
    OpenPalletTransactionId INT NOT NULL PRIMARY KEY,
    PalletTrackingId INT NOT NULL,
    OpenPalletLocation NVARCHAR(50) NOT NULL,
    Remarks NVARCHAR(MAX),
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName VARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName VARCHAR(MAX)
);


CREATE TABLE Pallet (
    PalletId INT NOT NULL PRIMARY KEY,
    PalletNo VARCHAR(50) NOT NULL,
    Area VARCHAR(50) NOT NULL,
    IsActive INT,
    IsLocked INT NOT NULL,
    Status INT NOT NULL,
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName NVARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName NVARCHAR(MAX)
);

CREATE TABLE LoadPalletTransaction (
    LoadPalletTransactionId INT NOT NULL PRIMARY KEY,
    PalletTrackingId INT NOT NULL,
    LoadToTruck VARCHAR(50) NOT NULL,
    LoadBy NVARCHAR(MAX) NOT NULL,
    LoadOn DATETIME2(7),
    ReceiveBy NVARCHAR(MAX),
    LoadDocFilePath NVARCHAR(MAX),
    AssignedBy NVARCHAR(MAX) NOT NULL,
    AssignedOn DATETIME2(7) NOT NULL,
    IsDeleted INT NOT NULL,
    CreatedOn DATETIME2(7) NOT NULL,
    CreatedByUserGuid UNIQUEIDENTIFIER NOT NULL,
    CreatedByUserName VARCHAR(MAX) NOT NULL,
    ModifiedOn DATETIME2(7),
    ModifiedByUserGuid UNIQUEIDENTIFIER,
    ModifiedByUserName NVARCHAR(MAX)
);


s