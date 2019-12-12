SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
GO
CREATE OR ALTER PROCEDURE [dbo].[SP_QUESTION_BANK_UPDATE_ORDER] (
	@id INT, @serialNo INT
)
AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	IF OBJECT_ID('tempdb..#SortedQuestions') IS NOT NULL DROP TABLE [#SortedQuestions];

	CREATE TABLE [#SortedQuestions] ([serialNo] [INT] IDENTITY(1, 1), [id] [INT]);

	INSERT INTO [#SortedQuestions] ([id])
	SELECT [id]
	FROM [dbo].[QUESTION_BANK] WITH (NOLOCK)
	WHERE [isDeleted] = 0 AND [serialNo] < @serialNo AND [id] <> @id
	ORDER BY [serialNo] ASC, [id] ASC;

	INSERT INTO [#SortedQuestions] ([id]) VALUES (@id);

	INSERT INTO [#SortedQuestions] ([id])
	SELECT [id]
	FROM [dbo].[QUESTION_BANK] WITH (NOLOCK)
	WHERE [isDeleted] = 0 AND [serialNo] >= @serialNo AND [id] <> @id
	ORDER BY [serialNo] ASC, [id] ASC;

	UPDATE [I]
		SET [I].[serialNo] = [II].[serialNo]
	FROM [dbo].[QUESTION_BANK] AS [I] WITH (NOLOCK)
	INNER JOIN [#SortedQuestions] AS [II] WITH (NOLOCK) ON [I].[id] = [II].[id];

	IF OBJECT_ID('tempdb..#SortedQuestions') IS NOT NULL DROP TABLE [#SortedQuestions];
END;
--EXEC [dbo].[SP_QUESTION_BANK_UPDATE_ORDER] 3, 2;
--EXEC [dbo].[SP_QUESTION_BANK_UPDATE_ORDER] 1, 10000;
--SELECT [id], [serialNo], [text], [isDeleted], [createdOn], [createdBy], [updatedOn], [updatedBy] FROM [dbo].[QUESTION_BANK] WITH (NOLOCK) ORDER BY [serialNo] ASC;
