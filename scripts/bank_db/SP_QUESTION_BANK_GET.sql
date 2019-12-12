SET ANSI_NULLS ON;
SET QUOTED_IDENTIFIER ON;
GO
CREATE OR ALTER PROCEDURE [dbo].[SP_QUESTION_BANK_GET]
AS
BEGIN
	SET ANSI_NULLS ON;
	SET ANSI_PADDING ON;
	SET QUOTED_IDENTIFIER ON;
	SET NOCOUNT ON;

	SELECT [id], [serialNo], [text]
	FROM [dbo].[QUESTION_BANK] WITH (NOLOCK)
	WHERE [isDeleted] = 0
	ORDER BY [serialNo] ASC, [id] ASC;
END;
--EXEC [dbo].[SP_QUESTION_BANK_GET];
