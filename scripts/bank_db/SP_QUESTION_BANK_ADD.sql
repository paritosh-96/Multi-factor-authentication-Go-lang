SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
GO
CREATE OR ALTER PROCEDURE [dbo].[SP_QUESTION_BANK_ADD] (
	@text VARCHAR (512), @userID VARCHAR (64)
)
AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	DECLARE @serialNo INT = 0;

	IF EXISTS (SELECT TOP 1 1 FROM [dbo].[QUESTION_BANK] WITH (NOLOCK))
	BEGIN
		SELECT TOP 1 @serialNo = MAX([serialNo]) FROM [dbo].[QUESTION_BANK] WITH (NOLOCK) WHERE [isDeleted] = 0;
	END;

	INSERT INTO [dbo].[QUESTION_BANK] ([serialNo], [text], [createdBy], [updatedBy])
	VALUES ((@serialNo + 1), @text, @userID, @userID);
END;
--EXEC [dbo].[SP_QUESTION_BANK_ADD] 'Where is ur your hometown?', 'CXP_001';
--EXEC [dbo].[SP_QUESTION_BANK_ADD] 'Where is ur watch make?', 'CXP_002';
--EXEC [dbo].[SP_QUESTION_BANK_ADD] 'What is ur Mother''s Name?', 'CXP_003';
--SELECT [id], [serialNo], [text], [isDeleted], [createdOn], [createdBy], [updatedOn], [updatedBy] FROM [dbo].[QUESTION_BANK] WITH (NOLOCK) ORDER BY [serialNo] ASC;
