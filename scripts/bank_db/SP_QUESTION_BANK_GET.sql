SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
GO
CREATE OR ALTER PROCEDURE [dbo].[SP_QUESTION_BANK_GET] (
	@questionId int
)
AS
	BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;
	declare @listOfIDs table (id int);

	insert @listOfIDs(id) values(1),(2),(3);

	SELECT [id], [text]
	FROM [dbo].[QUESTION_BANK] WITH (NOLOCK)
	WHERE [isDeleted] = 0 AND [id] = @questionId
	ORDER BY [serialNo] ASC, [id] ASC;
END;
